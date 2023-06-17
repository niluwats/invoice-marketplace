package domain

type Issuer struct {
	ID          int    `json:"id,omitempty"`
	CompanyName string `json:"company_name,omitempty"`
	InvestorId  int    `json:"investor_id,omitempty"`
	Investor    Investor
}
