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
	UpdateInvoiceLockStatus(id int) error
	UpdateInvoiceInvestor(id, investorId int) error
}

type DefaultInvoiceService struct {
	repo repositories.InvoiceRepository
}

func NewInvoiceService(repo repositories.InvoiceRepository) DefaultInvoiceService {
	return DefaultInvoiceService{repo}
}

func (s DefaultInvoiceService) NewInvoice(invRequest dto.InvoiceRequest) error {
	invoice := domain.Invoice{
		InvoiceNumber:  invRequest.InvoiceNumber,
		CreatedOn:      time.Now(),
		AmountDue:      invRequest.AmountDue,
		AmountEnclosed: invRequest.AmountEnclosed,
		IssuerId:       invRequest.IssuerId,
	}
	err := s.repo.Insert(invoice)
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

func (s DefaultInvoiceService) UpdateInvoiceLockStatus(id int) error {
	err := s.repo.UpdateLockStatus(id)
	if err != nil {
		return err
	}
	return nil
}

func (s DefaultInvoiceService) UpdateInvoiceInvestor(id, investorId int) error {
	err := s.repo.UpdateInvoiceInvestor(id, investorId)
	if err != nil {
		return err
	}
	return nil
}
