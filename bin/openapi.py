import functools
import glob
import json
import os
import sys

import click
from jinja2 import Environment, FileSystemLoader

PROJECT = 'cloudforet'
BASE_DIR = os.path.dirname(os.path.dirname(os.path.abspath(__file__)))
OUTPUT_DIR = os.path.join(BASE_DIR, 'dist', 'openapi')
JSON_DIR = os.path.join(BASE_DIR, 'dist', 'json', PROJECT, 'api')
TEMPLATE_DIR = os.path.join(BASE_DIR, 'template')
DEFAULT_COMMON_SCHEMA_DIR = os.path.join(BASE_DIR, 'dist', 'json', PROJECT, 'api', 'core', 'v1')


def _error(msg):
    print()
    print('[ERROR] %s' % (msg))
    print()
    sys.exit(1)


def _get_enums(json_data) -> dict:
    enums = {}
    enum_list = json_data.get('files')[0].get('enums', [])
    for enum in enum_list:
        enums[enum['name']] = enum
    return enums


def _get_scalar_value_types(scalar_value_type_list) -> dict:
    scalar_value_types = {}
    for scalar_value_type in scalar_value_type_list:
        scalar_value_types[(scalar_value_type['protoType'])] = scalar_value_type['pythonType']
    return scalar_value_types


def _modify_message_field_type_proto_to_python(field_list, scalar_value_types):
    for field in field_list:
        if field['type'] in scalar_value_types:
            if scalar_value_types[field['type']] == 'str/unicode':
                field['type'] = 'string'
            else:
                field['type'] = scalar_value_types[field['type']]


def _modify_message_field_type_python_to_openapi(message) -> dict:
    field_list = message.get('fields')
    for field in field_list:
        if field['type'] == 'ListValue':
            field['type'] = 'array'
    return message


def _set_required_or_optional(field_list) -> list:
    fields = {}
    for field in field_list:
        field_description = field['description']
        if field_description.split('\n')[-1].lower().strip() == '+optional':
            field['required'] = ''
        else:
            field['required'] = 'True'
    return field_list


def _set_enum_field_as_string(field_list, enums) -> list:
    if enums:
        for field in field_list:
            if field.get('type') in enums:
                field['description'] = ', '.join([value.get('name') for value in enums[field.get('type')]['values']
                                                  if value.get('number') != '0']) + ' ' + field['description']
                field['type'] = 'string'
    return field_list


def _sort_fields_by_required(field_list) -> list:
    return sorted(field_list, key=lambda x: x['required'] != 'True')


def _make_table_description(field_list) -> str:
    _label_type_mapper = {
        'repeated': 'array',
        'Struct': 'object',
        'ListValue': 'array'
    }

    _table_header = "\n" \
                    "| Key               | Description                                                   | Type      | Required|\n" \
                    "|-------------------|---------------------------------------------------------------|-----------|-------|\n"
    _table_body_format = " |{key}|{description}|{type}|{required}|\n"
    _table_body = ""

    for filed in field_list:
        _description = filed['description'].replace('"', "'")
        _description = _description.replace('+optional', '')
        _type = _label_type_mapper.get(filed.get('label'), filed.get('type'))
        _table_body += _table_body_format.format(key=filed['name'], description=_description, type=_type,
                                                 required=filed['required'])

    table_description = _table_header + _table_body
    return table_description


def _get_tag_from_url(url):
    _split_url = url.split('/')
    tag = _split_url[1] + ' > ' + _split_url[3]
    return tag


def _set_auth_required_or_not(method):
    if method.get('description') is not None:
        method_description = method['description']
        if method_description.split('\n')[-1].lower().strip() == '+noauth':
            method['description'] = method_description.replace('+noauth', '')
            method['auth_required'] = 'False'
        else:
            method['auth_required'] = 'True'
    return method


def _transform_to_filtered_method(method):
    if method.get('options') is not None:
        url = method['options'].get('google.api.http').get('rules')[0].get('pattern')
        method['tag'] = _get_tag_from_url(url)
        method['summary'] = method.get('name').replace('_', ' ').title()
        method = _set_auth_required_or_not(method)
        method['description'] = "### Description \n" + method['description']
        return method


