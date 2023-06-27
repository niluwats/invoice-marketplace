package service

import (
	"context"
	"strconv"

	"github.com/niluwats/invoice-marketplace/internal/dto"
	"github.com/niluwats/invoice-marketplace/internal/repositories"
	appErr "github.com/niluwats/invoice-marketplace/pkg/errors"
)

type IssuerService interface {
	GetIssuer(ctx context.Context, id string) (*dto.IssuerResponse, *appErr.AppError)
	GetAllIssuers(ctx context.Context) ([]dto.IssuerResponse, *appErr.AppError)
}

type DefaultIssuerService struct {
	repo repositories.IssuerRepository
}

func NewIssuerService(repo repositories.IssuerRepository) DefaultIssuerService {
	return DefaultIssuerService{repo}
}

func (s DefaultIssuerService) GetIssuer(ctx context.Context, id string) (*dto.IssuerResponse, *appErr.AppError) {
	issuerId, _ := strconv.Atoi(id)

	issuer, err_ := s.repo.FindById(&ctx, issuerId)
	if err_ != nil {
		return nil, err_
	}

	response := dto.MapToIssuersResponse(*issuer)
	return &response, nil
}

func (s DefaultIssuerService) GetAllIssuers(ctx context.Context) ([]dto.IssuerResponse, *appErr.AppError) {
	issuers, err_ := s.repo.FindAll(&ctx)
	if err_ != nil {
		return nil, err_
	}

	response := make([]dto.IssuerResponse, 0)
	for _, v := range issuers {
		response = append(response, dto.MapToIssuersResponse(v))
	}

	return response, nil
}
