package mocks

import (
	"testing"

	"github.com/niluwats/invoice-marketplace/internal/domain"
	"github.com/niluwats/invoice-marketplace/internal/service"
	"github.com/stretchr/testify/assert"
	mock "github.com/stretchr/testify/mock"
)

func TestInvestorService_GetInvestor(t *testing.T) {
	investerId := "1"

	expectedInvestor := &domain.Investor{
		ID:        investerId,
		FirstName: "Jane",
		LastName:  "Daves",
		Balance:   8000,
		IsIssuer:  true,
	}

	repo := &InvestorRepository{}
	repo.On("FindById", mock.Anything).Return(expectedInvestor, nil).Once()

	service := service.NewInvestorService(repo)

	investor, err := service.GetInvestor(investerId)
	assert.Nil(t, err)
	assert.Equal(t, expectedInvestor, investor)
}

func TestInvestorService_GetAllInvestors(t *testing.T) {
	expectedInvestors := []domain.Investor{
		{
			ID:        "1",
			FirstName: "Jane",
			LastName:  "Daves",
			Balance:   8000,
			IsIssuer:  true,
		}, {
			ID:        "2",
			FirstName: "Will",
			LastName:  "Johnson",
			Balance:   10000,
			IsIssuer:  true,
		},
	}

	repo := &InvestorRepository{}
	repo.On("FindAll").Return(expectedInvestors, nil).Once()

	service := service.NewInvestorService(repo)

	investors, err := service.GetAllInvestors()
	assert.Nil(t, err)
	assert.Equal(t, expectedInvestors, investors)
}
