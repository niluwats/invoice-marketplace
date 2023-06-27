package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/niluwats/invoice-marketplace/internal/dto"
	"github.com/niluwats/invoice-marketplace/internal/service"
)

type AuthHandler struct {
	service service.DefaultAuthService
}

func NewAuthHandler(service service.DefaultAuthService) AuthHandler {
	return AuthHandler{service}
}

func (h AuthHandler) authenticate(w http.ResponseWriter, r *http.Request) {
	var request dto.AuthRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		writeResponse(w, http.StatusBadRequest, err.Error())
	} else {
		response, err := h.service.VerifyUser(r.Context(), request)
		if err != nil {
			writeResponse(w, err.Code, err.Message)
		} else {
			writeResponse(w, http.StatusOK, response)
		}
	}
}

func (h AuthHandler) register(w http.ResponseWriter, r *http.Request) {
	var request dto.NewUserRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		writeResponse(w, http.StatusBadRequest, err.Error())
	} else {
		err := h.service.Register(r.Context(), request)
		if err != nil {
			writeResponse(w, err.Code, err.Message)
		} else {
			writeResponse(w, http.StatusCreated, "New user created")
		}
	}
}
