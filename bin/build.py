#!/usr/bin/env python3
import functools
import shutil
import os
import sys
import subprocess
import glob
import click
import re
import yaml

from pathlib import Path

import openapi


PROJECT = 'spaceone'
OUTPUT_DOC_FORMAT = 'json'
BASE_DIR = os.path.dirname(os.path.dirname(os.path.abspath(__file__)))
TEMPLATE_DIR = os.path.join(BASE_DIR, 'template')
PROTO_DIR = os.path.join(BASE_DIR, 'proto')
OUTPUT_DIR = os.path.join(BASE_DIR, 'dist')
ARTIFACT_DIR = os.path.join(BASE_DIR, 'artifact')
VERSION = os.path.join(BASE_DIR, 'VERSION')
AVAILABLE_CODES = ['all', 'python', 'go', 'gateway', 'json', 'openapi']
DEFAULT_THIRD_PARTY_DIR = 'third_party/googleapis:third_party/protobuf/src'
DEFAULT_CODE = 'all'
REPOSITORY_URL = 'github.com/cloudforet-io/api'
GO_MODULE_PATH = f'{REPOSITORY_URL}/dist'
GO_PREFIX_IMPORT_PATH = f'{GO_MODULE_PATH}/go'

YAML_LOADER = yaml.Loader
YAML_LOADER.add_implicit_resolver(
    u'tag:yaml.org,2002:float',
    re.compile(u'''^(?:
     [-+]?(?:[0-9][0-9_]*)\\.[0-9_]*(?:[eE][-+]?[0-9]+)?
    |[-+]?(?:[0-9][0-9_]*)(?:[eE][-+]?[0-9]+)
    |\\.[0-9_]+(?:[eE][-+][0-9]+)?
    |[-+]?[0-9][0-9_]*(?::[0-5]?[0-9])+\\.[0-9_]*
    |[-+]?\\.(?:inf|Inf|INF)
    |\\.(?:nan|NaN|NAN))$''', re.X),
    list(u'-+0123456789.'))


def load_yaml(yaml_str: str) -> dict:
    try:
        return yaml.load(yaml_str, Loader=YAML_LOADER)
    except Exception:
        raise ValueError(f'YAML Load Error: {yaml_str}')


def load_yaml_from_file(yaml_file: str):
    try:
        with open(yaml_file, 'r') as f:
            return load_yaml(f.read())
    except Exception:
        raise Exception(f'YAML Load Error: {yaml_file}')


def _error(msg):
    print()
    print('[ERROR] %s' % (msg))
    print()
    sys.exit(1)


def _get_default_third_party_dir():
    default_third_party_dir = []

    for third_party_dir in os.environ.get('DEFAULT_THIRD_PARTY_DIR', DEFAULT_THIRD_PARTY_DIR).split(':'):
        if third_party_dir.strip() != '':
            default_third_party_dir.append(os.path.join(BASE_DIR, third_party_dir))

    return default_third_party_dir


def _get_services_from_target(target):
    _services = []

    api_path_in_proto = os.path.join(PROTO_DIR, 'spaceone', 'api')

    # Get all services name from proto path
    for service in os.listdir(api_path_in_proto):
        service_path_in_proto = os.path.join(api_path_in_proto, service)
        if os.path.isdir(service_path_in_proto):
            _services.append(service)

    if target == 'all':
        return _services
    elif target in _services:
        return [target]
    else:
        _error(f"Target({target}) is not found.")


def _get_proto_files(proto_path):
    return [proto_file for proto_file in glob.iglob(os.path.join(proto_path, '**', '*.proto'), recursive=True)]


def _get_proto_path_list(proto_dir, third_party_dir):
    proto_path_list = [proto_dir]
    proto_path_list.extend(_get_default_third_party_dir())
    proto_path_list.extend(third_party_dir)
    return proto_path_list


def _make_output_path(output_dir, code):
    output_path = os.path.join(output_dir, code)

    if not os.path.exists(output_path):
        os.makedirs(output_path, exist_ok=True)


def _get_package_path(proto_path, package_path_list):
    parent_dir, current_dir = os.path.split(proto_path)
    package_path_list.insert(0, current_dir)

    if current_dir == PROJECT:
        return package_path_list
    else:
        _get_package_path(parent_dir, package_path_list)

    return package_path_list


