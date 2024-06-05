// Code generated by mockery. DO NOT EDIT.

package sqlmanager

import (
	context "context"

	sqlmanager_shared "github.com/nucleuscloud/neosync/backend/pkg/sqlmanager/shared"
	mock "github.com/stretchr/testify/mock"
)

// MockSqlDatabase is an autogenerated mock type for the SqlDatabase type
type MockSqlDatabase struct {
	mock.Mock
}

type MockSqlDatabase_Expecter struct {
	mock *mock.Mock
}

func (_m *MockSqlDatabase) EXPECT() *MockSqlDatabase_Expecter {
	return &MockSqlDatabase_Expecter{mock: &_m.Mock}
}

// BatchExec provides a mock function with given fields: ctx, batchSize, statements, opts
func (_m *MockSqlDatabase) BatchExec(ctx context.Context, batchSize int, statements []string, opts *sqlmanager_shared.BatchExecOpts) error {
	ret := _m.Called(ctx, batchSize, statements, opts)

	if len(ret) == 0 {
		panic("no return value specified for BatchExec")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int, []string, *sqlmanager_shared.BatchExecOpts) error); ok {
		r0 = rf(ctx, batchSize, statements, opts)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockSqlDatabase_BatchExec_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'BatchExec'
type MockSqlDatabase_BatchExec_Call struct {
	*mock.Call
}

// BatchExec is a helper method to define mock.On call
//   - ctx context.Context
//   - batchSize int
//   - statements []string
//   - opts *sqlmanager_shared.BatchExecOpts
func (_e *MockSqlDatabase_Expecter) BatchExec(ctx interface{}, batchSize interface{}, statements interface{}, opts interface{}) *MockSqlDatabase_BatchExec_Call {
	return &MockSqlDatabase_BatchExec_Call{Call: _e.mock.On("BatchExec", ctx, batchSize, statements, opts)}
}

func (_c *MockSqlDatabase_BatchExec_Call) Run(run func(ctx context.Context, batchSize int, statements []string, opts *sqlmanager_shared.BatchExecOpts)) *MockSqlDatabase_BatchExec_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int), args[2].([]string), args[3].(*sqlmanager_shared.BatchExecOpts))
	})
	return _c
}

func (_c *MockSqlDatabase_BatchExec_Call) Return(_a0 error) *MockSqlDatabase_BatchExec_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockSqlDatabase_BatchExec_Call) RunAndReturn(run func(context.Context, int, []string, *sqlmanager_shared.BatchExecOpts) error) *MockSqlDatabase_BatchExec_Call {
	_c.Call.Return(run)
	return _c
}

// Close provides a mock function with given fields:
func (_m *MockSqlDatabase) Close() {
	_m.Called()
}

// MockSqlDatabase_Close_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Close'
type MockSqlDatabase_Close_Call struct {
	*mock.Call
}

// Close is a helper method to define mock.On call
func (_e *MockSqlDatabase_Expecter) Close() *MockSqlDatabase_Close_Call {
	return &MockSqlDatabase_Close_Call{Call: _e.mock.On("Close")}
}

func (_c *MockSqlDatabase_Close_Call) Run(run func()) *MockSqlDatabase_Close_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockSqlDatabase_Close_Call) Return() *MockSqlDatabase_Close_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockSqlDatabase_Close_Call) RunAndReturn(run func()) *MockSqlDatabase_Close_Call {
	_c.Call.Return(run)
	return _c
}

// Exec provides a mock function with given fields: ctx, statement
func (_m *MockSqlDatabase) Exec(ctx context.Context, statement string) error {
	ret := _m.Called(ctx, statement)

	if len(ret) == 0 {
		panic("no return value specified for Exec")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, statement)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockSqlDatabase_Exec_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Exec'
type MockSqlDatabase_Exec_Call struct {
	*mock.Call
}

// Exec is a helper method to define mock.On call
//   - ctx context.Context
//   - statement string
func (_e *MockSqlDatabase_Expecter) Exec(ctx interface{}, statement interface{}) *MockSqlDatabase_Exec_Call {
	return &MockSqlDatabase_Exec_Call{Call: _e.mock.On("Exec", ctx, statement)}
}

func (_c *MockSqlDatabase_Exec_Call) Run(run func(ctx context.Context, statement string)) *MockSqlDatabase_Exec_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockSqlDatabase_Exec_Call) Return(_a0 error) *MockSqlDatabase_Exec_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockSqlDatabase_Exec_Call) RunAndReturn(run func(context.Context, string) error) *MockSqlDatabase_Exec_Call {
	_c.Call.Return(run)
	return _c
}

