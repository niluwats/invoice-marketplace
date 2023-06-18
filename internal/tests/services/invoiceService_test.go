package mocks

import (
	"testing"

	"github.com/niluwats/invoice-marketplace/internal/domain"
	"github.com/niluwats/invoice-marketplace/internal/dto"
	"github.com/niluwats/invoice-marketplace/internal/service"
	"github.com/stretchr/testify/assert"
	mock "github.com/stretchr/testify/mock"
)

func TestInvoiceService_NewInvoice(t *testing.T) {
	repo := &InvoiceRepository{}
	repo.On("Insert", mock.AnythingOfType("domain.Invoice")).Return(nil).Once()

	invoiceRequest := dto.InvoiceRequest{
		InvoiceNumber: "RF-001",
		IssuerId:      1,
		AmountDue:     5000,
		AskingPrice:   5000,
		DueDate:       "2023-06-30",
	}

	service := service.NewInvoiceService(repo)

	err := service.NewInvoice(invoiceRequest)
	assert.Nil(t, err)
}

func TestInvoiceService_GetInvoice(t *testing.T) {
	repo := &InvoiceRepository{}

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
	repo.On("FindById", mock.Anything).Return(expectedInvoice, nil).Once()

	service := service.NewInvoiceService(repo)
	invoice, err := service.GetInvoice(id)
	assert.Nil(t, err)
	assert.Equal(t, expectedInvoice, invoice)
}
