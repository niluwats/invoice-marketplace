package service

import "github.com/niluwats/invoice-marketplace/pkg/domain"

type TradeService interface {
	PlaceBid()
	UpdateApproval()
}

type DefaultTradeService struct {
	repo domain.TradeRepository
}

func NewTradeService(repo domain.TradeRepository) DefaultTradeService {
	return DefaultTradeService{repo}
}

func (d DefaultTradeService) PlaceBid() {

}

func (d DefaultTradeService) UpdateApproval() {

}
