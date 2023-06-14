package dto

type InvoiceRequest struct {
	InvoiceNumber     string  `json:"invoice_number"`
	CustomerFirstName string  `json:"customer_first_name"`
	CustomerLastName  string  `json:"customer_last_name"`
	IssuerId          int     `json:"issuer"`
	AmountDue         float64 `json:"amount_due"`
	AmountEnclosed    float64 `json:"amount_enclosed"`
	DueDate           string  `json:"duedate"`
}