// GetCreateTableStatement provides a mock function with given fields: ctx, schema, table
func (_m *MockSqlDatabase) GetCreateTableStatement(ctx context.Context, schema string, table string) (string, error) {
	ret := _m.Called(ctx, schema, table)

	if len(ret) == 0 {
		panic("no return value specified for GetCreateTableStatement")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) (string, error)); ok {
		return rf(ctx, schema, table)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string) string); ok {
		r0 = rf(ctx, schema, table)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, schema, table)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockSqlDatabase_GetCreateTableStatement_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetCreateTableStatement'
type MockSqlDatabase_GetCreateTableStatement_Call struct {
	*mock.Call
}

// GetCreateTableStatement is a helper method to define mock.On call
//   - ctx context.Context
//   - schema string
//   - table string
func (_e *MockSqlDatabase_Expecter) GetCreateTableStatement(ctx interface{}, schema interface{}, table interface{}) *MockSqlDatabase_GetCreateTableStatement_Call {
	return &MockSqlDatabase_GetCreateTableStatement_Call{Call: _e.mock.On("GetCreateTableStatement", ctx, schema, table)}
}

func (_c *MockSqlDatabase_GetCreateTableStatement_Call) Run(run func(ctx context.Context, schema string, table string)) *MockSqlDatabase_GetCreateTableStatement_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string))
	})
	return _c
}

func (_c *MockSqlDatabase_GetCreateTableStatement_Call) Return(_a0 string, _a1 error) *MockSqlDatabase_GetCreateTableStatement_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockSqlDatabase_GetCreateTableStatement_Call) RunAndReturn(run func(context.Context, string, string) (string, error)) *MockSqlDatabase_GetCreateTableStatement_Call {
	_c.Call.Return(run)
	return _c
}

// GetDatabaseSchema provides a mock function with given fields: ctx
func (_m *MockSqlDatabase) GetDatabaseSchema(ctx context.Context) ([]*sqlmanager_shared.DatabaseSchemaRow, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for GetDatabaseSchema")
	}

	var r0 []*sqlmanager_shared.DatabaseSchemaRow
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]*sqlmanager_shared.DatabaseSchemaRow, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []*sqlmanager_shared.DatabaseSchemaRow); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*sqlmanager_shared.DatabaseSchemaRow)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockSqlDatabase_GetDatabaseSchema_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetDatabaseSchema'
type MockSqlDatabase_GetDatabaseSchema_Call struct {
	*mock.Call
}

// GetDatabaseSchema is a helper method to define mock.On call
//   - ctx context.Context
func (_e *MockSqlDatabase_Expecter) GetDatabaseSchema(ctx interface{}) *MockSqlDatabase_GetDatabaseSchema_Call {
	return &MockSqlDatabase_GetDatabaseSchema_Call{Call: _e.mock.On("GetDatabaseSchema", ctx)}
}

func (_c *MockSqlDatabase_GetDatabaseSchema_Call) Run(run func(ctx context.Context)) *MockSqlDatabase_GetDatabaseSchema_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *MockSqlDatabase_GetDatabaseSchema_Call) Return(_a0 []*sqlmanager_shared.DatabaseSchemaRow, _a1 error) *MockSqlDatabase_GetDatabaseSchema_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockSqlDatabase_GetDatabaseSchema_Call) RunAndReturn(run func(context.Context) ([]*sqlmanager_shared.DatabaseSchemaRow, error)) *MockSqlDatabase_GetDatabaseSchema_Call {
	_c.Call.Return(run)
	return _c
}

