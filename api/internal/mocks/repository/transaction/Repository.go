// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	context "context"
	model "pg/api/internal/model"

	mock "github.com/stretchr/testify/mock"
)

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

// CreateTransaction provides a mock function with given fields: ctx, _a1
func (_m *Repository) CreateTransaction(ctx context.Context, _a1 model.Transaction) (model.Transaction, error) {
	ret := _m.Called(ctx, _a1)

	var r0 model.Transaction
	if rf, ok := ret.Get(0).(func(context.Context, model.Transaction) model.Transaction); ok {
		r0 = rf(ctx, _a1)
	} else {
		r0 = ret.Get(0).(model.Transaction)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, model.Transaction) error); ok {
		r1 = rf(ctx, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteTransaction provides a mock function with given fields: ctx, transID
func (_m *Repository) DeleteTransaction(ctx context.Context, transID int64) error {
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
func (_m *Repository) FindTransactionByID(ctx context.Context, transID int64) (model.Transaction, error) {
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

// FindTransactionByOTP provides a mock function with given fields: ctx, otp
func (_m *Repository) FindTransactionByOTP(ctx context.Context, otp string) (model.Transaction, error) {
	ret := _m.Called(ctx, otp)

	var r0 model.Transaction
	if rf, ok := ret.Get(0).(func(context.Context, string) model.Transaction); ok {
		r0 = rf(ctx, otp)
	} else {
		r0 = ret.Get(0).(model.Transaction)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, otp)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateTransaction provides a mock function with given fields: ctx, input
func (_m *Repository) UpdateTransaction(ctx context.Context, input model.Transaction) (model.Transaction, error) {
	ret := _m.Called(ctx, input)

	var r0 model.Transaction
	if rf, ok := ret.Get(0).(func(context.Context, model.Transaction) model.Transaction); ok {
		r0 = rf(ctx, input)
	} else {
		r0 = ret.Get(0).(model.Transaction)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, model.Transaction) error); ok {
		r1 = rf(ctx, input)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewRepository creates a new instance of Repository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewRepository(t mockConstructorTestingTNewRepository) *Repository {
	mock := &Repository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
