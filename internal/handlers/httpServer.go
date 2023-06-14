package handlers

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/niluwats/invoice-marketplace/internal/db"
	"github.com/niluwats/invoice-marketplace/internal/repositories"
	"github.com/niluwats/invoice-marketplace/internal/service"
)

func StartServer() {
	log.Println("app started")
	dbClient := db.SetupDBConn()

	invoiceRepoDb := repositories.NewInvoiceRepositoryDb(dbClient)
	invoiceService := service.NewInvoiceService(invoiceRepoDb)
	invoiceHandler := InvoiceHandler{invoiceService}

	investorRepoDb := repositories.NewInvestorRepositoryDb(dbClient)
	investorService := service.NewInvestorService(investorRepoDb)
	investorHandler := InvestorHandler{investorService}

	issuerRepoDb := repositories.NewIssuerRepositoryDb(dbClient)
	issuerService := service.NewIssuerService(issuerRepoDb)
	issuerHandler := IssuerHandler{issuerService}

	bidRepoDb := repositories.NewBidRepositoryDb(dbClient)
	bidService := service.NewBidService(bidRepoDb)
	bidHandler := BidHandler{bidService}

	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://*", "https://*"},
		AllowedHeaders:   []string{"Accept", "Content-Type"},
		AllowedMethods:   []string{"GET", "POST", "PATCH"},
		AllowCredentials: true,
	}))

	router.Use(middleware.Heartbeat("/ping"))

	router.Get("/", home)

	router.Post("/invoice", invoiceHandler.createInvoice)
	router.Get("/invoice", invoiceHandler.viewInvoice)

	router.Get("/issuer", issuerHandler.viewAllIssuers)
	router.Get("/issuer/{id}", issuerHandler.viewIssuer)

	router.Get("/investor", investorHandler.viewAllInvestors)
	router.Get("/investor/{id}", investorHandler.viewInvestor)

	router.Post("/trade", bidHandler.placeBid)
	router.Get("/trade/{invoice_id}", bidHandler.viewAllBids)
	router.Patch("/trade/{bid_id}", bidHandler.approveTrade)

	http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("WEB_PORT")), router)
}