// GetForeignKeyConstraints provides a mock function with given fields: ctx, schemas
func (_m *MockSqlDatabase) GetForeignKeyConstraints(ctx context.Context, schemas []string) ([]*sqlmanager_shared.ForeignKeyConstraintsRow, error) {
	ret := _m.Called(ctx, schemas)

	if len(ret) == 0 {
		panic("no return value specified for GetForeignKeyConstraints")
	}

	var r0 []*sqlmanager_shared.ForeignKeyConstraintsRow
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, []string) ([]*sqlmanager_shared.ForeignKeyConstraintsRow, error)); ok {
		return rf(ctx, schemas)
	}
	if rf, ok := ret.Get(0).(func(context.Context, []string) []*sqlmanager_shared.ForeignKeyConstraintsRow); ok {
		r0 = rf(ctx, schemas)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*sqlmanager_shared.ForeignKeyConstraintsRow)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, []string) error); ok {
		r1 = rf(ctx, schemas)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockSqlDatabase_GetForeignKeyConstraints_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetForeignKeyConstraints'
type MockSqlDatabase_GetForeignKeyConstraints_Call struct {
	*mock.Call
}

// GetForeignKeyConstraints is a helper method to define mock.On call
//   - ctx context.Context
//   - schemas []string
func (_e *MockSqlDatabase_Expecter) GetForeignKeyConstraints(ctx interface{}, schemas interface{}) *MockSqlDatabase_GetForeignKeyConstraints_Call {
	return &MockSqlDatabase_GetForeignKeyConstraints_Call{Call: _e.mock.On("GetForeignKeyConstraints", ctx, schemas)}
}

func (_c *MockSqlDatabase_GetForeignKeyConstraints_Call) Run(run func(ctx context.Context, schemas []string)) *MockSqlDatabase_GetForeignKeyConstraints_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].([]string))
	})
	return _c
}

func (_c *MockSqlDatabase_GetForeignKeyConstraints_Call) Return(_a0 []*sqlmanager_shared.ForeignKeyConstraintsRow, _a1 error) *MockSqlDatabase_GetForeignKeyConstraints_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockSqlDatabase_GetForeignKeyConstraints_Call) RunAndReturn(run func(context.Context, []string) ([]*sqlmanager_shared.ForeignKeyConstraintsRow, error)) *MockSqlDatabase_GetForeignKeyConstraints_Call {
	_c.Call.Return(run)
	return _c
}

// GetForeignKeyConstraintsMap provides a mock function with given fields: ctx, schemas
func (_m *MockSqlDatabase) GetForeignKeyConstraintsMap(ctx context.Context, schemas []string) (map[string][]*sqlmanager_shared.ForeignConstraint, error) {
	ret := _m.Called(ctx, schemas)

	if len(ret) == 0 {
		panic("no return value specified for GetForeignKeyConstraintsMap")
	}

	var r0 map[string][]*sqlmanager_shared.ForeignConstraint
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, []string) (map[string][]*sqlmanager_shared.ForeignConstraint, error)); ok {
		return rf(ctx, schemas)
	}
	if rf, ok := ret.Get(0).(func(context.Context, []string) map[string][]*sqlmanager_shared.ForeignConstraint); ok {
		r0 = rf(ctx, schemas)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string][]*sqlmanager_shared.ForeignConstraint)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, []string) error); ok {
		r1 = rf(ctx, schemas)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockSqlDatabase_GetForeignKeyConstraintsMap_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetForeignKeyConstraintsMap'
type MockSqlDatabase_GetForeignKeyConstraintsMap_Call struct {
	*mock.Call
}

// GetForeignKeyConstraintsMap is a helper method to define mock.On call
//   - ctx context.Context
//   - schemas []string
func (_e *MockSqlDatabase_Expecter) GetForeignKeyConstraintsMap(ctx interface{}, schemas interface{}) *MockSqlDatabase_GetForeignKeyConstraintsMap_Call {
	return &MockSqlDatabase_GetForeignKeyConstraintsMap_Call{Call: _e.mock.On("GetForeignKeyConstraintsMap", ctx, schemas)}
}

func (_c *MockSqlDatabase_GetForeignKeyConstraintsMap_Call) Run(run func(ctx context.Context, schemas []string)) *MockSqlDatabase_GetForeignKeyConstraintsMap_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].([]string))
	})
	return _c
}

func (_c *MockSqlDatabase_GetForeignKeyConstraintsMap_Call) Return(_a0 map[string][]*sqlmanager_shared.ForeignConstraint, _a1 error) *MockSqlDatabase_GetForeignKeyConstraintsMap_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockSqlDatabase_GetForeignKeyConstraintsMap_Call) RunAndReturn(run func(context.Context, []string) (map[string][]*sqlmanager_shared.ForeignConstraint, error)) *MockSqlDatabase_GetForeignKeyConstraintsMap_Call {
	_c.Call.Return(run)
	return _c
}

