# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
"""Client and server classes corresponding to protobuf-defined services."""
import grpc
import warnings

from google.protobuf import empty_pb2 as google_dot_protobuf_dot_empty__pb2
from google.protobuf import struct_pb2 as google_dot_protobuf_dot_struct__pb2
from spaceone.api.cost_analysis.v1 import data_source_pb2 as spaceone_dot_api_dot_cost__analysis_dot_v1_dot_data__source__pb2
from spaceone.api.cost_analysis.v1 import job_pb2 as spaceone_dot_api_dot_cost__analysis_dot_v1_dot_job__pb2

GRPC_GENERATED_VERSION = '1.64.1'
GRPC_VERSION = grpc.__version__
EXPECTED_ERROR_RELEASE = '1.65.0'
SCHEDULED_RELEASE_DATE = 'June 25, 2024'
_version_not_supported = False

try:
    from grpc._utilities import first_version_is_lower
    _version_not_supported = first_version_is_lower(GRPC_VERSION, GRPC_GENERATED_VERSION)
except ImportError:
    _version_not_supported = True

if _version_not_supported:
    warnings.warn(
        f'The grpc package installed is at version {GRPC_VERSION},'
        + f' but the generated code in spaceone/api/cost_analysis/v1/data_source_pb2_grpc.py depends on'
        + f' grpcio>={GRPC_GENERATED_VERSION}.'
        + f' Please upgrade your grpc module to grpcio>={GRPC_GENERATED_VERSION}'
        + f' or downgrade your generated code using grpcio-tools<={GRPC_VERSION}.'
        + f' This warning will become an error in {EXPECTED_ERROR_RELEASE},'
        + f' scheduled for release on {SCHEDULED_RELEASE_DATE}.',
        RuntimeWarning
    )


