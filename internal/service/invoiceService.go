package service

import (
	"strconv"
	"time"

	"github.com/niluwats/invoice-marketplace/internal/domain"
	"github.com/niluwats/invoice-marketplace/internal/dto"
	"github.com/niluwats/invoice-marketplace/internal/repositories"
)

type InvoiceService interface {
	NewInvoice(invRequest dto.InvoiceRequest) error
	GetInvoice(id string) (*domain.Invoice, error)
}

type DefaultInvoiceService struct {
	repo repositories.InvoiceRepository
}

func NewInvoiceService(repo repositories.InvoiceRepository) DefaultInvoiceService {
	return DefaultInvoiceService{repo}
}

func (s DefaultInvoiceService) NewInvoice(invRequest dto.InvoiceRequest) error {
	layout := "2006-01-02"
	dueDate, err := time.Parse(layout, invRequest.DueDate)
	if err != nil {
		return err
	}

	invoice := domain.Invoice{
		InvoiceNumber: invRequest.InvoiceNumber,
		CreatedOn:     time.Now(),
		DueDate:       dueDate,
		AmountDue:     invRequest.AmountDue,
		AskingPrice:   invRequest.AskingPrice,
		IssuerId:      invRequest.IssuerId,
	}
	err = s.repo.Insert(invoice)
	if err != nil {
		return err
	}
	return nil
}

func (s DefaultInvoiceService) GetInvoice(id string) (*domain.Invoice, error) {
	invoiceId, _ := strconv.Atoi(id)
	invoice, err := s.repo.FindById(invoiceId)
	if err != nil {
		return nil, err
	}
	return invoice, nil
}
