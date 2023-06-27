package service

import (
	"context"
	"testing"
	"time"

	"github.com/niluwats/invoice-marketplace/internal/domain"
	"github.com/niluwats/invoice-marketplace/internal/dto"
	mocks "github.com/niluwats/invoice-marketplace/internal/mocks/repos"
	"github.com/stretchr/testify/assert"
	mock "github.com/stretchr/testify/mock"
)

func TestBidService_PlaceBid(t *testing.T) {
	bidRepo := &mocks.BidRepository{}
	investorRepo := &mocks.InvestorRepository{}
	invoiceRepo := &mocks.InvoiceRepository{}

	expectedInvoice := &domain.Invoice{
		ID:            40,
		InvoiceNumber: "RF-003",
		AmountDue:     5000,
		AskingPrice:   5000,
		IsLocked:      false,
		IsTraded:      false,
		IssuerId:      3,
	}

	expectedInvestor := &domain.Investor{
		ID:        "1",
		FirstName: "Jane",
		LastName:  "Daves",
		Balance:   8000,
		IsIssuer:  true,
	}

	expectedResponse := &domain.Bid{
		ID:         1,
		BidAmount:  1000,
		IsApproved: false,
		CreatedAt:  time.Date(2023, time.June, 10, 10, 20, 40, 45, time.Local),
		InvestorId: 4,
		InvoiceId:  1,
	}

	bidRepo.On("ProcessBid", mock.Anything, mock.AnythingOfType("domain.Bid"), mock.Anything).Return(expectedResponse, nil).Once()
	invoiceRepo.On("FindById", mock.Anything, mock.Anything).Return(expectedInvoice, nil).Once()
	invoiceRepo.On("FindTotalInvestment", mock.Anything, mock.Anything).Return(float64(0), nil).Once()
	investorRepo.On("FindById", mock.Anything, mock.Anything).Return(expectedInvestor, nil).Once()

	bidRequest := dto.BidRequest{
		InvoiceId:  41,
		BidAmount:  1000,
		InvestorId: 2,
	}

	service := NewBidService(bidRepo, investorRepo, invoiceRepo)

	bid, err := service.PlaceBid(context.Background(), bidRequest)
	assert.Nil(t, err)
	assert.Equal(t, expectedResponse, bid)
}

func TestBidService_ApproveTrade(t *testing.T) {
	repo := &mocks.BidRepository{}
	investorRepo := &mocks.InvestorRepository{}
	invoiceRepo := &mocks.InvoiceRepository{}

	expectedInvoice := &domain.Invoice{
		ID:            40,
		InvoiceNumber: "RF-003",
		AmountDue:     5000,
		AskingPrice:   5000,
		IsLocked:      true,
		IsTraded:      false,
		IssuerId:      3,
	}

	id := "40"
	repo.On("ProcessApproveBid", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil).Once()
	invoiceRepo.On("FindById", mock.Anything, mock.Anything).Return(expectedInvoice, nil).Once()

	service := NewBidService(repo, investorRepo, invoiceRepo)
	err := service.ApproveTrade(context.Background(), id)
	assert.Nil(t, err)
}

func TestBidService_RejectTrade(t *testing.T) {
	expectedInvoice := &domain.Invoice{
		ID:            1,
		InvoiceNumber: "RF-001",
		AmountDue:     5000,
		AskingPrice:   5000,
		IsLocked:      true,
		IsTraded:      false,
		IssuerId:      1,
	}

	bidRepo := &mocks.BidRepository{}
	investorRepo := &mocks.InvestorRepository{}
	invoiceRepo := &mocks.InvoiceRepository{}

	id := "1"
	bidRepo.On("ProcessCancelBid", mock.Anything, mock.Anything).Return(nil).Once()
	invoiceRepo.On("FindById", mock.Anything, mock.Anything).Return(expectedInvoice, nil).Once()

	service := NewBidService(bidRepo, investorRepo, invoiceRepo)
	err := service.RejectTrade(context.Background(), id)
	assert.Nil(t, err)
}

func TestBidService_GetAllBids(t *testing.T) {
	repo := &mocks.BidRepository{}
	investorRepo := &mocks.InvestorRepository{}
	invoiceRepo := &mocks.InvoiceRepository{}

	id := "41"

	expectedInvoice := []domain.Bid{
		{
			ID:         1,
			BidAmount:  1000,
			IsApproved: false,
			InvestorId: 2,
			InvoiceId:  41,
			CreatedAt:  time.Date(2023, time.June, 10, 10, 20, 40, 45, time.Local),
		},
		{
			ID:         2,
			BidAmount:  2000,
			IsApproved: false,
			InvestorId: 3,
			InvoiceId:  41,
			CreatedAt:  time.Date(2023, time.June, 10, 11, 20, 40, 45, time.Local),
		},
	}
	repo.On("GetAll", mock.Anything, mock.Anything).Return(expectedInvoice, nil).Once()

	service := NewBidService(repo, investorRepo, invoiceRepo)
	invoice, err := service.GetAllBids(context.Background(), id)
	assert.Nil(t, err)
	assert.Equal(t, expectedInvoice, invoice)
}
