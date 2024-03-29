# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
"""Client and server classes corresponding to protobuf-defined services."""
import grpc

from google.protobuf import empty_pb2 as google_dot_protobuf_dot_empty__pb2
from google.protobuf import struct_pb2 as google_dot_protobuf_dot_struct__pb2
from spaceone.api.inventory.v1 import cloud_service_report_pb2 as spaceone_dot_api_dot_inventory_dot_v1_dot_cloud__service__report__pb2


class CloudServiceReportStub(object):
    """Missing associated documentation comment in .proto file."""

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.create = channel.unary_unary(
                '/spaceone.api.inventory.v1.CloudServiceReport/create',
                request_serializer=spaceone_dot_api_dot_inventory_dot_v1_dot_cloud__service__report__pb2.CreateCloudServiceReportRequest.SerializeToString,
                response_deserializer=spaceone_dot_api_dot_inventory_dot_v1_dot_cloud__service__report__pb2.CloudServiceReportInfo.FromString,
                )
        self.update = channel.unary_unary(
                '/spaceone.api.inventory.v1.CloudServiceReport/update',
                request_serializer=spaceone_dot_api_dot_inventory_dot_v1_dot_cloud__service__report__pb2.UpdateCloudServiceReportRequest.SerializeToString,
                response_deserializer=spaceone_dot_api_dot_inventory_dot_v1_dot_cloud__service__report__pb2.CloudServiceReportInfo.FromString,
                )
        self.delete = channel.unary_unary(
                '/spaceone.api.inventory.v1.CloudServiceReport/delete',
                request_serializer=spaceone_dot_api_dot_inventory_dot_v1_dot_cloud__service__report__pb2.CloudServiceReportRequest.SerializeToString,
                response_deserializer=google_dot_protobuf_dot_empty__pb2.Empty.FromString,
                )
        self.send = channel.unary_unary(
                '/spaceone.api.inventory.v1.CloudServiceReport/send',
                request_serializer=spaceone_dot_api_dot_inventory_dot_v1_dot_cloud__service__report__pb2.CloudServiceReportRequest.SerializeToString,
                response_deserializer=google_dot_protobuf_dot_empty__pb2.Empty.FromString,
                )
        self.get = channel.unary_unary(
                '/spaceone.api.inventory.v1.CloudServiceReport/get',
                request_serializer=spaceone_dot_api_dot_inventory_dot_v1_dot_cloud__service__report__pb2.CloudServiceReportRequest.SerializeToString,
                response_deserializer=spaceone_dot_api_dot_inventory_dot_v1_dot_cloud__service__report__pb2.CloudServiceReportInfo.FromString,
                )
        self.list = channel.unary_unary(
                '/spaceone.api.inventory.v1.CloudServiceReport/list',
                request_serializer=spaceone_dot_api_dot_inventory_dot_v1_dot_cloud__service__report__pb2.CloudServiceReportQuery.SerializeToString,
                response_deserializer=spaceone_dot_api_dot_inventory_dot_v1_dot_cloud__service__report__pb2.CloudServiceReportsInfo.FromString,
                )
        self.stat = channel.unary_unary(
                '/spaceone.api.inventory.v1.CloudServiceReport/stat',
                request_serializer=spaceone_dot_api_dot_inventory_dot_v1_dot_cloud__service__report__pb2.CloudServiceReportStatQuery.SerializeToString,
                response_deserializer=google_dot_protobuf_dot_struct__pb2.Struct.FromString,
                )


class CloudServiceReportServicer(object):
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

    def send(self, request, context):
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


