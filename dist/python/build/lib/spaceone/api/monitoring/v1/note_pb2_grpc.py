# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
"""Client and server classes corresponding to protobuf-defined services."""
import grpc
import warnings

from google.protobuf import empty_pb2 as google_dot_protobuf_dot_empty__pb2
from google.protobuf import struct_pb2 as google_dot_protobuf_dot_struct__pb2
from spaceone.api.monitoring.v1 import note_pb2 as spaceone_dot_api_dot_monitoring_dot_v1_dot_note__pb2

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
        + f' but the generated code in spaceone/api/monitoring/v1/note_pb2_grpc.py depends on'
        + f' grpcio>={GRPC_GENERATED_VERSION}.'
        + f' Please upgrade your grpc module to grpcio>={GRPC_GENERATED_VERSION}'
        + f' or downgrade your generated code using grpcio-tools<={GRPC_VERSION}.'
        + f' This warning will become an error in {EXPECTED_ERROR_RELEASE},'
        + f' scheduled for release on {SCHEDULED_RELEASE_DATE}.',
        RuntimeWarning
    )


class NoteStub(object):
    """Missing associated documentation comment in .proto file."""

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.create = channel.unary_unary(
                '/spaceone.api.monitoring.v1.Note/create',
                request_serializer=spaceone_dot_api_dot_monitoring_dot_v1_dot_note__pb2.CreateNoteRequest.SerializeToString,
                response_deserializer=spaceone_dot_api_dot_monitoring_dot_v1_dot_note__pb2.NoteInfo.FromString,
                _registered_method=True)
        self.update = channel.unary_unary(
                '/spaceone.api.monitoring.v1.Note/update',
                request_serializer=spaceone_dot_api_dot_monitoring_dot_v1_dot_note__pb2.UpdateNoteRequest.SerializeToString,
                response_deserializer=spaceone_dot_api_dot_monitoring_dot_v1_dot_note__pb2.NoteInfo.FromString,
                _registered_method=True)
        self.delete = channel.unary_unary(
                '/spaceone.api.monitoring.v1.Note/delete',
                request_serializer=spaceone_dot_api_dot_monitoring_dot_v1_dot_note__pb2.NoteRequest.SerializeToString,
                response_deserializer=google_dot_protobuf_dot_empty__pb2.Empty.FromString,
                _registered_method=True)
        self.get = channel.unary_unary(
                '/spaceone.api.monitoring.v1.Note/get',
                request_serializer=spaceone_dot_api_dot_monitoring_dot_v1_dot_note__pb2.NoteRequest.SerializeToString,
                response_deserializer=spaceone_dot_api_dot_monitoring_dot_v1_dot_note__pb2.NoteInfo.FromString,
                _registered_method=True)
        self.list = channel.unary_unary(
                '/spaceone.api.monitoring.v1.Note/list',
                request_serializer=spaceone_dot_api_dot_monitoring_dot_v1_dot_note__pb2.NoteQuery.SerializeToString,
                response_deserializer=spaceone_dot_api_dot_monitoring_dot_v1_dot_note__pb2.NotesInfo.FromString,
                _registered_method=True)
        self.stat = channel.unary_unary(
                '/spaceone.api.monitoring.v1.Note/stat',
                request_serializer=spaceone_dot_api_dot_monitoring_dot_v1_dot_note__pb2.NoteStatQuery.SerializeToString,
                response_deserializer=google_dot_protobuf_dot_struct__pb2.Struct.FromString,
                _registered_method=True)


