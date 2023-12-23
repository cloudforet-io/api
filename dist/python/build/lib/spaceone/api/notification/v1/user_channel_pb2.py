# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: spaceone/api/notification/v1/user_channel.proto
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


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n/spaceone/api/notification/v1/user_channel.proto\x12\x1cspaceone.api.notification.v1\x1a\x1bgoogle/protobuf/empty.proto\x1a\x1cgoogle/protobuf/struct.proto\x1a\x1cgoogle/api/annotations.proto\x1a spaceone/api/core/v2/query.proto\"\xe3\x01\n\x13UserChannelSchedule\x12P\n\x0b\x64\x61y_of_week\x18\x01 \x03(\x0e\x32;.spaceone.api.notification.v1.UserChannelSchedule.DayOfWeek\x12\x12\n\nstart_hour\x18\x02 \x01(\x05\x12\x10\n\x08\x65nd_hour\x18\x03 \x01(\x05\"T\n\tDayOfWeek\x12\x08\n\x04NONE\x10\x00\x12\x07\n\x03MON\x10\x01\x12\x07\n\x03TUE\x10\x02\x12\x07\n\x03WED\x10\x03\x12\x07\n\x03THU\x10\x04\x12\x07\n\x03\x46RI\x10\x05\x12\x07\n\x03SAT\x10\x06\x12\x07\n\x03SUN\x10\x07\"\x93\x02\n\x18\x43reateUserChannelRequest\x12\x13\n\x0bprotocol_id\x18\x01 \x01(\t\x12\x0c\n\x04name\x18\x02 \x01(\t\x12%\n\x04\x64\x61ta\x18\x03 \x01(\x0b\x32\x17.google.protobuf.Struct\x12\x14\n\x0cis_subscribe\x18\x04 \x01(\x08\x12\x15\n\rsubscriptions\x18\x05 \x03(\t\x12\x14\n\x0cis_scheduled\x18\x06 \x01(\x08\x12\x43\n\x08schedule\x18\x07 \x01(\x0b\x32\x31.spaceone.api.notification.v1.UserChannelSchedule\x12%\n\x04tags\x18\x08 \x01(\x0b\x32\x17.google.protobuf.Struct\"\xd4\x01\n\x18UpdateUserChannelRequest\x12\x17\n\x0fuser_channel_id\x18\x01 \x01(\t\x12\x0c\n\x04name\x18\x02 \x01(\t\x12%\n\x04\x64\x61ta\x18\x03 \x01(\x0b\x32\x17.google.protobuf.Struct\x12\x43\n\x08schedule\x18\x04 \x01(\x0b\x32\x31.spaceone.api.notification.v1.UserChannelSchedule\x12%\n\x04tags\x18\x05 \x01(\x0b\x32\x17.google.protobuf.Struct\"\x96\x01\n UpdateUserChannelScheduleRequest\x12\x17\n\x0fuser_channel_id\x18\x01 \x01(\t\x12\x14\n\x0cis_scheduled\x18\x02 \x01(\x08\x12\x43\n\x08schedule\x18\x03 \x01(\x0b\x32\x31.spaceone.api.notification.v1.UserChannelSchedule\"l\n$UpdateUserChannelSubscriptionRequest\x12\x17\n\x0fuser_channel_id\x18\x01 \x01(\t\x12\x14\n\x0cis_subscribe\x18\x02 \x01(\x08\x12\x15\n\rsubscriptions\x18\x03 \x03(\t\"-\n\x12UserChannelRequest\x12\x17\n\x0fuser_channel_id\x18\x01 \x01(\t\"\xc2\x02\n\x10UserChannelQuery\x12*\n\x05query\x18\x01 \x01(\x0b\x32\x1b.spaceone.api.core.v2.Query\x12\x17\n\x0fuser_channel_id\x18\x02 \x01(\t\x12\x0c\n\x04name\x18\x03 \x01(\t\x12N\n\x05state\x18\x04 \x01(\x0e\x32?.spaceone.api.notification.v1.UserChannelQuery.UserChannelState\x12\x14\n\x0cis_subscribe\x18\x05 \x01(\x08\x12\x14\n\x0cis_scheduled\x18\x06 \x01(\x08\x12\x13\n\x0bprotocol_id\x18\x15 \x01(\t\x12\x11\n\tsecret_id\x18\x16 \x01(\t\"7\n\x10UserChannelState\x12\x08\n\x04NONE\x10\x00\x12\x0b\n\x07\x45NABLED\x10\x01\x12\x0c\n\x08\x44ISABLED\x10\x02\"\xf6\x03\n\x0fUserChannelInfo\x12\x17\n\x0fuser_channel_id\x18\x01 \x01(\t\x12\x0c\n\x04name\x18\x02 \x01(\t\x12M\n\x05state\x18\x03 \x01(\x0e\x32>.spaceone.api.notification.v1.UserChannelInfo.UserChannelState\x12%\n\x04\x64\x61ta\x18\x04 \x01(\x0b\x32\x17.google.protobuf.Struct\x12\x14\n\x0cis_subscribe\x18\x05 \x01(\x08\x12\x15\n\rsubscriptions\x18\x06 \x03(\t\x12\x14\n\x0cis_scheduled\x18\x07 \x01(\x08\x12\x43\n\x08schedule\x18\x08 \x01(\x0b\x32\x31.spaceone.api.notification.v1.UserChannelSchedule\x12%\n\x04tags\x18\t \x01(\x0b\x32\x17.google.protobuf.Struct\x12\x11\n\tdomain_id\x18\x15 \x01(\t\x12\x0f\n\x07user_id\x18\x16 \x01(\t\x12\x13\n\x0bprotocol_id\x18\x17 \x01(\t\x12\x11\n\tsecret_id\x18\x18 \x01(\t\x12\x12\n\ncreated_at\x18\x1f \x01(\t\"7\n\x10UserChannelState\x12\x08\n\x04NONE\x10\x00\x12\x0b\n\x07\x45NABLED\x10\x01\x12\x0c\n\x08\x44ISABLED\x10\x02\"g\n\x10UserChannelsInfo\x12>\n\x07results\x18\x01 \x03(\x0b\x32-.spaceone.api.notification.v1.UserChannelInfo\x12\x13\n\x0btotal_count\x18\x02 \x01(\x05\"L\n\x14UserChannelStatQuery\x12\x34\n\x05query\x18\x01 \x01(\x0b\x32%.spaceone.api.core.v2.StatisticsQuery2\xc3\x0c\n\x0bUserChannel\x12\xa0\x01\n\x06\x63reate\x12\x36.spaceone.api.notification.v1.CreateUserChannelRequest\x1a-.spaceone.api.notification.v1.UserChannelInfo\"/\x82\xd3\xe4\x93\x02)\"$/notification/v1/user-channel/create:\x01*\x12\xa0\x01\n\x06update\x12\x36.spaceone.api.notification.v1.UpdateUserChannelRequest\x1a-.spaceone.api.notification.v1.UserChannelInfo\"/\x82\xd3\xe4\x93\x02)\"$/notification/v1/user-channel/update:\x01*\x12\xb4\x01\n\x0cset_schedule\x12>.spaceone.api.notification.v1.UpdateUserChannelScheduleRequest\x1a-.spaceone.api.notification.v1.UserChannelInfo\"5\x82\xd3\xe4\x93\x02/\"*/notification/v1/user-channel/set-schedule:\x01*\x12\xc0\x01\n\x10set_subscription\x12\x42.spaceone.api.notification.v1.UpdateUserChannelSubscriptionRequest\x1a-.spaceone.api.notification.v1.UserChannelInfo\"9\x82\xd3\xe4\x93\x02\x33\"./notification/v1/user-channel/set-subscription:\x01*\x12\x9a\x01\n\x06\x65nable\x12\x30.spaceone.api.notification.v1.UserChannelRequest\x1a-.spaceone.api.notification.v1.UserChannelInfo\"/\x82\xd3\xe4\x93\x02)\"$/notification/v1/user-channel/enable:\x01*\x12\x9c\x01\n\x07\x64isable\x12\x30.spaceone.api.notification.v1.UserChannelRequest\x1a-.spaceone.api.notification.v1.UserChannelInfo\"0\x82\xd3\xe4\x93\x02*\"%/notification/v1/user-channel/disable:\x01*\x12\x83\x01\n\x06\x64\x65lete\x12\x30.spaceone.api.notification.v1.UserChannelRequest\x1a\x16.google.protobuf.Empty\"/\x82\xd3\xe4\x93\x02)\"$/notification/v1/user-channel/delete:\x01*\x12\x94\x01\n\x03get\x12\x30.spaceone.api.notification.v1.UserChannelRequest\x1a-.spaceone.api.notification.v1.UserChannelInfo\",\x82\xd3\xe4\x93\x02&\"!/notification/v1/user-channel/get:\x01*\x12\x95\x01\n\x04list\x12..spaceone.api.notification.v1.UserChannelQuery\x1a..spaceone.api.notification.v1.UserChannelsInfo\"-\x82\xd3\xe4\x93\x02\'\"\"/notification/v1/user-channel/list:\x01*\x12\x82\x01\n\x04stat\x12\x32.spaceone.api.notification.v1.UserChannelStatQuery\x1a\x17.google.protobuf.Struct\"-\x82\xd3\xe4\x93\x02\'\"\"/notification/v1/user-channel/stat:\x01*BCZAgithub.com/cloudforet-io/api/dist/go/spaceone/api/notification/v1b\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'spaceone.api.notification.v1.user_channel_pb2', _globals)
if _descriptor._USE_C_DESCRIPTORS == False:
  _globals['DESCRIPTOR']._options = None
  _globals['DESCRIPTOR']._serialized_options = b'ZAgithub.com/cloudforet-io/api/dist/go/spaceone/api/notification/v1'
  _globals['_USERCHANNEL'].methods_by_name['create']._options = None
  _globals['_USERCHANNEL'].methods_by_name['create']._serialized_options = b'\202\323\344\223\002)\"$/notification/v1/user-channel/create:\001*'
  _globals['_USERCHANNEL'].methods_by_name['update']._options = None
  _globals['_USERCHANNEL'].methods_by_name['update']._serialized_options = b'\202\323\344\223\002)\"$/notification/v1/user-channel/update:\001*'
  _globals['_USERCHANNEL'].methods_by_name['set_schedule']._options = None
  _globals['_USERCHANNEL'].methods_by_name['set_schedule']._serialized_options = b'\202\323\344\223\002/\"*/notification/v1/user-channel/set-schedule:\001*'
  _globals['_USERCHANNEL'].methods_by_name['set_subscription']._options = None
  _globals['_USERCHANNEL'].methods_by_name['set_subscription']._serialized_options = b'\202\323\344\223\0023\"./notification/v1/user-channel/set-subscription:\001*'
  _globals['_USERCHANNEL'].methods_by_name['enable']._options = None
  _globals['_USERCHANNEL'].methods_by_name['enable']._serialized_options = b'\202\323\344\223\002)\"$/notification/v1/user-channel/enable:\001*'
  _globals['_USERCHANNEL'].methods_by_name['disable']._options = None
  _globals['_USERCHANNEL'].methods_by_name['disable']._serialized_options = b'\202\323\344\223\002*\"%/notification/v1/user-channel/disable:\001*'
  _globals['_USERCHANNEL'].methods_by_name['delete']._options = None
  _globals['_USERCHANNEL'].methods_by_name['delete']._serialized_options = b'\202\323\344\223\002)\"$/notification/v1/user-channel/delete:\001*'
  _globals['_USERCHANNEL'].methods_by_name['get']._options = None
  _globals['_USERCHANNEL'].methods_by_name['get']._serialized_options = b'\202\323\344\223\002&\"!/notification/v1/user-channel/get:\001*'
  _globals['_USERCHANNEL'].methods_by_name['list']._options = None
  _globals['_USERCHANNEL'].methods_by_name['list']._serialized_options = b'\202\323\344\223\002\'\"\"/notification/v1/user-channel/list:\001*'
  _globals['_USERCHANNEL'].methods_by_name['stat']._options = None
  _globals['_USERCHANNEL'].methods_by_name['stat']._serialized_options = b'\202\323\344\223\002\'\"\"/notification/v1/user-channel/stat:\001*'
  _globals['_USERCHANNELSCHEDULE']._serialized_start=205
  _globals['_USERCHANNELSCHEDULE']._serialized_end=432
  _globals['_USERCHANNELSCHEDULE_DAYOFWEEK']._serialized_start=348
  _globals['_USERCHANNELSCHEDULE_DAYOFWEEK']._serialized_end=432
  _globals['_CREATEUSERCHANNELREQUEST']._serialized_start=435
  _globals['_CREATEUSERCHANNELREQUEST']._serialized_end=710
  _globals['_UPDATEUSERCHANNELREQUEST']._serialized_start=713
  _globals['_UPDATEUSERCHANNELREQUEST']._serialized_end=925
  _globals['_UPDATEUSERCHANNELSCHEDULEREQUEST']._serialized_start=928
  _globals['_UPDATEUSERCHANNELSCHEDULEREQUEST']._serialized_end=1078
  _globals['_UPDATEUSERCHANNELSUBSCRIPTIONREQUEST']._serialized_start=1080
  _globals['_UPDATEUSERCHANNELSUBSCRIPTIONREQUEST']._serialized_end=1188
  _globals['_USERCHANNELREQUEST']._serialized_start=1190
  _globals['_USERCHANNELREQUEST']._serialized_end=1235
  _globals['_USERCHANNELQUERY']._serialized_start=1238
  _globals['_USERCHANNELQUERY']._serialized_end=1560
  _globals['_USERCHANNELQUERY_USERCHANNELSTATE']._serialized_start=1505
  _globals['_USERCHANNELQUERY_USERCHANNELSTATE']._serialized_end=1560
  _globals['_USERCHANNELINFO']._serialized_start=1563
  _globals['_USERCHANNELINFO']._serialized_end=2065
  _globals['_USERCHANNELINFO_USERCHANNELSTATE']._serialized_start=1505
  _globals['_USERCHANNELINFO_USERCHANNELSTATE']._serialized_end=1560
  _globals['_USERCHANNELSINFO']._serialized_start=2067
  _globals['_USERCHANNELSINFO']._serialized_end=2170
  _globals['_USERCHANNELSTATQUERY']._serialized_start=2172
  _globals['_USERCHANNELSTATQUERY']._serialized_end=2248
  _globals['_USERCHANNEL']._serialized_start=2251
  _globals['_USERCHANNEL']._serialized_end=3854
# @@protoc_insertion_point(module_scope)
