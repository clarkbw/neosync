// Code generated by mockery. DO NOT EDIT.

package pg_queries

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// MockQuerier is an autogenerated mock type for the Querier type
type MockQuerier struct {
	mock.Mock
}

type MockQuerier_Expecter struct {
	mock *mock.Mock
}

func (_m *MockQuerier) EXPECT() *MockQuerier_Expecter {
	return &MockQuerier_Expecter{mock: &_m.Mock}
}

// GetDatabaseSchema provides a mock function with given fields: ctx, db
func (_m *MockQuerier) GetDatabaseSchema(ctx context.Context, db DBTX) ([]*GetDatabaseSchemaRow, error) {
	ret := _m.Called(ctx, db)

	if len(ret) == 0 {
		panic("no return value specified for GetDatabaseSchema")
	}

	var r0 []*GetDatabaseSchemaRow
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, DBTX) ([]*GetDatabaseSchemaRow, error)); ok {
		return rf(ctx, db)
	}
	if rf, ok := ret.Get(0).(func(context.Context, DBTX) []*GetDatabaseSchemaRow); ok {
		r0 = rf(ctx, db)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*GetDatabaseSchemaRow)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, DBTX) error); ok {
		r1 = rf(ctx, db)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockQuerier_GetDatabaseSchema_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetDatabaseSchema'
type MockQuerier_GetDatabaseSchema_Call struct {
	*mock.Call
}

// GetDatabaseSchema is a helper method to define mock.On call
//   - ctx context.Context
//   - db DBTX
func (_e *MockQuerier_Expecter) GetDatabaseSchema(ctx interface{}, db interface{}) *MockQuerier_GetDatabaseSchema_Call {
	return &MockQuerier_GetDatabaseSchema_Call{Call: _e.mock.On("GetDatabaseSchema", ctx, db)}
}

func (_c *MockQuerier_GetDatabaseSchema_Call) Run(run func(ctx context.Context, db DBTX)) *MockQuerier_GetDatabaseSchema_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(DBTX))
	})
	return _c
}

func (_c *MockQuerier_GetDatabaseSchema_Call) Return(_a0 []*GetDatabaseSchemaRow, _a1 error) *MockQuerier_GetDatabaseSchema_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockQuerier_GetDatabaseSchema_Call) RunAndReturn(run func(context.Context, DBTX) ([]*GetDatabaseSchemaRow, error)) *MockQuerier_GetDatabaseSchema_Call {
	_c.Call.Return(run)
	return _c
}

// GetDatabaseTableSchema provides a mock function with given fields: ctx, db, arg
func (_m *MockQuerier) GetDatabaseTableSchema(ctx context.Context, db DBTX, arg *GetDatabaseTableSchemaParams) ([]*GetDatabaseTableSchemaRow, error) {
	ret := _m.Called(ctx, db, arg)

	if len(ret) == 0 {
		panic("no return value specified for GetDatabaseTableSchema")
	}

	var r0 []*GetDatabaseTableSchemaRow
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, DBTX, *GetDatabaseTableSchemaParams) ([]*GetDatabaseTableSchemaRow, error)); ok {
		return rf(ctx, db, arg)
	}
	if rf, ok := ret.Get(0).(func(context.Context, DBTX, *GetDatabaseTableSchemaParams) []*GetDatabaseTableSchemaRow); ok {
		r0 = rf(ctx, db, arg)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*GetDatabaseTableSchemaRow)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, DBTX, *GetDatabaseTableSchemaParams) error); ok {
		r1 = rf(ctx, db, arg)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockQuerier_GetDatabaseTableSchema_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetDatabaseTableSchema'
type MockQuerier_GetDatabaseTableSchema_Call struct {
	*mock.Call
}

// GetDatabaseTableSchema is a helper method to define mock.On call
//   - ctx context.Context
//   - db DBTX
//   - arg *GetDatabaseTableSchemaParams
func (_e *MockQuerier_Expecter) GetDatabaseTableSchema(ctx interface{}, db interface{}, arg interface{}) *MockQuerier_GetDatabaseTableSchema_Call {
	return &MockQuerier_GetDatabaseTableSchema_Call{Call: _e.mock.On("GetDatabaseTableSchema", ctx, db, arg)}
}

func (_c *MockQuerier_GetDatabaseTableSchema_Call) Run(run func(ctx context.Context, db DBTX, arg *GetDatabaseTableSchemaParams)) *MockQuerier_GetDatabaseTableSchema_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(DBTX), args[2].(*GetDatabaseTableSchemaParams))
	})
	return _c
}

