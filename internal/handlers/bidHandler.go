package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"sync"

	"github.com/go-chi/chi/v5"
	"github.com/niluwats/invoice-marketplace/internal/dto"
	"github.com/niluwats/invoice-marketplace/internal/middleware"
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
	user, _ := r.Context().Value(middleware.UserKey).(middleware.User)

	var request dto.BidRequest

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		writeResponse(w, http.StatusBadRequest, err.Error())
	} else {
		invoiceMutex, _ := mutexMap.LoadOrStore(request.InvoiceId, &sync.Mutex{})
		mutex := invoiceMutex.(*sync.Mutex)

		mutex.Lock()
		defer mutex.Unlock()

		request.InvestorId, _ = strconv.Atoi(user.InvestorId)

		bid, err := h.service.PlaceBid(request)
		if err != nil {
			writeResponse(w, err.Code, err.Message)
		} else {
			writeResponse(w, http.StatusCreated, &bid)
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

func (h BidHandler) rejectTrade(w http.ResponseWriter, r *http.Request) {
	invoiceId := chi.URLParam(r, "invoice_id")
	err := h.service.RejectTrade(invoiceId)
	if err != nil {
		writeResponse(w, err.Code, err.Message)
	} else {
		writeResponse(w, http.StatusOK, "Trade canceled!")
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

func (h BidHandler) viewBid(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	bid, err := h.service.GetBid(id)
	if err != nil {
		writeResponse(w, err.Code, err.Message)
	} else {
		writeResponse(w, http.StatusOK, bid)
	}
}
