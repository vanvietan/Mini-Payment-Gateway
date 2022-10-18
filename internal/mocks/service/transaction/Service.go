// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	context "context"
	model "pg/internal/model"

	mock "github.com/stretchr/testify/mock"
)

// Service is an autogenerated mock type for the Service type
type Service struct {
	mock.Mock
}

// CreateTransaction provides a mock function with given fields: ctx, cardID, orderID
func (_m *Service) CreateTransaction(ctx context.Context, cardID int64, orderID int64) (string, error) {
	ret := _m.Called(ctx, cardID, orderID)

	var r0 string
	if rf, ok := ret.Get(0).(func(context.Context, int64, int64) string); ok {
		r0 = rf(ctx, cardID, orderID)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int64, int64) error); ok {
		r1 = rf(ctx, cardID, orderID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteTransaction provides a mock function with given fields: ctx, transID
func (_m *Service) DeleteTransaction(ctx context.Context, transID int64) error {
	ret := _m.Called(ctx, transID)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int64) error); ok {
		r0 = rf(ctx, transID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FindTransactionByID provides a mock function with given fields: ctx, transID
func (_m *Service) FindTransactionByID(ctx context.Context, transID int64) (model.Transaction, error) {
	ret := _m.Called(ctx, transID)

	var r0 model.Transaction
	if rf, ok := ret.Get(0).(func(context.Context, int64) model.Transaction); ok {
		r0 = rf(ctx, transID)
	} else {
		r0 = ret.Get(0).(model.Transaction)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(ctx, transID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindTransactionByOTP provides a mock function with given fields: ctx, input
func (_m *Service) FindTransactionByOTP(ctx context.Context, input string) (model.Transaction, error) {
	ret := _m.Called(ctx, input)

	var r0 model.Transaction
	if rf, ok := ret.Get(0).(func(context.Context, string) model.Transaction); ok {
		r0 = rf(ctx, input)
	} else {
		r0 = ret.Get(0).(model.Transaction)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, input)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// InitAuthentication provides a mock function with given fields: ctx, inputCard, inputOrder
func (_m *Service) InitAuthentication(ctx context.Context, inputCard model.Card, inputOrder model.Order) (model.Card, model.Order, error) {
	ret := _m.Called(ctx, inputCard, inputOrder)

	var r0 model.Card
	if rf, ok := ret.Get(0).(func(context.Context, model.Card, model.Order) model.Card); ok {
		r0 = rf(ctx, inputCard, inputOrder)
	} else {
		r0 = ret.Get(0).(model.Card)
	}

	var r1 model.Order
	if rf, ok := ret.Get(1).(func(context.Context, model.Card, model.Order) model.Order); ok {
		r1 = rf(ctx, inputCard, inputOrder)
	} else {
		r1 = ret.Get(1).(model.Order)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, model.Card, model.Order) error); ok {
		r2 = rf(ctx, inputCard, inputOrder)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

type mockConstructorTestingTNewService interface {
	mock.TestingT
	Cleanup(func())
}

// NewService creates a new instance of Service. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewService(t mockConstructorTestingTNewService) *Service {
	mock := &Service{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