// GetPrimaryKeyConstraints provides a mock function with given fields: ctx, schemas
func (_m *MockSqlDatabase) GetPrimaryKeyConstraints(ctx context.Context, schemas []string) ([]*sqlmanager_shared.PrimaryKey, error) {
	ret := _m.Called(ctx, schemas)

	if len(ret) == 0 {
		panic("no return value specified for GetPrimaryKeyConstraints")
	}

	var r0 []*sqlmanager_shared.PrimaryKey
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, []string) ([]*sqlmanager_shared.PrimaryKey, error)); ok {
		return rf(ctx, schemas)
	}
	if rf, ok := ret.Get(0).(func(context.Context, []string) []*sqlmanager_shared.PrimaryKey); ok {
		r0 = rf(ctx, schemas)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*sqlmanager_shared.PrimaryKey)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, []string) error); ok {
		r1 = rf(ctx, schemas)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockSqlDatabase_GetPrimaryKeyConstraints_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetPrimaryKeyConstraints'
type MockSqlDatabase_GetPrimaryKeyConstraints_Call struct {
	*mock.Call
}

// GetPrimaryKeyConstraints is a helper method to define mock.On call
//   - ctx context.Context
//   - schemas []string
func (_e *MockSqlDatabase_Expecter) GetPrimaryKeyConstraints(ctx interface{}, schemas interface{}) *MockSqlDatabase_GetPrimaryKeyConstraints_Call {
	return &MockSqlDatabase_GetPrimaryKeyConstraints_Call{Call: _e.mock.On("GetPrimaryKeyConstraints", ctx, schemas)}
}

func (_c *MockSqlDatabase_GetPrimaryKeyConstraints_Call) Run(run func(ctx context.Context, schemas []string)) *MockSqlDatabase_GetPrimaryKeyConstraints_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].([]string))
	})
	return _c
}

func (_c *MockSqlDatabase_GetPrimaryKeyConstraints_Call) Return(_a0 []*sqlmanager_shared.PrimaryKey, _a1 error) *MockSqlDatabase_GetPrimaryKeyConstraints_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockSqlDatabase_GetPrimaryKeyConstraints_Call) RunAndReturn(run func(context.Context, []string) ([]*sqlmanager_shared.PrimaryKey, error)) *MockSqlDatabase_GetPrimaryKeyConstraints_Call {
	_c.Call.Return(run)
	return _c
}

// GetPrimaryKeyConstraintsMap provides a mock function with given fields: ctx, schemas
func (_m *MockSqlDatabase) GetPrimaryKeyConstraintsMap(ctx context.Context, schemas []string) (map[string][]string, error) {
	ret := _m.Called(ctx, schemas)

	if len(ret) == 0 {
		panic("no return value specified for GetPrimaryKeyConstraintsMap")
	}

	var r0 map[string][]string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, []string) (map[string][]string, error)); ok {
		return rf(ctx, schemas)
	}
	if rf, ok := ret.Get(0).(func(context.Context, []string) map[string][]string); ok {
		r0 = rf(ctx, schemas)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string][]string)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, []string) error); ok {
		r1 = rf(ctx, schemas)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockSqlDatabase_GetPrimaryKeyConstraintsMap_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetPrimaryKeyConstraintsMap'
type MockSqlDatabase_GetPrimaryKeyConstraintsMap_Call struct {
	*mock.Call
}

// GetPrimaryKeyConstraintsMap is a helper method to define mock.On call
//   - ctx context.Context
//   - schemas []string
func (_e *MockSqlDatabase_Expecter) GetPrimaryKeyConstraintsMap(ctx interface{}, schemas interface{}) *MockSqlDatabase_GetPrimaryKeyConstraintsMap_Call {
	return &MockSqlDatabase_GetPrimaryKeyConstraintsMap_Call{Call: _e.mock.On("GetPrimaryKeyConstraintsMap", ctx, schemas)}
}

func (_c *MockSqlDatabase_GetPrimaryKeyConstraintsMap_Call) Run(run func(ctx context.Context, schemas []string)) *MockSqlDatabase_GetPrimaryKeyConstraintsMap_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].([]string))
	})
	return _c
}

func (_c *MockSqlDatabase_GetPrimaryKeyConstraintsMap_Call) Return(_a0 map[string][]string, _a1 error) *MockSqlDatabase_GetPrimaryKeyConstraintsMap_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockSqlDatabase_GetPrimaryKeyConstraintsMap_Call) RunAndReturn(run func(context.Context, []string) (map[string][]string, error)) *MockSqlDatabase_GetPrimaryKeyConstraintsMap_Call {
	_c.Call.Return(run)
	return _c
}

