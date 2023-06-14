package service

import (
	"strconv"

	"github.com/niluwats/invoice-marketplace/internal/domain"
	"github.com/niluwats/invoice-marketplace/internal/repositories"
)

type InvestorService interface {
	GetInvestor(id string) (*domain.Investor, error)
	GetAllInvestors() ([]domain.Investor, error)
	EditInvestorBalance(id int, amount float64) error
}

type DefaultInvestorService struct {
	repo repositories.InvestorRepository
}

func NewInvestorService(repo repositories.InvestorRepository) DefaultInvestorService {
	return DefaultInvestorService{repo}
}

func (s DefaultInvestorService) GetInvestor(id string) (*domain.Investor, error) {
	investorId, _ := strconv.Atoi(id)
	investor, err := s.repo.FindById(investorId)
	if err != nil {
		return nil, err
	}
	return investor, nil
}

func (s DefaultInvestorService) GetAllInvestors() ([]domain.Investor, error) {
	investors, err := s.repo.FindAll()
	if err != nil {
		return nil, err
	}
	return investors, nil
}

func (s DefaultInvestorService) EditInvestorBalance(id int, amount float64) error {
	err := s.repo.UpdateBalance(id, amount)
	if err != nil {
		return err
	}
	return nil
}
