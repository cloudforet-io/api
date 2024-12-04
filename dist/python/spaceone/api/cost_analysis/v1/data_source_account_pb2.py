# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: spaceone/api/cost_analysis/v1/data_source_account.proto
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


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n7spaceone/api/cost_analysis/v1/data_source_account.proto\x12\x1dspaceone.api.cost_analysis.v1\x1a\x1bgoogle/protobuf/empty.proto\x1a\x1cgoogle/protobuf/struct.proto\x1a\x1cgoogle/api/annotations.proto\x1a spaceone/api/core/v2/query.proto\"b\n\x1eUpdateDataSourceAccountRequest\x12\x16\n\x0e\x64\x61ta_source_id\x18\x01 \x01(\t\x12\x12\n\naccount_id\x18\x02 \x01(\t\x12\x14\n\x0cworkspace_id\x18\x15 \x01(\t\"_\n\x1dResetDataSourceAccountRequest\x12\x16\n\x0e\x64\x61ta_source_id\x18\x01 \x01(\t\x12\x12\n\naccount_id\x18\x02 \x01(\t\x12\x12\n\nreset_sync\x18\x03 \x01(\x08\"N\n\x18\x44\x61taSourceAccountRequest\x12\x1e\n\x16\x64\x61ta_source_account_id\x18\x01 \x01(\t\x12\x12\n\naccount_id\x18\x02 \x01(\t\"\xb6\x01\n\x16\x44\x61taSourceAccountQuery\x12*\n\x05query\x18\x01 \x01(\x0b\x32\x1b.spaceone.api.core.v2.Query\x12\x16\n\x0e\x64\x61ta_source_id\x18\x02 \x01(\t\x12\x12\n\naccount_id\x18\x03 \x01(\t\x12\x14\n\x0cworkspace_id\x18\x15 \x01(\t\x12\x12\n\nproject_id\x18\x16 \x01(\t\x12\x1a\n\x12service_account_id\x18\x17 \x01(\t\"\xf6\x01\n\x15\x44\x61taSourceAccountInfo\x12\x12\n\naccount_id\x18\x01 \x01(\t\x12\x16\n\x0e\x64\x61ta_source_id\x18\x02 \x01(\t\x12\x0c\n\x04name\x18\x03 \x01(\t\x12\x0f\n\x07is_sync\x18\x04 \x01(\x08\x12\x11\n\tis_linked\x18\x05 \x01(\x08\x12\x11\n\tdomain_id\x18\x15 \x01(\t\x12\x14\n\x0cworkspace_id\x18\x16 \x01(\t\x12\x16\n\x0ev_workspace_id\x18\x17 \x01(\t\x12\x12\n\ncreated_at\x18\x1f \x01(\t\x12\x12\n\nupdated_at\x18  \x01(\t\x12\x16\n\x0elast_synced_at\x18! \x01(\t\"t\n\x16\x44\x61taSourceAccountsInfo\x12\x45\n\x07results\x18\x01 \x03(\x0b\x32\x34.spaceone.api.cost_analysis.v1.DataSourceAccountInfo\x12\x13\n\x0btotal_count\x18\x02 \x01(\x05\"\x9e\x01\n\x1d\x44\x61taSourceAccountAnalyzeQuery\x12;\n\x05query\x18\x01 \x01(\x0b\x32,.spaceone.api.core.v2.TimeSeriesAnalyzeQuery\x12\x16\n\x0e\x64\x61ta_source_id\x18\x02 \x01(\t\x12\x12\n\naccount_id\x18\x03 \x01(\t\x12\x14\n\x0cworkspace_id\x18\x15 \x01(\t\"R\n\x1a\x44\x61taSourceAccountStatQuery\x12\x34\n\x05query\x18\x01 \x01(\x0b\x32%.spaceone.api.core.v2.StatisticsQuery2\xf0\x07\n\x11\x44\x61taSourceAccount\x12\xb6\x01\n\x06update\x12=.spaceone.api.cost_analysis.v1.UpdateDataSourceAccountRequest\x1a\x34.spaceone.api.cost_analysis.v1.DataSourceAccountInfo\"7\x82\xd3\xe4\x93\x02\x31\",/cost-analysis/v1/data-source-account/update:\x01*\x12\x95\x01\n\x05reset\x12<.spaceone.api.cost_analysis.v1.ResetDataSourceAccountRequest\x1a\x16.google.protobuf.Empty\"6\x82\xd3\xe4\x93\x02\x30\"+/cost-analysis/v1/data-source-account/reset:\x01*\x12\xaa\x01\n\x03get\x12\x37.spaceone.api.cost_analysis.v1.DataSourceAccountRequest\x1a\x34.spaceone.api.cost_analysis.v1.DataSourceAccountInfo\"4\x82\xd3\xe4\x93\x02.\")/cost-analysis/v1/data-source-account/get:\x01*\x12\xab\x01\n\x04list\x12\x35.spaceone.api.cost_analysis.v1.DataSourceAccountQuery\x1a\x35.spaceone.api.cost_analysis.v1.DataSourceAccountsInfo\"5\x82\xd3\xe4\x93\x02/\"*/cost-analysis/v1/data-source-account/list:\x01*\x12\x9a\x01\n\x07\x61nalyze\x12<.spaceone.api.cost_analysis.v1.DataSourceAccountAnalyzeQuery\x1a\x17.google.protobuf.Struct\"8\x82\xd3\xe4\x93\x02\x32\"-/cost-analysis/v1/data-source-account/analyze:\x01*\x12\x91\x01\n\x04stat\x12\x39.spaceone.api.cost_analysis.v1.DataSourceAccountStatQuery\x1a\x17.google.protobuf.Struct\"5\x82\xd3\xe4\x93\x02/\"*/cost-analysis/v1/data-source-account/stat:\x01*BDZBgithub.com/cloudforet-io/api/dist/go/spaceone/api/cost_analysis/v1b\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'spaceone.api.cost_analysis.v1.data_source_account_pb2', _globals)
if not _descriptor._USE_C_DESCRIPTORS:
  _globals['DESCRIPTOR']._loaded_options = None
  _globals['DESCRIPTOR']._serialized_options = b'ZBgithub.com/cloudforet-io/api/dist/go/spaceone/api/cost_analysis/v1'
  _globals['_DATASOURCEACCOUNT'].methods_by_name['update']._loaded_options = None
  _globals['_DATASOURCEACCOUNT'].methods_by_name['update']._serialized_options = b'\202\323\344\223\0021\",/cost-analysis/v1/data-source-account/update:\001*'
  _globals['_DATASOURCEACCOUNT'].methods_by_name['reset']._loaded_options = None
  _globals['_DATASOURCEACCOUNT'].methods_by_name['reset']._serialized_options = b'\202\323\344\223\0020\"+/cost-analysis/v1/data-source-account/reset:\001*'
  _globals['_DATASOURCEACCOUNT'].methods_by_name['get']._loaded_options = None
  _globals['_DATASOURCEACCOUNT'].methods_by_name['get']._serialized_options = b'\202\323\344\223\002.\")/cost-analysis/v1/data-source-account/get:\001*'
  _globals['_DATASOURCEACCOUNT'].methods_by_name['list']._loaded_options = None
  _globals['_DATASOURCEACCOUNT'].methods_by_name['list']._serialized_options = b'\202\323\344\223\002/\"*/cost-analysis/v1/data-source-account/list:\001*'
  _globals['_DATASOURCEACCOUNT'].methods_by_name['analyze']._loaded_options = None
  _globals['_DATASOURCEACCOUNT'].methods_by_name['analyze']._serialized_options = b'\202\323\344\223\0022\"-/cost-analysis/v1/data-source-account/analyze:\001*'
  _globals['_DATASOURCEACCOUNT'].methods_by_name['stat']._loaded_options = None
  _globals['_DATASOURCEACCOUNT'].methods_by_name['stat']._serialized_options = b'\202\323\344\223\002/\"*/cost-analysis/v1/data-source-account/stat:\001*'
  _globals['_UPDATEDATASOURCEACCOUNTREQUEST']._serialized_start=213
  _globals['_UPDATEDATASOURCEACCOUNTREQUEST']._serialized_end=311
  _globals['_RESETDATASOURCEACCOUNTREQUEST']._serialized_start=313
  _globals['_RESETDATASOURCEACCOUNTREQUEST']._serialized_end=408
  _globals['_DATASOURCEACCOUNTREQUEST']._serialized_start=410
  _globals['_DATASOURCEACCOUNTREQUEST']._serialized_end=488
  _globals['_DATASOURCEACCOUNTQUERY']._serialized_start=491
  _globals['_DATASOURCEACCOUNTQUERY']._serialized_end=673
  _globals['_DATASOURCEACCOUNTINFO']._serialized_start=676
  _globals['_DATASOURCEACCOUNTINFO']._serialized_end=922
  _globals['_DATASOURCEACCOUNTSINFO']._serialized_start=924
  _globals['_DATASOURCEACCOUNTSINFO']._serialized_end=1040
  _globals['_DATASOURCEACCOUNTANALYZEQUERY']._serialized_start=1043
  _globals['_DATASOURCEACCOUNTANALYZEQUERY']._serialized_end=1201
  _globals['_DATASOURCEACCOUNTSTATQUERY']._serialized_start=1203
  _globals['_DATASOURCEACCOUNTSTATQUERY']._serialized_end=1285
  _globals['_DATASOURCEACCOUNT']._serialized_start=1288
  _globals['_DATASOURCEACCOUNT']._serialized_end=2296
# @@protoc_insertion_point(module_scope)