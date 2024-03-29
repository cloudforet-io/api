# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
"""Client and server classes corresponding to protobuf-defined services."""
import grpc

from google.protobuf import struct_pb2 as google_dot_protobuf_dot_struct__pb2
from spaceone.api.statistics.v1 import resource_pb2 as spaceone_dot_api_dot_statistics_dot_v1_dot_resource__pb2


class ResourceStub(object):
    """Missing associated documentation comment in .proto file."""

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.stat = channel.unary_unary(
                '/spaceone.api.statistics.v1.Resource/stat',
                request_serializer=spaceone_dot_api_dot_statistics_dot_v1_dot_resource__pb2.ResourceStatRequest.SerializeToString,
                response_deserializer=google_dot_protobuf_dot_struct__pb2.Struct.FromString,
                )


class ResourceServicer(object):
    """Missing associated documentation comment in .proto file."""

    def stat(self, request, context):
        """Enables data preprocessing of different services. Although limited, it is possible to create not only basic queries but also data suitable for users' needs, such as joins between two tables created by the query, handling missing values, and sorting.
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_ResourceServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'stat': grpc.unary_unary_rpc_method_handler(
                    servicer.stat,
                    request_deserializer=spaceone_dot_api_dot_statistics_dot_v1_dot_resource__pb2.ResourceStatRequest.FromString,
                    response_serializer=google_dot_protobuf_dot_struct__pb2.Struct.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'spaceone.api.statistics.v1.Resource', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))


 # This class is part of an EXPERIMENTAL API.
class Resource(object):
    """Missing associated documentation comment in .proto file."""

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
        return grpc.experimental.unary_unary(request, target, '/spaceone.api.statistics.v1.Resource/stat',
            spaceone_dot_api_dot_statistics_dot_v1_dot_resource__pb2.ResourceStatRequest.SerializeToString,
            google_dot_protobuf_dot_struct__pb2.Struct.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)
