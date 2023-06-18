package service

import (
	"strconv"
	"time"

	"github.com/niluwats/invoice-marketplace/internal/domain"
	"github.com/niluwats/invoice-marketplace/internal/dto"
	"github.com/niluwats/invoice-marketplace/internal/repositories"
	appErr "github.com/niluwats/invoice-marketplace/pkg/errors"
)

type InvoiceService interface {
	NewInvoice(invRequest dto.InvoiceRequest) *appErr.AppError
	GetInvoice(id string) (*domain.Invoice, *appErr.AppError)
}

type DefaultInvoiceService struct {
	repo repositories.InvoiceRepository
}

func NewInvoiceService(repo repositories.InvoiceRepository) DefaultInvoiceService {
	return DefaultInvoiceService{repo}
}

func (s DefaultInvoiceService) NewInvoice(invRequest dto.InvoiceRequest) *appErr.AppError {
	layout := "2006-01-02"
	dueDate, err := time.Parse(layout, invRequest.DueDate)
	if err != nil {
		return appErr.NewUnexpectedError("Error parsing time format : " + err.Error())
	}

	if invRequest.IfInValidRequest() {
		return appErr.NewBadRequest("All fields required")
	}

	invoice := domain.Invoice{
		InvoiceNumber: invRequest.InvoiceNumber,
		CreatedOn:     time.Now(),
		DueDate:       dueDate,
		AmountDue:     invRequest.AmountDue,
		AskingPrice:   invRequest.AskingPrice,
		IssuerId:      invRequest.IssuerId,
	}
	err_ := s.repo.Insert(invoice)
	if err_ != nil {
		return err_
	}
	return nil
}

func (s DefaultInvoiceService) GetInvoice(id string) (*domain.Invoice, *appErr.AppError) {
	invoiceId, _ := strconv.Atoi(id)
	invoice, err_ := s.repo.FindById(invoiceId)
	if err_ != nil {
		return nil, err_
	}
	return invoice, nil
}
