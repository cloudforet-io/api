# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
"""Client and server classes corresponding to protobuf-defined services."""
import grpc

from spaceone.api.identity.plugin import external_auth_pb2 as spaceone_dot_api_dot_identity_dot_plugin_dot_external__auth__pb2


class ExternalAuthStub(object):
    """Missing associated documentation comment in .proto file."""

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.init = channel.unary_unary(
                '/spaceone.api.identity.plugin.ExternalAuth/init',
                request_serializer=spaceone_dot_api_dot_identity_dot_plugin_dot_external__auth__pb2.InitRequest.SerializeToString,
                response_deserializer=spaceone_dot_api_dot_identity_dot_plugin_dot_external__auth__pb2.PluginInfo.FromString,
                )
        self.authorize = channel.unary_unary(
                '/spaceone.api.identity.plugin.ExternalAuth/authorize',
                request_serializer=spaceone_dot_api_dot_identity_dot_plugin_dot_external__auth__pb2.AuthorizeRequest.SerializeToString,
                response_deserializer=spaceone_dot_api_dot_identity_dot_plugin_dot_external__auth__pb2.UserInfo.FromString,
                )


class ExternalAuthServicer(object):
    """Missing associated documentation comment in .proto file."""

    def init(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def authorize(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_ExternalAuthServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'init': grpc.unary_unary_rpc_method_handler(
                    servicer.init,
                    request_deserializer=spaceone_dot_api_dot_identity_dot_plugin_dot_external__auth__pb2.InitRequest.FromString,
                    response_serializer=spaceone_dot_api_dot_identity_dot_plugin_dot_external__auth__pb2.PluginInfo.SerializeToString,
            ),
            'authorize': grpc.unary_unary_rpc_method_handler(
                    servicer.authorize,
                    request_deserializer=spaceone_dot_api_dot_identity_dot_plugin_dot_external__auth__pb2.AuthorizeRequest.FromString,
                    response_serializer=spaceone_dot_api_dot_identity_dot_plugin_dot_external__auth__pb2.UserInfo.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'spaceone.api.identity.plugin.ExternalAuth', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))


 # This class is part of an EXPERIMENTAL API.
class ExternalAuth(object):
    """Missing associated documentation comment in .proto file."""

    @staticmethod
    def init(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/spaceone.api.identity.plugin.ExternalAuth/init',
            spaceone_dot_api_dot_identity_dot_plugin_dot_external__auth__pb2.InitRequest.SerializeToString,
            spaceone_dot_api_dot_identity_dot_plugin_dot_external__auth__pb2.PluginInfo.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def authorize(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/spaceone.api.identity.plugin.ExternalAuth/authorize',
            spaceone_dot_api_dot_identity_dot_plugin_dot_external__auth__pb2.AuthorizeRequest.SerializeToString,
            spaceone_dot_api_dot_identity_dot_plugin_dot_external__auth__pb2.UserInfo.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)