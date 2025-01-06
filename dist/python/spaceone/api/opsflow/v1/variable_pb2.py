# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: spaceone/api/opsflow/v1/variable.proto
# Protobuf Python Version: 5.26.1
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from google.protobuf import empty_pb2 as google_dot_protobuf_dot_empty__pb2
from google.protobuf import struct_pb2 as google_dot_protobuf_dot_struct__pb2
from google.api import annotations_pb2 as google_dot_api_dot_annotations__pb2
from spaceone.api.core.v2 import query_pb2 as spaceone_dot_api_dot_core_dot_v2_dot_query__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n&spaceone/api/opsflow/v1/variable.proto\x12\x17spaceone.api.opsflow.v1\x1a\x1bgoogle/protobuf/empty.proto\x1a\x1cgoogle/protobuf/struct.proto\x1a\x1cgoogle/api/annotations.proto\x1a spaceone/api/core/v2/query.proto\"q\n\x15VariableCreateRequest\x12\x0c\n\x04name\x18\x01 \x01(\t\x12\r\n\x05value\x18\x02 \x01(\t\x12%\n\x04tags\x18\x03 \x01(\x0b\x32\x17.google.protobuf.Struct\x12\x14\n\x0cworkspace_id\x18\x15 \x01(\t\"[\n\x15VariableUpdateRequest\x12\x0c\n\x04name\x18\x01 \x01(\t\x12\r\n\x05value\x18\x02 \x01(\t\x12%\n\x04tags\x18\x03 \x01(\x0b\x32\x17.google.protobuf.Struct\"\x1f\n\x0fVariableRequest\x12\x0c\n\x04name\x18\x01 \x01(\t\"e\n\x13VariableSearchQuery\x12*\n\x05query\x18\x01 \x01(\x0b\x32\x1b.spaceone.api.core.v2.Query\x12\x0c\n\x04name\x18\x02 \x01(\t\x12\x14\n\x0cworkspace_id\x18\x15 \x01(\t\"\xa3\x01\n\x0cVariableInfo\x12\x0c\n\x04name\x18\x01 \x01(\t\x12\r\n\x05value\x18\x02 \x01(\t\x12%\n\x04tags\x18\x03 \x01(\x0b\x32\x17.google.protobuf.Struct\x12\x11\n\tdomain_id\x18\x15 \x01(\t\x12\x14\n\x0cworkspace_id\x18\x16 \x01(\t\x12\x12\n\ncreated_at\x18\x1f \x01(\t\x12\x12\n\nupdated_at\x18  \x01(\t\"\\\n\rVariablesInfo\x12\x36\n\x07results\x18\x01 \x03(\x0b\x32%.spaceone.api.opsflow.v1.VariableInfo\x12\x13\n\x0btotal_count\x18\x02 \x01(\x05\"I\n\x11VariableStatQuery\x12\x34\n\x05query\x18\x01 \x01(\x0b\x32%.spaceone.api.core.v2.StatisticsQuery2\x87\x06\n\x08Variable\x12\x87\x01\n\x06\x63reate\x12..spaceone.api.opsflow.v1.VariableCreateRequest\x1a%.spaceone.api.opsflow.v1.VariableInfo\"&\x82\xd3\xe4\x93\x02 \"\x1b/opsflow/v1/variable/create:\x01*\x12\x87\x01\n\x06update\x12..spaceone.api.opsflow.v1.VariableUpdateRequest\x1a%.spaceone.api.opsflow.v1.VariableInfo\"&\x82\xd3\xe4\x93\x02 \"\x1b/opsflow/v1/variable/update:\x01*\x12r\n\x06\x64\x65lete\x12(.spaceone.api.opsflow.v1.VariableRequest\x1a\x16.google.protobuf.Empty\"&\x82\xd3\xe4\x93\x02 \"\x1b/opsflow/v1/variable/delete:\x01*\x12{\n\x03get\x12(.spaceone.api.opsflow.v1.VariableRequest\x1a%.spaceone.api.opsflow.v1.VariableInfo\"#\x82\xd3\xe4\x93\x02\x1d\"\x18/opsflow/v1/variable/get:\x01*\x12\x82\x01\n\x04list\x12,.spaceone.api.opsflow.v1.VariableSearchQuery\x1a&.spaceone.api.opsflow.v1.VariablesInfo\"$\x82\xd3\xe4\x93\x02\x1e\"\x19/opsflow/v1/variable/list:\x01*\x12q\n\x04stat\x12*.spaceone.api.opsflow.v1.VariableStatQuery\x1a\x17.google.protobuf.Struct\"$\x82\xd3\xe4\x93\x02\x1e\"\x19/opsflow/v1/variable/stat:\x01*B>Z<github.com/cloudforet-io/api/dist/go/spaceone/api/opsflow/v1b\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'spaceone.api.opsflow.v1.variable_pb2', _globals)
if not _descriptor._USE_C_DESCRIPTORS:
  _globals['DESCRIPTOR']._loaded_options = None
  _globals['DESCRIPTOR']._serialized_options = b'Z<github.com/cloudforet-io/api/dist/go/spaceone/api/opsflow/v1'
  _globals['_VARIABLE'].methods_by_name['create']._loaded_options = None
  _globals['_VARIABLE'].methods_by_name['create']._serialized_options = b'\202\323\344\223\002 \"\033/opsflow/v1/variable/create:\001*'
  _globals['_VARIABLE'].methods_by_name['update']._loaded_options = None
  _globals['_VARIABLE'].methods_by_name['update']._serialized_options = b'\202\323\344\223\002 \"\033/opsflow/v1/variable/update:\001*'
  _globals['_VARIABLE'].methods_by_name['delete']._loaded_options = None
  _globals['_VARIABLE'].methods_by_name['delete']._serialized_options = b'\202\323\344\223\002 \"\033/opsflow/v1/variable/delete:\001*'
  _globals['_VARIABLE'].methods_by_name['get']._loaded_options = None
  _globals['_VARIABLE'].methods_by_name['get']._serialized_options = b'\202\323\344\223\002\035\"\030/opsflow/v1/variable/get:\001*'
  _globals['_VARIABLE'].methods_by_name['list']._loaded_options = None
  _globals['_VARIABLE'].methods_by_name['list']._serialized_options = b'\202\323\344\223\002\036\"\031/opsflow/v1/variable/list:\001*'
  _globals['_VARIABLE'].methods_by_name['stat']._loaded_options = None
  _globals['_VARIABLE'].methods_by_name['stat']._serialized_options = b'\202\323\344\223\002\036\"\031/opsflow/v1/variable/stat:\001*'
  _globals['_VARIABLECREATEREQUEST']._serialized_start=190
  _globals['_VARIABLECREATEREQUEST']._serialized_end=303
  _globals['_VARIABLEUPDATEREQUEST']._serialized_start=305
  _globals['_VARIABLEUPDATEREQUEST']._serialized_end=396
  _globals['_VARIABLEREQUEST']._serialized_start=398
  _globals['_VARIABLEREQUEST']._serialized_end=429
  _globals['_VARIABLESEARCHQUERY']._serialized_start=431
  _globals['_VARIABLESEARCHQUERY']._serialized_end=532
  _globals['_VARIABLEINFO']._serialized_start=535
  _globals['_VARIABLEINFO']._serialized_end=698
  _globals['_VARIABLESINFO']._serialized_start=700
  _globals['_VARIABLESINFO']._serialized_end=792
  _globals['_VARIABLESTATQUERY']._serialized_start=794
  _globals['_VARIABLESTATQUERY']._serialized_end=867
  _globals['_VARIABLE']._serialized_start=870
  _globals['_VARIABLE']._serialized_end=1645
# @@protoc_insertion_point(module_scope)