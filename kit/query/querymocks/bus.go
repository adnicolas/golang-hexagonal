// Code generated by mockery v2.33.0. DO NOT EDIT.

package querymocks

import (
	context "context"

	query "github.com/adnicolas/golang-hexagonal/kit/query"
	mock "github.com/stretchr/testify/mock"
)

// Bus is an autogenerated mock type for the Bus type
type Bus struct {
	mock.Mock
}

// DispatchQuery provides a mock function with given fields: _a0, _a1
func (_m *Bus) DispatchQuery(_a0 context.Context, _a1 query.Query) (query.QueryResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 query.QueryResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, query.Query) (query.QueryResponse, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, query.Query) query.QueryResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(query.QueryResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, query.Query) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RegisterQuery provides a mock function with given fields: _a0, _a1
func (_m *Bus) RegisterQuery(_a0 query.Type, _a1 query.QueryHandler) {
	_m.Called(_a0, _a1)
}

// NewBus creates a new instance of Bus. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewBus(t interface {
	mock.TestingT
	Cleanup(func())
}) *Bus {
	mock := &Bus{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
