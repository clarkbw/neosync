package datasync_workflow

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"
	"time"

	"connectrpc.com/connect"
	mysql_queries "github.com/nucleuscloud/neosync/backend/gen/go/db/dbschemas/mysql"
	pg_queries "github.com/nucleuscloud/neosync/backend/gen/go/db/dbschemas/postgresql"
	mgmtv1alpha1 "github.com/nucleuscloud/neosync/backend/gen/go/protos/mgmt/v1alpha1"
	"github.com/nucleuscloud/neosync/backend/gen/go/protos/mgmt/v1alpha1/mgmtv1alpha1connect"
	"github.com/nucleuscloud/neosync/backend/pkg/sqlconnect"
	sql_manager "github.com/nucleuscloud/neosync/backend/pkg/sqlmanager"
	sqlmanager_shared "github.com/nucleuscloud/neosync/backend/pkg/sqlmanager/shared"
	mockTemporalClient "github.com/nucleuscloud/neosync/worker/internal/mocks/go.temporal.io/sdk/client"
	genbenthosconfigs_activity "github.com/nucleuscloud/neosync/worker/pkg/workflows/datasync/activities/gen-benthos-configs"
	runsqlinittablestmts_activity "github.com/nucleuscloud/neosync/worker/pkg/workflows/datasync/activities/run-sql-init-table-stmts"
	"github.com/nucleuscloud/neosync/worker/pkg/workflows/datasync/activities/shared"
	sync_activity "github.com/nucleuscloud/neosync/worker/pkg/workflows/datasync/activities/sync"
	syncactivityopts_activity "github.com/nucleuscloud/neosync/worker/pkg/workflows/datasync/activities/sync-activity-opts"
	syncrediscleanup_activity "github.com/nucleuscloud/neosync/worker/pkg/workflows/datasync/activities/sync-redis-clean-up"
	workflow_testdata "github.com/nucleuscloud/neosync/worker/pkg/workflows/datasync/workflow/testdata"
	testdata_doublereference "github.com/nucleuscloud/neosync/worker/pkg/workflows/datasync/workflow/testdata/double-reference"
	testdata_virtualforeignkeys "github.com/nucleuscloud/neosync/worker/pkg/workflows/datasync/workflow/testdata/virtual-foreign-keys"
	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/otel/metric"
	"go.temporal.io/sdk/testsuite"
)

func getAllSyncTests() []*workflow_testdata.IntegrationTest {
	allTests := []*workflow_testdata.IntegrationTest{}
	drTests := testdata_doublereference.GetSyncTests()
	vfkTests := testdata_virtualforeignkeys.GetSyncTests()
	allTests = append(allTests, drTests...)
	allTests = append(allTests, vfkTests...)
	return allTests
}

