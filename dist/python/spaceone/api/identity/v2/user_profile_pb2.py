# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: spaceone/api/identity/v2/user_profile.proto
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
from spaceone.api.identity.v2 import workspace_pb2 as spaceone_dot_api_dot_identity_dot_v2_dot_workspace__pb2
from spaceone.api.identity.v2 import user_pb2 as spaceone_dot_api_dot_identity_dot_v2_dot_user__pb2
from spaceone.api.identity.v2 import workspace_group_pb2 as spaceone_dot_api_dot_identity_dot_v2_dot_workspace__group__pb2
from spaceone.api.identity.v2 import role_binding_pb2 as spaceone_dot_api_dot_identity_dot_v2_dot_role__binding__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n+spaceone/api/identity/v2/user_profile.proto\x12\x18spaceone.api.identity.v2\x1a\x1bgoogle/protobuf/empty.proto\x1a\x1cgoogle/protobuf/struct.proto\x1a\x1cgoogle/api/annotations.proto\x1a spaceone/api/core/v2/query.proto\x1a(spaceone/api/identity/v2/workspace.proto\x1a#spaceone/api/identity/v2/user.proto\x1a.spaceone/api/identity/v2/workspace_group.proto\x1a+spaceone/api/identity/v2/role_binding.proto\"\x82\x01\n\x18UpdateUserProfileRequest\x12\x0c\n\x04name\x18\x01 \x01(\t\x12\r\n\x05\x65mail\x18\x02 \x01(\t\x12\x10\n\x08language\x18\x03 \x01(\t\x12\x10\n\x08timezone\x18\x04 \x01(\t\x12%\n\x04tags\x18\x05 \x01(\x0b\x32\x17.google.protobuf.Struct\"R\n UpdatePasswordUserProfileRequest\x12\x14\n\x0cnew_password\x18\x01 \x01(\t\x12\x18\n\x10\x63urrent_password\x18\x02 \x01(\t\"#\n\x12VerifyEmailRequest\x12\r\n\x05\x65mail\x18\x01 \x01(\t\"*\n\x13\x43onfirmEmailRequest\x12\x13\n\x0bverify_code\x18\x01 \x01(\t\"N\n\x10\x45nableMFARequest\x12\x10\n\x08mfa_type\x18\x01 \x01(\t\x12(\n\x07options\x18\x03 \x01(\x0b\x32\x17.google.protobuf.Struct\"\x13\n\x11\x44isableMFARequest\"(\n\x11\x43onfirmMFARequest\x12\x13\n\x0bverify_code\x18\x01 \x01(\t\"0\n\x12UserProfileRequest\x12\x1a\n\x12workspace_group_id\x18\x01 \x01(\t\"\"\n WorkspaceGroupUserProfileRequest\">\n\x18UserPasswordResetRequest\x12\x0f\n\x07user_id\x18\x01 \x01(\t\x12\x11\n\tdomain_id\x18\x15 \x01(\t\"\xbf\x04\n\x0fMyWorkspaceInfo\x12\x14\n\x0cworkspace_id\x18\x01 \x01(\t\x12\x0c\n\x04name\x18\x02 \x01(\t\x12>\n\x05state\x18\x03 \x01(\x0e\x32/.spaceone.api.identity.v2.MyWorkspaceInfo.State\x12\x11\n\trole_name\x18\x04 \x01(\t\x12\x45\n\trole_type\x18\x05 \x01(\x0e\x32\x32.spaceone.api.identity.v2.MyWorkspaceInfo.RoleType\x12%\n\x04tags\x18\x06 \x01(\x0b\x32\x17.google.protobuf.Struct\x12\x12\n\ncreated_by\x18\x07 \x01(\t\x12\x14\n\x0creference_id\x18\x08 \x01(\t\x12\x12\n\nis_managed\x18\t \x01(\x08\x12\x12\n\nis_dormant\x18\n \x01(\x08\x12\x11\n\tdomain_id\x18\x15 \x01(\t\x12\x0f\n\x07role_id\x18\x16 \x01(\t\x12\x12\n\ncreated_at\x18\x1f \x01(\t\x12\x16\n\x0elast_synced_at\x18  \x01(\t\x12\x1a\n\x12\x64ormant_updated_at\x18! \x01(\t\",\n\x05State\x12\x08\n\x04NONE\x10\x00\x12\x0b\n\x07\x45NABLED\x10\x01\x12\x0c\n\x08\x44ISABLED\x10\x02\"[\n\x08RoleType\x12\x12\n\x0eROLE_TYPE_NONE\x10\x00\x12\x10\n\x0c\x44OMAIN_ADMIN\x10\x01\x12\x13\n\x0fWORKSPACE_OWNER\x10\x02\x12\x14\n\x10WORKSPACE_MEMBER\x10\x03\"c\n\x10MyWorkspacesInfo\x12:\n\x07results\x18\x01 \x03(\x0b\x32).spaceone.api.identity.v2.MyWorkspaceInfo\x12\x13\n\x0btotal_count\x18\x02 \x01(\x05\"\xcd\x02\n\x14MyWorkspaceGroupInfo\x12\x1a\n\x12workspace_group_id\x18\x01 \x01(\t\x12\x0c\n\x04name\x18\x02 \x01(\t\x12;\n\x05users\x18\x03 \x03(\x0b\x32,.spaceone.api.identity.v2.UserWorkspaceGroup\x12%\n\x04tags\x18\x04 \x01(\x0b\x32\x17.google.protobuf.Struct\x12\x44\n\x11role_binding_info\x18\x05 \x01(\x0b\x32).spaceone.api.identity.v2.RoleBindingInfo\x12\x12\n\ncreated_by\x18\x06 \x01(\t\x12\x12\n\nupdated_by\x18\x07 \x01(\t\x12\x11\n\tdomain_id\x18\x15 \x01(\t\x12\x12\n\ncreated_at\x18\x1f \x01(\t\x12\x12\n\nupdated_at\x18  \x01(\t\"m\n\x15MyWorkspaceGroupsInfo\x12?\n\x07results\x18\x01 \x03(\x0b\x32..spaceone.api.identity.v2.MyWorkspaceGroupInfo\x12\x13\n\x0btotal_count\x18\x02 \x01(\x05\x32\x9b\r\n\x0bUserProfile\x12\x8d\x01\n\x06update\x12\x32.spaceone.api.identity.v2.UpdateUserProfileRequest\x1a\".spaceone.api.identity.v2.UserInfo\"+\x82\xd3\xe4\x93\x02%\" /identity/v2/user-profile/update:\x01*\x12\xa7\x01\n\x0fupdate_password\x12:.spaceone.api.identity.v2.UpdatePasswordUserProfileRequest\x1a\".spaceone.api.identity.v2.UserInfo\"4\x82\xd3\xe4\x93\x02.\")/identity/v2/user-profile/update-password:\x01*\x12\x87\x01\n\x0cverify_email\x12,.spaceone.api.identity.v2.VerifyEmailRequest\x1a\x16.google.protobuf.Empty\"1\x82\xd3\xe4\x93\x02+\"&/identity/v2/user-profile/verify-email:\x01*\x12\x96\x01\n\rconfirm_email\x12-.spaceone.api.identity.v2.ConfirmEmailRequest\x1a\".spaceone.api.identity.v2.UserInfo\"2\x82\xd3\xe4\x93\x02,\"\'/identity/v2/user-profile/confirm-email:\x01*\x12\x91\x01\n\x0ereset_password\x12\x32.spaceone.api.identity.v2.UserPasswordResetRequest\x1a\x16.google.protobuf.Empty\"3\x82\xd3\xe4\x93\x02-\"(/identity/v2/user-profile/reset-password:\x01*\x12\x8d\x01\n\nenable_mfa\x12*.spaceone.api.identity.v2.EnableMFARequest\x1a\".spaceone.api.identity.v2.UserInfo\"/\x82\xd3\xe4\x93\x02)\"$/identity/v2/user-profile/enable-mfa:\x01*\x12\x90\x01\n\x0b\x64isable_mfa\x12+.spaceone.api.identity.v2.DisableMFARequest\x1a\".spaceone.api.identity.v2.UserInfo\"0\x82\xd3\xe4\x93\x02*\"%/identity/v2/user-profile/disable-mfa:\x01*\x12\x90\x01\n\x0b\x63onfirm_mfa\x12+.spaceone.api.identity.v2.ConfirmMFARequest\x1a\".spaceone.api.identity.v2.UserInfo\"0\x82\xd3\xe4\x93\x02*\"%/identity/v2/user-profile/confirm-mfa:\x01*\x12\x81\x01\n\x03get\x12,.spaceone.api.identity.v2.UserProfileRequest\x1a\".spaceone.api.identity.v2.UserInfo\"(\x82\xd3\xe4\x93\x02\"\"\x1d/identity/v2/user-profile/get:\x01*\x12\x9f\x01\n\x0eget_workspaces\x12,.spaceone.api.identity.v2.UserProfileRequest\x1a*.spaceone.api.identity.v2.MyWorkspacesInfo\"3\x82\xd3\xe4\x93\x02-\"(/identity/v2/user-profile/get-workspaces:\x01*\x12\xbe\x01\n\x14get_workspace_groups\x12:.spaceone.api.identity.v2.WorkspaceGroupUserProfileRequest\x1a/.spaceone.api.identity.v2.MyWorkspaceGroupsInfo\"9\x82\xd3\xe4\x93\x02\x33\"./identity/v2/user-profile/get-workspace-groups:\x01*B?Z=github.com/cloudforet-io/api/dist/go/spaceone/api/identity/v2b\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'spaceone.api.identity.v2.user_profile_pb2', _globals)
if not _descriptor._USE_C_DESCRIPTORS:
  _globals['DESCRIPTOR']._loaded_options = None
  _globals['DESCRIPTOR']._serialized_options = b'Z=github.com/cloudforet-io/api/dist/go/spaceone/api/identity/v2'
  _globals['_USERPROFILE'].methods_by_name['update']._loaded_options = None
  _globals['_USERPROFILE'].methods_by_name['update']._serialized_options = b'\202\323\344\223\002%\" /identity/v2/user-profile/update:\001*'
  _globals['_USERPROFILE'].methods_by_name['update_password']._loaded_options = None
  _globals['_USERPROFILE'].methods_by_name['update_password']._serialized_options = b'\202\323\344\223\002.\")/identity/v2/user-profile/update-password:\001*'
  _globals['_USERPROFILE'].methods_by_name['verify_email']._loaded_options = None
  _globals['_USERPROFILE'].methods_by_name['verify_email']._serialized_options = b'\202\323\344\223\002+\"&/identity/v2/user-profile/verify-email:\001*'
  _globals['_USERPROFILE'].methods_by_name['confirm_email']._loaded_options = None
  _globals['_USERPROFILE'].methods_by_name['confirm_email']._serialized_options = b'\202\323\344\223\002,\"\'/identity/v2/user-profile/confirm-email:\001*'
  _globals['_USERPROFILE'].methods_by_name['reset_password']._loaded_options = None
  _globals['_USERPROFILE'].methods_by_name['reset_password']._serialized_options = b'\202\323\344\223\002-\"(/identity/v2/user-profile/reset-password:\001*'
  _globals['_USERPROFILE'].methods_by_name['enable_mfa']._loaded_options = None
  _globals['_USERPROFILE'].methods_by_name['enable_mfa']._serialized_options = b'\202\323\344\223\002)\"$/identity/v2/user-profile/enable-mfa:\001*'
  _globals['_USERPROFILE'].methods_by_name['disable_mfa']._loaded_options = None
  _globals['_USERPROFILE'].methods_by_name['disable_mfa']._serialized_options = b'\202\323\344\223\002*\"%/identity/v2/user-profile/disable-mfa:\001*'
  _globals['_USERPROFILE'].methods_by_name['confirm_mfa']._loaded_options = None
  _globals['_USERPROFILE'].methods_by_name['confirm_mfa']._serialized_options = b'\202\323\344\223\002*\"%/identity/v2/user-profile/confirm-mfa:\001*'
  _globals['_USERPROFILE'].methods_by_name['get']._loaded_options = None
  _globals['_USERPROFILE'].methods_by_name['get']._serialized_options = b'\202\323\344\223\002\"\"\035/identity/v2/user-profile/get:\001*'
  _globals['_USERPROFILE'].methods_by_name['get_workspaces']._loaded_options = None
  _globals['_USERPROFILE'].methods_by_name['get_workspaces']._serialized_options = b'\202\323\344\223\002-\"(/identity/v2/user-profile/get-workspaces:\001*'
  _globals['_USERPROFILE'].methods_by_name['get_workspace_groups']._loaded_options = None
  _globals['_USERPROFILE'].methods_by_name['get_workspace_groups']._serialized_options = b'\202\323\344\223\0023\"./identity/v2/user-profile/get-workspace-groups:\001*'
  _globals['_UPDATEUSERPROFILEREQUEST']._serialized_start=369
  _globals['_UPDATEUSERPROFILEREQUEST']._serialized_end=499
  _globals['_UPDATEPASSWORDUSERPROFILEREQUEST']._serialized_start=501
  _globals['_UPDATEPASSWORDUSERPROFILEREQUEST']._serialized_end=583
  _globals['_VERIFYEMAILREQUEST']._serialized_start=585
  _globals['_VERIFYEMAILREQUEST']._serialized_end=620
  _globals['_CONFIRMEMAILREQUEST']._serialized_start=622
  _globals['_CONFIRMEMAILREQUEST']._serialized_end=664
  _globals['_ENABLEMFAREQUEST']._serialized_start=666
  _globals['_ENABLEMFAREQUEST']._serialized_end=744
  _globals['_DISABLEMFAREQUEST']._serialized_start=746
  _globals['_DISABLEMFAREQUEST']._serialized_end=765
  _globals['_CONFIRMMFAREQUEST']._serialized_start=767
  _globals['_CONFIRMMFAREQUEST']._serialized_end=807
  _globals['_USERPROFILEREQUEST']._serialized_start=809
  _globals['_USERPROFILEREQUEST']._serialized_end=857
  _globals['_WORKSPACEGROUPUSERPROFILEREQUEST']._serialized_start=859
  _globals['_WORKSPACEGROUPUSERPROFILEREQUEST']._serialized_end=893
  _globals['_USERPASSWORDRESETREQUEST']._serialized_start=895
  _globals['_USERPASSWORDRESETREQUEST']._serialized_end=957
  _globals['_MYWORKSPACEINFO']._serialized_start=960
  _globals['_MYWORKSPACEINFO']._serialized_end=1535
  _globals['_MYWORKSPACEINFO_STATE']._serialized_start=1398
  _globals['_MYWORKSPACEINFO_STATE']._serialized_end=1442
  _globals['_MYWORKSPACEINFO_ROLETYPE']._serialized_start=1444
  _globals['_MYWORKSPACEINFO_ROLETYPE']._serialized_end=1535
  _globals['_MYWORKSPACESINFO']._serialized_start=1537
  _globals['_MYWORKSPACESINFO']._serialized_end=1636
  _globals['_MYWORKSPACEGROUPINFO']._serialized_start=1639
  _globals['_MYWORKSPACEGROUPINFO']._serialized_end=1972
  _globals['_MYWORKSPACEGROUPSINFO']._serialized_start=1974
  _globals['_MYWORKSPACEGROUPSINFO']._serialized_end=2083
  _globals['_USERPROFILE']._serialized_start=2086
  _globals['_USERPROFILE']._serialized_end=3777
# @@protoc_insertion_point(module_scope)
