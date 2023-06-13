package handlers

import (
	"encoding/json"
	"log"
	"net/http"
)

type IssuerHandler struct {
}

func viewIssuer(w http.ResponseWriter, r *http.Request) {

}

func viewAllIssuers(w http.ResponseWriter, r *http.Request) {

}

func writeResponse(w http.ResponseWriter, code int, data any) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Println("Error encoding response json : ", err)
	}
}
