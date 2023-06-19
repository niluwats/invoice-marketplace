package dto

type Investor struct {
	ID        int     `json:"id,omitempty"`
	FirstName string  `json:"first_name"`
	LastName  string  `json:"last_name"`
	Balance   float64 `json:"balance"`
}
