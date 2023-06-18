// Code generated by mockery v2.30.1. DO NOT EDIT.

package mocks

import (
	domain "github.com/niluwats/invoice-marketplace/internal/domain"
	errors "github.com/niluwats/invoice-marketplace/pkg/errors"

	mock "github.com/stretchr/testify/mock"
)

// BidRepository is an autogenerated mock type for the BidRepository type
type BidRepository struct {
	mock.Mock
}

// GetAll provides a mock function with given fields: invoiceId
func (_m *BidRepository) GetAll(invoiceId int) ([]domain.Bid, *errors.AppError) {
	ret := _m.Called(invoiceId)

	var r0 []domain.Bid
	var r1 *errors.AppError
	if rf, ok := ret.Get(0).(func(int) ([]domain.Bid, *errors.AppError)); ok {
		return rf(invoiceId)
	}
	if rf, ok := ret.Get(0).(func(int) []domain.Bid); ok {
		r0 = rf(invoiceId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.Bid)
		}
	}

	if rf, ok := ret.Get(1).(func(int) *errors.AppError); ok {
		r1 = rf(invoiceId)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*errors.AppError)
		}
	}

	return r0, r1
}

// ProcessApproveBid provides a mock function with given fields: invoiceid, issuerid, amount
func (_m *BidRepository) ProcessApproveBid(invoiceid int, issuerid int, amount float64) *errors.AppError {
	ret := _m.Called(invoiceid, issuerid, amount)

	var r0 *errors.AppError
	if rf, ok := ret.Get(0).(func(int, int, float64) *errors.AppError); ok {
		r0 = rf(invoiceid, issuerid, amount)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*errors.AppError)
		}
	}

	return r0
}

// ProcessBid provides a mock function with given fields: bid, restBalance
func (_m *BidRepository) ProcessBid(bid domain.Bid, restBalance float64) *errors.AppError {
	ret := _m.Called(bid, restBalance)

	var r0 *errors.AppError
	if rf, ok := ret.Get(0).(func(domain.Bid, float64) *errors.AppError); ok {
		r0 = rf(bid, restBalance)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*errors.AppError)
		}
	}

	return r0
}

// NewBidRepository creates a new instance of BidRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewBidRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *BidRepository {
	mock := &BidRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
