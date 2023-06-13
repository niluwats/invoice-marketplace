package handlers

import (
	"fmt"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/niluwats/invoice-marketplace/pkg/db"
)

func StartServer() {
	dbClient := db.SetupDBConn()

	fmt.Println(dbClient)

	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://*", "https://*"},
		AllowedHeaders:   []string{"Accept", "Content-Type"},
		AllowedMethods:   []string{"GET", "PUT", "POST", "DELETE"},
		AllowCredentials: true,
	}))

	router.Use(middleware.Heartbeat("/ping"))

	router.Get("/", home)

	router.Post("/invoice", createInvoice)
	router.Get("/invoice", viewInvoice)

	router.Get("/issuer", viewAllIssuers)
	router.Get("/issuer/{id}", viewIssuer)

	router.Get("/investor", viewAllInvestors)
	router.Get("/investor/{id}", viewInvestor)

	router.Post("/trade", placeBid)
	router.Get("/trade", viewAllTrades)

	router.Put("/approve", approveTrade)

	http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("WEB_PORT")), router)
}