func (_c *MockQuerier_GetDatabaseTableSchema_Call) Return(_a0 []*GetDatabaseTableSchemaRow, _a1 error) *MockQuerier_GetDatabaseTableSchema_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockQuerier_GetDatabaseTableSchema_Call) RunAndReturn(run func(context.Context, DBTX, *GetDatabaseTableSchemaParams) ([]*GetDatabaseTableSchemaRow, error)) *MockQuerier_GetDatabaseTableSchema_Call {
	_c.Call.Return(run)
	return _c
}

// GetForeignKeyConstraints provides a mock function with given fields: ctx, db, tableschema
func (_m *MockQuerier) GetForeignKeyConstraints(ctx context.Context, db DBTX, tableschema string) ([]*GetForeignKeyConstraintsRow, error) {
	ret := _m.Called(ctx, db, tableschema)

	if len(ret) == 0 {
		panic("no return value specified for GetForeignKeyConstraints")
	}

	var r0 []*GetForeignKeyConstraintsRow
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, DBTX, string) ([]*GetForeignKeyConstraintsRow, error)); ok {
		return rf(ctx, db, tableschema)
	}
	if rf, ok := ret.Get(0).(func(context.Context, DBTX, string) []*GetForeignKeyConstraintsRow); ok {
		r0 = rf(ctx, db, tableschema)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*GetForeignKeyConstraintsRow)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, DBTX, string) error); ok {
		r1 = rf(ctx, db, tableschema)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockQuerier_GetForeignKeyConstraints_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetForeignKeyConstraints'
type MockQuerier_GetForeignKeyConstraints_Call struct {
	*mock.Call
}

// GetForeignKeyConstraints is a helper method to define mock.On call
//   - ctx context.Context
//   - db DBTX
//   - tableschema string
func (_e *MockQuerier_Expecter) GetForeignKeyConstraints(ctx interface{}, db interface{}, tableschema interface{}) *MockQuerier_GetForeignKeyConstraints_Call {
	return &MockQuerier_GetForeignKeyConstraints_Call{Call: _e.mock.On("GetForeignKeyConstraints", ctx, db, tableschema)}
}

func (_c *MockQuerier_GetForeignKeyConstraints_Call) Run(run func(ctx context.Context, db DBTX, tableschema string)) *MockQuerier_GetForeignKeyConstraints_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(DBTX), args[2].(string))
	})
	return _c
}

func (_c *MockQuerier_GetForeignKeyConstraints_Call) Return(_a0 []*GetForeignKeyConstraintsRow, _a1 error) *MockQuerier_GetForeignKeyConstraints_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockQuerier_GetForeignKeyConstraints_Call) RunAndReturn(run func(context.Context, DBTX, string) ([]*GetForeignKeyConstraintsRow, error)) *MockQuerier_GetForeignKeyConstraints_Call {
	_c.Call.Return(run)
	return _c
}

