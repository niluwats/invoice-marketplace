package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/go-chi/chi/v5"
	"github.com/niluwats/invoice-marketplace/internal/dto"
	"github.com/niluwats/invoice-marketplace/internal/repositories"
	"github.com/niluwats/invoice-marketplace/internal/service"
)

var mutexMap sync.Map

type BidHandler struct {
	service      service.DefaultBidService
	invoiceRepo  repositories.InvoiceRepository
	investorRepo repositories.InvestorRepository
}

func NewBidHandler(service service.DefaultBidService, invoiceRepo repositories.InvoiceRepository, investorRepo repositories.InvestorRepository) BidHandler {
	return BidHandler{service: service, invoiceRepo: invoiceRepo, investorRepo: investorRepo}
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

		log.Println("placing bid ", request.InvoiceId, request.InvestorId)

		err := h.service.PlaceBid(request, h.invoiceRepo, h.investorRepo)
		if err != nil {
			writeResponse(w, http.StatusInternalServerError, err.Error())
		} else {
			writeResponse(w, http.StatusOK, "Bid placed successfully")
		}
	}
}

func (h BidHandler) approveTrade(w http.ResponseWriter, r *http.Request) {
	invoiceId := chi.URLParam(r, "invoice_id")
	err := h.service.UpdateApproval(invoiceId)
	if err != nil {
		writeResponse(w, http.StatusInternalServerError, err.Error())
	} else {
		writeResponse(w, http.StatusOK, "Trade approved!")
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
