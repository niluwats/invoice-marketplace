package dto

import (
	"strconv"

	"github.com/niluwats/invoice-marketplace/internal/domain"
)

type Issuer struct {
	ID          int    `json:"id,omitempty"`
	CompanyName string `json:"company_name"`
	InvestorId  int    `json:"investor_id"`
}

type IssuerResponse struct {
	ID          string  `json:"id,omitempty"`
	CompanyName string  `json:"company_name"`
	Balance     float64 `json:"balance"`
}

func MapToIssuersResponse(issuer domain.Issuer) IssuerResponse {
	strIssuerId := strconv.Itoa(issuer.ID)
	return IssuerResponse{
		ID:          strIssuerId,
		CompanyName: issuer.CompanyName,
		Balance:     issuer.Investor.Balance,
	}
}
