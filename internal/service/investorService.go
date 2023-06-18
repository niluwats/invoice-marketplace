package service

import (
	"strconv"

	"github.com/niluwats/invoice-marketplace/internal/domain"
	"github.com/niluwats/invoice-marketplace/internal/dto"
	"github.com/niluwats/invoice-marketplace/internal/repositories"
	appErr "github.com/niluwats/invoice-marketplace/pkg/errors"
)

type InvestorService interface {
	GetInvestor(id string) (*domain.Investor, *appErr.AppError)
	GetAllInvestors() ([]domain.Investor, *appErr.AppError)
	VerifyUser(dto.AuthRequest) (*dto.AuthResponse, *appErr.AppError)
}

type DefaultInvestorService struct {
	repo repositories.InvestorRepository
}

func NewInvestorService(repo repositories.InvestorRepository) DefaultInvestorService {
	return DefaultInvestorService{repo}
}

func (s DefaultInvestorService) GetInvestor(id string) (*domain.Investor, *appErr.AppError) {
	investorId, _ := strconv.Atoi(id)
	investor, err_ := s.repo.FindById(investorId)
	if err_ != nil {
		return nil, err_
	}
	return investor, nil
}

func (s DefaultInvestorService) GetAllInvestors() ([]domain.Investor, *appErr.AppError) {
	investors, err_ := s.repo.FindAll()
	if err_ != nil {
		return nil, err_
	}
	return investors, nil
}
