import glob
import json
import os
import sys

from jinja2 import Environment, FileSystemLoader

from build import load_yaml_from_file
from build import BASE_DIR, PROTO_DIR, TEMPLATE_DIR, PROJECT


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
    return sorted(field_list, key=lambda x: x['required'], reverse=True)


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


def _make_openapi_json(output_dir, code, service_name, core_message_list):
    openapi_file_name = f'{service_name}_openapi.json'
    output_path = os.path.join(BASE_DIR, output_dir, code, openapi_file_name)

    metadata_path = os.path.join(PROTO_DIR, PROJECT, 'api', service_name, 'metadata.yaml')
    output_dir_with_service_name = os.path.join(output_dir, 'json', 'cloudforet', 'api', service_name)
    template_dir = os.path.join(TEMPLATE_DIR, code)
    json_file_path_list = _get_json_file_paths(output_dir_with_service_name)

    filtered_method_list = []
    filtered_messages = {}

    for json_file_path in json_file_path_list:
        with open(json_file_path, 'r') as f:
            json_data = json.load(f)

            enums = _get_enums(json_data)
            scalar_value_types = _get_scalar_value_types(json_data.get('scalarValueTypes', []))
            service_list = json_data.get('files')[0].get('services', [])
            message_list = json_data.get('files')[0].get('messages', [])
            message_list.extend(core_message_list)

            for message in message_list:
                field_list = message.get('fields')

                _modify_message_field_type_proto_to_python(field_list, scalar_value_types)
                message = _modify_message_field_type_python_to_openapi(message)
                message['fields'] = _set_required_or_optional(field_list)
                message['fields'] = _sort_fields_by_required(field_list)
                message['fields'] = _set_enum_field_as_string(field_list, enums)
                message['table_description'] = _make_table_description(field_list)
                filtered_messages.update({message.get('name'): message})

            for service in service_list:
                method_list = service.get('methods')
                for method in method_list:
                    if filtered_method := _transform_to_filtered_method(method):
                        filtered_method_list.append(filtered_method)

    if filtered_method_list:
        metadata = load_yaml_from_file(metadata_path) if os.path.exists(metadata_path) else {}
        jinja_env = Environment(autoescape=False, loader=FileSystemLoader(template_dir), trim_blocks=True)
        template = jinja_env.get_template('openapi.json.tmpl')
        content = template.render(methods=filtered_method_list, messages=filtered_messages,
                                  scalar_value_types=scalar_value_types, metadata=metadata)

        with open(output_path, 'w') as f:
            f.write(content)


def _get_json_file_paths(dist_json_dir) -> list:
    return [json_file for json_file in glob.iglob(os.path.join(dist_json_dir, '**', '*.json'), recursive=True)
            if f'{dist_json_dir}/plugin' not in json_file]


def _get_core_messages(output_dir) -> list:
    core_message_list = []
    output_dir_with_service = os.path.join(output_dir, 'json', 'cloudforet', 'api', 'core')
    json_file_path_list = _get_json_file_paths(output_dir_with_service)

    for json_file_path in json_file_path_list:
        with open(json_file_path, 'r') as f:
            json_data = json.load(f)
            messages = json_data.get('files')[0].get('messages')
            core_message_list.extend(messages)

    return core_message_list


def openapi_compile(output_dir, code, service, debug):
    json_file_path = os.path.join(output_dir, 'json', service)
    try:
        core_message_list = _get_core_messages(output_dir)
        _make_openapi_json(output_dir, code, service, core_message_list)

        if debug:
            print()
            print(f'{json_file_path}')

    except Exception as e:
        _error(f"Failed to OpenAPI Compile : {json_file_path}\n{e}")

    print(f"[SUCCESS] OpenAPI Compile : {json_file_path}")


def make_build_environment_openapi(output_dir, code):
    pass
