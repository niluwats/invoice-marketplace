package service

import (
	"errors"
	"strconv"
	"time"

	"github.com/niluwats/invoice-marketplace/internal/domain"
	"github.com/niluwats/invoice-marketplace/internal/dto"
	"github.com/niluwats/invoice-marketplace/internal/repositories"
)

// types
type BidService interface {
	PlaceBid(bidRequest dto.BidRequest) error
	ApproveTrade(invoiceId string) error
	GetAllBids(invoiceId string) ([]domain.Bid, error)
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
func (s DefaultBidService) PlaceBid(bidRequest dto.BidRequest) error {
	invoiceId := bidRequest.InvoiceId
	bidAmount := bidRequest.BidAmount
	investorId := bidRequest.InvestorId

	invoice, err := s.invoiceRepo.FindById(invoiceId)
	if err != nil {
		return err
	}

	//check if invoice is valid to bid on
	if invoice.IsLocked {
		return errors.New("invoice is locked")
	}

	//check if investor's balance is sufficient
	err = s.checkIfInvestorBalanceSufficient(investorId, bidAmount)
	if err != nil {
		return err
	}

	//trim if bid amount exceeds rest amount
	newBidAmount, restBalance, err := s.trimIfBidAmountExceeds(invoiceId, bidAmount, invoice.AskingPrice)
	if err != nil {
		return err
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

	err = s.bidRepo.ProcessBid(bid, restBalance)
	if err != nil {
		return err
	}

	return nil
}

func (s DefaultBidService) ApproveTrade(invoiceId string) error {
	intInvoiceId, _ := strconv.Atoi(invoiceId)

	invoice, err := s.invoiceRepo.FindById(intInvoiceId)
	if !invoice.IsLocked {
		return errors.New("Invoice is not locked yet")
	}

	if invoice.IsTraded {
		return errors.New("Invoice is already traded")
	}

	err = s.bidRepo.ProcessApproveBid(intInvoiceId, invoice.IssuerId, invoice.AskingPrice)
	if err != nil {
		return err
	}
	return nil
}

func (s DefaultBidService) GetAllBids(invoiceid string) ([]domain.Bid, error) {
	invId, _ := strconv.Atoi(invoiceid)
	bids, err := s.bidRepo.GetAll(invId)
	if err != nil {
		return nil, err
	}
	return bids, nil
}

// private methods
func (s DefaultBidService) checkIfInvestorBalanceSufficient(investorId int, bidAmount float64) error {
	investor, err := s.investorRepo.FindById(investorId)
	if err != nil {
		return err
	}

	if investor.Balance < bidAmount {
		return errors.New("Investor's balance is insufficient!")
	}
	return nil
}

func (s DefaultBidService) trimIfBidAmountExceeds(invoiceId int, bidAmount, invoicePrice float64) (newBidAmount float64, restBalance float64, err error) {
	investedSum, err := s.invoiceRepo.FindTotalInvestment(invoiceId)
	if err != nil {
		return 0, 0, err
	}
	restInvoiceBalance := invoicePrice - investedSum
	if restInvoiceBalance < bidAmount {
		bidAmount = restInvoiceBalance
	}
	return bidAmount, restInvoiceBalance, nil
}