// GetPrimaryKeyConstraints provides a mock function with given fields: ctx, db, tableschema
func (_m *MockQuerier) GetPrimaryKeyConstraints(ctx context.Context, db DBTX, tableschema string) ([]*GetPrimaryKeyConstraintsRow, error) {
	ret := _m.Called(ctx, db, tableschema)

	if len(ret) == 0 {
		panic("no return value specified for GetPrimaryKeyConstraints")
	}

	var r0 []*GetPrimaryKeyConstraintsRow
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, DBTX, string) ([]*GetPrimaryKeyConstraintsRow, error)); ok {
		return rf(ctx, db, tableschema)
	}
	if rf, ok := ret.Get(0).(func(context.Context, DBTX, string) []*GetPrimaryKeyConstraintsRow); ok {
		r0 = rf(ctx, db, tableschema)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*GetPrimaryKeyConstraintsRow)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, DBTX, string) error); ok {
		r1 = rf(ctx, db, tableschema)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockQuerier_GetPrimaryKeyConstraints_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetPrimaryKeyConstraints'
type MockQuerier_GetPrimaryKeyConstraints_Call struct {
	*mock.Call
}

// GetPrimaryKeyConstraints is a helper method to define mock.On call
//   - ctx context.Context
//   - db DBTX
//   - tableschema string
func (_e *MockQuerier_Expecter) GetPrimaryKeyConstraints(ctx interface{}, db interface{}, tableschema interface{}) *MockQuerier_GetPrimaryKeyConstraints_Call {
	return &MockQuerier_GetPrimaryKeyConstraints_Call{Call: _e.mock.On("GetPrimaryKeyConstraints", ctx, db, tableschema)}
}

func (_c *MockQuerier_GetPrimaryKeyConstraints_Call) Run(run func(ctx context.Context, db DBTX, tableschema string)) *MockQuerier_GetPrimaryKeyConstraints_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(DBTX), args[2].(string))
	})
	return _c
}

func (_c *MockQuerier_GetPrimaryKeyConstraints_Call) Return(_a0 []*GetPrimaryKeyConstraintsRow, _a1 error) *MockQuerier_GetPrimaryKeyConstraints_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockQuerier_GetPrimaryKeyConstraints_Call) RunAndReturn(run func(context.Context, DBTX, string) ([]*GetPrimaryKeyConstraintsRow, error)) *MockQuerier_GetPrimaryKeyConstraints_Call {
	_c.Call.Return(run)
	return _c
}

// GetTableConstraints provides a mock function with given fields: ctx, db, arg
func (_m *MockQuerier) GetTableConstraints(ctx context.Context, db DBTX, arg *GetTableConstraintsParams) ([]*GetTableConstraintsRow, error) {
	ret := _m.Called(ctx, db, arg)

	if len(ret) == 0 {
		panic("no return value specified for GetTableConstraints")
	}

	var r0 []*GetTableConstraintsRow
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, DBTX, *GetTableConstraintsParams) ([]*GetTableConstraintsRow, error)); ok {
		return rf(ctx, db, arg)
	}
	if rf, ok := ret.Get(0).(func(context.Context, DBTX, *GetTableConstraintsParams) []*GetTableConstraintsRow); ok {
		r0 = rf(ctx, db, arg)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*GetTableConstraintsRow)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, DBTX, *GetTableConstraintsParams) error); ok {
		r1 = rf(ctx, db, arg)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockQuerier_GetTableConstraints_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetTableConstraints'
type MockQuerier_GetTableConstraints_Call struct {
	*mock.Call
}

// GetTableConstraints is a helper method to define mock.On call
//   - ctx context.Context
//   - db DBTX
//   - arg *GetTableConstraintsParams
func (_e *MockQuerier_Expecter) GetTableConstraints(ctx interface{}, db interface{}, arg interface{}) *MockQuerier_GetTableConstraints_Call {
	return &MockQuerier_GetTableConstraints_Call{Call: _e.mock.On("GetTableConstraints", ctx, db, arg)}
}

func (_c *MockQuerier_GetTableConstraints_Call) Run(run func(ctx context.Context, db DBTX, arg *GetTableConstraintsParams)) *MockQuerier_GetTableConstraints_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(DBTX), args[2].(*GetTableConstraintsParams))
	})
	return _c
}

func (_c *MockQuerier_GetTableConstraints_Call) Return(_a0 []*GetTableConstraintsRow, _a1 error) *MockQuerier_GetTableConstraints_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockQuerier_GetTableConstraints_Call) RunAndReturn(run func(context.Context, DBTX, *GetTableConstraintsParams) ([]*GetTableConstraintsRow, error)) *MockQuerier_GetTableConstraints_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockQuerier creates a new instance of MockQuerier. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockQuerier(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockQuerier {
	mock := &MockQuerier{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
