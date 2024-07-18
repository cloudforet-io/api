# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
"""Client and server classes corresponding to protobuf-defined services."""
import grpc
import warnings

from google.protobuf import empty_pb2 as google_dot_protobuf_dot_empty__pb2
from google.protobuf import struct_pb2 as google_dot_protobuf_dot_struct__pb2
from spaceone.api.core.v1 import handler_pb2 as spaceone_dot_api_dot_core_dot_v1_dot_handler__pb2
from spaceone.api.identity.v1 import domain_pb2 as spaceone_dot_api_dot_identity_dot_v1_dot_domain__pb2

GRPC_GENERATED_VERSION = '1.65.1'
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
        + f' but the generated code in spaceone/api/identity/v1/domain_pb2_grpc.py depends on'
        + f' grpcio>={GRPC_GENERATED_VERSION}.'
        + f' Please upgrade your grpc module to grpcio>={GRPC_GENERATED_VERSION}'
        + f' or downgrade your generated code using grpcio-tools<={GRPC_VERSION}.'
        + f' This warning will become an error in {EXPECTED_ERROR_RELEASE},'
        + f' scheduled for release on {SCHEDULED_RELEASE_DATE}.',
        RuntimeWarning
    )


class DomainStub(object):
    """Missing associated documentation comment in .proto file."""

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.create = channel.unary_unary(
                '/spaceone.api.identity.v1.Domain/create',
                request_serializer=spaceone_dot_api_dot_identity_dot_v1_dot_domain__pb2.CreateDomainRequest.SerializeToString,
                response_deserializer=spaceone_dot_api_dot_identity_dot_v1_dot_domain__pb2.DomainInfo.FromString,
                _registered_method=True)
        self.update = channel.unary_unary(
                '/spaceone.api.identity.v1.Domain/update',
                request_serializer=spaceone_dot_api_dot_identity_dot_v1_dot_domain__pb2.UpdateDomainRequest.SerializeToString,
                response_deserializer=spaceone_dot_api_dot_identity_dot_v1_dot_domain__pb2.DomainInfo.FromString,
                _registered_method=True)
        self.change_auth_plugin = channel.unary_unary(
                '/spaceone.api.identity.v1.Domain/change_auth_plugin',
                request_serializer=spaceone_dot_api_dot_identity_dot_v1_dot_domain__pb2.ChangeAuthRequest.SerializeToString,
                response_deserializer=spaceone_dot_api_dot_identity_dot_v1_dot_domain__pb2.DomainInfo.FromString,
                _registered_method=True)
        self.update_plugin = channel.unary_unary(
                '/spaceone.api.identity.v1.Domain/update_plugin',
                request_serializer=spaceone_dot_api_dot_identity_dot_v1_dot_domain__pb2.UpdatePluginRequest.SerializeToString,
                response_deserializer=spaceone_dot_api_dot_identity_dot_v1_dot_domain__pb2.DomainInfo.FromString,
                _registered_method=True)
        self.verify_plugin = channel.unary_unary(
                '/spaceone.api.identity.v1.Domain/verify_plugin',
                request_serializer=spaceone_dot_api_dot_identity_dot_v1_dot_domain__pb2.DomainRequest.SerializeToString,
                response_deserializer=google_dot_protobuf_dot_empty__pb2.Empty.FromString,
                _registered_method=True)
        self.delete = channel.unary_unary(
                '/spaceone.api.identity.v1.Domain/delete',
                request_serializer=spaceone_dot_api_dot_identity_dot_v1_dot_domain__pb2.DomainRequest.SerializeToString,
                response_deserializer=google_dot_protobuf_dot_empty__pb2.Empty.FromString,
                _registered_method=True)
        self.enable = channel.unary_unary(
                '/spaceone.api.identity.v1.Domain/enable',
                request_serializer=spaceone_dot_api_dot_identity_dot_v1_dot_domain__pb2.DomainRequest.SerializeToString,
                response_deserializer=spaceone_dot_api_dot_identity_dot_v1_dot_domain__pb2.DomainInfo.FromString,
                _registered_method=True)
        self.disable = channel.unary_unary(
                '/spaceone.api.identity.v1.Domain/disable',
                request_serializer=spaceone_dot_api_dot_identity_dot_v1_dot_domain__pb2.DomainRequest.SerializeToString,
                response_deserializer=spaceone_dot_api_dot_identity_dot_v1_dot_domain__pb2.DomainInfo.FromString,
                _registered_method=True)
        self.get = channel.unary_unary(
                '/spaceone.api.identity.v1.Domain/get',
                request_serializer=spaceone_dot_api_dot_identity_dot_v1_dot_domain__pb2.GetDomainRequest.SerializeToString,
                response_deserializer=spaceone_dot_api_dot_identity_dot_v1_dot_domain__pb2.DomainInfo.FromString,
                _registered_method=True)
        self.list = channel.unary_unary(
                '/spaceone.api.identity.v1.Domain/list',
                request_serializer=spaceone_dot_api_dot_identity_dot_v1_dot_domain__pb2.DomainQuery.SerializeToString,
                response_deserializer=spaceone_dot_api_dot_identity_dot_v1_dot_domain__pb2.DomainsInfo.FromString,
                _registered_method=True)
        self.stat = channel.unary_unary(
                '/spaceone.api.identity.v1.Domain/stat',
                request_serializer=spaceone_dot_api_dot_identity_dot_v1_dot_domain__pb2.DomainStatQuery.SerializeToString,
                response_deserializer=google_dot_protobuf_dot_struct__pb2.Struct.FromString,
                _registered_method=True)
        self.get_public_key = channel.unary_unary(
                '/spaceone.api.identity.v1.Domain/get_public_key',
                request_serializer=spaceone_dot_api_dot_core_dot_v1_dot_handler__pb2.AuthenticationRequest.SerializeToString,
                response_deserializer=spaceone_dot_api_dot_core_dot_v1_dot_handler__pb2.AuthenticationResponse.FromString,
                _registered_method=True)


