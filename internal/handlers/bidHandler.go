package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"

	"github.com/go-chi/chi/v5"
	"github.com/niluwats/invoice-marketplace/internal/dto"
	"github.com/niluwats/invoice-marketplace/internal/service"
)

var mutexMap sync.Map

type BidHandler struct {
	service service.DefaultBidService
}

func NewBidHandler(service service.DefaultBidService) BidHandler {
	return BidHandler{service}
}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Invoice-marketplace HTTP Server")
}

func (h BidHandler) placeBid(w http.ResponseWriter, r *http.Request) {
	var request dto.BidRequest

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		writeResponse(w, http.StatusBadRequest, err.Error())
	} else {
		invoiceMutex, _ := mutexMap.LoadOrStore(request.InvoiceId, &sync.Mutex{})
		mutex := invoiceMutex.(*sync.Mutex)

		mutex.Lock()
		defer mutex.Unlock()

		err := h.service.PlaceBid(request)
		if err != nil {
			writeResponse(w, err.Code, err.Message)
		} else {
			writeResponse(w, http.StatusCreated, "Bid placed successfully")
		}
	}
}

func (h BidHandler) approveTrade(w http.ResponseWriter, r *http.Request) {
	invoiceId := chi.URLParam(r, "invoice_id")
	err := h.service.ApproveTrade(invoiceId)
	if err != nil {
		writeResponse(w, err.Code, err.Message)
	} else {
		writeResponse(w, http.StatusOK, "Trade approved!")
	}
}

func (h BidHandler) viewAllBids(w http.ResponseWriter, r *http.Request) {
	invoiceId := chi.URLParam(r, "invoice_id")
	bids, err := h.service.GetAllBids(invoiceId)
	if err != nil {
		writeResponse(w, err.Code, err.Message)
	} else {
		writeResponse(w, http.StatusOK, bids)
	}
}
