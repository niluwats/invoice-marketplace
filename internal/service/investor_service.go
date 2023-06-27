package service

import (
	"context"
	"strconv"

	"github.com/niluwats/invoice-marketplace/internal/domain"
	"github.com/niluwats/invoice-marketplace/internal/repositories"
	appErr "github.com/niluwats/invoice-marketplace/pkg/errors"
)

type InvestorService interface {
	GetInvestor(ctx context.Context, id string) (*domain.Investor, *appErr.AppError)
	GetAllInvestors() ([]domain.Investor, *appErr.AppError)
}

type DefaultInvestorService struct {
	repo repositories.InvestorRepository
}

func NewInvestorService(repo repositories.InvestorRepository) DefaultInvestorService {
	return DefaultInvestorService{repo}
}

func (s DefaultInvestorService) GetInvestor(ctx context.Context, id string) (*domain.Investor, *appErr.AppError) {
	investorId, _ := strconv.Atoi(id)
	investor, err_ := s.repo.FindById(&ctx, investorId)
	if err_ != nil {
		return nil, err_
	}
	return investor, nil
}

func (s DefaultInvestorService) GetAllInvestors(ctx context.Context) ([]domain.Investor, *appErr.AppError) {
	investors, err_ := s.repo.FindAll(&ctx)
	if err_ != nil {
		return nil, err_
	}
	return investors, nil
}
