package service

import "github.com/niluwats/invoice-marketplace/pkg/domain"

type InvestorService interface {
	GetInvestor()
	GetAllInvestors()
	EditInvestorBalance()
}

type DefaultInvestorService struct {
	repo domain.InvestorRepository
}

func NewInvestorService(repo domain.InvestorRepository) DefaultInvestorService {
	return DefaultInvestorService{repo}
}

func (s DefaultIssuerService) GetInvestor() {

}

func (s DefaultIssuerService) GetAllInvestors() {

}

func (s DefaultIssuerService) EditInvestorBalance() {

}
