package domain

type Issuer struct {
	ID          int    `json:"id,omitempty"`
	CompanyName string `json:"company_name"`
	InvestorId  int    `json:"investor_id"`
	Investor    Investor
}
