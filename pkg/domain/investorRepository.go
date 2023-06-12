package domain

type InvestorRepository interface {
	GetById(id int) (*Investor, error)
	GetAll() ([]Investor, error)
}