func (s *IntegrationTestSuite) Test_Workflow_Sync() {
	tests := getAllSyncTests()
	for _, tt := range tests {
		s.T().Run(tt.Name, func(t *testing.T) {
			s.T().Logf("running integration test: %s \n", tt.Name)
			s.SetupSourceDb(tt.Folder, tt.SourceFilePaths)
			s.SetupTargetDb(tt.Folder, tt.TargetFilePaths)

			schemas := []*mgmtv1alpha1.PostgresSourceSchemaOption{}
			subsetMap := map[string]*mgmtv1alpha1.PostgresSourceSchemaOption{}
			for table, where := range tt.SubsetMap {
				schema, table := sqlmanager_shared.SplitTableKey(table)
				if _, exists := subsetMap[schema]; !exists {
					subsetMap[schema] = &mgmtv1alpha1.PostgresSourceSchemaOption{
						Schema: schema,
						Tables: []*mgmtv1alpha1.PostgresSourceTableOption{},
					}
				}
				w := where
				subsetMap[schema].Tables = append(subsetMap[schema].Tables, &mgmtv1alpha1.PostgresSourceTableOption{
					Table:       table,
					WhereClause: &w,
				})
			}

			for _, s := range subsetMap {
				schemas = append(schemas, s)
			}

			var subsetByForeignKeyConstraints bool
			if tt.JobOptions != nil && tt.JobOptions.SubsetByForeignKeyConstraints {
				subsetByForeignKeyConstraints = true
			}

			mux := http.NewServeMux()
			mux.Handle(mgmtv1alpha1connect.JobServiceGetJobProcedure, connect.NewUnaryHandler(
				mgmtv1alpha1connect.JobServiceGetJobProcedure,
				func(ctx context.Context, r *connect.Request[mgmtv1alpha1.GetJobRequest]) (*connect.Response[mgmtv1alpha1.GetJobResponse], error) {
					return connect.NewResponse(&mgmtv1alpha1.GetJobResponse{
						Job: &mgmtv1alpha1.Job{
							Id: "115aaf2c-776e-4847-8268-d914e3c15968",
							Source: &mgmtv1alpha1.JobSource{
								Options: &mgmtv1alpha1.JobSourceOptions{
									Config: &mgmtv1alpha1.JobSourceOptions_Postgres{
										Postgres: &mgmtv1alpha1.PostgresSourceConnectionOptions{
											ConnectionId:                  "c9b6ce58-5c8e-4dce-870d-96841b19d988",
											Schemas:                       schemas,
											SubsetByForeignKeyConstraints: subsetByForeignKeyConstraints,
										},
									},
								},
							},

							Destinations: []*mgmtv1alpha1.JobDestination{
								{
									ConnectionId: "226add85-5751-4232-b085-a0ae93afc7ce",
								},
							},
							Mappings:           tt.JobMappings,
							VirtualForeignKeys: tt.VirtualForeignKeys,
						}}), nil
				},
			))

			mux.Handle(mgmtv1alpha1connect.ConnectionServiceGetConnectionProcedure, connect.NewUnaryHandler(
				mgmtv1alpha1connect.ConnectionServiceGetConnectionProcedure,
				func(ctx context.Context, r *connect.Request[mgmtv1alpha1.GetConnectionRequest]) (*connect.Response[mgmtv1alpha1.GetConnectionResponse], error) {
					if r.Msg.GetId() == "c9b6ce58-5c8e-4dce-870d-96841b19d988" {
						return connect.NewResponse(&mgmtv1alpha1.GetConnectionResponse{
							Connection: &mgmtv1alpha1.Connection{
								Id:   "c9b6ce58-5c8e-4dce-870d-96841b19d988",
								Name: "source",
								ConnectionConfig: &mgmtv1alpha1.ConnectionConfig{
									Config: &mgmtv1alpha1.ConnectionConfig_PgConfig{
										PgConfig: &mgmtv1alpha1.PostgresConnectionConfig{
											ConnectionConfig: &mgmtv1alpha1.PostgresConnectionConfig_Url{
												Url: s.sourceDsn,
											},
										},
									},
								},
							},
						}), nil
					}
					if r.Msg.GetId() == "226add85-5751-4232-b085-a0ae93afc7ce" {
						return connect.NewResponse(&mgmtv1alpha1.GetConnectionResponse{
							Connection: &mgmtv1alpha1.Connection{
								Id:   "226add85-5751-4232-b085-a0ae93afc7ce",
								Name: "target",
								ConnectionConfig: &mgmtv1alpha1.ConnectionConfig{
									Config: &mgmtv1alpha1.ConnectionConfig_PgConfig{
										PgConfig: &mgmtv1alpha1.PostgresConnectionConfig{
											ConnectionConfig: &mgmtv1alpha1.PostgresConnectionConfig_Url{
												Url: s.targetDsn,
											},
										},
									},
								},
							},
						}), nil
					}
					return nil, nil
				},
			))
			srv := startHTTPServer(s.T(), mux)
			executeWorkflow(s.T(), srv, s.redisUrl, "115aaf2c-776e-4847-8268-d914e3c15968", tt.Name)

			for table, expected := range tt.Expected {
				rows, err := s.targetPgPool.Query(s.ctx, fmt.Sprintf("select * from %s;", table))
				require.NoError(s.T(), err)
				count := 0
				for rows.Next() {
					count++
				}
				require.Equalf(s.T(), expected.RowCount, count, fmt.Sprintf("Test: %s Table: %s", tt.Name, table))
			}

			s.TearDownTestByFolder(tt.Folder)
		})
	}
}