class DataSourceStub(object):
    """Missing associated documentation comment in .proto file."""

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.register = channel.unary_unary(
                '/spaceone.api.cost_analysis.v1.DataSource/register',
                request_serializer=spaceone_dot_api_dot_cost__analysis_dot_v1_dot_data__source__pb2.RegisterDataSourceRequest.SerializeToString,
                response_deserializer=spaceone_dot_api_dot_cost__analysis_dot_v1_dot_data__source__pb2.DataSourceInfo.FromString,
                _registered_method=True)
        self.update = channel.unary_unary(
                '/spaceone.api.cost_analysis.v1.DataSource/update',
                request_serializer=spaceone_dot_api_dot_cost__analysis_dot_v1_dot_data__source__pb2.UpdateDataSourceRequest.SerializeToString,
                response_deserializer=spaceone_dot_api_dot_cost__analysis_dot_v1_dot_data__source__pb2.DataSourceInfo.FromString,
                _registered_method=True)
        self.update_permissions = channel.unary_unary(
                '/spaceone.api.cost_analysis.v1.DataSource/update_permissions',
                request_serializer=spaceone_dot_api_dot_cost__analysis_dot_v1_dot_data__source__pb2.UpdateDataSourcePermissionsRequest.SerializeToString,
                response_deserializer=spaceone_dot_api_dot_cost__analysis_dot_v1_dot_data__source__pb2.DataSourceInfo.FromString,
                _registered_method=True)
        self.update_plugin = channel.unary_unary(
                '/spaceone.api.cost_analysis.v1.DataSource/update_plugin',
                request_serializer=spaceone_dot_api_dot_cost__analysis_dot_v1_dot_data__source__pb2.UpdateDataSourcePluginRequest.SerializeToString,
                response_deserializer=spaceone_dot_api_dot_cost__analysis_dot_v1_dot_data__source__pb2.DataSourceInfo.FromString,
                _registered_method=True)
        self.update_secret_data = channel.unary_unary(
                '/spaceone.api.cost_analysis.v1.DataSource/update_secret_data',
                request_serializer=spaceone_dot_api_dot_cost__analysis_dot_v1_dot_data__source__pb2.UpdateSecretDataSourceRequest.SerializeToString,
                response_deserializer=spaceone_dot_api_dot_cost__analysis_dot_v1_dot_data__source__pb2.DataSourceInfo.FromString,
                _registered_method=True)
        self.verify_plugin = channel.unary_unary(
                '/spaceone.api.cost_analysis.v1.DataSource/verify_plugin',
                request_serializer=spaceone_dot_api_dot_cost__analysis_dot_v1_dot_data__source__pb2.DataSourceRequest.SerializeToString,
                response_deserializer=google_dot_protobuf_dot_empty__pb2.Empty.FromString,
                _registered_method=True)
        self.deregister = channel.unary_unary(
                '/spaceone.api.cost_analysis.v1.DataSource/deregister',
                request_serializer=spaceone_dot_api_dot_cost__analysis_dot_v1_dot_data__source__pb2.DeregisterDataSourceRequest.SerializeToString,
                response_deserializer=google_dot_protobuf_dot_empty__pb2.Empty.FromString,
                _registered_method=True)
        self.sync = channel.unary_unary(
                '/spaceone.api.cost_analysis.v1.DataSource/sync',
                request_serializer=spaceone_dot_api_dot_cost__analysis_dot_v1_dot_data__source__pb2.SyncDataSourceRequest.SerializeToString,
                response_deserializer=spaceone_dot_api_dot_cost__analysis_dot_v1_dot_job__pb2.JobInfo.FromString,
                _registered_method=True)
        self.get = channel.unary_unary(
                '/spaceone.api.cost_analysis.v1.DataSource/get',
                request_serializer=spaceone_dot_api_dot_cost__analysis_dot_v1_dot_data__source__pb2.DataSourceRequest.SerializeToString,
                response_deserializer=spaceone_dot_api_dot_cost__analysis_dot_v1_dot_data__source__pb2.DataSourceInfo.FromString,
                _registered_method=True)
        self.list = channel.unary_unary(
                '/spaceone.api.cost_analysis.v1.DataSource/list',
                request_serializer=spaceone_dot_api_dot_cost__analysis_dot_v1_dot_data__source__pb2.DataSourceQuery.SerializeToString,
                response_deserializer=spaceone_dot_api_dot_cost__analysis_dot_v1_dot_data__source__pb2.DataSourcesInfo.FromString,
                _registered_method=True)
        self.stat = channel.unary_unary(
                '/spaceone.api.cost_analysis.v1.DataSource/stat',
                request_serializer=spaceone_dot_api_dot_cost__analysis_dot_v1_dot_data__source__pb2.DataSourceStatQuery.SerializeToString,
                response_deserializer=google_dot_protobuf_dot_struct__pb2.Struct.FromString,
                _registered_method=True)


