package domain

type IssuerRepository interface {
	GetById(id int) (*Issuer, error)
	GetAll() ([]Issuer, error)
	UpdateBalance(id int, amount float64) error
}