func (s *IntegrationTestSuite) Test_Workflow_VirtualForeignKeys_Transform() {
	testFolder := "virtual-foreign-keys"
	s.SetupSourceDb(testFolder, []string{"source-setup.sql"})
	s.SetupTargetDb(testFolder, []string{"target-setup.sql"})
	virtualForeignKeys := testdata_virtualforeignkeys.GetVirtualForeignKeys()
	jobmappings := testdata_virtualforeignkeys.GetDefaultSyncJobMappings()

	for _, m := range jobmappings {
		if m.Table == "countries" && m.Column == "country_id" {
			m.Transformer = &mgmtv1alpha1.JobMappingTransformer{
				Source: mgmtv1alpha1.TransformerSource_TRANSFORMER_SOURCE_TRANSFORM_JAVASCRIPT,
				Config: &mgmtv1alpha1.TransformerConfig{
					Config: &mgmtv1alpha1.TransformerConfig_TransformJavascriptConfig{
						TransformJavascriptConfig: &mgmtv1alpha1.TransformJavascript{Code: `if (value == 'US') { return 'SU'; } return value;`},
					},
				},
			}
		}
	}
	// neosync api mocks
	mux := http.NewServeMux()
	mux.Handle(mgmtv1alpha1connect.JobServiceGetJobProcedure, connect.NewUnaryHandler(
		mgmtv1alpha1connect.JobServiceGetJobProcedure,
		func(ctx context.Context, r *connect.Request[mgmtv1alpha1.GetJobRequest]) (*connect.Response[mgmtv1alpha1.GetJobResponse], error) {
			return connect.NewResponse(&mgmtv1alpha1.GetJobResponse{
				Job: &mgmtv1alpha1.Job{
					Id: "fd4d8660-31a0-48b2-9adf-10f11b94898f",
					Source: &mgmtv1alpha1.JobSource{
						Options: &mgmtv1alpha1.JobSourceOptions{
							Config: &mgmtv1alpha1.JobSourceOptions_Postgres{
								Postgres: &mgmtv1alpha1.PostgresSourceConnectionOptions{
									ConnectionId: "c9b6ce58-5c8e-4dce-870d-96841b19d988",
								},
							},
						},
					},
					Mappings:           jobmappings,
					VirtualForeignKeys: virtualForeignKeys,
					Destinations: []*mgmtv1alpha1.JobDestination{
						{
							ConnectionId: "226add85-5751-4232-b085-a0ae93afc7ce",
						},
					},
				},
			}), nil
		},
	))

	mux.Handle(mgmtv1alpha1connect.ConnectionServiceGetConnectionProcedure, connect.NewUnaryHandler(
		mgmtv1alpha1connect.ConnectionServiceGetConnectionProcedure,
		func(ctx context.Context, r *connect.Request[mgmtv1alpha1.GetConnectionRequest]) (*connect.Response[mgmtv1alpha1.GetConnectionResponse], error) {
			if r.Msg.GetId() == "c9b6ce58-5c8e-4dce-870d-96841b19d988" {
				return connect.NewResponse(&mgmtv1alpha1.GetConnectionResponse{
					Connection: &mgmtv1alpha1.Connection{
						Id:   "c9b6ce58-5c8e-4dce-870d-96841b19d988",
						Name: "source",
						ConnectionConfig: &mgmtv1alpha1.ConnectionConfig{
							Config: &mgmtv1alpha1.ConnectionConfig_PgConfig{
								PgConfig: &mgmtv1alpha1.PostgresConnectionConfig{
									ConnectionConfig: &mgmtv1alpha1.PostgresConnectionConfig_Url{
										Url: s.sourceDsn,
									},
								},
							},
						},
					},
				}), nil
			}
			if r.Msg.GetId() == "226add85-5751-4232-b085-a0ae93afc7ce" {
				return connect.NewResponse(&mgmtv1alpha1.GetConnectionResponse{
					Connection: &mgmtv1alpha1.Connection{
						Id:   "226add85-5751-4232-b085-a0ae93afc7ce",
						Name: "target",
						ConnectionConfig: &mgmtv1alpha1.ConnectionConfig{
							Config: &mgmtv1alpha1.ConnectionConfig_PgConfig{
								PgConfig: &mgmtv1alpha1.PostgresConnectionConfig{
									ConnectionConfig: &mgmtv1alpha1.PostgresConnectionConfig_Url{
										Url: s.targetDsn,
									},
								},
							},
						},
					},
				}), nil
			}
			return nil, nil
		},
	))
	srv := startHTTPServer(s.T(), mux)
	executeWorkflow(s.T(), srv, s.redisUrl, "fd4d8660-31a0-48b2-9adf-10f11b94898f", "Virtual Foreign Key primary key transform")

	tables := []string{"regions", "countries", "locations", "departments", "dependents", "jobs", "employees"}
	for _, t := range tables {
		rows, err := s.targetPgPool.Query(s.ctx, fmt.Sprintf("select * from vfk_hr.%s;", t))
		require.NoError(s.T(), err)
		count := 0
		for rows.Next() {
			count++
		}
		require.Greater(s.T(), count, 0)
		require.NoError(s.T(), err)
	}

	rows := s.sourcePgPool.QueryRow(s.ctx, "select count(*) from vfk_hr.countries where country_id = 'US';")
	var rowCount int
	err := rows.Scan(&rowCount)
	require.NoError(s.T(), err)
	require.Equal(s.T(), 1, rowCount)

	rows = s.sourcePgPool.QueryRow(s.ctx, "select count(*) from vfk_hr.locations where country_id = 'US';")
	err = rows.Scan(&rowCount)
	require.NoError(s.T(), err)
	require.Equal(s.T(), 3, rowCount)

	rows = s.targetPgPool.QueryRow(s.ctx, "select count(*) from vfk_hr.countries where country_id = 'US';")
	err = rows.Scan(&rowCount)
	require.NoError(s.T(), err)
	require.Equal(s.T(), 0, rowCount)

	rows = s.targetPgPool.QueryRow(s.ctx, "select count(*) from vfk_hr.countries where country_id = 'SU';")
	err = rows.Scan(&rowCount)
	require.NoError(s.T(), err)
	require.Equal(s.T(), 1, rowCount)

	rows = s.targetPgPool.QueryRow(s.ctx, "select count(*) from vfk_hr.locations where country_id = 'SU';")
	err = rows.Scan(&rowCount)
	require.NoError(s.T(), err)
	require.Equal(s.T(), 3, rowCount)

	s.TearDownTestByFolder(testFolder)
}

