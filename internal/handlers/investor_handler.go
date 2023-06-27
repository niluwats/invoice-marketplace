package handlers

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/niluwats/invoice-marketplace/internal/service"
)

type InvestorHandler struct {
	service service.DefaultInvestorService
}

func (h InvestorHandler) viewInvestor(w http.ResponseWriter, r *http.Request) {
	investorId := chi.URLParam(r, "ID")

	resp, err := h.service.GetInvestor(r.Context(), investorId)
	if err != nil {
		writeResponse(w, err.Code, err.Message)
	} else {
		writeResponse(w, http.StatusOK, resp)
	}
}

func (h InvestorHandler) viewAllInvestors(w http.ResponseWriter, r *http.Request) {
	resp, err := h.service.GetAllInvestors(r.Context())
	if err != nil {
		writeResponse(w, err.Code, err.Message)
	} else {
		writeResponse(w, http.StatusOK, resp)
	}
}
