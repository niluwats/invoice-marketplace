package domain

import "time"

type Bid struct {
	ID         int       `json:"id"`
	BidAmount  float64   `json:"bid_amount"`
	IsApproved bool      `json:"is_approved"`
	TimeStamp  time.Time `json:"timestamp"`
	InvestorId int       `json:"investor_id"`
	InvoiceId  int       `json:"invoice_id"`
}
