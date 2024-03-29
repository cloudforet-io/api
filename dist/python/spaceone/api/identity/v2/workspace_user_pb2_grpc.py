# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
"""Client and server classes corresponding to protobuf-defined services."""
import grpc

from google.protobuf import struct_pb2 as google_dot_protobuf_dot_struct__pb2
from spaceone.api.identity.v2 import workspace_user_pb2 as spaceone_dot_api_dot_identity_dot_v2_dot_workspace__user__pb2


class WorkspaceUserStub(object):
    """Missing associated documentation comment in .proto file."""

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.create = channel.unary_unary(
                '/spaceone.api.identity.v2.WorkspaceUser/create',
                request_serializer=spaceone_dot_api_dot_identity_dot_v2_dot_workspace__user__pb2.CreateWorkspaceUserRequest.SerializeToString,
                response_deserializer=spaceone_dot_api_dot_identity_dot_v2_dot_workspace__user__pb2.WorkspaceUserInfo.FromString,
                )
        self.get = channel.unary_unary(
                '/spaceone.api.identity.v2.WorkspaceUser/get',
                request_serializer=spaceone_dot_api_dot_identity_dot_v2_dot_workspace__user__pb2.WorkspaceUserRequest.SerializeToString,
                response_deserializer=spaceone_dot_api_dot_identity_dot_v2_dot_workspace__user__pb2.WorkspaceUserInfo.FromString,
                )
        self.find = channel.unary_unary(
                '/spaceone.api.identity.v2.WorkspaceUser/find',
                request_serializer=spaceone_dot_api_dot_identity_dot_v2_dot_workspace__user__pb2.WorkspaceUserFindRequest.SerializeToString,
                response_deserializer=spaceone_dot_api_dot_identity_dot_v2_dot_workspace__user__pb2.UsersSummaryInfo.FromString,
                )
        self.list = channel.unary_unary(
                '/spaceone.api.identity.v2.WorkspaceUser/list',
                request_serializer=spaceone_dot_api_dot_identity_dot_v2_dot_workspace__user__pb2.WorkspaceUserSearchQuery.SerializeToString,
                response_deserializer=spaceone_dot_api_dot_identity_dot_v2_dot_workspace__user__pb2.WorkspaceUsersInfo.FromString,
                )
        self.stat = channel.unary_unary(
                '/spaceone.api.identity.v2.WorkspaceUser/stat',
                request_serializer=spaceone_dot_api_dot_identity_dot_v2_dot_workspace__user__pb2.WorkspaceUserStatQuery.SerializeToString,
                response_deserializer=google_dot_protobuf_dot_struct__pb2.Struct.FromString,
                )


class WorkspaceUserServicer(object):
    """Missing associated documentation comment in .proto file."""

    def create(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def get(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def find(self, request, context):
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


def add_WorkspaceUserServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'create': grpc.unary_unary_rpc_method_handler(
                    servicer.create,
                    request_deserializer=spaceone_dot_api_dot_identity_dot_v2_dot_workspace__user__pb2.CreateWorkspaceUserRequest.FromString,
                    response_serializer=spaceone_dot_api_dot_identity_dot_v2_dot_workspace__user__pb2.WorkspaceUserInfo.SerializeToString,
            ),
            'get': grpc.unary_unary_rpc_method_handler(
                    servicer.get,
                    request_deserializer=spaceone_dot_api_dot_identity_dot_v2_dot_workspace__user__pb2.WorkspaceUserRequest.FromString,
                    response_serializer=spaceone_dot_api_dot_identity_dot_v2_dot_workspace__user__pb2.WorkspaceUserInfo.SerializeToString,
            ),
            'find': grpc.unary_unary_rpc_method_handler(
                    servicer.find,
                    request_deserializer=spaceone_dot_api_dot_identity_dot_v2_dot_workspace__user__pb2.WorkspaceUserFindRequest.FromString,
                    response_serializer=spaceone_dot_api_dot_identity_dot_v2_dot_workspace__user__pb2.UsersSummaryInfo.SerializeToString,
            ),
            'list': grpc.unary_unary_rpc_method_handler(
                    servicer.list,
                    request_deserializer=spaceone_dot_api_dot_identity_dot_v2_dot_workspace__user__pb2.WorkspaceUserSearchQuery.FromString,
                    response_serializer=spaceone_dot_api_dot_identity_dot_v2_dot_workspace__user__pb2.WorkspaceUsersInfo.SerializeToString,
            ),
            'stat': grpc.unary_unary_rpc_method_handler(
                    servicer.stat,
                    request_deserializer=spaceone_dot_api_dot_identity_dot_v2_dot_workspace__user__pb2.WorkspaceUserStatQuery.FromString,
                    response_serializer=google_dot_protobuf_dot_struct__pb2.Struct.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'spaceone.api.identity.v2.WorkspaceUser', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))


 # This class is part of an EXPERIMENTAL API.
class WorkspaceUser(object):
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
        return grpc.experimental.unary_unary(request, target, '/spaceone.api.identity.v2.WorkspaceUser/create',
            spaceone_dot_api_dot_identity_dot_v2_dot_workspace__user__pb2.CreateWorkspaceUserRequest.SerializeToString,
            spaceone_dot_api_dot_identity_dot_v2_dot_workspace__user__pb2.WorkspaceUserInfo.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

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
        return grpc.experimental.unary_unary(request, target, '/spaceone.api.identity.v2.WorkspaceUser/get',
            spaceone_dot_api_dot_identity_dot_v2_dot_workspace__user__pb2.WorkspaceUserRequest.SerializeToString,
            spaceone_dot_api_dot_identity_dot_v2_dot_workspace__user__pb2.WorkspaceUserInfo.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def find(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/spaceone.api.identity.v2.WorkspaceUser/find',
            spaceone_dot_api_dot_identity_dot_v2_dot_workspace__user__pb2.WorkspaceUserFindRequest.SerializeToString,
            spaceone_dot_api_dot_identity_dot_v2_dot_workspace__user__pb2.UsersSummaryInfo.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

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
        return grpc.experimental.unary_unary(request, target, '/spaceone.api.identity.v2.WorkspaceUser/list',
            spaceone_dot_api_dot_identity_dot_v2_dot_workspace__user__pb2.WorkspaceUserSearchQuery.SerializeToString,
            spaceone_dot_api_dot_identity_dot_v2_dot_workspace__user__pb2.WorkspaceUsersInfo.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

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
        return grpc.experimental.unary_unary(request, target, '/spaceone.api.identity.v2.WorkspaceUser/stat',
            spaceone_dot_api_dot_identity_dot_v2_dot_workspace__user__pb2.WorkspaceUserStatQuery.SerializeToString,
            google_dot_protobuf_dot_struct__pb2.Struct.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)
