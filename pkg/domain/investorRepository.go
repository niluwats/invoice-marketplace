package domain

type InvestorRepository interface {
	GetById(id int) (*Investor, error)
	GetAll() ([]Investor, error)
	UpdateBalance(id int, amount float64) error
}
