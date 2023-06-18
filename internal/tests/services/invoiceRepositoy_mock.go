// Code generated by mockery v2.30.1. DO NOT EDIT.

package mocks

import (
	domain "github.com/niluwats/invoice-marketplace/internal/domain"
	errors "github.com/niluwats/invoice-marketplace/pkg/errors"

	mock "github.com/stretchr/testify/mock"
)

// InvoiceRepository is an autogenerated mock type for the InvoiceRepository type
type InvoiceRepository struct {
	mock.Mock
}

// FindById provides a mock function with given fields: id
func (_m *InvoiceRepository) FindById(id int) (*domain.Invoice, *errors.AppError) {
	ret := _m.Called(id)

	var r0 *domain.Invoice
	var r1 *errors.AppError
	if rf, ok := ret.Get(0).(func(int) (*domain.Invoice, *errors.AppError)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(int) *domain.Invoice); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Invoice)
		}
	}

	if rf, ok := ret.Get(1).(func(int) *errors.AppError); ok {
		r1 = rf(id)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*errors.AppError)
		}
	}

	return r0, r1
}

// FindTotalInvestment provides a mock function with given fields: id
func (_m *InvoiceRepository) FindTotalInvestment(id int) (float64, *errors.AppError) {
	ret := _m.Called(id)

	var r0 float64
	var r1 *errors.AppError
	if rf, ok := ret.Get(0).(func(int) (float64, *errors.AppError)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(int) float64); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(float64)
	}

	if rf, ok := ret.Get(1).(func(int) *errors.AppError); ok {
		r1 = rf(id)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*errors.AppError)
		}
	}

	return r0, r1
}

// Insert provides a mock function with given fields: invoice
func (_m *InvoiceRepository) Insert(invoice domain.Invoice) *errors.AppError {
	ret := _m.Called(invoice)

	var r0 *errors.AppError
	if rf, ok := ret.Get(0).(func(domain.Invoice) *errors.AppError); ok {
		r0 = rf(invoice)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*errors.AppError)
		}
	}

	return r0
}

// NewInvoiceRepository creates a new instance of InvoiceRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewInvoiceRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *InvoiceRepository {
	mock := &InvoiceRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
