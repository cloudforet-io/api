# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
"""Client and server classes corresponding to protobuf-defined services."""
import grpc
import warnings

from google.protobuf import empty_pb2 as google_dot_protobuf_dot_empty__pb2
from google.protobuf import struct_pb2 as google_dot_protobuf_dot_struct__pb2
from spaceone.api.identity.v2 import user_group_pb2 as spaceone_dot_api_dot_identity_dot_v2_dot_user__group__pb2

GRPC_GENERATED_VERSION = '1.65.1'
GRPC_VERSION = grpc.__version__
EXPECTED_ERROR_RELEASE = '1.66.0'
SCHEDULED_RELEASE_DATE = 'August 6, 2024'
_version_not_supported = False

try:
    from grpc._utilities import first_version_is_lower
    _version_not_supported = first_version_is_lower(GRPC_VERSION, GRPC_GENERATED_VERSION)
except ImportError:
    _version_not_supported = True

if _version_not_supported:
    warnings.warn(
        f'The grpc package installed is at version {GRPC_VERSION},'
        + f' but the generated code in spaceone/api/identity/v2/user_group_pb2_grpc.py depends on'
        + f' grpcio>={GRPC_GENERATED_VERSION}.'
        + f' Please upgrade your grpc module to grpcio>={GRPC_GENERATED_VERSION}'
        + f' or downgrade your generated code using grpcio-tools<={GRPC_VERSION}.'
        + f' This warning will become an error in {EXPECTED_ERROR_RELEASE},'
        + f' scheduled for release on {SCHEDULED_RELEASE_DATE}.',
        RuntimeWarning
    )


class UserGroupStub(object):
    """Missing associated documentation comment in .proto file."""

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.create = channel.unary_unary(
                '/spaceone.api.identity.v2.UserGroup/create',
                request_serializer=spaceone_dot_api_dot_identity_dot_v2_dot_user__group__pb2.CreateUserGroupRequest.SerializeToString,
                response_deserializer=spaceone_dot_api_dot_identity_dot_v2_dot_user__group__pb2.UserGroupInfo.FromString,
                _registered_method=True)
        self.update = channel.unary_unary(
                '/spaceone.api.identity.v2.UserGroup/update',
                request_serializer=spaceone_dot_api_dot_identity_dot_v2_dot_user__group__pb2.UpdateUserGroupRequest.SerializeToString,
                response_deserializer=spaceone_dot_api_dot_identity_dot_v2_dot_user__group__pb2.UserGroupInfo.FromString,
                _registered_method=True)
        self.delete = channel.unary_unary(
                '/spaceone.api.identity.v2.UserGroup/delete',
                request_serializer=spaceone_dot_api_dot_identity_dot_v2_dot_user__group__pb2.UserGroupRequest.SerializeToString,
                response_deserializer=google_dot_protobuf_dot_empty__pb2.Empty.FromString,
                _registered_method=True)
        self.add_users = channel.unary_unary(
                '/spaceone.api.identity.v2.UserGroup/add_users',
                request_serializer=spaceone_dot_api_dot_identity_dot_v2_dot_user__group__pb2.UsersUserGroupRequest.SerializeToString,
                response_deserializer=spaceone_dot_api_dot_identity_dot_v2_dot_user__group__pb2.UserGroupInfo.FromString,
                _registered_method=True)
        self.remove_users = channel.unary_unary(
                '/spaceone.api.identity.v2.UserGroup/remove_users',
                request_serializer=spaceone_dot_api_dot_identity_dot_v2_dot_user__group__pb2.UsersUserGroupRequest.SerializeToString,
                response_deserializer=spaceone_dot_api_dot_identity_dot_v2_dot_user__group__pb2.UserGroupInfo.FromString,
                _registered_method=True)
        self.get = channel.unary_unary(
                '/spaceone.api.identity.v2.UserGroup/get',
                request_serializer=spaceone_dot_api_dot_identity_dot_v2_dot_user__group__pb2.UserGroupRequest.SerializeToString,
                response_deserializer=spaceone_dot_api_dot_identity_dot_v2_dot_user__group__pb2.UserGroupInfo.FromString,
                _registered_method=True)
        self.list = channel.unary_unary(
                '/spaceone.api.identity.v2.UserGroup/list',
                request_serializer=spaceone_dot_api_dot_identity_dot_v2_dot_user__group__pb2.UserGroupSearchQuery.SerializeToString,
                response_deserializer=spaceone_dot_api_dot_identity_dot_v2_dot_user__group__pb2.UserGroupsInfo.FromString,
                _registered_method=True)
        self.stat = channel.unary_unary(
                '/spaceone.api.identity.v2.UserGroup/stat',
                request_serializer=spaceone_dot_api_dot_identity_dot_v2_dot_user__group__pb2.UserGroupStatQuery.SerializeToString,
                response_deserializer=google_dot_protobuf_dot_struct__pb2.Struct.FromString,
                _registered_method=True)


