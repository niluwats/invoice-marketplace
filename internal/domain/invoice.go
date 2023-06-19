package domain

import (
	"time"
)

type Invoice struct {
	ID            int       `json:"id"`
	InvoiceNumber string    `json:"invoice_number,omitempty"`
	AmountDue     float64   `json:"amount_due,omitempty"`
	AskingPrice   float64   `json:"asking_price,omitempty"`
	DueDate       time.Time `json:"duedate,omitempty"`
	CreatedOn     time.Time `json:"created_on,omitempty"`
	IsLocked      bool      `json:"is_locked"`
	IsTraded      bool      `json:"is_traded"`
	IssuerId      int       `json:"issuer_id,omitempty"`
	InvestorIds   []int     `json:"investors,omitempty"`
}