class NoteServicer(object):
    """Missing associated documentation comment in .proto file."""

    def create(self, request, context):
        """Creates a new Note. You can create Notes for each Alert to record the information needed to manage the Alerts.
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def update(self, request, context):
        """Updates a specific Note. You must specify the `note_id` for Note validation check. If the Note exists, it is updated.
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def delete(self, request, context):
        """Deletes a specific Note. You must specify the `note_id` of the Note to delete.
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def get(self, request, context):
        """Gets a specific Note. You must specify the `note_id` and `domain_id`.
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def list(self, request, context):
        """Gets a list of all Notes. You can use a query to get a filtered list of Notes.
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def stat(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_NoteServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'create': grpc.unary_unary_rpc_method_handler(
                    servicer.create,
                    request_deserializer=spaceone_dot_api_dot_monitoring_dot_v1_dot_note__pb2.CreateNoteRequest.FromString,
                    response_serializer=spaceone_dot_api_dot_monitoring_dot_v1_dot_note__pb2.NoteInfo.SerializeToString,
            ),
            'update': grpc.unary_unary_rpc_method_handler(
                    servicer.update,
                    request_deserializer=spaceone_dot_api_dot_monitoring_dot_v1_dot_note__pb2.UpdateNoteRequest.FromString,
                    response_serializer=spaceone_dot_api_dot_monitoring_dot_v1_dot_note__pb2.NoteInfo.SerializeToString,
            ),
            'delete': grpc.unary_unary_rpc_method_handler(
                    servicer.delete,
                    request_deserializer=spaceone_dot_api_dot_monitoring_dot_v1_dot_note__pb2.NoteRequest.FromString,
                    response_serializer=google_dot_protobuf_dot_empty__pb2.Empty.SerializeToString,
            ),
            'get': grpc.unary_unary_rpc_method_handler(
                    servicer.get,
                    request_deserializer=spaceone_dot_api_dot_monitoring_dot_v1_dot_note__pb2.NoteRequest.FromString,
                    response_serializer=spaceone_dot_api_dot_monitoring_dot_v1_dot_note__pb2.NoteInfo.SerializeToString,
            ),
            'list': grpc.unary_unary_rpc_method_handler(
                    servicer.list,
                    request_deserializer=spaceone_dot_api_dot_monitoring_dot_v1_dot_note__pb2.NoteQuery.FromString,
                    response_serializer=spaceone_dot_api_dot_monitoring_dot_v1_dot_note__pb2.NotesInfo.SerializeToString,
            ),
            'stat': grpc.unary_unary_rpc_method_handler(
                    servicer.stat,
                    request_deserializer=spaceone_dot_api_dot_monitoring_dot_v1_dot_note__pb2.NoteStatQuery.FromString,
                    response_serializer=google_dot_protobuf_dot_struct__pb2.Struct.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'spaceone.api.monitoring.v1.Note', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))
    server.add_registered_method_handlers('spaceone.api.monitoring.v1.Note', rpc_method_handlers)


 # This class is part of an EXPERIMENTAL API.
class Note(object):
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
            '/spaceone.api.monitoring.v1.Note/create',
            spaceone_dot_api_dot_monitoring_dot_v1_dot_note__pb2.CreateNoteRequest.SerializeToString,
            spaceone_dot_api_dot_monitoring_dot_v1_dot_note__pb2.NoteInfo.FromString,
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
            '/spaceone.api.monitoring.v1.Note/update',
            spaceone_dot_api_dot_monitoring_dot_v1_dot_note__pb2.UpdateNoteRequest.SerializeToString,
            spaceone_dot_api_dot_monitoring_dot_v1_dot_note__pb2.NoteInfo.FromString,
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
            '/spaceone.api.monitoring.v1.Note/delete',
            spaceone_dot_api_dot_monitoring_dot_v1_dot_note__pb2.NoteRequest.SerializeToString,
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
            '/spaceone.api.monitoring.v1.Note/get',
            spaceone_dot_api_dot_monitoring_dot_v1_dot_note__pb2.NoteRequest.SerializeToString,
            spaceone_dot_api_dot_monitoring_dot_v1_dot_note__pb2.NoteInfo.FromString,
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
            '/spaceone.api.monitoring.v1.Note/list',
            spaceone_dot_api_dot_monitoring_dot_v1_dot_note__pb2.NoteQuery.SerializeToString,
            spaceone_dot_api_dot_monitoring_dot_v1_dot_note__pb2.NotesInfo.FromString,
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
            '/spaceone.api.monitoring.v1.Note/stat',
            spaceone_dot_api_dot_monitoring_dot_v1_dot_note__pb2.NoteStatQuery.SerializeToString,
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
