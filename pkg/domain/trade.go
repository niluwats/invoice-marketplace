package domain

import "time"

type Trade struct {
	ID         int       `json:"id"`
	InvoiceId  string    `json:"invoice_id"`
	BidAmount  float64   `json:"bid_amount"`
	TimeStamp  time.Time `json:"timestamp"`
	InvestorId int       `json:"investor_id"`
	IsApproved bool      `json:"is_approved"`
}