def add_CloudServiceReportServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'create': grpc.unary_unary_rpc_method_handler(
                    servicer.create,
                    request_deserializer=spaceone_dot_api_dot_inventory_dot_v1_dot_cloud__service__report__pb2.CreateCloudServiceReportRequest.FromString,
                    response_serializer=spaceone_dot_api_dot_inventory_dot_v1_dot_cloud__service__report__pb2.CloudServiceReportInfo.SerializeToString,
            ),
            'update': grpc.unary_unary_rpc_method_handler(
                    servicer.update,
                    request_deserializer=spaceone_dot_api_dot_inventory_dot_v1_dot_cloud__service__report__pb2.UpdateCloudServiceReportRequest.FromString,
                    response_serializer=spaceone_dot_api_dot_inventory_dot_v1_dot_cloud__service__report__pb2.CloudServiceReportInfo.SerializeToString,
            ),
            'delete': grpc.unary_unary_rpc_method_handler(
                    servicer.delete,
                    request_deserializer=spaceone_dot_api_dot_inventory_dot_v1_dot_cloud__service__report__pb2.CloudServiceReportRequest.FromString,
                    response_serializer=google_dot_protobuf_dot_empty__pb2.Empty.SerializeToString,
            ),
            'send': grpc.unary_unary_rpc_method_handler(
                    servicer.send,
                    request_deserializer=spaceone_dot_api_dot_inventory_dot_v1_dot_cloud__service__report__pb2.CloudServiceReportRequest.FromString,
                    response_serializer=google_dot_protobuf_dot_empty__pb2.Empty.SerializeToString,
            ),
            'get': grpc.unary_unary_rpc_method_handler(
                    servicer.get,
                    request_deserializer=spaceone_dot_api_dot_inventory_dot_v1_dot_cloud__service__report__pb2.CloudServiceReportRequest.FromString,
                    response_serializer=spaceone_dot_api_dot_inventory_dot_v1_dot_cloud__service__report__pb2.CloudServiceReportInfo.SerializeToString,
            ),
            'list': grpc.unary_unary_rpc_method_handler(
                    servicer.list,
                    request_deserializer=spaceone_dot_api_dot_inventory_dot_v1_dot_cloud__service__report__pb2.CloudServiceReportQuery.FromString,
                    response_serializer=spaceone_dot_api_dot_inventory_dot_v1_dot_cloud__service__report__pb2.CloudServiceReportsInfo.SerializeToString,
            ),
            'stat': grpc.unary_unary_rpc_method_handler(
                    servicer.stat,
                    request_deserializer=spaceone_dot_api_dot_inventory_dot_v1_dot_cloud__service__report__pb2.CloudServiceReportStatQuery.FromString,
                    response_serializer=google_dot_protobuf_dot_struct__pb2.Struct.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'spaceone.api.inventory.v1.CloudServiceReport', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))


 # This class is part of an EXPERIMENTAL API.
class CloudServiceReport(object):
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
        return grpc.experimental.unary_unary(request, target, '/spaceone.api.inventory.v1.CloudServiceReport/create',
            spaceone_dot_api_dot_inventory_dot_v1_dot_cloud__service__report__pb2.CreateCloudServiceReportRequest.SerializeToString,
            spaceone_dot_api_dot_inventory_dot_v1_dot_cloud__service__report__pb2.CloudServiceReportInfo.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

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
        return grpc.experimental.unary_unary(request, target, '/spaceone.api.inventory.v1.CloudServiceReport/update',
            spaceone_dot_api_dot_inventory_dot_v1_dot_cloud__service__report__pb2.UpdateCloudServiceReportRequest.SerializeToString,
            spaceone_dot_api_dot_inventory_dot_v1_dot_cloud__service__report__pb2.CloudServiceReportInfo.FromString,
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
        return grpc.experimental.unary_unary(request, target, '/spaceone.api.inventory.v1.CloudServiceReport/delete',
            spaceone_dot_api_dot_inventory_dot_v1_dot_cloud__service__report__pb2.CloudServiceReportRequest.SerializeToString,
            google_dot_protobuf_dot_empty__pb2.Empty.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def send(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/spaceone.api.inventory.v1.CloudServiceReport/send',
            spaceone_dot_api_dot_inventory_dot_v1_dot_cloud__service__report__pb2.CloudServiceReportRequest.SerializeToString,
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
        return grpc.experimental.unary_unary(request, target, '/spaceone.api.inventory.v1.CloudServiceReport/get',
            spaceone_dot_api_dot_inventory_dot_v1_dot_cloud__service__report__pb2.CloudServiceReportRequest.SerializeToString,
            spaceone_dot_api_dot_inventory_dot_v1_dot_cloud__service__report__pb2.CloudServiceReportInfo.FromString,
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
        return grpc.experimental.unary_unary(request, target, '/spaceone.api.inventory.v1.CloudServiceReport/list',
            spaceone_dot_api_dot_inventory_dot_v1_dot_cloud__service__report__pb2.CloudServiceReportQuery.SerializeToString,
            spaceone_dot_api_dot_inventory_dot_v1_dot_cloud__service__report__pb2.CloudServiceReportsInfo.FromString,
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
        return grpc.experimental.unary_unary(request, target, '/spaceone.api.inventory.v1.CloudServiceReport/stat',
            spaceone_dot_api_dot_inventory_dot_v1_dot_cloud__service__report__pb2.CloudServiceReportStatQuery.SerializeToString,
            google_dot_protobuf_dot_struct__pb2.Struct.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)
