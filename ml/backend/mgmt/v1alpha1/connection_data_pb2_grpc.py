# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
"""Client and server classes corresponding to protobuf-defined services."""
import grpc

from mgmt.v1alpha1 import connection_data_pb2 as mgmt_dot_v1alpha1_dot_connection__data__pb2


class ConnectionDataServiceStub(object):
    """Service for managing connection data.
    This is used in handle data from a connection
    """

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.GetConnectionDataStream = channel.unary_stream(
                '/mgmt.v1alpha1.ConnectionDataService/GetConnectionDataStream',
                request_serializer=mgmt_dot_v1alpha1_dot_connection__data__pb2.GetConnectionDataStreamRequest.SerializeToString,
                response_deserializer=mgmt_dot_v1alpha1_dot_connection__data__pb2.GetConnectionDataStreamResponse.FromString,
                )
        self.GetConnectionSchema = channel.unary_unary(
                '/mgmt.v1alpha1.ConnectionDataService/GetConnectionSchema',
                request_serializer=mgmt_dot_v1alpha1_dot_connection__data__pb2.GetConnectionSchemaRequest.SerializeToString,
                response_deserializer=mgmt_dot_v1alpha1_dot_connection__data__pb2.GetConnectionSchemaResponse.FromString,
                )
        self.GetConnectionForeignConstraints = channel.unary_unary(
                '/mgmt.v1alpha1.ConnectionDataService/GetConnectionForeignConstraints',
                request_serializer=mgmt_dot_v1alpha1_dot_connection__data__pb2.GetConnectionForeignConstraintsRequest.SerializeToString,
                response_deserializer=mgmt_dot_v1alpha1_dot_connection__data__pb2.GetConnectionForeignConstraintsResponse.FromString,
                )
        self.GetConnectionInitStatements = channel.unary_unary(
                '/mgmt.v1alpha1.ConnectionDataService/GetConnectionInitStatements',
                request_serializer=mgmt_dot_v1alpha1_dot_connection__data__pb2.GetConnectionInitStatementsRequest.SerializeToString,
                response_deserializer=mgmt_dot_v1alpha1_dot_connection__data__pb2.GetConnectionInitStatementsResponse.FromString,
                )


class ConnectionDataServiceServicer(object):
    """Service for managing connection data.
    This is used in handle data from a connection
    """

    def GetConnectionDataStream(self, request, context):
        """Streaming endpoint that will stream the data available from the Connection to the client.
        Used primarily by the CLI sync command.
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def GetConnectionSchema(self, request, context):
        """Returns the schema for a specific connection. Used mostly for SQL-based connections
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def GetConnectionForeignConstraints(self, request, context):
        """For a specific connection, returns the foreign key constraints. Mostly useful for SQL-based Connections.
        Used primarily by the CLI sync command to determine stream order.
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def GetConnectionInitStatements(self, request, context):
        """For a specific connection, returns the init table statements. Mostly useful for SQL-based Connections.
        Used primarily by the CLI sync command to create table schema init statement.
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_ConnectionDataServiceServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'GetConnectionDataStream': grpc.unary_stream_rpc_method_handler(
                    servicer.GetConnectionDataStream,
                    request_deserializer=mgmt_dot_v1alpha1_dot_connection__data__pb2.GetConnectionDataStreamRequest.FromString,
                    response_serializer=mgmt_dot_v1alpha1_dot_connection__data__pb2.GetConnectionDataStreamResponse.SerializeToString,
            ),
            'GetConnectionSchema': grpc.unary_unary_rpc_method_handler(
                    servicer.GetConnectionSchema,
                    request_deserializer=mgmt_dot_v1alpha1_dot_connection__data__pb2.GetConnectionSchemaRequest.FromString,
                    response_serializer=mgmt_dot_v1alpha1_dot_connection__data__pb2.GetConnectionSchemaResponse.SerializeToString,
            ),
            'GetConnectionForeignConstraints': grpc.unary_unary_rpc_method_handler(
                    servicer.GetConnectionForeignConstraints,
                    request_deserializer=mgmt_dot_v1alpha1_dot_connection__data__pb2.GetConnectionForeignConstraintsRequest.FromString,
                    response_serializer=mgmt_dot_v1alpha1_dot_connection__data__pb2.GetConnectionForeignConstraintsResponse.SerializeToString,
            ),
            'GetConnectionInitStatements': grpc.unary_unary_rpc_method_handler(
                    servicer.GetConnectionInitStatements,
                    request_deserializer=mgmt_dot_v1alpha1_dot_connection__data__pb2.GetConnectionInitStatementsRequest.FromString,
                    response_serializer=mgmt_dot_v1alpha1_dot_connection__data__pb2.GetConnectionInitStatementsResponse.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'mgmt.v1alpha1.ConnectionDataService', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))


 # This class is part of an EXPERIMENTAL API.
class ConnectionDataService(object):
    """Service for managing connection data.
    This is used in handle data from a connection
    """

    @staticmethod
    def GetConnectionDataStream(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_stream(request, target, '/mgmt.v1alpha1.ConnectionDataService/GetConnectionDataStream',
            mgmt_dot_v1alpha1_dot_connection__data__pb2.GetConnectionDataStreamRequest.SerializeToString,
            mgmt_dot_v1alpha1_dot_connection__data__pb2.GetConnectionDataStreamResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def GetConnectionSchema(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/mgmt.v1alpha1.ConnectionDataService/GetConnectionSchema',
            mgmt_dot_v1alpha1_dot_connection__data__pb2.GetConnectionSchemaRequest.SerializeToString,
            mgmt_dot_v1alpha1_dot_connection__data__pb2.GetConnectionSchemaResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def GetConnectionForeignConstraints(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/mgmt.v1alpha1.ConnectionDataService/GetConnectionForeignConstraints',
            mgmt_dot_v1alpha1_dot_connection__data__pb2.GetConnectionForeignConstraintsRequest.SerializeToString,
            mgmt_dot_v1alpha1_dot_connection__data__pb2.GetConnectionForeignConstraintsResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def GetConnectionInitStatements(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/mgmt.v1alpha1.ConnectionDataService/GetConnectionInitStatements',
            mgmt_dot_v1alpha1_dot_connection__data__pb2.GetConnectionInitStatementsRequest.SerializeToString,
            mgmt_dot_v1alpha1_dot_connection__data__pb2.GetConnectionInitStatementsResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)