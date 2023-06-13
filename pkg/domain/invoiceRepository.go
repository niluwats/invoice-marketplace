package domain

type InvoiceRepository interface {
	Insert(invoice Invoice) error
	FindById(id int) (*Invoice, error)
	UpdateLockStatus(id int) error
	UpdateInvoiceInvestor(id int) error
}
