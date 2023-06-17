package service

import (
	"strconv"

	"github.com/niluwats/invoice-marketplace/internal/dto"
	"github.com/niluwats/invoice-marketplace/internal/repositories"
)

type IssuerService interface {
	GetIssuer(id string) (*dto.IssuerResponse, error)
	GetAllIssuers() ([]dto.IssuerResponse, error)
}

type DefaultIssuerService struct {
	repo repositories.IssuerRepository
}

func NewIssuerService(repo repositories.IssuerRepository) DefaultIssuerService {
	return DefaultIssuerService{repo}
}

func (s DefaultIssuerService) GetIssuer(id string) (*dto.IssuerResponse, error) {
	issuerId, _ := strconv.Atoi(id)

	issuer, err := s.repo.FindById(issuerId)
	if err != nil {
		return nil, err
	}

	response := dto.MapToIssuersResponse(*issuer)
	return &response, nil
}

func (s DefaultIssuerService) GetAllIssuers() ([]dto.IssuerResponse, error) {
	issuers, err := s.repo.FindAll()
	if err != nil {
		return nil, err
	}

	response := make([]dto.IssuerResponse, 0)
	for _, v := range issuers {
		response = append(response, dto.MapToIssuersResponse(v))
	}

	return response, nil
}
