package mocks

import (
	"testing"
	"time"

	"github.com/niluwats/invoice-marketplace/internal/domain"
	"github.com/niluwats/invoice-marketplace/internal/dto"
	"github.com/niluwats/invoice-marketplace/internal/service"
	"github.com/stretchr/testify/assert"
	mock "github.com/stretchr/testify/mock"
)

func TestBidService_PlaceBid(t *testing.T) {
	bidRepo := &BidRepository{}
	investorRepo := &InvestorRepository{}
	invoiceRepo := &InvoiceRepository{}

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

	bidRepo.On("ProcessBid", mock.AnythingOfType("domain.Bid"), mock.Anything).Return(nil).Once()
	invoiceRepo.On("FindById", mock.Anything).Return(expectedInvoice, nil).Once()
	invoiceRepo.On("FindTotalInvestment", mock.Anything).Return(float64(0), nil).Once()
	investorRepo.On("FindById", mock.Anything).Return(expectedInvestor, nil).Once()

	bidRequest := dto.BidRequest{
		InvoiceId:  41,
		BidAmount:  1000,
		InvestorId: 2,
	}

	service := service.NewBidService(bidRepo, investorRepo, invoiceRepo)

	err := service.PlaceBid(bidRequest)
	assert.Nil(t, err)
}

func TestBidService_ApproveTrade(t *testing.T) {
	repo := &BidRepository{}
	investorRepo := &InvestorRepository{}
	invoiceRepo := &InvoiceRepository{}

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
	repo.On("ProcessApproveBid", mock.Anything, mock.Anything, mock.Anything).Return(nil).Once()
	invoiceRepo.On("FindById", mock.Anything).Return(expectedInvoice, nil).Once()

	service := service.NewBidService(repo, investorRepo, invoiceRepo)
	err := service.ApproveTrade(id)
	assert.Nil(t, err)
}

func TestBidService_GetAllBids(t *testing.T) {
	repo := &BidRepository{}
	investorRepo := &InvestorRepository{}
	invoiceRepo := &InvoiceRepository{}

	id := "41"

	expectedInvoice := []domain.Bid{
		{
			ID:         1,
			BidAmount:  1000,
			IsApproved: false,
			InvestorId: 2,
			InvoiceId:  41,
			TimeStamp:  time.Date(2023, time.June, 10, 10, 20, 40, 45, time.Local),
		},
		{
			ID:         2,
			BidAmount:  2000,
			IsApproved: false,
			InvestorId: 3,
			InvoiceId:  41,
			TimeStamp:  time.Date(2023, time.June, 10, 11, 20, 40, 45, time.Local),
		},
	}
	repo.On("GetAll", mock.Anything).Return(expectedInvoice, nil).Once()

	service := service.NewBidService(repo, investorRepo, invoiceRepo)
	invoice, err := service.GetAllBids(id)
	assert.Nil(t, err)
	assert.Equal(t, expectedInvoice, invoice)
}
