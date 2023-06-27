package domain

import "golang.org/x/crypto/bcrypt"

type Investor struct {
	ID        string  `json:"id,omitempty"`
	FirstName string  `json:"first_name,omitempty"`
	LastName  string  `json:"last_name,omitempty"`
	Balance   float64 `json:"balance,omitempty"`
	Email     string  `json:"email,omitempty"`
	Password  string  `json:"password,omitempty"`
	Status    bool    `json:"status,omitempty"`
	IsIssuer  bool    `json:"is_issuer,omitempty"`
}

func (i *Investor) GetBalance() float64 {
	return i.Balance
}

func (user *Investor) CheckPassword(providedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(providedPassword))
	if err != nil {
		return false
	}
	return true
}
