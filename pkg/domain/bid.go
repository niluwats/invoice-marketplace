package domain

import "time"

type Bid struct {
	ID            int       `json:"id"`
	InvoiceNumber string    `json:"invoice_number"`
	BidAmount     float64   `json:"bid_amount"`
	TimeStamp     time.Time `json:"timestamp"`
	InvestorId    int       `json:"investor_id"`
	IsApproved    bool      `json:"is_approved"`
}
