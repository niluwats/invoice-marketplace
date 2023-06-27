package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/niluwats/invoice-marketplace/internal/service"
)

type IssuerHandler struct {
	service service.DefaultIssuerService
}

func (h IssuerHandler) viewIssuer(w http.ResponseWriter, r *http.Request) {
	issuerId := chi.URLParam(r, "ID")

	resp, err := h.service.GetIssuer(r.Context(), issuerId)
	if err != nil {
		writeResponse(w, err.Code, err.Message)
	} else {
		writeResponse(w, http.StatusOK, resp)
	}
}

func (h IssuerHandler) viewAllIssuers(w http.ResponseWriter, r *http.Request) {
	resp, err := h.service.GetAllIssuers(r.Context())
	if err != nil {
		writeResponse(w, err.Code, err.Message)
	} else {
		writeResponse(w, http.StatusOK, resp)
	}
}

func writeResponse(w http.ResponseWriter, code int, data any) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Println("Error encoding response json : ", err)
	}
}
