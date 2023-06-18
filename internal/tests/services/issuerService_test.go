package mocks

import (
	"testing"

	"github.com/niluwats/invoice-marketplace/internal/domain"
	"github.com/niluwats/invoice-marketplace/internal/dto"
	"github.com/niluwats/invoice-marketplace/internal/service"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestIssuerService_GetIssuer(t *testing.T) {
	repo := &IssuerRepository{}

	id := "1"
	expectedIssuer := &domain.Issuer{
		ID:          1,
		CompanyName: "test1",
		Investor: domain.Investor{
			Balance: 8000,
		},
	}

	expectedResponse := &dto.IssuerResponse{
		ID:          "1",
		CompanyName: "test1",
		Balance:     8000,
	}

	repo.On("FindById", mock.Anything).Return(expectedIssuer, nil).Once()

	service := service.NewIssuerService(repo)
	issuer, err := service.GetIssuer(id)
	assert.Nil(t, err)
	assert.Equal(t, expectedResponse, issuer)
}

func TestIssuerService_GetAllIssuers(t *testing.T) {
	expectedIssuers := []domain.Issuer{
		{
			ID:          1,
			CompanyName: "test1",
			Investor: domain.Investor{
				Balance: 8000,
			},
		}, {
			ID:          2,
			CompanyName: "test2",
			Investor: domain.Investor{
				Balance: 10000,
			},
		},
	}

	expectedResponse := []dto.IssuerResponse{
		{
			ID:          "1",
			CompanyName: "test1",
			Balance:     8000,
		},
		{
			ID:          "2",
			CompanyName: "test2",
			Balance:     10000,
		},
	}

	repo := &IssuerRepository{}
	repo.On("FindAll").Return(expectedIssuers, nil).Once()

	service := service.NewIssuerService(repo)

	issuers, err := service.GetAllIssuers()
	assert.Nil(t, err)
	assert.Equal(t, expectedResponse, issuers)
}
