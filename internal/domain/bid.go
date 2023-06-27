package domain

import "time"

type Bid struct {
	ID         int       `json:"id"`
	BidAmount  float64   `json:"bid_amount"`
	IsApproved bool      `json:"is_approved"`
	CreatedAt  time.Time `json:"created_at"`
	InvestorId int       `json:"investor_id"`
	InvoiceId  int       `json:"invoice_id"`
	Status     int8      `json:"status"`
}