class DomainServicer(object):
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

    def change_auth_plugin(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def update_plugin(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def verify_plugin(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def delete(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def enable(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def disable(self, request, context):
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

    def get_public_key(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_DomainServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'create': grpc.unary_unary_rpc_method_handler(
                    servicer.create,
                    request_deserializer=spaceone_dot_api_dot_identity_dot_v1_dot_domain__pb2.CreateDomainRequest.FromString,
                    response_serializer=spaceone_dot_api_dot_identity_dot_v1_dot_domain__pb2.DomainInfo.SerializeToString,
            ),
            'update': grpc.unary_unary_rpc_method_handler(
                    servicer.update,
                    request_deserializer=spaceone_dot_api_dot_identity_dot_v1_dot_domain__pb2.UpdateDomainRequest.FromString,
                    response_serializer=spaceone_dot_api_dot_identity_dot_v1_dot_domain__pb2.DomainInfo.SerializeToString,
            ),
            'change_auth_plugin': grpc.unary_unary_rpc_method_handler(
                    servicer.change_auth_plugin,
                    request_deserializer=spaceone_dot_api_dot_identity_dot_v1_dot_domain__pb2.ChangeAuthRequest.FromString,
                    response_serializer=spaceone_dot_api_dot_identity_dot_v1_dot_domain__pb2.DomainInfo.SerializeToString,
            ),
            'update_plugin': grpc.unary_unary_rpc_method_handler(
                    servicer.update_plugin,
                    request_deserializer=spaceone_dot_api_dot_identity_dot_v1_dot_domain__pb2.UpdatePluginRequest.FromString,
                    response_serializer=spaceone_dot_api_dot_identity_dot_v1_dot_domain__pb2.DomainInfo.SerializeToString,
            ),
            'verify_plugin': grpc.unary_unary_rpc_method_handler(
                    servicer.verify_plugin,
                    request_deserializer=spaceone_dot_api_dot_identity_dot_v1_dot_domain__pb2.DomainRequest.FromString,
                    response_serializer=google_dot_protobuf_dot_empty__pb2.Empty.SerializeToString,
            ),
            'delete': grpc.unary_unary_rpc_method_handler(
                    servicer.delete,
                    request_deserializer=spaceone_dot_api_dot_identity_dot_v1_dot_domain__pb2.DomainRequest.FromString,
                    response_serializer=google_dot_protobuf_dot_empty__pb2.Empty.SerializeToString,
            ),
            'enable': grpc.unary_unary_rpc_method_handler(
                    servicer.enable,
                    request_deserializer=spaceone_dot_api_dot_identity_dot_v1_dot_domain__pb2.DomainRequest.FromString,
                    response_serializer=spaceone_dot_api_dot_identity_dot_v1_dot_domain__pb2.DomainInfo.SerializeToString,
            ),
            'disable': grpc.unary_unary_rpc_method_handler(
                    servicer.disable,
                    request_deserializer=spaceone_dot_api_dot_identity_dot_v1_dot_domain__pb2.DomainRequest.FromString,
                    response_serializer=spaceone_dot_api_dot_identity_dot_v1_dot_domain__pb2.DomainInfo.SerializeToString,
            ),
            'get': grpc.unary_unary_rpc_method_handler(
                    servicer.get,
                    request_deserializer=spaceone_dot_api_dot_identity_dot_v1_dot_domain__pb2.GetDomainRequest.FromString,
                    response_serializer=spaceone_dot_api_dot_identity_dot_v1_dot_domain__pb2.DomainInfo.SerializeToString,
            ),
            'list': grpc.unary_unary_rpc_method_handler(
                    servicer.list,
                    request_deserializer=spaceone_dot_api_dot_identity_dot_v1_dot_domain__pb2.DomainQuery.FromString,
                    response_serializer=spaceone_dot_api_dot_identity_dot_v1_dot_domain__pb2.DomainsInfo.SerializeToString,
            ),
            'stat': grpc.unary_unary_rpc_method_handler(
                    servicer.stat,
                    request_deserializer=spaceone_dot_api_dot_identity_dot_v1_dot_domain__pb2.DomainStatQuery.FromString,
                    response_serializer=google_dot_protobuf_dot_struct__pb2.Struct.SerializeToString,
            ),
            'get_public_key': grpc.unary_unary_rpc_method_handler(
                    servicer.get_public_key,
                    request_deserializer=spaceone_dot_api_dot_core_dot_v1_dot_handler__pb2.AuthenticationRequest.FromString,
                    response_serializer=spaceone_dot_api_dot_core_dot_v1_dot_handler__pb2.AuthenticationResponse.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'spaceone.api.identity.v1.Domain', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))
    server.add_registered_method_handlers('spaceone.api.identity.v1.Domain', rpc_method_handlers)


 # This class is part of an EXPERIMENTAL API.
class Domain(object):
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
            '/spaceone.api.identity.v1.Domain/create',
            spaceone_dot_api_dot_identity_dot_v1_dot_domain__pb2.CreateDomainRequest.SerializeToString,
            spaceone_dot_api_dot_identity_dot_v1_dot_domain__pb2.DomainInfo.FromString,
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
            '/spaceone.api.identity.v1.Domain/update',
            spaceone_dot_api_dot_identity_dot_v1_dot_domain__pb2.UpdateDomainRequest.SerializeToString,
            spaceone_dot_api_dot_identity_dot_v1_dot_domain__pb2.DomainInfo.FromString,
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
    def change_auth_plugin(request,
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
            '/spaceone.api.identity.v1.Domain/change_auth_plugin',
            spaceone_dot_api_dot_identity_dot_v1_dot_domain__pb2.ChangeAuthRequest.SerializeToString,
            spaceone_dot_api_dot_identity_dot_v1_dot_domain__pb2.DomainInfo.FromString,
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
    def update_plugin(request,
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
            '/spaceone.api.identity.v1.Domain/update_plugin',
            spaceone_dot_api_dot_identity_dot_v1_dot_domain__pb2.UpdatePluginRequest.SerializeToString,
            spaceone_dot_api_dot_identity_dot_v1_dot_domain__pb2.DomainInfo.FromString,
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
    def verify_plugin(request,
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
            '/spaceone.api.identity.v1.Domain/verify_plugin',
            spaceone_dot_api_dot_identity_dot_v1_dot_domain__pb2.DomainRequest.SerializeToString,
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
            '/spaceone.api.identity.v1.Domain/delete',
            spaceone_dot_api_dot_identity_dot_v1_dot_domain__pb2.DomainRequest.SerializeToString,
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
    def enable(request,
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
            '/spaceone.api.identity.v1.Domain/enable',
            spaceone_dot_api_dot_identity_dot_v1_dot_domain__pb2.DomainRequest.SerializeToString,
            spaceone_dot_api_dot_identity_dot_v1_dot_domain__pb2.DomainInfo.FromString,
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
    def disable(request,
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
            '/spaceone.api.identity.v1.Domain/disable',
            spaceone_dot_api_dot_identity_dot_v1_dot_domain__pb2.DomainRequest.SerializeToString,
            spaceone_dot_api_dot_identity_dot_v1_dot_domain__pb2.DomainInfo.FromString,
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
            '/spaceone.api.identity.v1.Domain/get',
            spaceone_dot_api_dot_identity_dot_v1_dot_domain__pb2.GetDomainRequest.SerializeToString,
            spaceone_dot_api_dot_identity_dot_v1_dot_domain__pb2.DomainInfo.FromString,
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
            '/spaceone.api.identity.v1.Domain/list',
            spaceone_dot_api_dot_identity_dot_v1_dot_domain__pb2.DomainQuery.SerializeToString,
            spaceone_dot_api_dot_identity_dot_v1_dot_domain__pb2.DomainsInfo.FromString,
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
            '/spaceone.api.identity.v1.Domain/stat',
            spaceone_dot_api_dot_identity_dot_v1_dot_domain__pb2.DomainStatQuery.SerializeToString,
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

    @staticmethod
    def get_public_key(request,
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
            '/spaceone.api.identity.v1.Domain/get_public_key',
            spaceone_dot_api_dot_core_dot_v1_dot_handler__pb2.AuthenticationRequest.SerializeToString,
            spaceone_dot_api_dot_core_dot_v1_dot_handler__pb2.AuthenticationResponse.FromString,
            options,
            channel_credentials,
            insecure,
            call_credentials,
            compression,
            wait_for_ready,
            timeout,
            metadata,
            _registered_method=True)