class DataSourceServicer(object):
    """Missing associated documentation comment in .proto file."""

    def register(self, request, context):
        """Registers a DataSource with information of the plugin to use. Information of the plugin includes `version`, `provider`, and `upgrade_mode`.
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def update(self, request, context):
        """Updates a specific DataSource. You can make changes in DataSource settings, including `name` and `tags`.
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def update_permissions(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def update_plugin(self, request, context):
        """Updates the plugin of a specific DataSource. This method resets the plugin data in the DataSource to update the `metadata`.
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def update_secret_data(self, request, context):
        """Updates the secret data of plugin for DataSource. This method updates the secret data in the DataSource to update the `secret_data`.
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def verify_plugin(self, request, context):
        """Verifies the plugin of a specific DataSource. This method validates the plugin data, `version` and `endpoint`.
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def deregister(self, request, context):
        """Deregisters and deletes a specific DataSource. You must specify the `data_source_id` of the DataSource to deregister.
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def sync(self, request, context):
        """Manually runs a specific DataSource to collect the cost data. This method is to get up-to-date cost data.
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def get(self, request, context):
        """Gets a specific DataSource. Prints detailed information about the DataSource, including `name`, `plugin_info`, and `schedule`.
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def list(self, request, context):
        """Gets a list of all DataSources. You can use a query to get a filtered list of DataSources.
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def stat(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_DataSourceServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'register': grpc.unary_unary_rpc_method_handler(
                    servicer.register,
                    request_deserializer=spaceone_dot_api_dot_cost__analysis_dot_v1_dot_data__source__pb2.RegisterDataSourceRequest.FromString,
                    response_serializer=spaceone_dot_api_dot_cost__analysis_dot_v1_dot_data__source__pb2.DataSourceInfo.SerializeToString,
            ),
            'update': grpc.unary_unary_rpc_method_handler(
                    servicer.update,
                    request_deserializer=spaceone_dot_api_dot_cost__analysis_dot_v1_dot_data__source__pb2.UpdateDataSourceRequest.FromString,
                    response_serializer=spaceone_dot_api_dot_cost__analysis_dot_v1_dot_data__source__pb2.DataSourceInfo.SerializeToString,
            ),
            'update_permissions': grpc.unary_unary_rpc_method_handler(
                    servicer.update_permissions,
                    request_deserializer=spaceone_dot_api_dot_cost__analysis_dot_v1_dot_data__source__pb2.UpdateDataSourcePermissionsRequest.FromString,
                    response_serializer=spaceone_dot_api_dot_cost__analysis_dot_v1_dot_data__source__pb2.DataSourceInfo.SerializeToString,
            ),
            'update_plugin': grpc.unary_unary_rpc_method_handler(
                    servicer.update_plugin,
                    request_deserializer=spaceone_dot_api_dot_cost__analysis_dot_v1_dot_data__source__pb2.UpdateDataSourcePluginRequest.FromString,
                    response_serializer=spaceone_dot_api_dot_cost__analysis_dot_v1_dot_data__source__pb2.DataSourceInfo.SerializeToString,
            ),
            'update_secret_data': grpc.unary_unary_rpc_method_handler(
                    servicer.update_secret_data,
                    request_deserializer=spaceone_dot_api_dot_cost__analysis_dot_v1_dot_data__source__pb2.UpdateSecretDataSourceRequest.FromString,
                    response_serializer=spaceone_dot_api_dot_cost__analysis_dot_v1_dot_data__source__pb2.DataSourceInfo.SerializeToString,
            ),
            'verify_plugin': grpc.unary_unary_rpc_method_handler(
                    servicer.verify_plugin,
                    request_deserializer=spaceone_dot_api_dot_cost__analysis_dot_v1_dot_data__source__pb2.DataSourceRequest.FromString,
                    response_serializer=google_dot_protobuf_dot_empty__pb2.Empty.SerializeToString,
            ),
            'deregister': grpc.unary_unary_rpc_method_handler(
                    servicer.deregister,
                    request_deserializer=spaceone_dot_api_dot_cost__analysis_dot_v1_dot_data__source__pb2.DeregisterDataSourceRequest.FromString,
                    response_serializer=google_dot_protobuf_dot_empty__pb2.Empty.SerializeToString,
            ),
            'sync': grpc.unary_unary_rpc_method_handler(
                    servicer.sync,
                    request_deserializer=spaceone_dot_api_dot_cost__analysis_dot_v1_dot_data__source__pb2.SyncDataSourceRequest.FromString,
                    response_serializer=spaceone_dot_api_dot_cost__analysis_dot_v1_dot_job__pb2.JobInfo.SerializeToString,
            ),
            'get': grpc.unary_unary_rpc_method_handler(
                    servicer.get,
                    request_deserializer=spaceone_dot_api_dot_cost__analysis_dot_v1_dot_data__source__pb2.DataSourceRequest.FromString,
                    response_serializer=spaceone_dot_api_dot_cost__analysis_dot_v1_dot_data__source__pb2.DataSourceInfo.SerializeToString,
            ),
            'list': grpc.unary_unary_rpc_method_handler(
                    servicer.list,
                    request_deserializer=spaceone_dot_api_dot_cost__analysis_dot_v1_dot_data__source__pb2.DataSourceQuery.FromString,
                    response_serializer=spaceone_dot_api_dot_cost__analysis_dot_v1_dot_data__source__pb2.DataSourcesInfo.SerializeToString,
            ),
            'stat': grpc.unary_unary_rpc_method_handler(
                    servicer.stat,
                    request_deserializer=spaceone_dot_api_dot_cost__analysis_dot_v1_dot_data__source__pb2.DataSourceStatQuery.FromString,
                    response_serializer=google_dot_protobuf_dot_struct__pb2.Struct.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'spaceone.api.cost_analysis.v1.DataSource', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))
    server.add_registered_method_handlers('spaceone.api.cost_analysis.v1.DataSource', rpc_method_handlers)


 # This class is part of an EXPERIMENTAL API.
class DataSource(object):
    """Missing associated documentation comment in .proto file."""

    @staticmethod
    def register(request,
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
            '/spaceone.api.cost_analysis.v1.DataSource/register',
            spaceone_dot_api_dot_cost__analysis_dot_v1_dot_data__source__pb2.RegisterDataSourceRequest.SerializeToString,
            spaceone_dot_api_dot_cost__analysis_dot_v1_dot_data__source__pb2.DataSourceInfo.FromString,
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
            '/spaceone.api.cost_analysis.v1.DataSource/update',
            spaceone_dot_api_dot_cost__analysis_dot_v1_dot_data__source__pb2.UpdateDataSourceRequest.SerializeToString,
            spaceone_dot_api_dot_cost__analysis_dot_v1_dot_data__source__pb2.DataSourceInfo.FromString,
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
    def update_permissions(request,
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
            '/spaceone.api.cost_analysis.v1.DataSource/update_permissions',
            spaceone_dot_api_dot_cost__analysis_dot_v1_dot_data__source__pb2.UpdateDataSourcePermissionsRequest.SerializeToString,
            spaceone_dot_api_dot_cost__analysis_dot_v1_dot_data__source__pb2.DataSourceInfo.FromString,
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
            '/spaceone.api.cost_analysis.v1.DataSource/update_plugin',
            spaceone_dot_api_dot_cost__analysis_dot_v1_dot_data__source__pb2.UpdateDataSourcePluginRequest.SerializeToString,
            spaceone_dot_api_dot_cost__analysis_dot_v1_dot_data__source__pb2.DataSourceInfo.FromString,
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
    def update_secret_data(request,
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
            '/spaceone.api.cost_analysis.v1.DataSource/update_secret_data',
            spaceone_dot_api_dot_cost__analysis_dot_v1_dot_data__source__pb2.UpdateSecretDataSourceRequest.SerializeToString,
            spaceone_dot_api_dot_cost__analysis_dot_v1_dot_data__source__pb2.DataSourceInfo.FromString,
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
            '/spaceone.api.cost_analysis.v1.DataSource/verify_plugin',
            spaceone_dot_api_dot_cost__analysis_dot_v1_dot_data__source__pb2.DataSourceRequest.SerializeToString,
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
    def deregister(request,
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
            '/spaceone.api.cost_analysis.v1.DataSource/deregister',
            spaceone_dot_api_dot_cost__analysis_dot_v1_dot_data__source__pb2.DeregisterDataSourceRequest.SerializeToString,
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
    def sync(request,
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
            '/spaceone.api.cost_analysis.v1.DataSource/sync',
            spaceone_dot_api_dot_cost__analysis_dot_v1_dot_data__source__pb2.SyncDataSourceRequest.SerializeToString,
            spaceone_dot_api_dot_cost__analysis_dot_v1_dot_job__pb2.JobInfo.FromString,
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
            '/spaceone.api.cost_analysis.v1.DataSource/get',
            spaceone_dot_api_dot_cost__analysis_dot_v1_dot_data__source__pb2.DataSourceRequest.SerializeToString,
            spaceone_dot_api_dot_cost__analysis_dot_v1_dot_data__source__pb2.DataSourceInfo.FromString,
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
            '/spaceone.api.cost_analysis.v1.DataSource/list',
            spaceone_dot_api_dot_cost__analysis_dot_v1_dot_data__source__pb2.DataSourceQuery.SerializeToString,
            spaceone_dot_api_dot_cost__analysis_dot_v1_dot_data__source__pb2.DataSourcesInfo.FromString,
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
            '/spaceone.api.cost_analysis.v1.DataSource/stat',
            spaceone_dot_api_dot_cost__analysis_dot_v1_dot_data__source__pb2.DataSourceStatQuery.SerializeToString,
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
