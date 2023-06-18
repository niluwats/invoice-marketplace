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
	auth_middleware "github.com/niluwats/invoice-marketplace/internal/middleware"
	"github.com/niluwats/invoice-marketplace/internal/repositories"
	"github.com/niluwats/invoice-marketplace/internal/service"
)

func StartServer() {
	log.Println("app started")
	dbClient := db.SetupDBConn()

	invoiceRepoDb := repositories.NewInvoiceRepositoryDb(dbClient)
	invoiceHandler := InvoiceHandler{service.NewInvoiceService(invoiceRepoDb)}

	investorRepoDb := repositories.NewInvestorRepositoryDb(dbClient)
	investorHandler := InvestorHandler{service.NewInvestorService(investorRepoDb)}

	issuerRepoDb := repositories.NewIssuerRepositoryDb(dbClient)
	issuerHandler := IssuerHandler{service.NewIssuerService(issuerRepoDb)}

	bidRepoDb := repositories.NewBidRepositoryDb(dbClient)
	bidHandler := NewBidHandler(service.NewBidService(bidRepoDb, investorRepoDb, invoiceRepoDb))

	authHandler := NewAuthHandler(service.NewAuthService(investorRepoDb))

	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://*", "https://*"},
		AllowedHeaders:   []string{"Accept", "Content-Type"},
		AllowedMethods:   []string{"GET", "POST", "PATCH"},
		AllowCredentials: true,
	}))

	router.Use(middleware.Heartbeat("/ping"))

	router.Get("/", home)
	router.Post("/register", authHandler.register)
	router.Post("/auth", authHandler.authenticate)

	router.Group(func(r chi.Router) {
		r.Use(auth_middleware.JWTMiddleware)

		r.Get("/investor", investorHandler.viewAllInvestors)
		r.Get("/investor/{id}", investorHandler.viewInvestor)

		r.Get("/issuer", issuerHandler.viewAllIssuers)
		r.Get("/issuer/{id}", issuerHandler.viewIssuer)

		r.Post("/bid", bidHandler.placeBid)
		r.Get("/bid/{invoice_id}", bidHandler.viewAllBids)

		r.Get("/invoice/{id}", invoiceHandler.viewInvoice)
		r.With(auth_middleware.PermissionMiddleware).Post("/invoice", invoiceHandler.createInvoice)
		r.With(auth_middleware.PermissionMiddleware).Patch("/invoice/{invoice_id}", bidHandler.approveTrade)
	})

	http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("WEB_PORT")), router)
}