// GetRolePermissionsMap provides a mock function with given fields: ctx, role
func (_m *MockSqlDatabase) GetRolePermissionsMap(ctx context.Context, role string) (map[string][]string, error) {
	ret := _m.Called(ctx, role)

	if len(ret) == 0 {
		panic("no return value specified for GetRolePermissionsMap")
	}

	var r0 map[string][]string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (map[string][]string, error)); ok {
		return rf(ctx, role)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) map[string][]string); ok {
		r0 = rf(ctx, role)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string][]string)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, role)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockSqlDatabase_GetRolePermissionsMap_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetRolePermissionsMap'
type MockSqlDatabase_GetRolePermissionsMap_Call struct {
	*mock.Call
}

// GetRolePermissionsMap is a helper method to define mock.On call
//   - ctx context.Context
//   - role string
func (_e *MockSqlDatabase_Expecter) GetRolePermissionsMap(ctx interface{}, role interface{}) *MockSqlDatabase_GetRolePermissionsMap_Call {
	return &MockSqlDatabase_GetRolePermissionsMap_Call{Call: _e.mock.On("GetRolePermissionsMap", ctx, role)}
}

func (_c *MockSqlDatabase_GetRolePermissionsMap_Call) Run(run func(ctx context.Context, role string)) *MockSqlDatabase_GetRolePermissionsMap_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockSqlDatabase_GetRolePermissionsMap_Call) Return(_a0 map[string][]string, _a1 error) *MockSqlDatabase_GetRolePermissionsMap_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockSqlDatabase_GetRolePermissionsMap_Call) RunAndReturn(run func(context.Context, string) (map[string][]string, error)) *MockSqlDatabase_GetRolePermissionsMap_Call {
	_c.Call.Return(run)
	return _c
}

// GetSchemaColumnMap provides a mock function with given fields: ctx
func (_m *MockSqlDatabase) GetSchemaColumnMap(ctx context.Context) (map[string]map[string]*sqlmanager_shared.ColumnInfo, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for GetSchemaColumnMap")
	}

	var r0 map[string]map[string]*sqlmanager_shared.ColumnInfo
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (map[string]map[string]*sqlmanager_shared.ColumnInfo, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) map[string]map[string]*sqlmanager_shared.ColumnInfo); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]map[string]*sqlmanager_shared.ColumnInfo)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockSqlDatabase_GetSchemaColumnMap_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetSchemaColumnMap'
type MockSqlDatabase_GetSchemaColumnMap_Call struct {
	*mock.Call
}

// GetSchemaColumnMap is a helper method to define mock.On call
//   - ctx context.Context
func (_e *MockSqlDatabase_Expecter) GetSchemaColumnMap(ctx interface{}) *MockSqlDatabase_GetSchemaColumnMap_Call {
	return &MockSqlDatabase_GetSchemaColumnMap_Call{Call: _e.mock.On("GetSchemaColumnMap", ctx)}
}

func (_c *MockSqlDatabase_GetSchemaColumnMap_Call) Run(run func(ctx context.Context)) *MockSqlDatabase_GetSchemaColumnMap_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *MockSqlDatabase_GetSchemaColumnMap_Call) Return(_a0 map[string]map[string]*sqlmanager_shared.ColumnInfo, _a1 error) *MockSqlDatabase_GetSchemaColumnMap_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockSqlDatabase_GetSchemaColumnMap_Call) RunAndReturn(run func(context.Context) (map[string]map[string]*sqlmanager_shared.ColumnInfo, error)) *MockSqlDatabase_GetSchemaColumnMap_Call {
	_c.Call.Return(run)
	return _c
}

