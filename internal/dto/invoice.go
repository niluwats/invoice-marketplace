package dto

type InvoiceRequest struct {
	InvoiceNumber  string  `json:"invoice_number"`
	IssuerId       int     `json:"issuer_id"`
	AmountDue      float64 `json:"amount_due"`
	AmountEnclosed float64 `json:"amount_enclosed"`
	DueDate        string  `json:"duedate"`
}
