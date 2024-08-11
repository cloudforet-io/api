# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
"""Client and server classes corresponding to protobuf-defined services."""
import grpc
import warnings

from spaceone.api.sample.v1 import helloworld_pb2 as spaceone_dot_api_dot_sample_dot_v1_dot_helloworld__pb2

GRPC_GENERATED_VERSION = '1.65.4'
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
        + f' but the generated code in spaceone/api/sample/v1/helloworld_pb2_grpc.py depends on'
        + f' grpcio>={GRPC_GENERATED_VERSION}.'
        + f' Please upgrade your grpc module to grpcio>={GRPC_GENERATED_VERSION}'
        + f' or downgrade your generated code using grpcio-tools<={GRPC_VERSION}.'
        + f' This warning will become an error in {EXPECTED_ERROR_RELEASE},'
        + f' scheduled for release on {SCHEDULED_RELEASE_DATE}.',
        RuntimeWarning
    )


class HelloWorldStub(object):
    """The greeting service definition.
    """

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.say_hello = channel.unary_unary(
                '/spaceone.api.sample.v1.HelloWorld/say_hello',
                request_serializer=spaceone_dot_api_dot_sample_dot_v1_dot_helloworld__pb2.HelloRequest.SerializeToString,
                response_deserializer=spaceone_dot_api_dot_sample_dot_v1_dot_helloworld__pb2.HelloReply.FromString,
                _registered_method=True)


class HelloWorldServicer(object):
    """The greeting service definition.
    """

    def say_hello(self, request, context):
        """Sends a greeting
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_HelloWorldServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'say_hello': grpc.unary_unary_rpc_method_handler(
                    servicer.say_hello,
                    request_deserializer=spaceone_dot_api_dot_sample_dot_v1_dot_helloworld__pb2.HelloRequest.FromString,
                    response_serializer=spaceone_dot_api_dot_sample_dot_v1_dot_helloworld__pb2.HelloReply.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'spaceone.api.sample.v1.HelloWorld', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))
    server.add_registered_method_handlers('spaceone.api.sample.v1.HelloWorld', rpc_method_handlers)


 # This class is part of an EXPERIMENTAL API.
class HelloWorld(object):
    """The greeting service definition.
    """

    @staticmethod
    def say_hello(request,
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
            '/spaceone.api.sample.v1.HelloWorld/say_hello',
            spaceone_dot_api_dot_sample_dot_v1_dot_helloworld__pb2.HelloRequest.SerializeToString,
            spaceone_dot_api_dot_sample_dot_v1_dot_helloworld__pb2.HelloReply.FromString,
            options,
            channel_credentials,
            insecure,
            call_credentials,
            compression,
            wait_for_ready,
            timeout,
            metadata,
            _registered_method=True)
