// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	context "context"
	model "pg/api/internal/model"

	mock "github.com/stretchr/testify/mock"
)

// Service is an autogenerated mock type for the Service type
type Service struct {
	mock.Mock
}

// AddCard provides a mock function with given fields: ctx, input
func (_m *Service) AddCard(ctx context.Context, input model.Card) (model.Card, error) {
	ret := _m.Called(ctx, input)

	var r0 model.Card
	if rf, ok := ret.Get(0).(func(context.Context, model.Card) model.Card); ok {
		r0 = rf(ctx, input)
	} else {
		r0 = ret.Get(0).(model.Card)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, model.Card) error); ok {
		r1 = rf(ctx, input)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeductCard provides a mock function with given fields: ctx, id, amount
func (_m *Service) DeductCard(ctx context.Context, id int64, amount int64) (model.Card, error) {
	ret := _m.Called(ctx, id, amount)

	var r0 model.Card
	if rf, ok := ret.Get(0).(func(context.Context, int64, int64) model.Card); ok {
		r0 = rf(ctx, id, amount)
	} else {
		r0 = ret.Get(0).(model.Card)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int64, int64) error); ok {
		r1 = rf(ctx, id, amount)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteCard provides a mock function with given fields: ctx, cardID
func (_m *Service) DeleteCard(ctx context.Context, cardID int64) error {
	ret := _m.Called(ctx, cardID)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int64) error); ok {
		r0 = rf(ctx, cardID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetCardByID provides a mock function with given fields: ctx, cardID
func (_m *Service) GetCardByID(ctx context.Context, cardID int64) (model.Card, error) {
	ret := _m.Called(ctx, cardID)

	var r0 model.Card
	if rf, ok := ret.Get(0).(func(context.Context, int64) model.Card); ok {
		r0 = rf(ctx, cardID)
	} else {
		r0 = ret.Get(0).(model.Card)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(ctx, cardID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetCardByNumber provides a mock function with given fields: ctx, cardNumber
func (_m *Service) GetCardByNumber(ctx context.Context, cardNumber string) (model.Card, error) {
	ret := _m.Called(ctx, cardNumber)

	var r0 model.Card
	if rf, ok := ret.Get(0).(func(context.Context, string) model.Card); ok {
		r0 = rf(ctx, cardNumber)
	} else {
		r0 = ret.Get(0).(model.Card)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, cardNumber)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetCards provides a mock function with given fields: ctx, limit, lastID
func (_m *Service) GetCards(ctx context.Context, limit int, lastID int64) ([]model.Card, error) {
	ret := _m.Called(ctx, limit, lastID)

	var r0 []model.Card
	if rf, ok := ret.Get(0).(func(context.Context, int, int64) []model.Card); ok {
		r0 = rf(ctx, limit, lastID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.Card)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int, int64) error); ok {
		r1 = rf(ctx, limit, lastID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateCard provides a mock function with given fields: ctx, input, cardID
func (_m *Service) UpdateCard(ctx context.Context, input model.Card, cardID int64) (model.Card, error) {
	ret := _m.Called(ctx, input, cardID)

	var r0 model.Card
	if rf, ok := ret.Get(0).(func(context.Context, model.Card, int64) model.Card); ok {
		r0 = rf(ctx, input, cardID)
	} else {
		r0 = ret.Get(0).(model.Card)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, model.Card, int64) error); ok {
		r1 = rf(ctx, input, cardID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
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
