package service

import "github.com/niluwats/invoice-marketplace/pkg/domain"

type IssuerService interface {
	GetIssuer()
	GetAllIssuers()
	EditIssuerBalance()
}

type DefaultIssuerService struct {
	repo domain.IssuerRepository
}

func NewIssuerService(repo domain.IssuerRepository) DefaultIssuerService {
	return DefaultIssuerService{repo}
}

func (s DefaultIssuerService) GetIssuer() {

}

func (s DefaultIssuerService) GetAllIssuers() {

}

func (s DefaultIssuerService) EditIssuerBalance() {

}
