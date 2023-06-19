package service

import (
	"strconv"
	"time"

	"github.com/niluwats/invoice-marketplace/internal/domain"
	"github.com/niluwats/invoice-marketplace/internal/dto"
	"github.com/niluwats/invoice-marketplace/internal/repositories"
	appErr "github.com/niluwats/invoice-marketplace/pkg/errors"
)

// types
type BidService interface {
	PlaceBid(bidRequest dto.BidRequest) (*domain.Bid, *appErr.AppError)
	ApproveTrade(invoiceId string) *appErr.AppError
	GetAllBids(invoiceId string) ([]domain.Bid, *appErr.AppError)
	GetBid(id string) (*domain.Bid, *appErr.AppError)
	checkIfInvestorBalanceSufficient(investorId int, bidAmount float64) *appErr.AppError
	trimIfBidAmountExceeds(invoiceId int, bidAmount, invoicePrice float64) (newBidAmount float64, restBalance float64, err *appErr.AppError)
}

type DefaultBidService struct {
	bidRepo      repositories.BidRepository
	investorRepo repositories.InvestorRepository
	invoiceRepo  repositories.InvoiceRepository
}

func NewBidService(bidRepo repositories.BidRepository, investorRepo repositories.InvestorRepository, invoiceRepo repositories.InvoiceRepository) DefaultBidService {
	return DefaultBidService{bidRepo, investorRepo, invoiceRepo}
}

// public methods
func (s DefaultBidService) PlaceBid(bidRequest dto.BidRequest) (*domain.Bid, *appErr.AppError) {
	invoiceId := bidRequest.InvoiceId
	bidAmount := bidRequest.BidAmount
	investorId := bidRequest.InvestorId

	if bidRequest.IfInValidRequest() {
		return nil, appErr.NewValidationError("All fields required")
	}

	invoice, err_ := s.invoiceRepo.FindById(invoiceId)
	if err_ != nil {
		return nil, err_
	}

	//check if invoice is valid to bid on
	if invoice.IsLocked {
		return nil, appErr.NewForbiddenError("Invoice is locked")
	}

	//check if investor's balance is sufficient
	err_ = s.checkIfInvestorBalanceSufficient(investorId, bidAmount)
	if err_ != nil {
		return nil, err_
	}

	//trim if bid amount exceeds rest amount
	newBidAmount, restBalance, err_ := s.trimIfBidAmountExceeds(invoiceId, bidAmount, invoice.AskingPrice)
	if err_ != nil {
		return nil, err_
	}

	bidAmount = newBidAmount

	//save bid
	bid := domain.Bid{
		InvoiceId:  invoiceId,
		BidAmount:  bidAmount,
		InvestorId: investorId,
		TimeStamp:  time.Now(),
		IsApproved: false,
	}

	resp, err_ := s.bidRepo.ProcessBid(bid, restBalance)
	if err_ != nil {
		return nil, err_
	}

	return resp, nil
}

func (s DefaultBidService) ApproveTrade(invoiceId string) *appErr.AppError {
	intInvoiceId, _ := strconv.Atoi(invoiceId)

	invoice, err_ := s.invoiceRepo.FindById(intInvoiceId)
	if err_ != nil {
		return err_
	}

	if !invoice.IsLocked {
		return appErr.NewForbiddenError("Invoice is not locked yet")
	}

	if invoice.IsTraded {
		return appErr.NewForbiddenError("Invoice is already traded")
	}

	err_ = s.bidRepo.ProcessApproveBid(intInvoiceId, invoice.IssuerId, invoice.AskingPrice)
	if err_ != nil {
		return err_
	}
	return nil
}

func (s DefaultBidService) GetAllBids(invoiceid string) ([]domain.Bid, *appErr.AppError) {
	invId, _ := strconv.Atoi(invoiceid)
	bids, err_ := s.bidRepo.GetAll(invId)
	if err_ != nil {
		return nil, err_
	}
	return bids, nil
}

func (s DefaultBidService) GetBid(id string) (*domain.Bid, *appErr.AppError) {
	bidId, _ := strconv.Atoi(id)

	bid, err_ := s.bidRepo.GetBid(bidId)
	if err_ != nil {
		return nil, err_
	}
	return bid, nil
}

// private methods
func (s DefaultBidService) checkIfInvestorBalanceSufficient(investorId int, bidAmount float64) *appErr.AppError {
	investor, err_ := s.investorRepo.FindById(investorId)
	if err_ != nil {
		return err_
	}

	if investor.Balance < bidAmount {
		return appErr.NewBadRequest("Investor's balance is insufficient!")
	}
	return nil
}

func (s DefaultBidService) trimIfBidAmountExceeds(invoiceId int, bidAmount, invoicePrice float64) (newBidAmount float64, restBalance float64, err *appErr.AppError) {
	investedSum, err_ := s.invoiceRepo.FindTotalInvestment(invoiceId)
	if err_ != nil {
		return 0, 0, err_
	}
	restInvoiceBalance := invoicePrice - investedSum
	if restInvoiceBalance < bidAmount {
		bidAmount = restInvoiceBalance
	}
	return bidAmount, restInvoiceBalance, nil
}