def _make_package_path(proto_file, output_path):
    package_path_list = _get_package_path(os.path.dirname(proto_file), [])
    output_path = os.path.join(output_path, *package_path_list)
    os.makedirs(output_path, exist_ok=True)

    return output_path


def _get_generate_codes(code):
    if 'all' in code:
        return list(filter(lambda x: x != 'all', AVAILABLE_CODES))
    else:
        return list(code)


def _execute_command(cmd, cwd=None):
    try:
        subprocess.check_output(cmd, cwd=cwd)
    except Exception:
        _error(f"Failed to execute {cmd} command.")


def _make_go_mod(output_dir):
    _execute_command(['go', 'mod', 'init', GO_MODULE_PATH], cwd=output_dir)
    _execute_command(['go', 'mod', 'edit', '-replace', f"{REPOSITORY_URL}=./"], cwd=output_dir)
    _execute_command(['go', 'mod', 'tidy'], cwd=output_dir)


def _make_build_environment_python(output_dir, code):
    # Copy setup.py
    src = os.path.join(TEMPLATE_DIR, 'python', 'setup.py.tmpl')
    dst = os.path.join(output_dir, code, 'setup.py')

    if not os.path.exists(dst):
        shutil.copyfile(src, dst)

    # Copy spaceone.__init__.py
    src = os.path.join(TEMPLATE_DIR, 'python', 'spaceone.tmpl')
    dst = os.path.join(output_dir, code, 'spaceone', '__init__.py')

    if not os.path.exists(dst):
        shutil.copyfile(src, dst)

    # Copy spaceone.api.__init__.py
    src = os.path.join(TEMPLATE_DIR, 'python', 'spaceone.api.tmpl')
    dst = os.path.join(output_dir, code, 'spaceone', 'api', '__init__.py')

    if not os.path.exists(dst):
        shutil.copyfile(src, dst)

    api_root_dir = os.path.join(output_dir, code, 'spaceone', 'api')
    for root_dir, sub_dirs, files in os.walk(api_root_dir):
        if os.path.basename(root_dir) != '__pycache__':
            init_path = os.path.join(root_dir, '__init__.py')

            if not os.path.exists(init_path):
                Path(init_path).touch()


def _make_build_environment_go(output_dir, code):
    api_root_dir = os.path.join(output_dir, code, 'spaceone', 'api')


def _make_build_environment_gateway(output_dir, code):
    _make_go_mod(output_dir)


def _make_build_environment_json(output_dir, code):
    pass


def _make_build_environment(output_dir, code):
    if code == 'python':
        _make_build_environment_python(output_dir, code)
    elif code == 'go':
        _make_build_environment_go(output_dir, code)
    elif code == 'gateway':
        _make_build_environment_gateway(output_dir, code)
    elif code == 'json':
        _make_build_environment_json(output_dir, code)
    elif code == 'openapi':
        openapi.make_build_environment_openapi(output_dir, code)


def _python_compile(proto_file, output_path, proto_path_list, debug):
    cmd = ['python3', '-m', 'grpc_tools.protoc', f'--python_out={output_path}', f'--grpc_python_out={output_path}']

    for proto_path in proto_path_list:
        cmd.append(f'--proto_path={proto_path}')
    cmd.append(proto_file)

    if debug:
        print()
        print(' '.join(cmd))

    try:
        subprocess.check_output(cmd)
    except Exception as e:
        _error(f"[ERROR] Failed to Python compile : {proto_file}\n{e}")

    print(f"[SUCCESS] Python Compile : {proto_file}")


def _go_compile(proto_file, output_path, proto_path_list, debug):
    cmd = ['protoc', f'--go_out={output_path}', f'--go_opt=module={GO_PREFIX_IMPORT_PATH}',
           f'--go-grpc_out={output_path}', f'--go-grpc_opt=module={GO_PREFIX_IMPORT_PATH}']

    for proto_path in proto_path_list:
        cmd.append(f'--proto_path={proto_path}')

    cmd.append(proto_file)

    if debug:
        print()
        print(' '.join(cmd))

    try:
        subprocess.check_output(cmd)
    except Exception as e:
        _error(f"[ERROR] Failed to Go compile : {proto_file}\n{e}")

    print(f"[SUCCESS] Go Compile : {proto_file}")


