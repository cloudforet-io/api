# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: spaceone/api/monitoring/v1/event_rule.proto
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


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n+spaceone/api/monitoring/v1/event_rule.proto\x12\x1aspaceone.api.monitoring.v1\x1a\x1bgoogle/protobuf/empty.proto\x1a\x1cgoogle/protobuf/struct.proto\x1a\x1cgoogle/api/annotations.proto\x1a spaceone/api/core/v2/query.proto\"B\n\x12\x45ventRuleCondition\x12\x0b\n\x03key\x18\x01 \x01(\t\x12\r\n\x05value\x18\x02 \x01(\t\x12\x10\n\x08operator\x18\x03 \x01(\t\"F\n\x18\x45ventRuleActionResponder\x12\x15\n\rresource_type\x18\x01 \x01(\t\x12\x13\n\x0bresource_id\x18\x02 \x01(\t\"+\n\tMatchRule\x12\x0e\n\x06source\x18\x01 \x01(\t\x12\x0e\n\x06target\x18\x02 \x01(\t\"\xdd\x02\n\x10\x45ventRuleActions\x12\x17\n\x0f\x63hange_assignee\x18\x01 \x01(\t\x12\x16\n\x0e\x63hange_urgency\x18\x02 \x01(\t\x12\x16\n\x0e\x63hange_project\x18\x03 \x01(\t\x12\x1e\n\x16\x61\x64\x64_project_dependency\x18\x04 \x03(\t\x12K\n\radd_responder\x18\x05 \x03(\x0b\x32\x34.spaceone.api.monitoring.v1.EventRuleActionResponder\x12\x44\n\x15match_service_account\x18\x06 \x01(\x0b\x32%.spaceone.api.monitoring.v1.MatchRule\x12\x34\n\x13\x61\x64\x64_additional_info\x18\x07 \x01(\x0b\x32\x17.google.protobuf.Struct\x12\x17\n\x0fno_notification\x18\x08 \x01(\x08\"+\n\x10\x45ventRuleOptions\x12\x17\n\x0fstop_processing\x18\x01 \x01(\x08\"\xdf\x04\n\x16\x43reateEventRuleRequest\x12\x0c\n\x04name\x18\x01 \x01(\t\x12\x42\n\nconditions\x18\x02 \x03(\x0b\x32..spaceone.api.monitoring.v1.EventRuleCondition\x12^\n\x11\x63onditions_policy\x18\x03 \x01(\x0e\x32\x43.spaceone.api.monitoring.v1.CreateEventRuleRequest.ConditionsPolicy\x12=\n\x07\x61\x63tions\x18\x04 \x01(\x0b\x32,.spaceone.api.monitoring.v1.EventRuleActions\x12=\n\x07options\x18\x05 \x01(\x0b\x32,.spaceone.api.monitoring.v1.EventRuleOptions\x12%\n\x04tags\x18\x06 \x01(\x0b\x32\x17.google.protobuf.Struct\x12X\n\x0eresource_group\x18\x14 \x01(\x0e\x32@.spaceone.api.monitoring.v1.CreateEventRuleRequest.ResourceGroup\x12\x12\n\nproject_id\x18\x15 \x01(\t\":\n\x10\x43onditionsPolicy\x12\x08\n\x04NONE\x10\x00\x12\x07\n\x03\x41LL\x10\x01\x12\x07\n\x03\x41NY\x10\x02\x12\n\n\x06\x41LWAYS\x10\x03\"D\n\rResourceGroup\x12\x17\n\x13RESOURCE_GROUP_NONE\x10\x00\x12\r\n\tWORKSPACE\x10\x01\x12\x0b\n\x07PROJECT\x10\x02\"\xc2\x03\n\x16UpdateEventRuleRequest\x12\x15\n\revent_rule_id\x18\x01 \x01(\t\x12\x0c\n\x04name\x18\x02 \x01(\t\x12\x42\n\nconditions\x18\x03 \x03(\x0b\x32..spaceone.api.monitoring.v1.EventRuleCondition\x12^\n\x11\x63onditions_policy\x18\x04 \x01(\x0e\x32\x43.spaceone.api.monitoring.v1.UpdateEventRuleRequest.ConditionsPolicy\x12=\n\x07\x61\x63tions\x18\x05 \x01(\x0b\x32,.spaceone.api.monitoring.v1.EventRuleActions\x12=\n\x07options\x18\x06 \x01(\x0b\x32,.spaceone.api.monitoring.v1.EventRuleOptions\x12%\n\x04tags\x18\x07 \x01(\x0b\x32\x17.google.protobuf.Struct\":\n\x10\x43onditionsPolicy\x12\x08\n\x04NONE\x10\x00\x12\x07\n\x03\x41LL\x10\x01\x12\x07\n\x03\x41NY\x10\x02\x12\n\n\x06\x41LWAYS\x10\x03\"C\n\x1b\x43hangeEventRuleOrderRequest\x12\x15\n\revent_rule_id\x18\x01 \x01(\t\x12\r\n\x05order\x18\x02 \x01(\x05\")\n\x10\x45ventRuleRequest\x12\x15\n\revent_rule_id\x18\x01 \x01(\t\"\x8b\x01\n\x0e\x45ventRuleQuery\x12*\n\x05query\x18\x01 \x01(\x0b\x32\x1b.spaceone.api.core.v2.Query\x12\x15\n\revent_rule_id\x18\x02 \x01(\t\x12\x0c\n\x04name\x18\x03 \x01(\t\x12\x14\n\x0cworkspace_id\x18\x15 \x01(\t\x12\x12\n\nproject_id\x18\x16 \x01(\t\"\xa7\x05\n\rEventRuleInfo\x12\x15\n\revent_rule_id\x18\x01 \x01(\t\x12\x0c\n\x04name\x18\x02 \x01(\t\x12\r\n\x05order\x18\x03 \x01(\x05\x12\x42\n\nconditions\x18\x04 \x03(\x0b\x32..spaceone.api.monitoring.v1.EventRuleCondition\x12U\n\x11\x63onditions_policy\x18\x05 \x01(\x0e\x32:.spaceone.api.monitoring.v1.EventRuleInfo.ConditionsPolicy\x12=\n\x07\x61\x63tions\x18\x06 \x01(\x0b\x32,.spaceone.api.monitoring.v1.EventRuleActions\x12=\n\x07options\x18\x07 \x01(\x0b\x32,.spaceone.api.monitoring.v1.EventRuleOptions\x12%\n\x04tags\x18\x08 \x01(\x0b\x32\x17.google.protobuf.Struct\x12O\n\x0eresource_group\x18\x14 \x01(\x0e\x32\x37.spaceone.api.monitoring.v1.EventRuleInfo.ResourceGroup\x12\x11\n\tdomain_id\x18\x15 \x01(\t\x12\x14\n\x0cworkspace_id\x18\x16 \x01(\t\x12\x12\n\nproject_id\x18\x17 \x01(\t\x12\x12\n\ncreated_at\x18\x1f \x01(\t\":\n\x10\x43onditionsPolicy\x12\x08\n\x04NONE\x10\x00\x12\x07\n\x03\x41LL\x10\x01\x12\x07\n\x03\x41NY\x10\x02\x12\n\n\x06\x41LWAYS\x10\x03\"D\n\rResourceGroup\x12\x17\n\x13RESOURCE_GROUP_NONE\x10\x00\x12\r\n\tWORKSPACE\x10\x01\x12\x0b\n\x07PROJECT\x10\x02\"a\n\x0e\x45ventRulesInfo\x12:\n\x07results\x18\x01 \x03(\x0b\x32).spaceone.api.monitoring.v1.EventRuleInfo\x12\x13\n\x0btotal_count\x18\x02 \x01(\x05\"J\n\x12\x45ventRuleStatQuery\x12\x34\n\x05query\x18\x01 \x01(\x0b\x32%.spaceone.api.core.v2.StatisticsQuery2\xf1\x07\n\tEventRule\x12\x94\x01\n\x06\x63reate\x12\x32.spaceone.api.monitoring.v1.CreateEventRuleRequest\x1a).spaceone.api.monitoring.v1.EventRuleInfo\"+\x82\xd3\xe4\x93\x02%\" /monitoring/v1/event-rule/create:\x01*\x12\x94\x01\n\x06update\x12\x32.spaceone.api.monitoring.v1.UpdateEventRuleRequest\x1a).spaceone.api.monitoring.v1.EventRuleInfo\"+\x82\xd3\xe4\x93\x02%\" /monitoring/v1/event-rule/update:\x01*\x12\xa5\x01\n\x0c\x63hange_order\x12\x37.spaceone.api.monitoring.v1.ChangeEventRuleOrderRequest\x1a).spaceone.api.monitoring.v1.EventRuleInfo\"1\x82\xd3\xe4\x93\x02+\"&/monitoring/v1/event-rule/change-order:\x01*\x12{\n\x06\x64\x65lete\x12,.spaceone.api.monitoring.v1.EventRuleRequest\x1a\x16.google.protobuf.Empty\"+\x82\xd3\xe4\x93\x02%\" /monitoring/v1/event-rule/delete:\x01*\x12\x88\x01\n\x03get\x12,.spaceone.api.monitoring.v1.EventRuleRequest\x1a).spaceone.api.monitoring.v1.EventRuleInfo\"(\x82\xd3\xe4\x93\x02\"\"\x1d/monitoring/v1/event-rule/get:\x01*\x12\x89\x01\n\x04list\x12*.spaceone.api.monitoring.v1.EventRuleQuery\x1a*.spaceone.api.monitoring.v1.EventRulesInfo\")\x82\xd3\xe4\x93\x02#\"\x1e/monitoring/v1/event-rule/list:\x01*\x12z\n\x04stat\x12..spaceone.api.monitoring.v1.EventRuleStatQuery\x1a\x17.google.protobuf.Struct\")\x82\xd3\xe4\x93\x02#\"\x1e/monitoring/v1/event-rule/stat:\x01*BAZ?github.com/cloudforet-io/api/dist/go/spaceone/api/monitoring/v1b\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'spaceone.api.monitoring.v1.event_rule_pb2', _globals)
if _descriptor._USE_C_DESCRIPTORS == False:
  _globals['DESCRIPTOR']._options = None
  _globals['DESCRIPTOR']._serialized_options = b'Z?github.com/cloudforet-io/api/dist/go/spaceone/api/monitoring/v1'
  _globals['_EVENTRULE'].methods_by_name['create']._options = None
  _globals['_EVENTRULE'].methods_by_name['create']._serialized_options = b'\202\323\344\223\002%\" /monitoring/v1/event-rule/create:\001*'
  _globals['_EVENTRULE'].methods_by_name['update']._options = None
  _globals['_EVENTRULE'].methods_by_name['update']._serialized_options = b'\202\323\344\223\002%\" /monitoring/v1/event-rule/update:\001*'
  _globals['_EVENTRULE'].methods_by_name['change_order']._options = None
  _globals['_EVENTRULE'].methods_by_name['change_order']._serialized_options = b'\202\323\344\223\002+\"&/monitoring/v1/event-rule/change-order:\001*'
  _globals['_EVENTRULE'].methods_by_name['delete']._options = None
  _globals['_EVENTRULE'].methods_by_name['delete']._serialized_options = b'\202\323\344\223\002%\" /monitoring/v1/event-rule/delete:\001*'
  _globals['_EVENTRULE'].methods_by_name['get']._options = None
  _globals['_EVENTRULE'].methods_by_name['get']._serialized_options = b'\202\323\344\223\002\"\"\035/monitoring/v1/event-rule/get:\001*'
  _globals['_EVENTRULE'].methods_by_name['list']._options = None
  _globals['_EVENTRULE'].methods_by_name['list']._serialized_options = b'\202\323\344\223\002#\"\036/monitoring/v1/event-rule/list:\001*'
  _globals['_EVENTRULE'].methods_by_name['stat']._options = None
  _globals['_EVENTRULE'].methods_by_name['stat']._serialized_options = b'\202\323\344\223\002#\"\036/monitoring/v1/event-rule/stat:\001*'
  _globals['_EVENTRULECONDITION']._serialized_start=198
  _globals['_EVENTRULECONDITION']._serialized_end=264
  _globals['_EVENTRULEACTIONRESPONDER']._serialized_start=266
  _globals['_EVENTRULEACTIONRESPONDER']._serialized_end=336
  _globals['_MATCHRULE']._serialized_start=338
  _globals['_MATCHRULE']._serialized_end=381
  _globals['_EVENTRULEACTIONS']._serialized_start=384
  _globals['_EVENTRULEACTIONS']._serialized_end=733
  _globals['_EVENTRULEOPTIONS']._serialized_start=735
  _globals['_EVENTRULEOPTIONS']._serialized_end=778
  _globals['_CREATEEVENTRULEREQUEST']._serialized_start=781
  _globals['_CREATEEVENTRULEREQUEST']._serialized_end=1388
  _globals['_CREATEEVENTRULEREQUEST_CONDITIONSPOLICY']._serialized_start=1260
  _globals['_CREATEEVENTRULEREQUEST_CONDITIONSPOLICY']._serialized_end=1318
  _globals['_CREATEEVENTRULEREQUEST_RESOURCEGROUP']._serialized_start=1320
  _globals['_CREATEEVENTRULEREQUEST_RESOURCEGROUP']._serialized_end=1388
  _globals['_UPDATEEVENTRULEREQUEST']._serialized_start=1391
  _globals['_UPDATEEVENTRULEREQUEST']._serialized_end=1841
  _globals['_UPDATEEVENTRULEREQUEST_CONDITIONSPOLICY']._serialized_start=1260
  _globals['_UPDATEEVENTRULEREQUEST_CONDITIONSPOLICY']._serialized_end=1318
  _globals['_CHANGEEVENTRULEORDERREQUEST']._serialized_start=1843
  _globals['_CHANGEEVENTRULEORDERREQUEST']._serialized_end=1910
  _globals['_EVENTRULEREQUEST']._serialized_start=1912
  _globals['_EVENTRULEREQUEST']._serialized_end=1953
  _globals['_EVENTRULEQUERY']._serialized_start=1956
  _globals['_EVENTRULEQUERY']._serialized_end=2095
  _globals['_EVENTRULEINFO']._serialized_start=2098
  _globals['_EVENTRULEINFO']._serialized_end=2777
  _globals['_EVENTRULEINFO_CONDITIONSPOLICY']._serialized_start=1260
  _globals['_EVENTRULEINFO_CONDITIONSPOLICY']._serialized_end=1318
  _globals['_EVENTRULEINFO_RESOURCEGROUP']._serialized_start=1320
  _globals['_EVENTRULEINFO_RESOURCEGROUP']._serialized_end=1388
  _globals['_EVENTRULESINFO']._serialized_start=2779
  _globals['_EVENTRULESINFO']._serialized_end=2876
  _globals['_EVENTRULESTATQUERY']._serialized_start=2878
  _globals['_EVENTRULESTATQUERY']._serialized_end=2952
  _globals['_EVENTRULE']._serialized_start=2955
  _globals['_EVENTRULE']._serialized_end=3964
# @@protoc_insertion_point(module_scope)