// GetSchemaTableDataTypes provides a mock function with given fields: ctx, tables
func (_m *MockSqlDatabase) GetSchemaTableDataTypes(ctx context.Context, tables []*sqlmanager_shared.SchemaTable) (*sqlmanager_shared.SchemaTableDataTypeResponse, error) {
	ret := _m.Called(ctx, tables)

	if len(ret) == 0 {
		panic("no return value specified for GetSchemaTableDataTypes")
	}

	var r0 *sqlmanager_shared.SchemaTableDataTypeResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, []*sqlmanager_shared.SchemaTable) (*sqlmanager_shared.SchemaTableDataTypeResponse, error)); ok {
		return rf(ctx, tables)
	}
	if rf, ok := ret.Get(0).(func(context.Context, []*sqlmanager_shared.SchemaTable) *sqlmanager_shared.SchemaTableDataTypeResponse); ok {
		r0 = rf(ctx, tables)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*sqlmanager_shared.SchemaTableDataTypeResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, []*sqlmanager_shared.SchemaTable) error); ok {
		r1 = rf(ctx, tables)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockSqlDatabase_GetSchemaTableDataTypes_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetSchemaTableDataTypes'
type MockSqlDatabase_GetSchemaTableDataTypes_Call struct {
	*mock.Call
}

// GetSchemaTableDataTypes is a helper method to define mock.On call
//   - ctx context.Context
//   - tables []*sqlmanager_shared.SchemaTable
func (_e *MockSqlDatabase_Expecter) GetSchemaTableDataTypes(ctx interface{}, tables interface{}) *MockSqlDatabase_GetSchemaTableDataTypes_Call {
	return &MockSqlDatabase_GetSchemaTableDataTypes_Call{Call: _e.mock.On("GetSchemaTableDataTypes", ctx, tables)}
}

func (_c *MockSqlDatabase_GetSchemaTableDataTypes_Call) Run(run func(ctx context.Context, tables []*sqlmanager_shared.SchemaTable)) *MockSqlDatabase_GetSchemaTableDataTypes_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].([]*sqlmanager_shared.SchemaTable))
	})
	return _c
}

func (_c *MockSqlDatabase_GetSchemaTableDataTypes_Call) Return(_a0 *sqlmanager_shared.SchemaTableDataTypeResponse, _a1 error) *MockSqlDatabase_GetSchemaTableDataTypes_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockSqlDatabase_GetSchemaTableDataTypes_Call) RunAndReturn(run func(context.Context, []*sqlmanager_shared.SchemaTable) (*sqlmanager_shared.SchemaTableDataTypeResponse, error)) *MockSqlDatabase_GetSchemaTableDataTypes_Call {
	_c.Call.Return(run)
	return _c
}

// GetSchemaTableTriggers provides a mock function with given fields: ctx, tables
func (_m *MockSqlDatabase) GetSchemaTableTriggers(ctx context.Context, tables []*sqlmanager_shared.SchemaTable) ([]*sqlmanager_shared.TableTrigger, error) {
	ret := _m.Called(ctx, tables)

	if len(ret) == 0 {
		panic("no return value specified for GetSchemaTableTriggers")
	}

	var r0 []*sqlmanager_shared.TableTrigger
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, []*sqlmanager_shared.SchemaTable) ([]*sqlmanager_shared.TableTrigger, error)); ok {
		return rf(ctx, tables)
	}
	if rf, ok := ret.Get(0).(func(context.Context, []*sqlmanager_shared.SchemaTable) []*sqlmanager_shared.TableTrigger); ok {
		r0 = rf(ctx, tables)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*sqlmanager_shared.TableTrigger)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, []*sqlmanager_shared.SchemaTable) error); ok {
		r1 = rf(ctx, tables)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockSqlDatabase_GetSchemaTableTriggers_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetSchemaTableTriggers'
type MockSqlDatabase_GetSchemaTableTriggers_Call struct {
	*mock.Call
}

// GetSchemaTableTriggers is a helper method to define mock.On call
//   - ctx context.Context
//   - tables []*sqlmanager_shared.SchemaTable
func (_e *MockSqlDatabase_Expecter) GetSchemaTableTriggers(ctx interface{}, tables interface{}) *MockSqlDatabase_GetSchemaTableTriggers_Call {
	return &MockSqlDatabase_GetSchemaTableTriggers_Call{Call: _e.mock.On("GetSchemaTableTriggers", ctx, tables)}
}

func (_c *MockSqlDatabase_GetSchemaTableTriggers_Call) Run(run func(ctx context.Context, tables []*sqlmanager_shared.SchemaTable)) *MockSqlDatabase_GetSchemaTableTriggers_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].([]*sqlmanager_shared.SchemaTable))
	})
	return _c
}

func (_c *MockSqlDatabase_GetSchemaTableTriggers_Call) Return(_a0 []*sqlmanager_shared.TableTrigger, _a1 error) *MockSqlDatabase_GetSchemaTableTriggers_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockSqlDatabase_GetSchemaTableTriggers_Call) RunAndReturn(run func(context.Context, []*sqlmanager_shared.SchemaTable) ([]*sqlmanager_shared.TableTrigger, error)) *MockSqlDatabase_GetSchemaTableTriggers_Call {
	_c.Call.Return(run)
	return _c
}

