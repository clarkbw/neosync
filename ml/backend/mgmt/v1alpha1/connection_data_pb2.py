# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: mgmt/v1alpha1/connection_data.proto
# Protobuf Python Version: 4.25.2
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from buf.validate import validate_pb2 as buf_dot_validate_dot_validate__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n#mgmt/v1alpha1/connection_data.proto\x12\rmgmt.v1alpha1\x1a\x1b\x62uf/validate/validate.proto\"\x16\n\x14PostgresStreamConfig\"\x13\n\x11MysqlStreamConfig\"e\n\x11\x41wsS3StreamConfig\x12!\n\x06job_id\x18\x01 \x01(\tB\x08\xbaH\x05r\x03\xb0\x01\x01H\x00R\x05jobId\x12\'\n\njob_run_id\x18\x02 \x01(\tB\x07\xbaH\x04r\x02\x10\x01H\x00R\x08jobRunIdB\x04\n\x02id\"\xfc\x01\n\x16\x43onnectionStreamConfig\x12\x42\n\tpg_config\x18\x01 \x01(\x0b\x32#.mgmt.v1alpha1.PostgresStreamConfigH\x00R\x08pgConfig\x12\x46\n\raws_s3_config\x18\x02 \x01(\x0b\x32 .mgmt.v1alpha1.AwsS3StreamConfigH\x00R\x0b\x61wsS3Config\x12\x45\n\x0cmysql_config\x18\x03 \x01(\x0b\x32 .mgmt.v1alpha1.MysqlStreamConfigH\x00R\x0bmysqlConfigB\x0f\n\x06\x63onfig\x12\x05\xbaH\x02\x08\x01\"\xc9\x01\n\x1eGetConnectionDataStreamRequest\x12-\n\rconnection_id\x18\x01 \x01(\tB\x08\xbaH\x05r\x03\xb0\x01\x01R\x0c\x63onnectionId\x12J\n\rstream_config\x18\x02 \x01(\x0b\x32%.mgmt.v1alpha1.ConnectionStreamConfigR\x0cstreamConfig\x12\x16\n\x06schema\x18\x03 \x01(\tR\x06schema\x12\x14\n\x05table\x18\x04 \x01(\tR\x05table\"\xa4\x01\n\x1fGetConnectionDataStreamResponse\x12I\n\x03row\x18\x01 \x03(\x0b\x32\x37.mgmt.v1alpha1.GetConnectionDataStreamResponse.RowEntryR\x03row\x1a\x36\n\x08RowEntry\x12\x10\n\x03key\x18\x01 \x01(\tR\x03key\x12\x14\n\x05value\x18\x02 \x01(\x0cR\x05value:\x02\x38\x01\"\x16\n\x14PostgresSchemaConfig\"\x13\n\x11MysqlSchemaConfig\"e\n\x11\x41wsS3SchemaConfig\x12!\n\x06job_id\x18\x01 \x01(\tB\x08\xbaH\x05r\x03\xb0\x01\x01H\x00R\x05jobId\x12\'\n\njob_run_id\x18\x02 \x01(\tB\x07\xbaH\x04r\x02\x10\x01H\x00R\x08jobRunIdB\x04\n\x02id\"\xfc\x01\n\x16\x43onnectionSchemaConfig\x12\x42\n\tpg_config\x18\x01 \x01(\x0b\x32#.mgmt.v1alpha1.PostgresSchemaConfigH\x00R\x08pgConfig\x12\x46\n\raws_s3_config\x18\x02 \x01(\x0b\x32 .mgmt.v1alpha1.AwsS3SchemaConfigH\x00R\x0b\x61wsS3Config\x12\x45\n\x0cmysql_config\x18\x03 \x01(\x0b\x32 .mgmt.v1alpha1.MysqlSchemaConfigH\x00R\x0bmysqlConfigB\x0f\n\x06\x63onfig\x12\x05\xbaH\x02\x08\x01\"s\n\x0e\x44\x61tabaseColumn\x12\x16\n\x06schema\x18\x01 \x01(\tR\x06schema\x12\x14\n\x05table\x18\x02 \x01(\tR\x05table\x12\x16\n\x06\x63olumn\x18\x03 \x01(\tR\x06\x63olumn\x12\x1b\n\tdata_type\x18\x04 \x01(\tR\x08\x64\x61taType\"\x97\x01\n\x1aGetConnectionSchemaRequest\x12-\n\rconnection_id\x18\x01 \x01(\tB\x08\xbaH\x05r\x03\xb0\x01\x01R\x0c\x63onnectionId\x12J\n\rschema_config\x18\x02 \x01(\x0b\x32%.mgmt.v1alpha1.ConnectionSchemaConfigR\x0cschemaConfig\"V\n\x1bGetConnectionSchemaResponse\x12\x37\n\x07schemas\x18\x01 \x03(\x0b\x32\x1d.mgmt.v1alpha1.DatabaseColumnR\x07schemas\"W\n&GetConnectionForeignConstraintsRequest\x12-\n\rconnection_id\x18\x01 \x01(\tB\x08\xbaH\x05r\x03\xb0\x01\x01R\x0c\x63onnectionId\":\n\nForeignKey\x12\x14\n\x05table\x18\x01 \x01(\tR\x05table\x12\x16\n\x06\x63olumn\x18\x02 \x01(\tR\x06\x63olumn\"\x88\x01\n\x11\x46oreignConstraint\x12\x16\n\x06\x63olumn\x18\x01 \x01(\tR\x06\x63olumn\x12\x1f\n\x0bis_nullable\x18\x02 \x01(\x08R\nisNullable\x12:\n\x0b\x66oreign_key\x18\x03 \x01(\x0b\x32\x19.mgmt.v1alpha1.ForeignKeyR\nforeignKey\"]\n\x17\x46oreignConstraintTables\x12\x42\n\x0b\x63onstraints\x18\x01 \x03(\x0b\x32 .mgmt.v1alpha1.ForeignConstraintR\x0b\x63onstraints\"\x91\x02\n\'GetConnectionForeignConstraintsResponse\x12y\n\x11table_constraints\x18\x01 \x03(\x0b\x32L.mgmt.v1alpha1.GetConnectionForeignConstraintsResponse.TableConstraintsEntryR\x10tableConstraints\x1ak\n\x15TableConstraintsEntry\x12\x10\n\x03key\x18\x01 \x01(\tR\x03key\x12<\n\x05value\x18\x02 \x01(\x0b\x32&.mgmt.v1alpha1.ForeignConstraintTablesR\x05value:\x02\x38\x01\"\x98\x01\n\x14InitStatementOptions\x12\x1f\n\x0binit_schema\x18\x01 \x01(\x08R\ninitSchema\x12\x34\n\x16truncate_before_insert\x18\x02 \x01(\x08R\x14truncateBeforeInsert\x12)\n\x10truncate_cascade\x18\x03 \x01(\x08R\x0ftruncateCascade\"\x92\x01\n\"GetConnectionInitStatementsRequest\x12-\n\rconnection_id\x18\x01 \x01(\tB\x08\xbaH\x05r\x03\xb0\x01\x01R\x0c\x63onnectionId\x12=\n\x07options\x18\x02 \x01(\x0b\x32#.mgmt.v1alpha1.InitStatementOptionsR\x07options\"\xee\x01\n#GetConnectionInitStatementsResponse\x12\x7f\n\x15table_init_statements\x18\x01 \x03(\x0b\x32K.mgmt.v1alpha1.GetConnectionInitStatementsResponse.TableInitStatementsEntryR\x13tableInitStatements\x1a\x46\n\x18TableInitStatementsEntry\x12\x10\n\x03key\x18\x01 \x01(\tR\x03key\x12\x14\n\x05value\x18\x02 \x01(\tR\x05value:\x02\x38\x01\x32\xa3\x04\n\x15\x43onnectionDataService\x12|\n\x17GetConnectionDataStream\x12-.mgmt.v1alpha1.GetConnectionDataStreamRequest\x1a..mgmt.v1alpha1.GetConnectionDataStreamResponse\"\x00\x30\x01\x12n\n\x13GetConnectionSchema\x12).mgmt.v1alpha1.GetConnectionSchemaRequest\x1a*.mgmt.v1alpha1.GetConnectionSchemaResponse\"\x00\x12\x92\x01\n\x1fGetConnectionForeignConstraints\x12\x35.mgmt.v1alpha1.GetConnectionForeignConstraintsRequest\x1a\x36.mgmt.v1alpha1.GetConnectionForeignConstraintsResponse\"\x00\x12\x86\x01\n\x1bGetConnectionInitStatements\x12\x31.mgmt.v1alpha1.GetConnectionInitStatementsRequest\x1a\x32.mgmt.v1alpha1.GetConnectionInitStatementsResponse\"\x00\x42\xcf\x01\n\x11\x63om.mgmt.v1alpha1B\x13\x43onnectionDataProtoP\x01ZPgithub.com/nucleuscloud/neosync/backend/gen/go/protos/mgmt/v1alpha1;mgmtv1alpha1\xa2\x02\x03MXX\xaa\x02\rMgmt.V1alpha1\xca\x02\rMgmt\\V1alpha1\xe2\x02\x19Mgmt\\V1alpha1\\GPBMetadata\xea\x02\x0eMgmt::V1alpha1b\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'mgmt.v1alpha1.connection_data_pb2', _globals)
if _descriptor._USE_C_DESCRIPTORS == False:
  _globals['DESCRIPTOR']._options = None
  _globals['DESCRIPTOR']._serialized_options = b'\n\021com.mgmt.v1alpha1B\023ConnectionDataProtoP\001ZPgithub.com/nucleuscloud/neosync/backend/gen/go/protos/mgmt/v1alpha1;mgmtv1alpha1\242\002\003MXX\252\002\rMgmt.V1alpha1\312\002\rMgmt\\V1alpha1\342\002\031Mgmt\\V1alpha1\\GPBMetadata\352\002\016Mgmt::V1alpha1'
  _globals['_AWSS3STREAMCONFIG'].fields_by_name['job_id']._options = None
  _globals['_AWSS3STREAMCONFIG'].fields_by_name['job_id']._serialized_options = b'\272H\005r\003\260\001\001'
  _globals['_AWSS3STREAMCONFIG'].fields_by_name['job_run_id']._options = None
  _globals['_AWSS3STREAMCONFIG'].fields_by_name['job_run_id']._serialized_options = b'\272H\004r\002\020\001'
  _globals['_CONNECTIONSTREAMCONFIG'].oneofs_by_name['config']._options = None
  _globals['_CONNECTIONSTREAMCONFIG'].oneofs_by_name['config']._serialized_options = b'\272H\002\010\001'
  _globals['_GETCONNECTIONDATASTREAMREQUEST'].fields_by_name['connection_id']._options = None
  _globals['_GETCONNECTIONDATASTREAMREQUEST'].fields_by_name['connection_id']._serialized_options = b'\272H\005r\003\260\001\001'
  _globals['_GETCONNECTIONDATASTREAMRESPONSE_ROWENTRY']._options = None
  _globals['_GETCONNECTIONDATASTREAMRESPONSE_ROWENTRY']._serialized_options = b'8\001'
  _globals['_AWSS3SCHEMACONFIG'].fields_by_name['job_id']._options = None
  _globals['_AWSS3SCHEMACONFIG'].fields_by_name['job_id']._serialized_options = b'\272H\005r\003\260\001\001'
  _globals['_AWSS3SCHEMACONFIG'].fields_by_name['job_run_id']._options = None
  _globals['_AWSS3SCHEMACONFIG'].fields_by_name['job_run_id']._serialized_options = b'\272H\004r\002\020\001'
  _globals['_CONNECTIONSCHEMACONFIG'].oneofs_by_name['config']._options = None
  _globals['_CONNECTIONSCHEMACONFIG'].oneofs_by_name['config']._serialized_options = b'\272H\002\010\001'
  _globals['_GETCONNECTIONSCHEMAREQUEST'].fields_by_name['connection_id']._options = None
  _globals['_GETCONNECTIONSCHEMAREQUEST'].fields_by_name['connection_id']._serialized_options = b'\272H\005r\003\260\001\001'
  _globals['_GETCONNECTIONFOREIGNCONSTRAINTSREQUEST'].fields_by_name['connection_id']._options = None
  _globals['_GETCONNECTIONFOREIGNCONSTRAINTSREQUEST'].fields_by_name['connection_id']._serialized_options = b'\272H\005r\003\260\001\001'
  _globals['_GETCONNECTIONFOREIGNCONSTRAINTSRESPONSE_TABLECONSTRAINTSENTRY']._options = None
  _globals['_GETCONNECTIONFOREIGNCONSTRAINTSRESPONSE_TABLECONSTRAINTSENTRY']._serialized_options = b'8\001'
  _globals['_GETCONNECTIONINITSTATEMENTSREQUEST'].fields_by_name['connection_id']._options = None
  _globals['_GETCONNECTIONINITSTATEMENTSREQUEST'].fields_by_name['connection_id']._serialized_options = b'\272H\005r\003\260\001\001'
  _globals['_GETCONNECTIONINITSTATEMENTSRESPONSE_TABLEINITSTATEMENTSENTRY']._options = None
  _globals['_GETCONNECTIONINITSTATEMENTSRESPONSE_TABLEINITSTATEMENTSENTRY']._serialized_options = b'8\001'
  _globals['_POSTGRESSTREAMCONFIG']._serialized_start=83
  _globals['_POSTGRESSTREAMCONFIG']._serialized_end=105
  _globals['_MYSQLSTREAMCONFIG']._serialized_start=107
  _globals['_MYSQLSTREAMCONFIG']._serialized_end=126
  _globals['_AWSS3STREAMCONFIG']._serialized_start=128
  _globals['_AWSS3STREAMCONFIG']._serialized_end=229
  _globals['_CONNECTIONSTREAMCONFIG']._serialized_start=232
  _globals['_CONNECTIONSTREAMCONFIG']._serialized_end=484
  _globals['_GETCONNECTIONDATASTREAMREQUEST']._serialized_start=487
  _globals['_GETCONNECTIONDATASTREAMREQUEST']._serialized_end=688
  _globals['_GETCONNECTIONDATASTREAMRESPONSE']._serialized_start=691
  _globals['_GETCONNECTIONDATASTREAMRESPONSE']._serialized_end=855
  _globals['_GETCONNECTIONDATASTREAMRESPONSE_ROWENTRY']._serialized_start=801
  _globals['_GETCONNECTIONDATASTREAMRESPONSE_ROWENTRY']._serialized_end=855
  _globals['_POSTGRESSCHEMACONFIG']._serialized_start=857
  _globals['_POSTGRESSCHEMACONFIG']._serialized_end=879
  _globals['_MYSQLSCHEMACONFIG']._serialized_start=881
  _globals['_MYSQLSCHEMACONFIG']._serialized_end=900
  _globals['_AWSS3SCHEMACONFIG']._serialized_start=902
  _globals['_AWSS3SCHEMACONFIG']._serialized_end=1003
  _globals['_CONNECTIONSCHEMACONFIG']._serialized_start=1006
  _globals['_CONNECTIONSCHEMACONFIG']._serialized_end=1258
  _globals['_DATABASECOLUMN']._serialized_start=1260
  _globals['_DATABASECOLUMN']._serialized_end=1375
  _globals['_GETCONNECTIONSCHEMAREQUEST']._serialized_start=1378
  _globals['_GETCONNECTIONSCHEMAREQUEST']._serialized_end=1529
  _globals['_GETCONNECTIONSCHEMARESPONSE']._serialized_start=1531
  _globals['_GETCONNECTIONSCHEMARESPONSE']._serialized_end=1617
  _globals['_GETCONNECTIONFOREIGNCONSTRAINTSREQUEST']._serialized_start=1619
  _globals['_GETCONNECTIONFOREIGNCONSTRAINTSREQUEST']._serialized_end=1706
  _globals['_FOREIGNKEY']._serialized_start=1708
  _globals['_FOREIGNKEY']._serialized_end=1766
  _globals['_FOREIGNCONSTRAINT']._serialized_start=1769
  _globals['_FOREIGNCONSTRAINT']._serialized_end=1905
  _globals['_FOREIGNCONSTRAINTTABLES']._serialized_start=1907
  _globals['_FOREIGNCONSTRAINTTABLES']._serialized_end=2000
  _globals['_GETCONNECTIONFOREIGNCONSTRAINTSRESPONSE']._serialized_start=2003
  _globals['_GETCONNECTIONFOREIGNCONSTRAINTSRESPONSE']._serialized_end=2276
  _globals['_GETCONNECTIONFOREIGNCONSTRAINTSRESPONSE_TABLECONSTRAINTSENTRY']._serialized_start=2169
  _globals['_GETCONNECTIONFOREIGNCONSTRAINTSRESPONSE_TABLECONSTRAINTSENTRY']._serialized_end=2276
  _globals['_INITSTATEMENTOPTIONS']._serialized_start=2279
  _globals['_INITSTATEMENTOPTIONS']._serialized_end=2431
  _globals['_GETCONNECTIONINITSTATEMENTSREQUEST']._serialized_start=2434
  _globals['_GETCONNECTIONINITSTATEMENTSREQUEST']._serialized_end=2580
  _globals['_GETCONNECTIONINITSTATEMENTSRESPONSE']._serialized_start=2583
  _globals['_GETCONNECTIONINITSTATEMENTSRESPONSE']._serialized_end=2821
  _globals['_GETCONNECTIONINITSTATEMENTSRESPONSE_TABLEINITSTATEMENTSENTRY']._serialized_start=2751
  _globals['_GETCONNECTIONINITSTATEMENTSRESPONSE_TABLEINITSTATEMENTSENTRY']._serialized_end=2821
  _globals['_CONNECTIONDATASERVICE']._serialized_start=2824
  _globals['_CONNECTIONDATASERVICE']._serialized_end=3371
# @@protoc_insertion_point(module_scope)