# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: spaceone/api/dashboard/v1/public_data_table.proto
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


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n1spaceone/api/dashboard/v1/public_data_table.proto\x12\x19spaceone.api.dashboard.v1\x1a\x1bgoogle/protobuf/empty.proto\x1a\x1cgoogle/protobuf/struct.proto\x1a\x1cgoogle/api/annotations.proto\x1a spaceone/api/core/v2/query.proto\" \n\x0b\x41ssetSource\x12\x11\n\tmetric_id\x18\x01 \x01(\t\"6\n\nCostSource\x12\x16\n\x0e\x64\x61ta_source_id\x18\x01 \x01(\t\x12\x10\n\x08\x64\x61ta_key\x18\x02 \x01(\t\"\xac\x01\n\x0cJoinOperator\x12\x13\n\x0b\x64\x61ta_tables\x18\x01 \x03(\t\x12<\n\x02on\x18\x02 \x01(\x0e\x32\x30.spaceone.api.dashboard.v1.JoinOperator.JoinType\"I\n\x08JoinType\x12\x12\n\x0eJOIN_TYPE_NONE\x10\x00\x12\t\n\x05INNER\x10\x01\x12\x08\n\x04LEFT\x10\x02\x12\t\n\x05RIGHT\x10\x03\x12\t\n\x05OUTER\x10\x04\"%\n\x0e\x43oncatOperator\x12\x13\n\x0b\x64\x61ta_tables\x18\x01 \x03(\t\"\xd8\x01\n\x11\x41ggregateOperator\x12\x15\n\rdata_table_id\x18\x01 \x01(\t\x12L\n\x08operator\x18\x02 \x01(\x0e\x32:.spaceone.api.dashboard.v1.AggregateOperator.AggregateType\x12\x10\n\x08group_by\x18\x03 \x01(\t\"L\n\rAggregateType\x12\x17\n\x13\x41GGREGATE_TYPE_NONE\x10\x00\x12\x07\n\x03SUM\x10\x01\x12\x07\n\x03\x41VG\x10\x02\x12\x07\n\x03MAX\x10\x03\x12\x07\n\x03MIN\x10\x04\":\n\rWhereOperator\x12\x15\n\rdata_table_id\x18\x01 \x01(\t\x12\x12\n\nconditions\x18\x02 \x03(\t\">\n\x10\x45valuateOperator\x12\x15\n\rdata_table_id\x18\x01 \x01(\t\x12\x13\n\x0b\x65xpressions\x18\x02 \x03(\t\"I\n\x08TimeDiff\x12\x0f\n\x05years\x18\x01 \x01(\x05H\x00\x12\x10\n\x06months\x18\x02 \x01(\x05H\x00\x12\x0e\n\x04\x64\x61ys\x18\x03 \x01(\x05H\x00\x42\n\n\x08timediff\"\xa9\x04\n\nAddOptions\x12\x37\n\x05\x41SSET\x18\x01 \x01(\x0b\x32&.spaceone.api.dashboard.v1.AssetSourceH\x00\x12\x35\n\x04\x43OST\x18\x02 \x01(\x0b\x32%.spaceone.api.dashboard.v1.CostSourceH\x00\x12\x11\n\tdata_name\x18\x03 \x01(\t\x12\x11\n\tdata_unit\x18\x04 \x01(\t\x12,\n\x08group_by\x18\x05 \x01(\x0b\x32\x1a.google.protobuf.ListValue\x12,\n\x06\x66ilter\x18\x06 \x03(\x0b\x32\x1c.spaceone.api.core.v2.Filter\x12/\n\tfilter_or\x18\x07 \x03(\x0b\x32\x1c.spaceone.api.core.v2.Filter\x12\x45\n\x0b\x64\x61te_format\x18\x08 \x01(\x0e\x32\x30.spaceone.api.dashboard.v1.AddOptions.DateFormat\x12\x32\n\x11\x61\x64\x64itional_labels\x18\t \x01(\x0b\x32\x17.google.protobuf.Struct\x12\x35\n\x08timediff\x18\n \x01(\x0b\x32#.spaceone.api.dashboard.v1.TimeDiff\"<\n\nDateFormat\x12\x14\n\x10\x44\x41TE_FORMAT_NONE\x10\x00\x12\n\n\x06SINGLE\x10\x01\x12\x0c\n\x08SEPARATE\x10\x02\x42\x08\n\x06source\"\xd3\x02\n\x10TransformOptions\x12\x37\n\x04JOIN\x18\x01 \x01(\x0b\x32\'.spaceone.api.dashboard.v1.JoinOperatorH\x00\x12;\n\x06\x43ONCAT\x18\x02 \x01(\x0b\x32).spaceone.api.dashboard.v1.ConcatOperatorH\x00\x12\x41\n\tAGGREGATE\x18\x03 \x01(\x0b\x32,.spaceone.api.dashboard.v1.AggregateOperatorH\x00\x12\x39\n\x05WHERE\x18\x04 \x01(\x0b\x32(.spaceone.api.dashboard.v1.WhereOperatorH\x00\x12?\n\x08\x45VALUATE\x18\x05 \x01(\x0b\x32+.spaceone.api.dashboard.v1.EvaluateOperatorH\x00\x42\n\n\x08operator\"\xd7\x01\n\x19\x41\x64\x64PublicDataTableRequest\x12\x11\n\twidget_id\x18\x01 \x01(\t\x12\x0c\n\x04name\x18\x02 \x01(\t\x12:\n\x0bsource_type\x18\x03 \x01(\x0e\x32%.spaceone.api.dashboard.v1.SourceType\x12\x36\n\x07options\x18\x04 \x01(\x0b\x32%.spaceone.api.dashboard.v1.AddOptions\x12%\n\x04tags\x18\x05 \x01(\x0b\x32\x17.google.protobuf.Struct\"\xde\x01\n\x1fTransformPublicDataTableRequest\x12\x11\n\twidget_id\x18\x01 \x01(\t\x12\x0c\n\x04name\x18\x02 \x01(\t\x12\x35\n\x08operator\x18\x03 \x01(\x0e\x32#.spaceone.api.dashboard.v1.Operator\x12<\n\x07options\x18\x04 \x01(\x0b\x32+.spaceone.api.dashboard.v1.TransformOptions\x12%\n\x04tags\x18\x05 \x01(\x0b\x32\x17.google.protobuf.Struct\"\x94\x01\n\x1cUpdatePublicDataTableRequest\x12\x15\n\rdata_table_id\x18\x01 \x01(\t\x12\x0c\n\x04name\x18\x02 \x01(\t\x12(\n\x07options\x18\x03 \x01(\x0b\x32\x17.google.protobuf.Struct\x12%\n\x04tags\x18\x04 \x01(\x0b\x32\x17.google.protobuf.Struct\"/\n\x16PublicDataTableRequest\x12\x15\n\rdata_table_id\x18\x01 \x01(\t\"\xc4\x02\n\x1aLoadPublicDataTableRequest\x12\x15\n\rdata_table_id\x18\x01 \x01(\t\x12V\n\x0bgranularity\x18\x02 \x01(\x0e\x32\x41.spaceone.api.dashboard.v1.LoadPublicDataTableRequest.Granularity\x12\r\n\x05start\x18\x03 \x01(\t\x12\x0b\n\x03\x65nd\x18\x04 \x01(\t\x12(\n\x04sort\x18\x05 \x03(\x0b\x32\x1a.spaceone.api.core.v2.Sort\x12(\n\x04page\x18\x06 \x01(\x0b\x32\x1a.spaceone.api.core.v2.Page\"G\n\x0bGranularity\x12\x14\n\x10GRANULARITY_NONE\x10\x00\x12\t\n\x05\x44\x41ILY\x10\x01\x12\x0b\n\x07MONTHLY\x10\x02\x12\n\n\x06YEARLY\x10\x03\"\xa5\x02\n\x14PublicDataTableQuery\x12*\n\x05query\x18\x01 \x01(\x0b\x32\x1b.spaceone.api.core.v2.Query\x12\x11\n\twidget_id\x18\x02 \x01(\t\x12\x15\n\rdata_table_id\x18\x03 \x01(\t\x12\x0c\n\x04name\x18\x04 \x01(\t\x12\x36\n\tdata_type\x18\x05 \x01(\x0e\x32#.spaceone.api.dashboard.v1.DataType\x12:\n\x0bsource_type\x18\x06 \x01(\x0e\x32%.spaceone.api.dashboard.v1.SourceType\x12\x35\n\x08operator\x18\x07 \x01(\x0e\x32#.spaceone.api.dashboard.v1.Operator\"\xc6\x05\n\x13PublicDataTableInfo\x12\x15\n\rdata_table_id\x18\x01 \x01(\t\x12\x0c\n\x04name\x18\x02 \x01(\t\x12\x36\n\tdata_type\x18\x03 \x01(\x0e\x32#.spaceone.api.dashboard.v1.DataType\x12:\n\x0bsource_type\x18\x04 \x01(\x0e\x32%.spaceone.api.dashboard.v1.SourceType\x12\x35\n\x08operator\x18\x05 \x01(\x0e\x32#.spaceone.api.dashboard.v1.Operator\x12(\n\x07options\x18\x06 \x01(\x0b\x32\x17.google.protobuf.Struct\x12%\n\x04tags\x18\x07 \x01(\x0b\x32\x17.google.protobuf.Struct\x12,\n\x0blabels_info\x18\x08 \x01(\x0b\x32\x17.google.protobuf.Struct\x12*\n\tdata_info\x18\t \x01(\x0b\x32\x17.google.protobuf.Struct\x12T\n\x0eresource_group\x18\x14 \x01(\x0e\x32<.spaceone.api.dashboard.v1.PublicDataTableInfo.ResourceGroup\x12\x11\n\tdomain_id\x18\x15 \x01(\t\x12\x14\n\x0cworkspace_id\x18\x16 \x01(\t\x12\x12\n\nproject_id\x18\x17 \x01(\t\x12\x14\n\x0c\x64\x61shboard_id\x18\x18 \x01(\t\x12\x11\n\twidget_id\x18\x19 \x01(\t\x12\x12\n\ncreated_at\x18\x1f \x01(\t\x12\x12\n\nupdated_at\x18  \x01(\t\"P\n\rResourceGroup\x12\x17\n\x13RESOURCE_GROUP_NONE\x10\x00\x12\n\n\x06\x44OMAIN\x10\x01\x12\r\n\tWORKSPACE\x10\x02\x12\x0b\n\x07PROJECT\x10\x03\"l\n\x14PublicDataTablesInfo\x12?\n\x07results\x18\x01 \x03(\x0b\x32..spaceone.api.dashboard.v1.PublicDataTableInfo\x12\x13\n\x0btotal_count\x18\x02 \x01(\x05*:\n\x08\x44\x61taType\x12\x12\n\x0e\x44\x41TA_TYPE_NONE\x10\x00\x12\t\n\x05\x41\x44\x44\x45\x44\x10\x01\x12\x0f\n\x0bTRANSFORMED\x10\x02*7\n\nSourceType\x12\x14\n\x10SOURCE_TYPE_NONE\x10\x00\x12\t\n\x05\x41SSET\x10\x01\x12\x08\n\x04\x43OST\x10\x02*[\n\x08Operator\x12\x11\n\rOPERATOR_NONE\x10\x00\x12\x08\n\x04JOIN\x10\x01\x12\n\n\x06\x43ONCAT\x10\x02\x12\r\n\tAGGREGATE\x10\x03\x12\t\n\x05WHERE\x10\x04\x12\x0c\n\x08\x45VALUATE\x10\x05\x32\xd0\x08\n\x0fPublicDataTable\x12\x9b\x01\n\x03\x61\x64\x64\x12\x34.spaceone.api.dashboard.v1.AddPublicDataTableRequest\x1a..spaceone.api.dashboard.v1.PublicDataTableInfo\".\x82\xd3\xe4\x93\x02(\"#/dashboard/v1/public-data-table/add:\x01*\x12\xad\x01\n\ttransform\x12:.spaceone.api.dashboard.v1.TransformPublicDataTableRequest\x1a..spaceone.api.dashboard.v1.PublicDataTableInfo\"4\x82\xd3\xe4\x93\x02.\")/dashboard/v1/public-data-table/transform:\x01*\x12\xa4\x01\n\x06update\x12\x37.spaceone.api.dashboard.v1.UpdatePublicDataTableRequest\x1a..spaceone.api.dashboard.v1.PublicDataTableInfo\"1\x82\xd3\xe4\x93\x02+\"&/dashboard/v1/public-data-table/update:\x01*\x12\x86\x01\n\x06\x64\x65lete\x12\x31.spaceone.api.dashboard.v1.PublicDataTableRequest\x1a\x16.google.protobuf.Empty\"1\x82\xd3\xe4\x93\x02+\"&/dashboard/v1/public-data-table/delete:\x01*\x12\x87\x01\n\x04load\x12\x35.spaceone.api.dashboard.v1.LoadPublicDataTableRequest\x1a\x17.google.protobuf.Struct\"/\x82\xd3\xe4\x93\x02)\"$/dashboard/v1/public-data-table/load:\x01*\x12\x98\x01\n\x03get\x12\x31.spaceone.api.dashboard.v1.PublicDataTableRequest\x1a..spaceone.api.dashboard.v1.PublicDataTableInfo\".\x82\xd3\xe4\x93\x02(\"#/dashboard/v1/public-data-table/get:\x01*\x12\x99\x01\n\x04list\x12/.spaceone.api.dashboard.v1.PublicDataTableQuery\x1a/.spaceone.api.dashboard.v1.PublicDataTablesInfo\"/\x82\xd3\xe4\x93\x02)\"$/dashboard/v1/public-data-table/list:\x01*B@Z>github.com/cloudforet-io/api/dist/go/spaceone/api/dashboard/v1b\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'spaceone.api.dashboard.v1.public_data_table_pb2', _globals)
if not _descriptor._USE_C_DESCRIPTORS:
  _globals['DESCRIPTOR']._loaded_options = None
  _globals['DESCRIPTOR']._serialized_options = b'Z>github.com/cloudforet-io/api/dist/go/spaceone/api/dashboard/v1'
  _globals['_PUBLICDATATABLE'].methods_by_name['add']._loaded_options = None
  _globals['_PUBLICDATATABLE'].methods_by_name['add']._serialized_options = b'\202\323\344\223\002(\"#/dashboard/v1/public-data-table/add:\001*'
  _globals['_PUBLICDATATABLE'].methods_by_name['transform']._loaded_options = None
  _globals['_PUBLICDATATABLE'].methods_by_name['transform']._serialized_options = b'\202\323\344\223\002.\")/dashboard/v1/public-data-table/transform:\001*'
  _globals['_PUBLICDATATABLE'].methods_by_name['update']._loaded_options = None
  _globals['_PUBLICDATATABLE'].methods_by_name['update']._serialized_options = b'\202\323\344\223\002+\"&/dashboard/v1/public-data-table/update:\001*'
  _globals['_PUBLICDATATABLE'].methods_by_name['delete']._loaded_options = None
  _globals['_PUBLICDATATABLE'].methods_by_name['delete']._serialized_options = b'\202\323\344\223\002+\"&/dashboard/v1/public-data-table/delete:\001*'
  _globals['_PUBLICDATATABLE'].methods_by_name['load']._loaded_options = None
  _globals['_PUBLICDATATABLE'].methods_by_name['load']._serialized_options = b'\202\323\344\223\002)\"$/dashboard/v1/public-data-table/load:\001*'
  _globals['_PUBLICDATATABLE'].methods_by_name['get']._loaded_options = None
  _globals['_PUBLICDATATABLE'].methods_by_name['get']._serialized_options = b'\202\323\344\223\002(\"#/dashboard/v1/public-data-table/get:\001*'
  _globals['_PUBLICDATATABLE'].methods_by_name['list']._loaded_options = None
  _globals['_PUBLICDATATABLE'].methods_by_name['list']._serialized_options = b'\202\323\344\223\002)\"$/dashboard/v1/public-data-table/list:\001*'
  _globals['_DATATYPE']._serialized_start=3912
  _globals['_DATATYPE']._serialized_end=3970
  _globals['_SOURCETYPE']._serialized_start=3972
  _globals['_SOURCETYPE']._serialized_end=4027
  _globals['_OPERATOR']._serialized_start=4029
  _globals['_OPERATOR']._serialized_end=4120
  _globals['_ASSETSOURCE']._serialized_start=203
  _globals['_ASSETSOURCE']._serialized_end=235
  _globals['_COSTSOURCE']._serialized_start=237
  _globals['_COSTSOURCE']._serialized_end=291
  _globals['_JOINOPERATOR']._serialized_start=294
  _globals['_JOINOPERATOR']._serialized_end=466
  _globals['_JOINOPERATOR_JOINTYPE']._serialized_start=393
  _globals['_JOINOPERATOR_JOINTYPE']._serialized_end=466
  _globals['_CONCATOPERATOR']._serialized_start=468
  _globals['_CONCATOPERATOR']._serialized_end=505
  _globals['_AGGREGATEOPERATOR']._serialized_start=508
  _globals['_AGGREGATEOPERATOR']._serialized_end=724
  _globals['_AGGREGATEOPERATOR_AGGREGATETYPE']._serialized_start=648
  _globals['_AGGREGATEOPERATOR_AGGREGATETYPE']._serialized_end=724
  _globals['_WHEREOPERATOR']._serialized_start=726
  _globals['_WHEREOPERATOR']._serialized_end=784
  _globals['_EVALUATEOPERATOR']._serialized_start=786
  _globals['_EVALUATEOPERATOR']._serialized_end=848
  _globals['_TIMEDIFF']._serialized_start=850
  _globals['_TIMEDIFF']._serialized_end=923
  _globals['_ADDOPTIONS']._serialized_start=926
  _globals['_ADDOPTIONS']._serialized_end=1479
  _globals['_ADDOPTIONS_DATEFORMAT']._serialized_start=1409
  _globals['_ADDOPTIONS_DATEFORMAT']._serialized_end=1469
  _globals['_TRANSFORMOPTIONS']._serialized_start=1482
  _globals['_TRANSFORMOPTIONS']._serialized_end=1821
  _globals['_ADDPUBLICDATATABLEREQUEST']._serialized_start=1824
  _globals['_ADDPUBLICDATATABLEREQUEST']._serialized_end=2039
  _globals['_TRANSFORMPUBLICDATATABLEREQUEST']._serialized_start=2042
  _globals['_TRANSFORMPUBLICDATATABLEREQUEST']._serialized_end=2264
  _globals['_UPDATEPUBLICDATATABLEREQUEST']._serialized_start=2267
  _globals['_UPDATEPUBLICDATATABLEREQUEST']._serialized_end=2415
  _globals['_PUBLICDATATABLEREQUEST']._serialized_start=2417
  _globals['_PUBLICDATATABLEREQUEST']._serialized_end=2464
  _globals['_LOADPUBLICDATATABLEREQUEST']._serialized_start=2467
  _globals['_LOADPUBLICDATATABLEREQUEST']._serialized_end=2791
  _globals['_LOADPUBLICDATATABLEREQUEST_GRANULARITY']._serialized_start=2720
  _globals['_LOADPUBLICDATATABLEREQUEST_GRANULARITY']._serialized_end=2791
  _globals['_PUBLICDATATABLEQUERY']._serialized_start=2794
  _globals['_PUBLICDATATABLEQUERY']._serialized_end=3087
  _globals['_PUBLICDATATABLEINFO']._serialized_start=3090
  _globals['_PUBLICDATATABLEINFO']._serialized_end=3800
  _globals['_PUBLICDATATABLEINFO_RESOURCEGROUP']._serialized_start=3720
  _globals['_PUBLICDATATABLEINFO_RESOURCEGROUP']._serialized_end=3800
  _globals['_PUBLICDATATABLESINFO']._serialized_start=3802
  _globals['_PUBLICDATATABLESINFO']._serialized_end=3910
  _globals['_PUBLICDATATABLE']._serialized_start=4123
  _globals['_PUBLICDATATABLE']._serialized_end=5227
# @@protoc_insertion_point(module_scope)
