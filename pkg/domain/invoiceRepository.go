package domain

type InvoiceRepository interface {
	Insert(invoice Invoice) error
	GetInvoice(id int) (*Invoice, error)
	UpdateLockStatus(id int) error
	UpdateInvoiceInvestor(id int) error
}
