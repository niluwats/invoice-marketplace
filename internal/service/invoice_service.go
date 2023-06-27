package service

import (
	"context"
	"strconv"
	"time"

	"github.com/niluwats/invoice-marketplace/internal/domain"
	"github.com/niluwats/invoice-marketplace/internal/dto"
	"github.com/niluwats/invoice-marketplace/internal/repositories"
	appErr "github.com/niluwats/invoice-marketplace/pkg/errors"
)

type InvoiceService interface {
	NewInvoice(ctx context.Context, invRequest dto.InvoiceRequest) (*domain.Invoice, *appErr.AppError)
	GetInvoice(ctx context.Context, id string) (*domain.Invoice, *appErr.AppError)
	GetAllInvoices(ctx context.Context) ([]domain.Invoice, *appErr.AppError)
}

type DefaultInvoiceService struct {
	repo repositories.InvoiceRepository
}

func NewInvoiceService(repo repositories.InvoiceRepository) DefaultInvoiceService {
	return DefaultInvoiceService{repo}
}

func (s DefaultInvoiceService) NewInvoice(ctx context.Context, invRequest dto.InvoiceRequest) (*domain.Invoice, *appErr.AppError) {
	dueDate, err := time.Parse("2006-01-02", invRequest.DueDate)
	if err != nil {
		return nil, appErr.NewUnexpectedError("Error parsing time format : " + err.Error())
	}

	if invRequest.IfInValidRequest() {
		return nil, appErr.NewBadRequest("All fields required")
	}

	if s.repo.FindIfExistsByNo(&ctx, invRequest.InvoiceNumber) {
		return nil, appErr.NewConflictError("Invoice number already exists")
	}

	invoice := domain.Invoice{
		InvoiceNumber: invRequest.InvoiceNumber,
		CreatedOn:     time.Now(),
		DueDate:       dueDate,
		AmountDue:     invRequest.AmountDue,
		AskingPrice:   invRequest.AskingPrice,
		IssuerId:      invRequest.IssuerId,
	}
	resp, err_ := s.repo.Insert(&ctx, invoice)
	if err_ != nil {
		return nil, err_
	}
	return resp, nil
}

func (s DefaultInvoiceService) GetInvoice(ctx context.Context, id string) (*domain.Invoice, *appErr.AppError) {
	invoiceId, _ := strconv.Atoi(id)
	invoice, err_ := s.repo.FindById(&ctx, invoiceId)
	if err_ != nil {
		return nil, err_
	}
	return invoice, nil
}

func (s DefaultInvoiceService) GetAllInvoices(ctx context.Context) ([]domain.Invoice, *appErr.AppError) {
	invoices, err_ := s.repo.FindAll(&ctx)
	if err_ != nil {
		return nil, err_
	}
	return invoices, nil
}
