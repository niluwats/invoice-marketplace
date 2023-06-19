package dto

import (
	"fmt"
	"strconv"
)

type InvoiceRequest struct {
	InvoiceNumber string  `json:"invoice_number"`
	IssuerId      int     `json:"issuer_id"`
	AmountDue     float64 `json:"amount_due"`
	AskingPrice   float64 `json:"asking_price"`
	DueDate       string  `json:"duedate"`
}

func (req *InvoiceRequest) IfInValidRequest() bool {
	return (req.InvoiceNumber == "" || strconv.Itoa(req.IssuerId) == "" || req.DueDate == "" || fmt.Sprintf("%f", req.AskingPrice) == "" || fmt.Sprintf("%f", req.AmountDue) == "")
}
