package service

import (
	"strconv"

	"github.com/niluwats/invoice-marketplace/internal/domain"
	"github.com/niluwats/invoice-marketplace/internal/repositories"
)

type IssuerService interface {
	GetIssuer(id string) (*domain.Issuer, error)
	GetAllIssuers() ([]domain.Issuer, error)
	EditIssuerBalance(id string, amount float64) error
}

type DefaultIssuerService struct {
	repo repositories.IssuerRepository
}

func NewIssuerService(repo repositories.IssuerRepository) DefaultIssuerService {
	return DefaultIssuerService{repo}
}

func (s DefaultIssuerService) GetIssuer(id string) (*domain.Issuer, error) {
	issuerId, _ := strconv.Atoi(id)

	issuer, err := s.repo.FindById(issuerId)
	if err != nil {
		return nil, err
	}

	return issuer, nil
}

func (s DefaultIssuerService) GetAllIssuers() ([]domain.Issuer, error) {
	issuers, err := s.repo.FindAll()
	if err != nil {
		return nil, err
	}

	return issuers, nil
}

func (s DefaultIssuerService) EditIssuerBalance(id string, amount float64) error {
	issuerId, _ := strconv.Atoi(id)

	err := s.repo.UpdateBalance(issuerId, amount)
	if err != nil {
		return err
	}
	return nil
}
