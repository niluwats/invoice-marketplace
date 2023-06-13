package service

import "github.com/niluwats/invoice-marketplace/pkg/domain"

type InvoiceService interface {
	NewInvoice()
	GetInvoice()
	UpdateInvoiceLockStatus()
	UpdateInvoiceInvestor()
}

type DefaultInvoiceService struct {
	repo domain.InvoiceRepository
}

func NewInvoiceService(repo domain.InvoiceRepository) DefaultInvoiceService {
	return DefaultInvoiceService{repo}
}

func (s DefaultInvoiceService) NewInvoice() {

}

func (s DefaultInvoiceService) GetInvoice() {

}

func (s DefaultInvoiceService) UpdateInvoiceLockStatus() {

}

func (s DefaultInvoiceService) UpdateInvoiceInvestor() {

}
