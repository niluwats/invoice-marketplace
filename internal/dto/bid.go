package dto

import (
	"fmt"
	"strconv"
)

type BidRequest struct {
	InvoiceId  int     `json:"invoice_id"`
	BidAmount  float64 `json:"bid_amount"`
	InvestorId int     `json:"investor_id"`
}

func (req *BidRequest) IfInValidRequest() bool {
	return fmt.Sprintf("%f", req.BidAmount) == "" || req.BidAmount <= 0 || strconv.Itoa(req.InvestorId) == "" || strconv.Itoa(req.InvoiceId) == ""
}
