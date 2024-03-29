# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
"""Client and server classes corresponding to protobuf-defined services."""
import grpc

from spaceone.api.repository.v1 import repository_pb2 as spaceone_dot_api_dot_repository_dot_v1_dot_repository__pb2


class RepositoryStub(object):
    """Missing associated documentation comment in .proto file."""

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.list = channel.unary_unary(
                '/spaceone.api.repository.v1.Repository/list',
                request_serializer=spaceone_dot_api_dot_repository_dot_v1_dot_repository__pb2.RepositoryQuery.SerializeToString,
                response_deserializer=spaceone_dot_api_dot_repository_dot_v1_dot_repository__pb2.RepositoriesInfo.FromString,
                )


class RepositoryServicer(object):
    """Missing associated documentation comment in .proto file."""

    def list(self, request, context):
        """Gets a list of all Repositories regardless of `domain`. You can use a query to get a filtered list of Repositories.
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_RepositoryServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'list': grpc.unary_unary_rpc_method_handler(
                    servicer.list,
                    request_deserializer=spaceone_dot_api_dot_repository_dot_v1_dot_repository__pb2.RepositoryQuery.FromString,
                    response_serializer=spaceone_dot_api_dot_repository_dot_v1_dot_repository__pb2.RepositoriesInfo.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'spaceone.api.repository.v1.Repository', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))


 # This class is part of an EXPERIMENTAL API.
class Repository(object):
    """Missing associated documentation comment in .proto file."""

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
        return grpc.experimental.unary_unary(request, target, '/spaceone.api.repository.v1.Repository/list',
            spaceone_dot_api_dot_repository_dot_v1_dot_repository__pb2.RepositoryQuery.SerializeToString,
            spaceone_dot_api_dot_repository_dot_v1_dot_repository__pb2.RepositoriesInfo.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)
