# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: spaceone/api/file_manager/v1/file.proto
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


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\'spaceone/api/file_manager/v1/file.proto\x12\x1cspaceone.api.file_manager.v1\x1a\x1bgoogle/protobuf/empty.proto\x1a\x1cgoogle/protobuf/struct.proto\x1a\x1cgoogle/api/annotations.proto\x1a spaceone/api/core/v2/query.proto\";\n\rFileReference\x12\x15\n\rresource_type\x18\x01 \x01(\t\x12\x13\n\x0bresource_id\x18\x02 \x01(\t\"\xfa\x02\n\x11\x43reateFileRequest\x12\x0c\n\x04name\x18\x01 \x01(\t\x12%\n\x04tags\x18\x02 \x01(\x0b\x32\x17.google.protobuf.Struct\x12>\n\treference\x18\x03 \x01(\x0b\x32+.spaceone.api.file_manager.v1.FileReference\x12U\n\x0eresource_group\x18\x14 \x01(\x0e\x32=.spaceone.api.file_manager.v1.CreateFileRequest.ResourceGroup\x12\x11\n\tdomain_id\x18\x15 \x01(\t\x12\x14\n\x0cworkspace_id\x18\x16 \x01(\t\x12\x12\n\nproject_id\x18\x17 \x01(\t\"\\\n\rResourceGroup\x12\x17\n\x13RESOURCE_GROUP_NONE\x10\x00\x12\n\n\x06SYSTEM\x10\x01\x12\n\n\x06\x44OMAIN\x10\x02\x12\r\n\tWORKSPACE\x10\x03\x12\x0b\n\x07PROJECT\x10\x04\"\x8b\x01\n\x11UpdateFileRequest\x12\x0f\n\x07\x66ile_id\x18\x01 \x01(\t\x12%\n\x04tags\x18\x02 \x01(\x0b\x32\x17.google.protobuf.Struct\x12>\n\treference\x18\x03 \x01(\x0b\x32+.spaceone.api.file_manager.v1.FileReference\"\x1e\n\x0b\x46ileRequest\x12\x0f\n\x07\x66ile_id\x18\x01 \x01(\t\"\xc5\x01\n\x0f\x46ileSearchQuery\x12*\n\x05query\x18\x01 \x01(\x0b\x32\x1b.spaceone.api.core.v2.Query\x12\x0f\n\x07\x66ile_id\x18\x02 \x01(\t\x12\x0c\n\x04name\x18\x03 \x01(\t\x12\x15\n\rresource_type\x18\x04 \x01(\t\x12\x13\n\x0bresource_id\x18\x05 \x01(\t\x12\x11\n\tdomain_id\x18\x15 \x01(\t\x12\x14\n\x0cworkspace_id\x18\x16 \x01(\t\x12\x12\n\nproject_id\x18\x17 \x01(\t\"\xa3\x03\n\x08\x46ileInfo\x12\x0f\n\x07\x66ile_id\x18\x01 \x01(\t\x12\x0c\n\x04name\x18\x02 \x01(\t\x12\x14\n\x0c\x64ownload_url\x18\x03 \x01(\t\x12>\n\treference\x18\x04 \x01(\x0b\x32+.spaceone.api.file_manager.v1.FileReference\x12%\n\x04tags\x18\x05 \x01(\x0b\x32\x17.google.protobuf.Struct\x12L\n\x0eresource_group\x18\x14 \x01(\x0e\x32\x34.spaceone.api.file_manager.v1.FileInfo.ResourceGroup\x12\x11\n\tdomain_id\x18\x15 \x01(\t\x12\x14\n\x0cworkspace_id\x18\x16 \x01(\t\x12\x12\n\nproject_id\x18\x17 \x01(\t\x12\x12\n\ncreated_at\x18\x1f \x01(\t\"\\\n\rResourceGroup\x12\x17\n\x13RESOURCE_GROUP_NONE\x10\x00\x12\n\n\x06SYSTEM\x10\x01\x12\n\n\x06\x44OMAIN\x10\x02\x12\r\n\tWORKSPACE\x10\x03\x12\x0b\n\x07PROJECT\x10\x04\"Y\n\tFilesInfo\x12\x37\n\x07results\x18\x01 \x03(\x0b\x32&.spaceone.api.file_manager.v1.FileInfo\x12\x13\n\x0btotal_count\x18\x02 \x01(\x05\"E\n\rFileStatQuery\x12\x34\n\x05query\x18\x01 \x01(\x0b\x32%.spaceone.api.core.v2.StatisticsQuery2\x86\x05\n\x04\x46ile\x12\x8a\x01\n\x06update\x12/.spaceone.api.file_manager.v1.UpdateFileRequest\x1a&.spaceone.api.file_manager.v1.FileInfo\"\'\x82\xd3\xe4\x93\x02!\"\x1c/file-manager/v1/file/update:\x01*\x12t\n\x06\x64\x65lete\x12).spaceone.api.file_manager.v1.FileRequest\x1a\x16.google.protobuf.Empty\"\'\x82\xd3\xe4\x93\x02!\"\x1c/file-manager/v1/file/delete:\x01*\x12~\n\x03get\x12).spaceone.api.file_manager.v1.FileRequest\x1a&.spaceone.api.file_manager.v1.FileInfo\"$\x82\xd3\xe4\x93\x02\x1e\"\x19/file-manager/v1/file/get:\x01*\x12\x85\x01\n\x04list\x12-.spaceone.api.file_manager.v1.FileSearchQuery\x1a\'.spaceone.api.file_manager.v1.FilesInfo\"%\x82\xd3\xe4\x93\x02\x1f\"\x1a/file-manager/v1/file/list:\x01*\x12s\n\x04stat\x12+.spaceone.api.file_manager.v1.FileStatQuery\x1a\x17.google.protobuf.Struct\"%\x82\xd3\xe4\x93\x02\x1f\"\x1a/file-manager/v1/file/stat:\x01*BCZAgithub.com/cloudforet-io/api/dist/go/spaceone/api/file_manager/v1b\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'spaceone.api.file_manager.v1.file_pb2', _globals)
if not _descriptor._USE_C_DESCRIPTORS:
  _globals['DESCRIPTOR']._loaded_options = None
  _globals['DESCRIPTOR']._serialized_options = b'ZAgithub.com/cloudforet-io/api/dist/go/spaceone/api/file_manager/v1'
  _globals['_FILE'].methods_by_name['update']._loaded_options = None
  _globals['_FILE'].methods_by_name['update']._serialized_options = b'\202\323\344\223\002!\"\034/file-manager/v1/file/update:\001*'
  _globals['_FILE'].methods_by_name['delete']._loaded_options = None
  _globals['_FILE'].methods_by_name['delete']._serialized_options = b'\202\323\344\223\002!\"\034/file-manager/v1/file/delete:\001*'
  _globals['_FILE'].methods_by_name['get']._loaded_options = None
  _globals['_FILE'].methods_by_name['get']._serialized_options = b'\202\323\344\223\002\036\"\031/file-manager/v1/file/get:\001*'
  _globals['_FILE'].methods_by_name['list']._loaded_options = None
  _globals['_FILE'].methods_by_name['list']._serialized_options = b'\202\323\344\223\002\037\"\032/file-manager/v1/file/list:\001*'
  _globals['_FILE'].methods_by_name['stat']._loaded_options = None
  _globals['_FILE'].methods_by_name['stat']._serialized_options = b'\202\323\344\223\002\037\"\032/file-manager/v1/file/stat:\001*'
  _globals['_FILEREFERENCE']._serialized_start=196
  _globals['_FILEREFERENCE']._serialized_end=255
  _globals['_CREATEFILEREQUEST']._serialized_start=258
  _globals['_CREATEFILEREQUEST']._serialized_end=636
  _globals['_CREATEFILEREQUEST_RESOURCEGROUP']._serialized_start=544
  _globals['_CREATEFILEREQUEST_RESOURCEGROUP']._serialized_end=636
  _globals['_UPDATEFILEREQUEST']._serialized_start=639
  _globals['_UPDATEFILEREQUEST']._serialized_end=778
  _globals['_FILEREQUEST']._serialized_start=780
  _globals['_FILEREQUEST']._serialized_end=810
  _globals['_FILESEARCHQUERY']._serialized_start=813
  _globals['_FILESEARCHQUERY']._serialized_end=1010
  _globals['_FILEINFO']._serialized_start=1013
  _globals['_FILEINFO']._serialized_end=1432
  _globals['_FILEINFO_RESOURCEGROUP']._serialized_start=544
  _globals['_FILEINFO_RESOURCEGROUP']._serialized_end=636
  _globals['_FILESINFO']._serialized_start=1434
  _globals['_FILESINFO']._serialized_end=1523
  _globals['_FILESTATQUERY']._serialized_start=1525
  _globals['_FILESTATQUERY']._serialized_end=1594
  _globals['_FILE']._serialized_start=1597
  _globals['_FILE']._serialized_end=2243
# @@protoc_insertion_point(module_scope)