def _go_grpc_gateway_compile(proto_file, output_path, proto_path_list, debug):
    cmd = ['protoc', '--grpc-gateway_opt=logtostderr=true', '--grpc-gateway_opt=generate_unbound_methods=true',
           '--grpc-gateway_opt=standalone=true', f'--grpc-gateway_opt=module={GO_PREFIX_IMPORT_PATH}',
           f'--grpc-gateway_out={output_path}']

    for proto_path in proto_path_list:
        cmd.append(f'--proto_path={proto_path}')

    cmd.append(proto_file)

    if debug:
        print()
        print(' '.join(cmd))

    try:
        subprocess.check_output(cmd)
    except Exception as e:
        _error(f"Failed to gRPC Gateway Compile : {proto_file}\n{e}")

    print(f"[SUCCESS] gRPC Gateway Compile : {proto_file}")


def _doc_compile(proto_file, output_path, proto_path_list, debug):
    # check GOPATH
    if not os.environ.get('GOBIN'):
        _error('\'protoc-gen-doc\' does not exist.')

    # if migration to cloudforet end this will be changed
    # output_path = _make_package_path(proto_file, output_path)
    output_path = os.path.join(output_path, 'cloudforet', '/'.join(proto_file.split('/')[-4:-1]))
    os.makedirs(output_path, exist_ok=True)

    gobin = os.path.join(os.environ.get('GOBIN'), 'protoc-gen-doc')
    cmd = ['protoc', f'--plugin=protoc-gen-doc={gobin}', f'--doc_out={output_path}']
    for proto_path in proto_path_list:
        cmd.append(f'--proto_path={proto_path}')

    opt_file_name = proto_file.split('/')[-1].split('.')[0].capitalize()

    cmd.append(f'--doc_opt={OUTPUT_DOC_FORMAT},{opt_file_name}.{OUTPUT_DOC_FORMAT}')
    cmd.append(proto_file)

    if debug:
        print()
        print(' '.join(cmd))

    try:
        subprocess.check_output(cmd)
    except Exception as e:

        _error(f" Failed to Document Compile : {proto_file}\n{e}")

    print(f"[SUCCESS] Document Compile : {proto_file}")


def _compile_code(params, code, proto_file=None, service=None):
    output_path = os.path.join(params['output_dir'], code)

    if code == 'python':
        _python_compile(proto_file, output_path, params['proto_path_list'], debug=params['debug'])

    elif code == 'go':
        _go_compile(proto_file, output_path, params['proto_path_list'], debug=params['debug'])

    elif code == 'gateway':
        _go_grpc_gateway_compile(proto_file, output_path, params['proto_path_list'], debug=params['debug'])

    elif code == 'json':
        _doc_compile(proto_file, output_path, params['proto_path_list'], debug=params['debug'])

    elif code == 'openapi':
        openapi.openapi_compile(params['output_dir'], code, service, debug=params['debug'])


@click.command()
@click.argument('target', default='all')
@click.option('-p', '--proto-dir', type=str, help='Protocol Buffers Directory.', default=PROTO_DIR)
@click.option('-t', '--third-party-dir', type=str, help='Third Party Protocol Buffers Directory.', multiple=True, default=[])
@click.option('-o', '--output-dir', type=str, help='Output Directory.', default=OUTPUT_DIR)
@click.option('-c', '--code', type=click.Choice(AVAILABLE_CODES), help='Generate Code.', multiple=True, default=[DEFAULT_CODE])
@click.option('-d', '--debug', help='Debug Mode.', is_flag=True)
def build(**params):
    """
    Cloudforet API Builder\n
    TARGET, 'all or specific service. (core, identity, inventory, etc.)'
    """

    params['target'] = _get_services_from_target(params['target'])
    params['proto_path_list'] = _get_proto_path_list(params['proto_dir'], params['third_party_dir'])

    for code in _get_generate_codes(params['code']):
        _make_output_path(params['output_dir'], code)

        for _tg in params['target']:
            if code == 'openapi':
                _compile_code(params, code, service=_tg)
            else:
                proto_files = _get_proto_files(os.path.join(params['proto_dir'], PROJECT, 'api', _tg))
                list(map(functools.partial(_compile_code, params, code), proto_files))

        _make_build_environment(params['output_dir'], code)


if __name__ == '__main__':
    build()
