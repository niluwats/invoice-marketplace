package domain

import (
	"time"
)

type Invoice struct {
	ID             int       `json:"id"`
	InvoiceNumber  string    `json:"invoice_number"`
	AmountDue      float64   `json:"amount_due"`
	AmountEnclosed float64   `json:"amount_enclosed"`
	DueDate        time.Time `json:"duedate"`
	CreatedOn      time.Time `json:"created_on"`
	IsLocked       bool      `json:"is_locked"`
	IsTraded       bool      `json:"is_traded"`
	IssuerId       int       `json:"issuer_id"`
}
