# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: spaceone/api/alert_manager/v1/escalation_policy.proto
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


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n5spaceone/api/alert_manager/v1/escalation_policy.proto\x12\x1dspaceone.api.alert_manager.v1\x1a\x1bgoogle/protobuf/empty.proto\x1a\x1cgoogle/protobuf/struct.proto\x1a\x1cgoogle/api/annotations.proto\x1a spaceone/api/core/v2/query.proto\"<\n\x0e\x45scalationRule\x12\x10\n\x08\x63hannels\x18\x01 \x03(\t\x12\x18\n\x10\x65scalate_minutes\x18\x02 \x01(\x05\"\x9a\x01\n\x06Repeat\x12@\n\x05state\x18\x01 \x01(\x0e\x32\x31.spaceone.api.alert_manager.v1.Repeat.RepeatState\x12\r\n\x05\x63ount\x18\x02 \x01(\x05\"?\n\x0bRepeatState\x12\x15\n\x11REPEAT_STATE_NONE\x10\x00\x12\x0b\n\x07\x45NABLED\x10\x01\x12\x0c\n\x08\x44ISABLED\x10\x02\"\xa7\x02\n\x1d\x45scalationPolicyCreateRequest\x12\x0c\n\x04name\x18\x01 \x01(\t\x12<\n\x05rules\x18\x02 \x03(\x0b\x32-.spaceone.api.alert_manager.v1.EscalationRule\x12\x35\n\x06repeat\x18\x03 \x01(\x0b\x32%.spaceone.api.alert_manager.v1.Repeat\x12H\n\x10\x66inish_condition\x18\x04 \x01(\x0e\x32..spaceone.api.alert_manager.v1.FinishCondition\x12%\n\x04tags\x18\x0b \x01(\x0b\x32\x17.google.protobuf.Struct\x12\x12\n\nservice_id\x18\x15 \x01(\t\"\xb1\x02\n\x1d\x45scalationPolicyUpdateRequest\x12\x1c\n\x14\x65scalation_policy_id\x18\x01 \x01(\t\x12\x0c\n\x04name\x18\x02 \x01(\t\x12<\n\x05rules\x18\x03 \x03(\x0b\x32-.spaceone.api.alert_manager.v1.EscalationRule\x12H\n\x10\x66inish_condition\x18\x04 \x01(\x0e\x32..spaceone.api.alert_manager.v1.FinishCondition\x12\x35\n\x06repeat\x18\x05 \x01(\x0b\x32%.spaceone.api.alert_manager.v1.Repeat\x12%\n\x04tags\x18\x0b \x01(\x0b\x32\x17.google.protobuf.Struct\"7\n\x17\x45scalationPolicyRequest\x12\x1c\n\x14\x65scalation_policy_id\x18\x01 \x01(\t\"\xd3\x01\n\x1b\x45scalationPolicySearchQuery\x12*\n\x05query\x18\x01 \x01(\x0b\x32\x1b.spaceone.api.core.v2.Query\x12\x1c\n\x14\x65scalation_policy_id\x18\x02 \x01(\t\x12\x0c\n\x04name\x18\x03 \x01(\t\x12H\n\x10\x66inish_condition\x18\x04 \x01(\x0e\x32..spaceone.api.alert_manager.v1.FinishCondition\x12\x12\n\nservice_id\x18\x15 \x01(\t\"Q\n\x19\x45scalationPolicyStatQuery\x12\x34\n\x05query\x18\x01 \x01(\x0b\x32%.spaceone.api.core.v2.StatisticsQuery\"\x8d\x03\n\x14\x45scalationPolicyInfo\x12\x1c\n\x14\x65scalation_policy_id\x18\x01 \x01(\t\x12\x0c\n\x04name\x18\x02 \x01(\t\x12<\n\x05rules\x18\x03 \x03(\x0b\x32-.spaceone.api.alert_manager.v1.EscalationRule\x12\x35\n\x06repeat\x18\x04 \x01(\x0b\x32%.spaceone.api.alert_manager.v1.Repeat\x12H\n\x10\x66inish_condition\x18\x05 \x01(\x0e\x32..spaceone.api.alert_manager.v1.FinishCondition\x12%\n\x04tags\x18\x0b \x01(\x0b\x32\x17.google.protobuf.Struct\x12\x11\n\tdomain_id\x18\x15 \x01(\t\x12\x14\n\x0cworkspace_id\x18\x16 \x01(\t\x12\x12\n\nservice_id\x18\x17 \x01(\t\x12\x12\n\ncreated_at\x18\x1f \x01(\t\x12\x12\n\nupdated_at\x18  \x01(\t\"s\n\x16\x45scalationPoliciesInfo\x12\x44\n\x07results\x18\x01 \x03(\x0b\x32\x33.spaceone.api.alert_manager.v1.EscalationPolicyInfo\x12\x13\n\x0btotal_count\x18\x02 \x01(\x05*L\n\x0f\x46inishCondition\x12\x19\n\x15\x46INISH_CONDITION_NONE\x10\x00\x12\x10\n\x0c\x41\x43KNOWLEDGED\x10\x01\x12\x0c\n\x08RESOLVED\x10\x02\x32\xf9\x07\n\x10\x45scalationPolicy\x12\xb2\x01\n\x06\x63reate\x12<.spaceone.api.alert_manager.v1.EscalationPolicyCreateRequest\x1a\x33.spaceone.api.alert_manager.v1.EscalationPolicyInfo\"5\x82\xd3\xe4\x93\x02/\"*/alert-manager/v1/escalation-policy/create:\x01*\x12\xb2\x01\n\x06update\x12<.spaceone.api.alert_manager.v1.EscalationPolicyUpdateRequest\x1a\x33.spaceone.api.alert_manager.v1.EscalationPolicyInfo\"5\x82\xd3\xe4\x93\x02/\"*/alert-manager/v1/escalation-policy/update:\x01*\x12\x8f\x01\n\x06\x64\x65lete\x12\x36.spaceone.api.alert_manager.v1.EscalationPolicyRequest\x1a\x16.google.protobuf.Empty\"5\x82\xd3\xe4\x93\x02/\"*/alert-manager/v1/escalation-policy/delete:\x01*\x12\xa6\x01\n\x03get\x12\x36.spaceone.api.alert_manager.v1.EscalationPolicyRequest\x1a\x33.spaceone.api.alert_manager.v1.EscalationPolicyInfo\"2\x82\xd3\xe4\x93\x02,\"\'/alert-manager/v1/escalation-policy/get:\x01*\x12\xae\x01\n\x04list\x12:.spaceone.api.alert_manager.v1.EscalationPolicySearchQuery\x1a\x35.spaceone.api.alert_manager.v1.EscalationPoliciesInfo\"3\x82\xd3\xe4\x93\x02-\"(/alert-manager/v1/escalation-policy/list:\x01*\x12\x8e\x01\n\x04stat\x12\x38.spaceone.api.alert_manager.v1.EscalationPolicyStatQuery\x1a\x17.google.protobuf.Struct\"3\x82\xd3\xe4\x93\x02-\"(/alert-manager/v1/escalation-policy/stat:\x01*BDZBgithub.com/cloudforet-io/api/dist/go/spaceone/api/alert-manager/v1b\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'spaceone.api.alert_manager.v1.escalation_policy_pb2', _globals)
if not _descriptor._USE_C_DESCRIPTORS:
  _globals['DESCRIPTOR']._loaded_options = None
  _globals['DESCRIPTOR']._serialized_options = b'ZBgithub.com/cloudforet-io/api/dist/go/spaceone/api/alert-manager/v1'
  _globals['_ESCALATIONPOLICY'].methods_by_name['create']._loaded_options = None
  _globals['_ESCALATIONPOLICY'].methods_by_name['create']._serialized_options = b'\202\323\344\223\002/\"*/alert-manager/v1/escalation-policy/create:\001*'
  _globals['_ESCALATIONPOLICY'].methods_by_name['update']._loaded_options = None
  _globals['_ESCALATIONPOLICY'].methods_by_name['update']._serialized_options = b'\202\323\344\223\002/\"*/alert-manager/v1/escalation-policy/update:\001*'
  _globals['_ESCALATIONPOLICY'].methods_by_name['delete']._loaded_options = None
  _globals['_ESCALATIONPOLICY'].methods_by_name['delete']._serialized_options = b'\202\323\344\223\002/\"*/alert-manager/v1/escalation-policy/delete:\001*'
  _globals['_ESCALATIONPOLICY'].methods_by_name['get']._loaded_options = None
  _globals['_ESCALATIONPOLICY'].methods_by_name['get']._serialized_options = b'\202\323\344\223\002,\"\'/alert-manager/v1/escalation-policy/get:\001*'
  _globals['_ESCALATIONPOLICY'].methods_by_name['list']._loaded_options = None
  _globals['_ESCALATIONPOLICY'].methods_by_name['list']._serialized_options = b'\202\323\344\223\002-\"(/alert-manager/v1/escalation-policy/list:\001*'
  _globals['_ESCALATIONPOLICY'].methods_by_name['stat']._loaded_options = None
  _globals['_ESCALATIONPOLICY'].methods_by_name['stat']._serialized_options = b'\202\323\344\223\002-\"(/alert-manager/v1/escalation-policy/stat:\001*'
  _globals['_FINISHCONDITION']._serialized_start=1907
  _globals['_FINISHCONDITION']._serialized_end=1983
  _globals['_ESCALATIONRULE']._serialized_start=211
  _globals['_ESCALATIONRULE']._serialized_end=271
  _globals['_REPEAT']._serialized_start=274
  _globals['_REPEAT']._serialized_end=428
  _globals['_REPEAT_REPEATSTATE']._serialized_start=365
  _globals['_REPEAT_REPEATSTATE']._serialized_end=428
  _globals['_ESCALATIONPOLICYCREATEREQUEST']._serialized_start=431
  _globals['_ESCALATIONPOLICYCREATEREQUEST']._serialized_end=726
  _globals['_ESCALATIONPOLICYUPDATEREQUEST']._serialized_start=729
  _globals['_ESCALATIONPOLICYUPDATEREQUEST']._serialized_end=1034
  _globals['_ESCALATIONPOLICYREQUEST']._serialized_start=1036
  _globals['_ESCALATIONPOLICYREQUEST']._serialized_end=1091
  _globals['_ESCALATIONPOLICYSEARCHQUERY']._serialized_start=1094
  _globals['_ESCALATIONPOLICYSEARCHQUERY']._serialized_end=1305
  _globals['_ESCALATIONPOLICYSTATQUERY']._serialized_start=1307
  _globals['_ESCALATIONPOLICYSTATQUERY']._serialized_end=1388
  _globals['_ESCALATIONPOLICYINFO']._serialized_start=1391
  _globals['_ESCALATIONPOLICYINFO']._serialized_end=1788
  _globals['_ESCALATIONPOLICIESINFO']._serialized_start=1790
  _globals['_ESCALATIONPOLICIESINFO']._serialized_end=1905
  _globals['_ESCALATIONPOLICY']._serialized_start=1986
  _globals['_ESCALATIONPOLICY']._serialized_end=3003
# @@protoc_insertion_point(module_scope)
