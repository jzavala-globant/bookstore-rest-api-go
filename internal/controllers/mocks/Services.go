// Code generated by mockery v2.30.1. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	models "github.com/jzavala-globant/bookstore-rest-api-go/internal/models"
)

// Services is an autogenerated mock type for the Services type
type Services struct {
	mock.Mock
}

// ListBooks provides a mock function with given fields: _a0
func (_m *Services) ListBooks(_a0 context.Context) (*models.APIResponse, error) {
	ret := _m.Called(_a0)

	var r0 *models.APIResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (*models.APIResponse, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(context.Context) *models.APIResponse); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.APIResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewServices creates a new instance of Services. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewServices(t interface {
	mock.TestingT
	Cleanup(func())
}) *Services {
	mock := &Services{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
