// Code generated by mockery. DO NOT EDIT.

package clientcredtokenprovider

import (
	context "context"

	auth_client "github.com/nucleuscloud/neosync/backend/internal/auth/client"

	mock "github.com/stretchr/testify/mock"
)

// MocktokenProvider is an autogenerated mock type for the tokenProvider type
type MocktokenProvider struct {
	mock.Mock
}

type MocktokenProvider_Expecter struct {
	mock *mock.Mock
}

func (_m *MocktokenProvider) EXPECT() *MocktokenProvider_Expecter {
	return &MocktokenProvider_Expecter{mock: &_m.Mock}
}

// GetToken provides a mock function with given fields: _a0
func (_m *MocktokenProvider) GetToken(_a0 context.Context) (*auth_client.AuthTokenResponse, error) {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for GetToken")
	}

	var r0 *auth_client.AuthTokenResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (*auth_client.AuthTokenResponse, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(context.Context) *auth_client.AuthTokenResponse); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*auth_client.AuthTokenResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MocktokenProvider_GetToken_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetToken'
type MocktokenProvider_GetToken_Call struct {
	*mock.Call
}

// GetToken is a helper method to define mock.On call
//   - _a0 context.Context
func (_e *MocktokenProvider_Expecter) GetToken(_a0 interface{}) *MocktokenProvider_GetToken_Call {
	return &MocktokenProvider_GetToken_Call{Call: _e.mock.On("GetToken", _a0)}
}

func (_c *MocktokenProvider_GetToken_Call) Run(run func(_a0 context.Context)) *MocktokenProvider_GetToken_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *MocktokenProvider_GetToken_Call) Return(_a0 *auth_client.AuthTokenResponse, _a1 error) *MocktokenProvider_GetToken_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MocktokenProvider_GetToken_Call) RunAndReturn(run func(context.Context) (*auth_client.AuthTokenResponse, error)) *MocktokenProvider_GetToken_Call {
	_c.Call.Return(run)
	return _c
}

// NewMocktokenProvider creates a new instance of MocktokenProvider. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMocktokenProvider(t interface {
	mock.TestingT
	Cleanup(func())
}) *MocktokenProvider {
	mock := &MocktokenProvider{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
