// Code generated by mockery v2.39.1. DO NOT EDIT.

package csvrepository_mocks

import (
	csvrepository "github.com/dstolbert/directory-generator/internal/repository/csvrepository"
	mock "github.com/stretchr/testify/mock"
)

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

type Repository_Expecter struct {
	mock *mock.Mock
}

func (_m *Repository) EXPECT() *Repository_Expecter {
	return &Repository_Expecter{mock: &_m.Mock}
}

// Get provides a mock function with given fields:
func (_m *Repository) Get() []csvrepository.Entry {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Get")
	}

	var r0 []csvrepository.Entry
	if rf, ok := ret.Get(0).(func() []csvrepository.Entry); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]csvrepository.Entry)
		}
	}

	return r0
}

// Repository_Get_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Get'
type Repository_Get_Call struct {
	*mock.Call
}

// Get is a helper method to define mock.On call
func (_e *Repository_Expecter) Get() *Repository_Get_Call {
	return &Repository_Get_Call{Call: _e.mock.On("Get")}
}

func (_c *Repository_Get_Call) Run(run func()) *Repository_Get_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Repository_Get_Call) Return(_a0 []csvrepository.Entry) *Repository_Get_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Repository_Get_Call) RunAndReturn(run func() []csvrepository.Entry) *Repository_Get_Call {
	_c.Call.Return(run)
	return _c
}

// NewRepository creates a new instance of Repository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *Repository {
	mock := &Repository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

// Code below was generated by components. DO NOT EDIT.
// Component version: v0.1.0

type Repository_ExpecterChain[M any] func(*M) *Repository_Expecter

func Create_Repository_ExpecterChain[M any](fetch func(*M) *Repository) Repository_ExpecterChain[M] {
	return func(m *M) *Repository_Expecter {
		c := fetch(m)
		return c.EXPECT()
	}
}

type Repository_GetChain[M any] func(*M) *Repository_Get_Call

func (_c Repository_ExpecterChain[M]) Get() Repository_GetChain[M] {
	return func(m *M) *Repository_Get_Call {
		expecter := _c(m)
		return expecter.Get()
	}
}

func (_c Repository_GetChain[M]) Run(run func()) Repository_GetChain[M] {
	return func(m *M) *Repository_Get_Call {
		call := _c(m)
		return call.Run(run)
	}
}

func (_c Repository_GetChain[M]) Return(_a0 []Entry) Repository_GetChain[M] {
	return func(m *M) *Repository_Get_Call {
		call := _c(m)
		return call.Return(_a0)
	}
}

func (_c Repository_GetChain[M]) RunAndReturn(run func() []Entry) Repository_GetChain[M] {
	return func(m *M) *Repository_Get_Call {
		call := _c(m)
		return call.RunAndReturn(run)
	}
}

func (_c Repository_ExpecterChain[M]) Get_Pointer() Repository_GetChain[M] {
	return func(m *M) *Repository_Get_Call {
		expecter := _c(m)
		return expecter.Get()
	}
}

func (_c Repository_GetChain[M]) Return_Pointer(_a0 *[]Entry) Repository_GetChain[M] {
	return func(m *M) *Repository_Get_Call {
		call := _c(m)
		return call.Return(*_a0)
	}
}
