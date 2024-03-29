# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
"""Client and server classes corresponding to protobuf-defined services."""
import grpc

from google.protobuf import empty_pb2 as google_dot_protobuf_dot_empty__pb2
from google.protobuf import struct_pb2 as google_dot_protobuf_dot_struct__pb2
from spaceone.api.identity.v2 import role_binding_pb2 as spaceone_dot_api_dot_identity_dot_v2_dot_role__binding__pb2


class RoleBindingStub(object):
    """Missing associated documentation comment in .proto file."""

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.create = channel.unary_unary(
                '/spaceone.api.identity.v2.RoleBinding/create',
                request_serializer=spaceone_dot_api_dot_identity_dot_v2_dot_role__binding__pb2.CreateRoleBindingRequest.SerializeToString,
                response_deserializer=spaceone_dot_api_dot_identity_dot_v2_dot_role__binding__pb2.RoleBindingInfo.FromString,
                )
        self.update_role = channel.unary_unary(
                '/spaceone.api.identity.v2.RoleBinding/update_role',
                request_serializer=spaceone_dot_api_dot_identity_dot_v2_dot_role__binding__pb2.UpdateRoleBindingRequest.SerializeToString,
                response_deserializer=spaceone_dot_api_dot_identity_dot_v2_dot_role__binding__pb2.RoleBindingInfo.FromString,
                )
        self.delete = channel.unary_unary(
                '/spaceone.api.identity.v2.RoleBinding/delete',
                request_serializer=spaceone_dot_api_dot_identity_dot_v2_dot_role__binding__pb2.RoleBindingRequest.SerializeToString,
                response_deserializer=google_dot_protobuf_dot_empty__pb2.Empty.FromString,
                )
        self.get = channel.unary_unary(
                '/spaceone.api.identity.v2.RoleBinding/get',
                request_serializer=spaceone_dot_api_dot_identity_dot_v2_dot_role__binding__pb2.RoleBindingRequest.SerializeToString,
                response_deserializer=spaceone_dot_api_dot_identity_dot_v2_dot_role__binding__pb2.RoleBindingInfo.FromString,
                )
        self.list = channel.unary_unary(
                '/spaceone.api.identity.v2.RoleBinding/list',
                request_serializer=spaceone_dot_api_dot_identity_dot_v2_dot_role__binding__pb2.RoleBindingSearchQuery.SerializeToString,
                response_deserializer=spaceone_dot_api_dot_identity_dot_v2_dot_role__binding__pb2.RoleBindingsInfo.FromString,
                )
        self.stat = channel.unary_unary(
                '/spaceone.api.identity.v2.RoleBinding/stat',
                request_serializer=spaceone_dot_api_dot_identity_dot_v2_dot_role__binding__pb2.RoleBindingStatQuery.SerializeToString,
                response_deserializer=google_dot_protobuf_dot_struct__pb2.Struct.FromString,
                )


class RoleBindingServicer(object):
    """Missing associated documentation comment in .proto file."""

    def create(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def update_role(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def delete(self, request, context):
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


def add_RoleBindingServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'create': grpc.unary_unary_rpc_method_handler(
                    servicer.create,
                    request_deserializer=spaceone_dot_api_dot_identity_dot_v2_dot_role__binding__pb2.CreateRoleBindingRequest.FromString,
                    response_serializer=spaceone_dot_api_dot_identity_dot_v2_dot_role__binding__pb2.RoleBindingInfo.SerializeToString,
            ),
            'update_role': grpc.unary_unary_rpc_method_handler(
                    servicer.update_role,
                    request_deserializer=spaceone_dot_api_dot_identity_dot_v2_dot_role__binding__pb2.UpdateRoleBindingRequest.FromString,
                    response_serializer=spaceone_dot_api_dot_identity_dot_v2_dot_role__binding__pb2.RoleBindingInfo.SerializeToString,
            ),
            'delete': grpc.unary_unary_rpc_method_handler(
                    servicer.delete,
                    request_deserializer=spaceone_dot_api_dot_identity_dot_v2_dot_role__binding__pb2.RoleBindingRequest.FromString,
                    response_serializer=google_dot_protobuf_dot_empty__pb2.Empty.SerializeToString,
            ),
            'get': grpc.unary_unary_rpc_method_handler(
                    servicer.get,
                    request_deserializer=spaceone_dot_api_dot_identity_dot_v2_dot_role__binding__pb2.RoleBindingRequest.FromString,
                    response_serializer=spaceone_dot_api_dot_identity_dot_v2_dot_role__binding__pb2.RoleBindingInfo.SerializeToString,
            ),
            'list': grpc.unary_unary_rpc_method_handler(
                    servicer.list,
                    request_deserializer=spaceone_dot_api_dot_identity_dot_v2_dot_role__binding__pb2.RoleBindingSearchQuery.FromString,
                    response_serializer=spaceone_dot_api_dot_identity_dot_v2_dot_role__binding__pb2.RoleBindingsInfo.SerializeToString,
            ),
            'stat': grpc.unary_unary_rpc_method_handler(
                    servicer.stat,
                    request_deserializer=spaceone_dot_api_dot_identity_dot_v2_dot_role__binding__pb2.RoleBindingStatQuery.FromString,
                    response_serializer=google_dot_protobuf_dot_struct__pb2.Struct.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'spaceone.api.identity.v2.RoleBinding', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))


 # This class is part of an EXPERIMENTAL API.
class RoleBinding(object):
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
        return grpc.experimental.unary_unary(request, target, '/spaceone.api.identity.v2.RoleBinding/create',
            spaceone_dot_api_dot_identity_dot_v2_dot_role__binding__pb2.CreateRoleBindingRequest.SerializeToString,
            spaceone_dot_api_dot_identity_dot_v2_dot_role__binding__pb2.RoleBindingInfo.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def update_role(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/spaceone.api.identity.v2.RoleBinding/update_role',
            spaceone_dot_api_dot_identity_dot_v2_dot_role__binding__pb2.UpdateRoleBindingRequest.SerializeToString,
            spaceone_dot_api_dot_identity_dot_v2_dot_role__binding__pb2.RoleBindingInfo.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

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
        return grpc.experimental.unary_unary(request, target, '/spaceone.api.identity.v2.RoleBinding/delete',
            spaceone_dot_api_dot_identity_dot_v2_dot_role__binding__pb2.RoleBindingRequest.SerializeToString,
            google_dot_protobuf_dot_empty__pb2.Empty.FromString,
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
        return grpc.experimental.unary_unary(request, target, '/spaceone.api.identity.v2.RoleBinding/get',
            spaceone_dot_api_dot_identity_dot_v2_dot_role__binding__pb2.RoleBindingRequest.SerializeToString,
            spaceone_dot_api_dot_identity_dot_v2_dot_role__binding__pb2.RoleBindingInfo.FromString,
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
        return grpc.experimental.unary_unary(request, target, '/spaceone.api.identity.v2.RoleBinding/list',
            spaceone_dot_api_dot_identity_dot_v2_dot_role__binding__pb2.RoleBindingSearchQuery.SerializeToString,
            spaceone_dot_api_dot_identity_dot_v2_dot_role__binding__pb2.RoleBindingsInfo.FromString,
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
        return grpc.experimental.unary_unary(request, target, '/spaceone.api.identity.v2.RoleBinding/stat',
            spaceone_dot_api_dot_identity_dot_v2_dot_role__binding__pb2.RoleBindingStatQuery.SerializeToString,
            google_dot_protobuf_dot_struct__pb2.Struct.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)
