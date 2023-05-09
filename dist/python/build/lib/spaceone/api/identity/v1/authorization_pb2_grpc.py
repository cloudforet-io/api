# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
"""Client and server classes corresponding to protobuf-defined services."""
import grpc

from spaceone.api.core.v1 import handler_pb2 as spaceone_dot_api_dot_core_dot_v1_dot_handler__pb2


class AuthorizationStub(object):
    """Missing associated documentation comment in .proto file."""

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.verify = channel.unary_unary(
                '/spaceone.api.identity.v1.Authorization/verify',
                request_serializer=spaceone_dot_api_dot_core_dot_v1_dot_handler__pb2.AuthorizationRequest.SerializeToString,
                response_deserializer=spaceone_dot_api_dot_core_dot_v1_dot_handler__pb2.AuthorizationResponse.FromString,
                )


class AuthorizationServicer(object):
    """Missing associated documentation comment in .proto file."""

    def verify(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_AuthorizationServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'verify': grpc.unary_unary_rpc_method_handler(
                    servicer.verify,
                    request_deserializer=spaceone_dot_api_dot_core_dot_v1_dot_handler__pb2.AuthorizationRequest.FromString,
                    response_serializer=spaceone_dot_api_dot_core_dot_v1_dot_handler__pb2.AuthorizationResponse.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'spaceone.api.identity.v1.Authorization', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))


 # This class is part of an EXPERIMENTAL API.
class Authorization(object):
    """Missing associated documentation comment in .proto file."""

    @staticmethod
    def verify(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/spaceone.api.identity.v1.Authorization/verify',
            spaceone_dot_api_dot_core_dot_v1_dot_handler__pb2.AuthorizationRequest.SerializeToString,
            spaceone_dot_api_dot_core_dot_v1_dot_handler__pb2.AuthorizationResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)