func executeWorkflow(
	t *testing.T,
	srv *httptest.Server,
	redisUrl string,
	jobId string,
	testName string,
) {
	connclient := mgmtv1alpha1connect.NewConnectionServiceClient(srv.Client(), srv.URL)
	jobclient := mgmtv1alpha1connect.NewJobServiceClient(srv.Client(), srv.URL)
	transformerclient := mgmtv1alpha1connect.NewTransformersServiceClient(srv.Client(), srv.URL)
	sqlconnector := &sqlconnect.SqlOpenConnector{}
	redisconfig := &shared.RedisConfig{
		Url:  redisUrl,
		Kind: "simple",
		Tls: &shared.RedisTlsConfig{
			Enabled: false,
		},
	}
	temporalClientMock := mockTemporalClient.NewMockClient(t)
	pgpoolmap := &sync.Map{}
	mysqlpoolmap := &sync.Map{}
	pgquerier := pg_queries.New()
	mysqlquerier := mysql_queries.New()
	sqlmanager := sql_manager.NewSqlManager(pgpoolmap, pgquerier, mysqlpoolmap, mysqlquerier, sqlconnector)

	// temporal workflow
	testSuite := &testsuite.WorkflowTestSuite{}
	env := testSuite.NewTestWorkflowEnvironment()

	// register activities
	genbenthosActivity := genbenthosconfigs_activity.New(
		jobclient,
		connclient,
		transformerclient,
		sqlmanager,
		redisconfig,
		false,
	)
	var activityMeter metric.Meter
	disableReaper := true
	syncActivity := sync_activity.New(connclient, &sync.Map{}, temporalClientMock, activityMeter, sync_activity.NewBenthosStreamManager(), disableReaper)
	retrieveActivityOpts := syncactivityopts_activity.New(jobclient)
	runSqlInitTableStatements := runsqlinittablestmts_activity.New(jobclient, connclient, sqlmanager)
	env.RegisterWorkflow(Workflow)
	env.RegisterActivity(syncActivity.Sync)
	env.RegisterActivity(retrieveActivityOpts.RetrieveActivityOptions)
	env.RegisterActivity(runSqlInitTableStatements.RunSqlInitTableStatements)
	env.RegisterActivity(syncrediscleanup_activity.DeleteRedisHash)
	env.RegisterActivity(genbenthosActivity.GenerateBenthosConfigs)
	env.SetTestTimeout(120 * time.Second) // increase the test timeout

	env.ExecuteWorkflow(Workflow, &WorkflowRequest{JobId: jobId})
	require.Truef(t, env.IsWorkflowCompleted(), fmt.Sprintf("Workflow did not complete. Test: %s", testName))

	err := env.GetWorkflowError()
	require.Nilf(t, err, fmt.Sprintf("Workflow error. Test: %s", testName))
}

func startHTTPServer(tb testing.TB, h http.Handler) *httptest.Server {
	tb.Helper()
	srv := httptest.NewUnstartedServer(h)
	srv.EnableHTTP2 = true
	srv.Start()
	tb.Cleanup(srv.Close)
	return srv
}
