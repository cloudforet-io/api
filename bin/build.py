#!/usr/bin/env python3
# -*- coding: utf-8 -*-

import shutil
import os
import sys
import argparse
import subprocess
import glob
from pathlib import Path
from shutil import copyfile

PROJECT = 'spaceone'
EXTENSION = ['json', 'markdown']
OUTPUT_TO_DOC = EXTENSION[0]
BASE_DIR = os.path.dirname(os.path.dirname(os.path.abspath(__file__)))
TEMPLATE_DIR = os.path.join(BASE_DIR, 'template')
PROTO_DIR = os.path.join(BASE_DIR, 'proto')
OUTPUT_DIR = os.path.join(BASE_DIR, 'dist')
ARTIFACT_DIR = os.path.join(BASE_DIR, 'artifact')
VERSION = os.path.join(BASE_DIR, 'VERSION')
AVAILABLE_CODES = ['all', 'python', 'json']
DEFAULT_THIRD_PARTY_DIR = 'third_party/googleapis:third_party/protobuf/src'
DEFAULT_CODE = 'all'


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


def _get_env():
    return {
        'proto_dir': os.environ.get('PROTO_DIR', PROTO_DIR),
        'output_dir': os.environ.get('OUTPUT_DIR', OUTPUT_DIR),
        'artifact_dir': os.environ.get('OUTPUT_DIR', ARTIFACT_DIR),
        'default_third_party_dir': _get_default_third_party_dir(),
        'default_code': os.environ.get('DEFAULT_CODE', DEFAULT_CODE),
        'protoc-gen-doc': shutil.which('protoc-gen-doc')
    }


def _get_args():
    env = _get_env()

    parser = argparse.ArgumentParser(description='Cloud One API Builder', formatter_class=argparse.ArgumentDefaultsHelpFormatter)

    parser.add_argument('target', metavar='<target>', help='Cloud One Service. (core, identity, inventory, etc.)')
    parser.add_argument('-p', '--proto-dir', type=str, help='Protocol Buffers Directory.', default=env['proto_dir'])
    parser.add_argument('-t', '--third-party-dir', type=str, help='Third Pary Protocol Buffers Directory.', default=[],action='append')
    parser.add_argument('-o', '--output-dir', type=str, help='Output Directory.', default=env['output_dir'])
    parser.add_argument('-a', '--artifact-dir', type=str, help='Artifact JSON Directory.', default=env['artifact_dir'])
    parser.add_argument('-c', '--code', type=str, help='Generate Code.', choices=AVAILABLE_CODES, default=env['default_code'])
    parser.add_argument('-d', '--debug', help='Debug Mode.', action='store_true')

    args = parser.parse_args()
    params = {}
    params['target'] = args.target
    params['proto_path'] = os.path.join(args.proto_dir, PROJECT, 'api', params['target'])
    params['output_dir'] = args.output_dir
    params['artifact_dir'] = args.artifact_dir
    params['version'] = _get_api_version()
    params['code'] = args.code
    params['debug'] = args.debug

    if not os.path.exists(params['proto_path']):
        _error(f"Protocol buffer's directory({params['proto_path']}) is not found.")

    params['proto_path_list'] = [args.proto_dir]
    params['proto_path_list'].extend(env['default_third_party_dir'])
    params['proto_path_list'].extend(args.third_party_dir)

    return params


def _get_proto_files(proto_path):
    return [proto_file for proto_file in glob.iglob(os.path.join(proto_path, '**', '*.proto'), recursive=True)]

def _get_api_version():
    with open(VERSION, 'r') as f:
        version = f.read().strip()
        f.close()
        return version

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
    if code == 'all':
        return list(filter(lambda x: x != 'all', AVAILABLE_CODES))
    else:
        return [code]


def _make_build_environment(output_dir, code):
    if code == 'python':
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
    except Exception:
        _error(f"Failed to compile : {proto_file}")

    print(f"[SUCCESS] Python Compile : {proto_file}")

def _doc_compile(proto_file, output_path, proto_path_list, debug):
    # check GOPATH
    if not os.environ.get('GOBIN'):
        _error('\'protoc-gen-doc\' does not exist.')

    output_path = _make_package_path(proto_file, output_path)

    gobin = os.path.join(os.environ.get('GOBIN'), 'protoc-gen-doc')
    cmd = ['protoc', f'--plugin=protoc-gen-doc={gobin}', f'--doc_out={output_path}']
    for proto_path in proto_path_list:
        cmd.append(f'--proto_path={proto_path}')

    # proto name
    temp = os.path.split(proto_file)
    temporary_name = temp[1].split('.')
    cap_name = temporary_name[0].capitalize()

    cmd.append(f'--doc_opt={OUTPUT_TO_DOC},{cap_name}.{OUTPUT_TO_DOC}')
    cmd.append(proto_file)

    if debug:
        print()
        print(' '.join(cmd))

    try:
        subprocess.check_output(cmd)
    except Exception:
        _error(f"Failed to compile : {proto_file}")

    print(f"[SUCCESS] Document Compile : {proto_file}")

def _version_factoring():
    ver_path = os.path.join(BASE_DIR, 'VERSION')
    art_path = os.path.join(ARTIFACT_DIR, 'json')
    art_file = os.path.join(art_path, 'VERSION')
    if (os.path.exists(art_path) and not os.path.isfile(art_file)):
        copyfile(ver_path, art_file)


def _compile_code(params, code, proto_file):
    output_path = os.path.join(params['output_dir'], code)

    if code == 'python':
        _python_compile(proto_file, output_path, params['proto_path_list'], debug=params['debug'])

    elif code == 'json':
        output_path = os.path.join(params['artifact_dir'], code, params['version'])
        _doc_compile(proto_file, output_path, params['proto_path_list'], debug=params['debug'])


def build(params):
    proto_files = _get_proto_files(params['proto_path'])
    for code in _get_generate_codes(params['code']):
        if (code == 'json'):
            _make_output_path(params['artifact_dir'], code)
        else:
            _make_output_path(params['output_dir'], code)

        for proto_file in proto_files:
            _compile_code(params, code, proto_file)

        if(code == 'json'):
            _make_build_environment(params['artifact_dir'], code)
        else:
            _make_build_environment(params['output_dir'], code)

    if (params['code'] == 'json'):
        _version_factoring()


if __name__ == '__main__':
    params = _get_args()
    build(params)



