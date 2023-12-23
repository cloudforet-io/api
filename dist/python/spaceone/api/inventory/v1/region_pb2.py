# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: spaceone/api/inventory/v1/region.proto
# Protobuf Python Version: 4.25.0
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


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n&spaceone/api/inventory/v1/region.proto\x12\x19spaceone.api.inventory.v1\x1a\x1bgoogle/protobuf/empty.proto\x1a\x1cgoogle/protobuf/struct.proto\x1a\x1cgoogle/api/annotations.proto\x1a spaceone/api/core/v2/query.proto\"q\n\x13\x43reateRegionRequest\x12\x0c\n\x04name\x18\x01 \x01(\t\x12\x13\n\x0bregion_code\x18\x02 \x01(\t\x12\x10\n\x08provider\x18\x03 \x01(\t\x12%\n\x04tags\x18\x04 \x01(\x0b\x32\x17.google.protobuf.Struct\"]\n\x13UpdateRegionRequest\x12\x11\n\tregion_id\x18\x01 \x01(\t\x12\x0c\n\x04name\x18\x02 \x01(\t\x12%\n\x04tags\x18\x03 \x01(\x0b\x32\x17.google.protobuf.Struct\"\"\n\rRegionRequest\x12\x11\n\tregion_id\x18\x01 \x01(\t\"\xab\x01\n\x0bRegionQuery\x12*\n\x05query\x18\x01 \x01(\x0b\x32\x1b.spaceone.api.core.v2.Query\x12\x11\n\tregion_id\x18\x02 \x01(\t\x12\x0c\n\x04name\x18\x03 \x01(\t\x12\x12\n\nregion_key\x18\x04 \x01(\t\x12\x13\n\x0bregion_code\x18\x05 \x01(\t\x12\x10\n\x08provider\x18\x06 \x01(\t\x12\x14\n\x0cworkspace_id\x18\x15 \x01(\t\"\xe0\x01\n\nRegionInfo\x12\x11\n\tregion_id\x18\x01 \x01(\t\x12\x0c\n\x04name\x18\x02 \x01(\t\x12\x12\n\nregion_key\x18\x03 \x01(\t\x12\x13\n\x0bregion_code\x18\x04 \x01(\t\x12\x10\n\x08provider\x18\x05 \x01(\t\x12%\n\x04tags\x18\x06 \x01(\x0b\x32\x17.google.protobuf.Struct\x12\x11\n\tdomain_id\x18\x15 \x01(\t\x12\x14\n\x0cworkspace_id\x18\x16 \x01(\t\x12\x12\n\ncreated_at\x18\x1f \x01(\t\x12\x12\n\nupdated_at\x18  \x01(\t\"Z\n\x0bRegionsInfo\x12\x36\n\x07results\x18\x01 \x03(\x0b\x32%.spaceone.api.inventory.v1.RegionInfo\x12\x13\n\x0btotal_count\x18\x02 \x01(\x05\"G\n\x0fRegionStatQuery\x12\x34\n\x05query\x18\x01 \x01(\x0b\x32%.spaceone.api.core.v2.StatisticsQuery2\xfe\x05\n\x06Region\x12\x87\x01\n\x06\x63reate\x12..spaceone.api.inventory.v1.CreateRegionRequest\x1a%.spaceone.api.inventory.v1.RegionInfo\"&\x82\xd3\xe4\x93\x02 \"\x1b/inventory/v1/region/create:\x01*\x12\x87\x01\n\x06update\x12..spaceone.api.inventory.v1.UpdateRegionRequest\x1a%.spaceone.api.inventory.v1.RegionInfo\"&\x82\xd3\xe4\x93\x02 \"\x1b/inventory/v1/region/update:\x01*\x12r\n\x06\x64\x65lete\x12(.spaceone.api.inventory.v1.RegionRequest\x1a\x16.google.protobuf.Empty\"&\x82\xd3\xe4\x93\x02 \"\x1b/inventory/v1/region/delete:\x01*\x12{\n\x03get\x12(.spaceone.api.inventory.v1.RegionRequest\x1a%.spaceone.api.inventory.v1.RegionInfo\"#\x82\xd3\xe4\x93\x02\x1d\"\x18/inventory/v1/region/get:\x01*\x12|\n\x04list\x12&.spaceone.api.inventory.v1.RegionQuery\x1a&.spaceone.api.inventory.v1.RegionsInfo\"$\x82\xd3\xe4\x93\x02\x1e\"\x19/inventory/v1/region/list:\x01*\x12q\n\x04stat\x12*.spaceone.api.inventory.v1.RegionStatQuery\x1a\x17.google.protobuf.Struct\"$\x82\xd3\xe4\x93\x02\x1e\"\x19/inventory/v1/region/stat:\x01*B@Z>github.com/cloudforet-io/api/dist/go/spaceone/api/inventory/v1b\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'spaceone.api.inventory.v1.region_pb2', _globals)
if _descriptor._USE_C_DESCRIPTORS == False:
  _globals['DESCRIPTOR']._options = None
  _globals['DESCRIPTOR']._serialized_options = b'Z>github.com/cloudforet-io/api/dist/go/spaceone/api/inventory/v1'
  _globals['_REGION'].methods_by_name['create']._options = None
  _globals['_REGION'].methods_by_name['create']._serialized_options = b'\202\323\344\223\002 \"\033/inventory/v1/region/create:\001*'
  _globals['_REGION'].methods_by_name['update']._options = None
  _globals['_REGION'].methods_by_name['update']._serialized_options = b'\202\323\344\223\002 \"\033/inventory/v1/region/update:\001*'
  _globals['_REGION'].methods_by_name['delete']._options = None
  _globals['_REGION'].methods_by_name['delete']._serialized_options = b'\202\323\344\223\002 \"\033/inventory/v1/region/delete:\001*'
  _globals['_REGION'].methods_by_name['get']._options = None
  _globals['_REGION'].methods_by_name['get']._serialized_options = b'\202\323\344\223\002\035\"\030/inventory/v1/region/get:\001*'
  _globals['_REGION'].methods_by_name['list']._options = None
  _globals['_REGION'].methods_by_name['list']._serialized_options = b'\202\323\344\223\002\036\"\031/inventory/v1/region/list:\001*'
  _globals['_REGION'].methods_by_name['stat']._options = None
  _globals['_REGION'].methods_by_name['stat']._serialized_options = b'\202\323\344\223\002\036\"\031/inventory/v1/region/stat:\001*'
  _globals['_CREATEREGIONREQUEST']._serialized_start=192
  _globals['_CREATEREGIONREQUEST']._serialized_end=305
  _globals['_UPDATEREGIONREQUEST']._serialized_start=307
  _globals['_UPDATEREGIONREQUEST']._serialized_end=400
  _globals['_REGIONREQUEST']._serialized_start=402
  _globals['_REGIONREQUEST']._serialized_end=436
  _globals['_REGIONQUERY']._serialized_start=439
  _globals['_REGIONQUERY']._serialized_end=610
  _globals['_REGIONINFO']._serialized_start=613
  _globals['_REGIONINFO']._serialized_end=837
  _globals['_REGIONSINFO']._serialized_start=839
  _globals['_REGIONSINFO']._serialized_end=929
  _globals['_REGIONSTATQUERY']._serialized_start=931
  _globals['_REGIONSTATQUERY']._serialized_end=1002
  _globals['_REGION']._serialized_start=1005
  _globals['_REGION']._serialized_end=1771
# @@protoc_insertion_point(module_scope)
