package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/niluwats/invoice-marketplace/internal/dto"
	"github.com/niluwats/invoice-marketplace/internal/service"
)

type BidHandler struct {
	service service.DefaultBidService
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
		err := h.service.PlaceBid(request)
		if err != nil {
			writeResponse(w, http.StatusInternalServerError, err.Error())
		} else {
			writeResponse(w, http.StatusOK, "Bid placed successfully")
		}
	}
}

func (h BidHandler) approveTrade(w http.ResponseWriter, r *http.Request) {
	bidId := chi.URLParam(r, "id")
	err := h.service.UpdateApproval(bidId)
	if err != nil {
		writeResponse(w, http.StatusInternalServerError, err.Error())
	} else {
		writeResponse(w, http.StatusOK, "Bid approved!")
	}
}

func (h BidHandler) viewAllBids(w http.ResponseWriter, r *http.Request) {
	invoiceId := chi.URLParam(r, "invoice_id")
	bids, err := h.service.GetAllBids(invoiceId)
	if err != nil {
		writeResponse(w, http.StatusInternalServerError, err.Error())
	} else {
		writeResponse(w, http.StatusOK, bids)
	}
}
