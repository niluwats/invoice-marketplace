package handlers

import (
	"fmt"
	"net/http"
)

type TradeHandler struct {
}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Invoice-marketplace HTTP Server")
}

func placeBid(w http.ResponseWriter, r *http.Request) {

}

func approveTrade(w http.ResponseWriter, r *http.Request) {

}

func viewAllTrades(w http.ResponseWriter, r *http.Request) {

}
