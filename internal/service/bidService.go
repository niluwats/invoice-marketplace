package service

import (
	"errors"
	"log"
	"strconv"
	"time"

	"github.com/niluwats/invoice-marketplace/internal/domain"
	"github.com/niluwats/invoice-marketplace/internal/dto"
	"github.com/niluwats/invoice-marketplace/internal/repositories"
)

type BidService interface {
	PlaceBid(bidRequest dto.BidRequest, invoiceRepo repositories.InvoiceRepository, investorRepo repositories.InvestorRepository) error
	UpdateApproval(invoiceId string) error
	GetAllBids(invoiceId string) ([]domain.Bid, error)
}

type DefaultBidService struct {
	repo repositories.BidRepository
}

func NewBidService(repo repositories.BidRepository) DefaultBidService {
	return DefaultBidService{repo}
}

func (s DefaultBidService) PlaceBid(bidRequest dto.BidRequest, invoiceRepo repositories.InvoiceRepository, investorRepo repositories.InvestorRepository) error {
	invoiceId := bidRequest.InvoiceId
	bidAmount := bidRequest.BidAmount
	investorId := bidRequest.InvestorId

	log.Println("1", invoiceId)

	//check if invoice is locked
	err := checkIfInvoiceLocked(invoiceId, invoiceRepo)
	if err != nil {
		return err
	}
	log.Println("2", invoiceId)

	//check if investor's balance is sufficient
	err = checkIfInvestorBalanceSufficient(investorId, bidAmount, investorRepo)
	if err != nil {
		return err
	}

	log.Println("3", invoiceId)

	//trim if bid amount exceeds rest amount
	newBidAmount, restBalance, err := trimIfBidAmountExceeds(invoiceId, bidAmount, invoiceRepo)
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

	log.Println(bid)
	err = s.repo.Insert(bid, restBalance)
	if err != nil {
		return err
	}

	return nil
}

func (s DefaultBidService) UpdateApproval(invoiceId string) error {
	intInvoiceId, _ := strconv.Atoi(invoiceId)

	invoice, err := s.repo.GetTrade(intInvoiceId)
	if !invoice.IsLocked {
		return errors.New("Invoice is not locked yet")
	}

	if invoice.IsTraded {
		return errors.New("Invoice is already traded")
	}

	err = s.repo.UpdateApproval(intInvoiceId, invoice.IssuerId, invoice.AmountEnclosed)
	if err != nil {
		return err
	}
	return nil
}

func (s DefaultBidService) GetAllBids(invoiceid string) ([]domain.Bid, error) {
	invId, _ := strconv.Atoi(invoiceid)
	bids, err := s.repo.GetAll(invId)
	if err != nil {
		return nil, err
	}
	return bids, nil
}

func checkIfInvoiceLocked(invoiceId int, invoiceRepo repositories.InvoiceRepository) error {
	invoice, err := invoiceRepo.FindById(invoiceId)
	if err != nil {
		return err
	}

	if invoice.IsLocked {
		return errors.New("Invoice is locked! can't bid anymore")
	}
	return nil
}

func checkIfInvestorBalanceSufficient(investorId int, bidAmount float64, investorRepo repositories.InvestorRepository) error {
	balance, err := investorRepo.FindInvestorBalance(investorId)
	if err != nil {
		return err
	}

	if balance < bidAmount {
		return errors.New("Investor's balance is insufficient!")
	}
	return nil
}

func trimIfBidAmountExceeds(invoiceId int, bidAmount float64, invoiceRepo repositories.InvoiceRepository) (newBidAmount float64, restBalance float64, err error) {
	invoice, err := invoiceRepo.FindById(invoiceId)
	if err != nil {
		return 0, 0, err
	}

	investedSum, err := invoiceRepo.FindSum(invoiceId)
	if err != nil {
		return 0, 0, err
	}
	restInvoiceBalance := invoice.AmountEnclosed - investedSum
	if restInvoiceBalance < bidAmount {
		bidAmount = restInvoiceBalance
	}
	return bidAmount, restInvoiceBalance, nil
}
