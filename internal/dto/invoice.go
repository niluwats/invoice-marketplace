package dto

type InvoiceRequest struct {
	InvoiceNumber string  `json:"invoice_number"`
	IssuerId      int     `json:"issuer_id"`
	AmountDue     float64 `json:"amount_due"`
	AskingPrice   float64 `json:"asking_price"`
	DueDate       string  `json:"duedate"`
}