// GetTableConstraintsBySchema provides a mock function with given fields: ctx, schemas
func (_m *MockSqlDatabase) GetTableConstraintsBySchema(ctx context.Context, schemas []string) (*sqlmanager_shared.TableConstraints, error) {
	ret := _m.Called(ctx, schemas)

	if len(ret) == 0 {
		panic("no return value specified for GetTableConstraintsBySchema")
	}

	var r0 *sqlmanager_shared.TableConstraints
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, []string) (*sqlmanager_shared.TableConstraints, error)); ok {
		return rf(ctx, schemas)
	}
	if rf, ok := ret.Get(0).(func(context.Context, []string) *sqlmanager_shared.TableConstraints); ok {
		r0 = rf(ctx, schemas)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*sqlmanager_shared.TableConstraints)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, []string) error); ok {
		r1 = rf(ctx, schemas)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockSqlDatabase_GetTableConstraintsBySchema_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetTableConstraintsBySchema'
type MockSqlDatabase_GetTableConstraintsBySchema_Call struct {
	*mock.Call
}

// GetTableConstraintsBySchema is a helper method to define mock.On call
//   - ctx context.Context
//   - schemas []string
func (_e *MockSqlDatabase_Expecter) GetTableConstraintsBySchema(ctx interface{}, schemas interface{}) *MockSqlDatabase_GetTableConstraintsBySchema_Call {
	return &MockSqlDatabase_GetTableConstraintsBySchema_Call{Call: _e.mock.On("GetTableConstraintsBySchema", ctx, schemas)}
}

func (_c *MockSqlDatabase_GetTableConstraintsBySchema_Call) Run(run func(ctx context.Context, schemas []string)) *MockSqlDatabase_GetTableConstraintsBySchema_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].([]string))
	})
	return _c
}

func (_c *MockSqlDatabase_GetTableConstraintsBySchema_Call) Return(_a0 *sqlmanager_shared.TableConstraints, _a1 error) *MockSqlDatabase_GetTableConstraintsBySchema_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockSqlDatabase_GetTableConstraintsBySchema_Call) RunAndReturn(run func(context.Context, []string) (*sqlmanager_shared.TableConstraints, error)) *MockSqlDatabase_GetTableConstraintsBySchema_Call {
	_c.Call.Return(run)
	return _c
}

// GetTableInitStatements provides a mock function with given fields: ctx, tables
func (_m *MockSqlDatabase) GetTableInitStatements(ctx context.Context, tables []*sqlmanager_shared.SchemaTable) ([]*sqlmanager_shared.TableInitStatement, error) {
	ret := _m.Called(ctx, tables)

	if len(ret) == 0 {
		panic("no return value specified for GetTableInitStatements")
	}

	var r0 []*sqlmanager_shared.TableInitStatement
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, []*sqlmanager_shared.SchemaTable) ([]*sqlmanager_shared.TableInitStatement, error)); ok {
		return rf(ctx, tables)
	}
	if rf, ok := ret.Get(0).(func(context.Context, []*sqlmanager_shared.SchemaTable) []*sqlmanager_shared.TableInitStatement); ok {
		r0 = rf(ctx, tables)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*sqlmanager_shared.TableInitStatement)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, []*sqlmanager_shared.SchemaTable) error); ok {
		r1 = rf(ctx, tables)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockSqlDatabase_GetTableInitStatements_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetTableInitStatements'
type MockSqlDatabase_GetTableInitStatements_Call struct {
	*mock.Call
}

// GetTableInitStatements is a helper method to define mock.On call
//   - ctx context.Context
//   - tables []*sqlmanager_shared.SchemaTable
func (_e *MockSqlDatabase_Expecter) GetTableInitStatements(ctx interface{}, tables interface{}) *MockSqlDatabase_GetTableInitStatements_Call {
	return &MockSqlDatabase_GetTableInitStatements_Call{Call: _e.mock.On("GetTableInitStatements", ctx, tables)}
}

func (_c *MockSqlDatabase_GetTableInitStatements_Call) Run(run func(ctx context.Context, tables []*sqlmanager_shared.SchemaTable)) *MockSqlDatabase_GetTableInitStatements_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].([]*sqlmanager_shared.SchemaTable))
	})
	return _c
}

func (_c *MockSqlDatabase_GetTableInitStatements_Call) Return(_a0 []*sqlmanager_shared.TableInitStatement, _a1 error) *MockSqlDatabase_GetTableInitStatements_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockSqlDatabase_GetTableInitStatements_Call) RunAndReturn(run func(context.Context, []*sqlmanager_shared.SchemaTable) ([]*sqlmanager_shared.TableInitStatement, error)) *MockSqlDatabase_GetTableInitStatements_Call {
	_c.Call.Return(run)
	return _c
}

