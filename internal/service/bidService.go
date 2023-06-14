package service

import (
	"strconv"
	"time"

	"github.com/niluwats/invoice-marketplace/internal/domain"
	"github.com/niluwats/invoice-marketplace/internal/dto"
	"github.com/niluwats/invoice-marketplace/internal/repositories"
)

type BidService interface {
	PlaceBid(bidRequest dto.BidRequest) error
	UpdateApproval(id string) error
	GetAllBids(invoideId string) ([]domain.Bid, error)
}

type DefaultBidService struct {
	repo repositories.BidRepository
}

func NewBidService(repo repositories.BidRepository) DefaultBidService {
	return DefaultBidService{repo}
}

func (s DefaultBidService) PlaceBid(bidRequest dto.BidRequest) error {
	trade := domain.Bid{
		InvoiceId:  bidRequest.InvoiceId,
		BidAmount:  bidRequest.BidAmount,
		InvestorId: bidRequest.InvestorId,
		TimeStamp:  time.Now(),
		IsApproved: false,
	}
	err := s.repo.Insert(trade)
	if err != nil {
		return err
	}
	return nil
}

func (s DefaultBidService) UpdateApproval(id string) error {
	tradeId, _ := strconv.Atoi(id)
	err := s.repo.UpdateApproval(tradeId)
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
