# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
"""Client and server classes corresponding to protobuf-defined services."""
import grpc

from google.protobuf import empty_pb2 as google_dot_protobuf_dot_empty__pb2
from google.protobuf import struct_pb2 as google_dot_protobuf_dot_struct__pb2
from spaceone.api.statistics.v1 import history_pb2 as spaceone_dot_api_dot_statistics_dot_v1_dot_history__pb2


class HistoryStub(object):
    """Missing associated documentation comment in .proto file."""

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.create = channel.unary_unary(
                '/spaceone.api.statistics.v1.History/create',
                request_serializer=spaceone_dot_api_dot_statistics_dot_v1_dot_history__pb2.CreateHistoryRequest.SerializeToString,
                response_deserializer=google_dot_protobuf_dot_empty__pb2.Empty.FromString,
                )
        self.list = channel.unary_unary(
                '/spaceone.api.statistics.v1.History/list',
                request_serializer=spaceone_dot_api_dot_statistics_dot_v1_dot_history__pb2.QueryHistoryRequest.SerializeToString,
                response_deserializer=spaceone_dot_api_dot_statistics_dot_v1_dot_history__pb2.HistoryInfo.FromString,
                )
        self.stat = channel.unary_unary(
                '/spaceone.api.statistics.v1.History/stat',
                request_serializer=spaceone_dot_api_dot_statistics_dot_v1_dot_history__pb2.HistoryStatRequest.SerializeToString,
                response_deserializer=google_dot_protobuf_dot_struct__pb2.Struct.FromString,
                )


class HistoryServicer(object):
    """Missing associated documentation comment in .proto file."""

    def create(self, request, context):
        """Creates a new History. Gets a Schedule as an input and creates a History as an output. You can use this method to manually run a specific Schedule.
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def list(self, request, context):
        """Gets a list of all Histories. You can use a query to get a filtered list of Histories.
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def stat(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_HistoryServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'create': grpc.unary_unary_rpc_method_handler(
                    servicer.create,
                    request_deserializer=spaceone_dot_api_dot_statistics_dot_v1_dot_history__pb2.CreateHistoryRequest.FromString,
                    response_serializer=google_dot_protobuf_dot_empty__pb2.Empty.SerializeToString,
            ),
            'list': grpc.unary_unary_rpc_method_handler(
                    servicer.list,
                    request_deserializer=spaceone_dot_api_dot_statistics_dot_v1_dot_history__pb2.QueryHistoryRequest.FromString,
                    response_serializer=spaceone_dot_api_dot_statistics_dot_v1_dot_history__pb2.HistoryInfo.SerializeToString,
            ),
            'stat': grpc.unary_unary_rpc_method_handler(
                    servicer.stat,
                    request_deserializer=spaceone_dot_api_dot_statistics_dot_v1_dot_history__pb2.HistoryStatRequest.FromString,
                    response_serializer=google_dot_protobuf_dot_struct__pb2.Struct.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'spaceone.api.statistics.v1.History', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))


 # This class is part of an EXPERIMENTAL API.
class History(object):
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
        return grpc.experimental.unary_unary(request, target, '/spaceone.api.statistics.v1.History/create',
            spaceone_dot_api_dot_statistics_dot_v1_dot_history__pb2.CreateHistoryRequest.SerializeToString,
            google_dot_protobuf_dot_empty__pb2.Empty.FromString,
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
        return grpc.experimental.unary_unary(request, target, '/spaceone.api.statistics.v1.History/list',
            spaceone_dot_api_dot_statistics_dot_v1_dot_history__pb2.QueryHistoryRequest.SerializeToString,
            spaceone_dot_api_dot_statistics_dot_v1_dot_history__pb2.HistoryInfo.FromString,
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
        return grpc.experimental.unary_unary(request, target, '/spaceone.api.statistics.v1.History/stat',
            spaceone_dot_api_dot_statistics_dot_v1_dot_history__pb2.HistoryStatRequest.SerializeToString,
            google_dot_protobuf_dot_struct__pb2.Struct.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)
