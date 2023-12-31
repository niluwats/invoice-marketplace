package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/niluwats/invoice-marketplace/internal/dto"
	"github.com/niluwats/invoice-marketplace/internal/service"
)

type InvoiceHandler struct {
	service service.DefaultInvoiceService
}

func (h InvoiceHandler) createInvoice(w http.ResponseWriter, r *http.Request) {
	var request dto.InvoiceRequest

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		writeResponse(w, http.StatusBadRequest, err.Error())
	} else {
		resp, err := h.service.NewInvoice(request)
		if err != nil {
			writeResponse(w, err.Code, err.Message)
		} else {
			writeResponse(w, http.StatusCreated, resp)
		}
	}
}

func (h InvoiceHandler) viewInvoice(w http.ResponseWriter, r *http.Request) {
	invoiceId := chi.URLParam(r, "id")

	resp, err := h.service.GetInvoice(invoiceId)
	if err != nil {
		writeResponse(w, err.Code, err.Message)
	} else {
		writeResponse(w, http.StatusOK, resp)
	}
}

func (h InvoiceHandler) viewAllInvoices(w http.ResponseWriter, r *http.Request) {
	resp, err := h.service.GetAllInvoices()
	if err != nil {
		writeResponse(w, err.Code, err.Message)
	} else {
		writeResponse(w, http.StatusOK, resp)
	}
}
