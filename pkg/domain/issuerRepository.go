package domain

type IssuerRepository interface {
	FindById(id int) (*Issuer, error)
	FindAll() ([]Issuer, error)
	UpdateBalance(id int, amount float64) error
}