class UserGroupServicer(object):
    """Missing associated documentation comment in .proto file."""

    def create(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def update(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def delete(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def add_users(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def remove_users(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def get(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def list(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def stat(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_UserGroupServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'create': grpc.unary_unary_rpc_method_handler(
                    servicer.create,
                    request_deserializer=spaceone_dot_api_dot_identity_dot_v2_dot_user__group__pb2.CreateUserGroupRequest.FromString,
                    response_serializer=spaceone_dot_api_dot_identity_dot_v2_dot_user__group__pb2.UserGroupInfo.SerializeToString,
            ),
            'update': grpc.unary_unary_rpc_method_handler(
                    servicer.update,
                    request_deserializer=spaceone_dot_api_dot_identity_dot_v2_dot_user__group__pb2.UpdateUserGroupRequest.FromString,
                    response_serializer=spaceone_dot_api_dot_identity_dot_v2_dot_user__group__pb2.UserGroupInfo.SerializeToString,
            ),
            'delete': grpc.unary_unary_rpc_method_handler(
                    servicer.delete,
                    request_deserializer=spaceone_dot_api_dot_identity_dot_v2_dot_user__group__pb2.UserGroupRequest.FromString,
                    response_serializer=google_dot_protobuf_dot_empty__pb2.Empty.SerializeToString,
            ),
            'add_users': grpc.unary_unary_rpc_method_handler(
                    servicer.add_users,
                    request_deserializer=spaceone_dot_api_dot_identity_dot_v2_dot_user__group__pb2.UsersUserGroupRequest.FromString,
                    response_serializer=spaceone_dot_api_dot_identity_dot_v2_dot_user__group__pb2.UserGroupInfo.SerializeToString,
            ),
            'remove_users': grpc.unary_unary_rpc_method_handler(
                    servicer.remove_users,
                    request_deserializer=spaceone_dot_api_dot_identity_dot_v2_dot_user__group__pb2.UsersUserGroupRequest.FromString,
                    response_serializer=spaceone_dot_api_dot_identity_dot_v2_dot_user__group__pb2.UserGroupInfo.SerializeToString,
            ),
            'get': grpc.unary_unary_rpc_method_handler(
                    servicer.get,
                    request_deserializer=spaceone_dot_api_dot_identity_dot_v2_dot_user__group__pb2.UserGroupRequest.FromString,
                    response_serializer=spaceone_dot_api_dot_identity_dot_v2_dot_user__group__pb2.UserGroupInfo.SerializeToString,
            ),
            'list': grpc.unary_unary_rpc_method_handler(
                    servicer.list,
                    request_deserializer=spaceone_dot_api_dot_identity_dot_v2_dot_user__group__pb2.UserGroupSearchQuery.FromString,
                    response_serializer=spaceone_dot_api_dot_identity_dot_v2_dot_user__group__pb2.UserGroupsInfo.SerializeToString,
            ),
            'stat': grpc.unary_unary_rpc_method_handler(
                    servicer.stat,
                    request_deserializer=spaceone_dot_api_dot_identity_dot_v2_dot_user__group__pb2.UserGroupStatQuery.FromString,
                    response_serializer=google_dot_protobuf_dot_struct__pb2.Struct.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'spaceone.api.identity.v2.UserGroup', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))
    server.add_registered_method_handlers('spaceone.api.identity.v2.UserGroup', rpc_method_handlers)


 # This class is part of an EXPERIMENTAL API.
class UserGroup(object):
    """Missing associated documentation comment in .proto file."""

    @staticmethod
    def create(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(
            request,
            target,
            '/spaceone.api.identity.v2.UserGroup/create',
            spaceone_dot_api_dot_identity_dot_v2_dot_user__group__pb2.CreateUserGroupRequest.SerializeToString,
            spaceone_dot_api_dot_identity_dot_v2_dot_user__group__pb2.UserGroupInfo.FromString,
            options,
            channel_credentials,
            insecure,
            call_credentials,
            compression,
            wait_for_ready,
            timeout,
            metadata,
            _registered_method=True)

    @staticmethod
    def update(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(
            request,
            target,
            '/spaceone.api.identity.v2.UserGroup/update',
            spaceone_dot_api_dot_identity_dot_v2_dot_user__group__pb2.UpdateUserGroupRequest.SerializeToString,
            spaceone_dot_api_dot_identity_dot_v2_dot_user__group__pb2.UserGroupInfo.FromString,
            options,
            channel_credentials,
            insecure,
            call_credentials,
            compression,
            wait_for_ready,
            timeout,
            metadata,
            _registered_method=True)

    @staticmethod
    def delete(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(
            request,
            target,
            '/spaceone.api.identity.v2.UserGroup/delete',
            spaceone_dot_api_dot_identity_dot_v2_dot_user__group__pb2.UserGroupRequest.SerializeToString,
            google_dot_protobuf_dot_empty__pb2.Empty.FromString,
            options,
            channel_credentials,
            insecure,
            call_credentials,
            compression,
            wait_for_ready,
            timeout,
            metadata,
            _registered_method=True)

    @staticmethod
    def add_users(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(
            request,
            target,
            '/spaceone.api.identity.v2.UserGroup/add_users',
            spaceone_dot_api_dot_identity_dot_v2_dot_user__group__pb2.UsersUserGroupRequest.SerializeToString,
            spaceone_dot_api_dot_identity_dot_v2_dot_user__group__pb2.UserGroupInfo.FromString,
            options,
            channel_credentials,
            insecure,
            call_credentials,
            compression,
            wait_for_ready,
            timeout,
            metadata,
            _registered_method=True)

    @staticmethod
    def remove_users(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(
            request,
            target,
            '/spaceone.api.identity.v2.UserGroup/remove_users',
            spaceone_dot_api_dot_identity_dot_v2_dot_user__group__pb2.UsersUserGroupRequest.SerializeToString,
            spaceone_dot_api_dot_identity_dot_v2_dot_user__group__pb2.UserGroupInfo.FromString,
            options,
            channel_credentials,
            insecure,
            call_credentials,
            compression,
            wait_for_ready,
            timeout,
            metadata,
            _registered_method=True)

    @staticmethod
    def get(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(
            request,
            target,
            '/spaceone.api.identity.v2.UserGroup/get',
            spaceone_dot_api_dot_identity_dot_v2_dot_user__group__pb2.UserGroupRequest.SerializeToString,
            spaceone_dot_api_dot_identity_dot_v2_dot_user__group__pb2.UserGroupInfo.FromString,
            options,
            channel_credentials,
            insecure,
            call_credentials,
            compression,
            wait_for_ready,
            timeout,
            metadata,
            _registered_method=True)

    @staticmethod
    def list(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(
            request,
            target,
            '/spaceone.api.identity.v2.UserGroup/list',
            spaceone_dot_api_dot_identity_dot_v2_dot_user__group__pb2.UserGroupSearchQuery.SerializeToString,
            spaceone_dot_api_dot_identity_dot_v2_dot_user__group__pb2.UserGroupsInfo.FromString,
            options,
            channel_credentials,
            insecure,
            call_credentials,
            compression,
            wait_for_ready,
            timeout,
            metadata,
            _registered_method=True)

    @staticmethod
    def stat(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(
            request,
            target,
            '/spaceone.api.identity.v2.UserGroup/stat',
            spaceone_dot_api_dot_identity_dot_v2_dot_user__group__pb2.UserGroupStatQuery.SerializeToString,
            google_dot_protobuf_dot_struct__pb2.Struct.FromString,
            options,
            channel_credentials,
            insecure,
            call_credentials,
            compression,
            wait_for_ready,
            timeout,
            metadata,
            _registered_method=True)
