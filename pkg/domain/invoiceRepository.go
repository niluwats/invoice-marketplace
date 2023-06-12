package domain

type InvoiceRepository interface {
	Insert(invoice Invoice) error
	Update(invoice Invoice) error
	GetInvoice(id int) (*Invoice, error)
	UpdateLockStatus(id int) error
	UpdateInvestor(id int) error
}
