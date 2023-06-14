package dto

type BidRequest struct {
	InvoiceId  int     `json:"invoice_id"`
	BidAmount  float64 `json:"bid_amount"`
	InvestorId int     `json:"investor_id"`
}
