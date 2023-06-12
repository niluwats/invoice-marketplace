package domain

type IssuerRepository interface {
	GetById(id int) (*Issuer, error)
	GetAll() ([]Issuer, error)
}