// GetTableRowCount provides a mock function with given fields: ctx, schema, table, whereClause
func (_m *MockSqlDatabase) GetTableRowCount(ctx context.Context, schema string, table string, whereClause *string) (int64, error) {
	ret := _m.Called(ctx, schema, table, whereClause)

	if len(ret) == 0 {
		panic("no return value specified for GetTableRowCount")
	}

	var r0 int64
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, *string) (int64, error)); ok {
		return rf(ctx, schema, table, whereClause)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string, *string) int64); ok {
		r0 = rf(ctx, schema, table, whereClause)
	} else {
		r0 = ret.Get(0).(int64)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string, *string) error); ok {
		r1 = rf(ctx, schema, table, whereClause)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockSqlDatabase_GetTableRowCount_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetTableRowCount'
type MockSqlDatabase_GetTableRowCount_Call struct {
	*mock.Call
}

// GetTableRowCount is a helper method to define mock.On call
//   - ctx context.Context
//   - schema string
//   - table string
//   - whereClause *string
func (_e *MockSqlDatabase_Expecter) GetTableRowCount(ctx interface{}, schema interface{}, table interface{}, whereClause interface{}) *MockSqlDatabase_GetTableRowCount_Call {
	return &MockSqlDatabase_GetTableRowCount_Call{Call: _e.mock.On("GetTableRowCount", ctx, schema, table, whereClause)}
}

func (_c *MockSqlDatabase_GetTableRowCount_Call) Run(run func(ctx context.Context, schema string, table string, whereClause *string)) *MockSqlDatabase_GetTableRowCount_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string), args[3].(*string))
	})
	return _c
}

func (_c *MockSqlDatabase_GetTableRowCount_Call) Return(_a0 int64, _a1 error) *MockSqlDatabase_GetTableRowCount_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockSqlDatabase_GetTableRowCount_Call) RunAndReturn(run func(context.Context, string, string, *string) (int64, error)) *MockSqlDatabase_GetTableRowCount_Call {
	_c.Call.Return(run)
	return _c
}

// GetUniqueConstraintsMap provides a mock function with given fields: ctx, schemas
func (_m *MockSqlDatabase) GetUniqueConstraintsMap(ctx context.Context, schemas []string) (map[string][][]string, error) {
	ret := _m.Called(ctx, schemas)

	if len(ret) == 0 {
		panic("no return value specified for GetUniqueConstraintsMap")
	}

	var r0 map[string][][]string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, []string) (map[string][][]string, error)); ok {
		return rf(ctx, schemas)
	}
	if rf, ok := ret.Get(0).(func(context.Context, []string) map[string][][]string); ok {
		r0 = rf(ctx, schemas)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string][][]string)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, []string) error); ok {
		r1 = rf(ctx, schemas)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockSqlDatabase_GetUniqueConstraintsMap_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetUniqueConstraintsMap'
type MockSqlDatabase_GetUniqueConstraintsMap_Call struct {
	*mock.Call
}

// GetUniqueConstraintsMap is a helper method to define mock.On call
//   - ctx context.Context
//   - schemas []string
func (_e *MockSqlDatabase_Expecter) GetUniqueConstraintsMap(ctx interface{}, schemas interface{}) *MockSqlDatabase_GetUniqueConstraintsMap_Call {
	return &MockSqlDatabase_GetUniqueConstraintsMap_Call{Call: _e.mock.On("GetUniqueConstraintsMap", ctx, schemas)}
}

func (_c *MockSqlDatabase_GetUniqueConstraintsMap_Call) Run(run func(ctx context.Context, schemas []string)) *MockSqlDatabase_GetUniqueConstraintsMap_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].([]string))
	})
	return _c
}

func (_c *MockSqlDatabase_GetUniqueConstraintsMap_Call) Return(_a0 map[string][][]string, _a1 error) *MockSqlDatabase_GetUniqueConstraintsMap_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockSqlDatabase_GetUniqueConstraintsMap_Call) RunAndReturn(run func(context.Context, []string) (map[string][][]string, error)) *MockSqlDatabase_GetUniqueConstraintsMap_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockSqlDatabase creates a new instance of MockSqlDatabase. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockSqlDatabase(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockSqlDatabase {
	mock := &MockSqlDatabase{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
