package domain

type Investor struct {
	ID        string  `json:"id,omitempty"`
	FirstName string  `json:"first_name,omitempty"`
	LastName  string  `json:"last_name,omitempty"`
	Balance   float64 `json:"balance,omitempty"`
}