def _make_openapi_json(target, common_schema_list, sub_dir, output_dir=None):
    template_dir = os.path.join(TEMPLATE_DIR, 'openapi')
    output_dir = output_dir if output_dir else sub_dir.replace('json', 'openapi')
    json_file_path_list = _get_json_file_paths(sub_dir)

    filtered_method_list = []
    filtered_messages = {}

    for json_file_path in json_file_path_list:
        with open(json_file_path, 'r') as f:
            json_data = json.load(f)

            enums = _get_enums(json_data)
            scalar_value_types = _get_scalar_value_types(json_data.get('scalarValueTypes', []))
            service_list = json_data.get('files')[0].get('services', [])
            message_list = json_data.get('files')[0].get('messages', [])
            message_list.extend(common_schema_list)

            for message in message_list:
                field_list = message.get('fields')

                _modify_message_field_type_proto_to_python(field_list, scalar_value_types)
                message = _modify_message_field_type_python_to_openapi(message)
                message['fields'] = _set_required_or_optional(field_list)
                message['fields'] = _set_enum_field_as_string(field_list, enums)
                sorted_field_list = _sort_fields_by_required(field_list)
                message['table_description'] = _make_table_description(sorted_field_list)
                filtered_messages.update({message.get('name'): message})

            for service in service_list:
                method_list = service.get('methods')
                for method in method_list:
                    if filtered_method := _transform_to_filtered_method(method):
                        filtered_method_list.append(filtered_method)

    if filtered_method_list:
        jinja_env = Environment(autoescape=False, loader=FileSystemLoader(template_dir), trim_blocks=True)
        template = jinja_env.get_template('openapi.json.tmpl')
        content = template.render(methods=filtered_method_list, messages=filtered_messages, service_name=target,
                                  scalar_value_types=scalar_value_types)

        if not os.path.exists(output_dir):
            os.makedirs(output_dir, exist_ok=True)

        with open(f'{output_dir}/openapi.json', 'w') as f:
            f.write(content)


def _get_json_file_paths(json_input_file_path) -> list:
    return [json_file for json_file in glob.iglob(os.path.join(json_input_file_path, '**', '*.json'), recursive=True)]


def _get_common_schemas(common_schema_dir_list) -> list:
    common_schema_list = []
    common_schema_path_list = []

    for common_schema_dir in common_schema_dir_list:
        common_schema_path_list.extend(_get_json_file_paths(common_schema_dir))

    for common_schema_path in common_schema_path_list:
        with open(common_schema_path, 'r') as f:
            json_data = json.load(f)
            messages = json_data.get('files')[0].get('messages')
            common_schema_list.extend(messages)

    return common_schema_list


def _get_json_input_sub_dirs(json_input_dir) -> list:
    json_input_sub_dir_list = []
    for dir_name in os.listdir(os.path.join(json_input_dir)):
        json_input_sub_dir_name = os.path.join(json_input_dir, dir_name)
        if os.path.isdir(json_input_sub_dir_name):
            json_input_sub_dir_list.append(json_input_sub_dir_name)
    return json_input_sub_dir_list


def build_openapi_json(**params):
    target = params['target'].replace('-', '_')
    debug = params.get('debug', False)
    output_dir = params.get('output_dir')
    json_input_dir = params.get('json_dir') if params.get('json_dir') else JSON_DIR

    common_schema_dir_list = params.get('common_schema_dirs', [DEFAULT_COMMON_SCHEMA_DIR])

    try:
        common_schema_list = _get_common_schemas(common_schema_dir_list)

        sub_dir_list = _get_json_input_sub_dirs(f'{json_input_dir}/{target}')
        list(map(functools.partial(_make_openapi_json, target, common_schema_list, output_dir=output_dir), sub_dir_list))

        if debug:
            print()
            print(f'[DEBUG] Directory list : {sub_dir_list}')

    except Exception as e:
        _error(f"Failed to OpenAPI Compile : {json_input_dir}\n{e}")

    print(f"[SUCCESS] OpenAPI Compile : {json_input_dir}/{target}")


@click.command()
@click.option('-t', '--target', type=str, help='Target service name (ex. identity, inventory, cost-analysis, etc.)')
@click.option('-j', '--json-dir', type=str, help='Input json directory', default=JSON_DIR)
@click.option('-o', '--output-dir', type=str, help='If not specified, the default output directory is decided by the json_dir')
@click.option('-s', '--common-schema-dirs', type=str, help='Common schema directories', multiple=True, default=[])
@click.option('-d', '--debug', is_flag=True, help='Debug mode')
def openapi(**params):
    build_openapi_json(**params)


if __name__ == '__main__':
    openapi()
