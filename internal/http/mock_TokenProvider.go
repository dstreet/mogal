// Code generated by mockery v2.32.4. DO NOT EDIT.

package http

import (
	time "time"

	mock "github.com/stretchr/testify/mock"

	user "github.com/dstreet/mogal/internal/user"
)

// MockTokenProvider is an autogenerated mock type for the TokenProvider type
type MockTokenProvider struct {
	mock.Mock
}

type MockTokenProvider_Expecter struct {
	mock *mock.Mock
}

func (_m *MockTokenProvider) EXPECT() *MockTokenProvider_Expecter {
	return &MockTokenProvider_Expecter{mock: &_m.Mock}
}

// CreateToken provides a mock function with given fields: u
func (_m *MockTokenProvider) CreateToken(u user.User) (string, error) {
	ret := _m.Called(u)

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(user.User) (string, error)); ok {
		return rf(u)
	}
	if rf, ok := ret.Get(0).(func(user.User) string); ok {
		r0 = rf(u)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(user.User) error); ok {
		r1 = rf(u)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockTokenProvider_CreateToken_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateToken'
type MockTokenProvider_CreateToken_Call struct {
	*mock.Call
}

// CreateToken is a helper method to define mock.On call
//   - u user.User
func (_e *MockTokenProvider_Expecter) CreateToken(u interface{}) *MockTokenProvider_CreateToken_Call {
	return &MockTokenProvider_CreateToken_Call{Call: _e.mock.On("CreateToken", u)}
}

func (_c *MockTokenProvider_CreateToken_Call) Run(run func(u user.User)) *MockTokenProvider_CreateToken_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(user.User))
	})
	return _c
}

func (_c *MockTokenProvider_CreateToken_Call) Return(_a0 string, _a1 error) *MockTokenProvider_CreateToken_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockTokenProvider_CreateToken_Call) RunAndReturn(run func(user.User) (string, error)) *MockTokenProvider_CreateToken_Call {
	_c.Call.Return(run)
	return _c
}

// TokenDuration provides a mock function with given fields:
func (_m *MockTokenProvider) TokenDuration() time.Duration {
	ret := _m.Called()

	var r0 time.Duration
	if rf, ok := ret.Get(0).(func() time.Duration); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(time.Duration)
	}

	return r0
}

// MockTokenProvider_TokenDuration_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'TokenDuration'
type MockTokenProvider_TokenDuration_Call struct {
	*mock.Call
}

// TokenDuration is a helper method to define mock.On call
func (_e *MockTokenProvider_Expecter) TokenDuration() *MockTokenProvider_TokenDuration_Call {
	return &MockTokenProvider_TokenDuration_Call{Call: _e.mock.On("TokenDuration")}
}

func (_c *MockTokenProvider_TokenDuration_Call) Run(run func()) *MockTokenProvider_TokenDuration_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockTokenProvider_TokenDuration_Call) Return(_a0 time.Duration) *MockTokenProvider_TokenDuration_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockTokenProvider_TokenDuration_Call) RunAndReturn(run func() time.Duration) *MockTokenProvider_TokenDuration_Call {
	_c.Call.Return(run)
	return _c
}

// VerifyToken provides a mock function with given fields: token
func (_m *MockTokenProvider) VerifyToken(token string) (string, error) {
	ret := _m.Called(token)

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (string, error)); ok {
		return rf(token)
	}
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(token)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(token)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockTokenProvider_VerifyToken_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'VerifyToken'
type MockTokenProvider_VerifyToken_Call struct {
	*mock.Call
}

// VerifyToken is a helper method to define mock.On call
//   - token string
func (_e *MockTokenProvider_Expecter) VerifyToken(token interface{}) *MockTokenProvider_VerifyToken_Call {
	return &MockTokenProvider_VerifyToken_Call{Call: _e.mock.On("VerifyToken", token)}
}

func (_c *MockTokenProvider_VerifyToken_Call) Run(run func(token string)) *MockTokenProvider_VerifyToken_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MockTokenProvider_VerifyToken_Call) Return(userID string, err error) *MockTokenProvider_VerifyToken_Call {
	_c.Call.Return(userID, err)
	return _c
}

func (_c *MockTokenProvider_VerifyToken_Call) RunAndReturn(run func(string) (string, error)) *MockTokenProvider_VerifyToken_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockTokenProvider creates a new instance of MockTokenProvider. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockTokenProvider(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockTokenProvider {
	mock := &MockTokenProvider{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}