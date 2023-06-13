package domain

type InvestorRepository interface {
	FindById(id int) (*Investor, error)
	FIndAll() ([]Investor, error)
	UpdateBalance(id int, amount float64) error
}
