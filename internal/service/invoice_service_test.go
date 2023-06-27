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

func TestInvoiceService_NewInvoice(t *testing.T) {
	repo := &mocks.InvoiceRepository{}
	expectedResponse := &domain.Invoice{
		InvoiceNumber: "RF-001",
		IssuerId:      1,
		AmountDue:     5000,
		AskingPrice:   5000,
		DueDate:       time.Date(2023, time.June, 30, 11, 20, 40, 45, time.Local),
	}
	repo.On("FindIfExistsByNo", mock.Anything, mock.Anything).Return(false).Once()
	repo.On("Insert", mock.Anything, mock.AnythingOfType("domain.Invoice")).Return(expectedResponse, nil).Once()

	invoiceRequest := dto.InvoiceRequest{
		InvoiceNumber: "RF-001",
		IssuerId:      1,
		AmountDue:     5000,
		AskingPrice:   5000,
		DueDate:       "2023-06-30",
	}

	service := NewInvoiceService(repo)

	invoiceCreated, err := service.NewInvoice(context.Background(), invoiceRequest)
	assert.Nil(t, err)
	assert.Equal(t, expectedResponse, invoiceCreated)
}

func TestInvoiceService_GetInvoice(t *testing.T) {
	repo := &mocks.InvoiceRepository{}

	id := "40"

	expectedInvoice := &domain.Invoice{
		ID:            40,
		InvoiceNumber: "RF-003",
		AmountDue:     5000,
		AskingPrice:   5000,
		IsLocked:      true,
		IsTraded:      true,
		IssuerId:      3,
	}
	repo.On("FindById", mock.Anything, mock.Anything).Return(expectedInvoice, nil).Once()

	service := NewInvoiceService(repo)
	invoice, err := service.GetInvoice(context.Background(), id)
	assert.Nil(t, err)
	assert.Equal(t, expectedInvoice, invoice)
}

func TestInvoiceService_GetAllInvoices(t *testing.T) {
	repo := &mocks.InvoiceRepository{}

	expectedInvoices := []domain.Invoice{
		{
			ID:            1,
			InvoiceNumber: "RF-001",
			AmountDue:     5000,
			AskingPrice:   5000,
			IsLocked:      true,
			IsTraded:      true,
			IssuerId:      1,
			CreatedOn:     time.Date(2023, time.June, 12, 0, 0, 0, 0, time.Local),
		},
		{
			ID:            2,
			InvoiceNumber: "RF-002",
			AmountDue:     5000,
			AskingPrice:   5000,
			IsLocked:      true,
			IsTraded:      true,
			IssuerId:      1,
			CreatedOn:     time.Date(2023, time.June, 11, 0, 0, 0, 0, time.Local),
		},
	}

	repo.On("FindAll", mock.Anything).Return(expectedInvoices, nil).Once()
	service := NewInvoiceService(repo)
	invoices, err := service.GetAllInvoices(context.Background())
	assert.Nil(t, err)
	assert.Equal(t, expectedInvoices, invoices)
}