package handlers

import (
	"bytes"
	"log"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
	"github.com/niluwats/invoice-marketplace/internal/db"
	auth_middleware "github.com/niluwats/invoice-marketplace/internal/middleware"
	"github.com/niluwats/invoice-marketplace/internal/repositories"
	"github.com/niluwats/invoice-marketplace/internal/service"
)

func Init() {
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	os.Getenv("DB_HOST")
}

func TestViewInvestorRoute(t *testing.T) {
	tests := []struct {
		description    string
		route          string
		expectedCode   int
		requestHeaders map[string]string
		requestBody    []byte
		method         string
		responseBody   string
	}{
		{description: "get HTTP status 404,when user Id is not found",
			route:  "/investor",
			method: "GET",
			requestHeaders: map[string]string{`Content-Type`: `application/json`,
				`Authorization`: `Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6IjEiLCJlbWFpbCI6IiIsInJvbGUiOiJpc3N1ZXIiLCJleHAiOjE2ODcxNDgxNjB9.wXcJSykp77l-lzvHuKUAFGgI1COQXy4BfT-G7RtWZyQ"`},
			expectedCode: 404,
			responseBody: `{"Investor not found "}`,
		},
		{
			description:    "get HTTP status 401, when user is not authorized",
			route:          "/investor",
			method:         "GET",
			requestHeaders: map[string]string{`Content-Type`: `application/json`},
			expectedCode:   401,
			responseBody:   `"Unauthorized"`,
		},
		{
			description: "get HTTP status 200, when request is correct",
			route:       "/investor",
			method:      "GET",
			requestHeaders: map[string]string{`Content-Type`: `application/json`,
				`Authorization`: `Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6IjEiLCJlbWFpbCI6IiIsInJvbGUiOiJpc3N1ZXIiLCJleHAiOjE2ODcxNDgxNjB9.wXcJSykp77l-lzvHuKUAFGgI1COQXy4BfT-G7RtWZyQ"`},
			expectedCode: 200,
			responseBody: `{""id": "1","first_name": "Jane","last_name": "Daves","balance": 8000,"is_issuer": true"}`,
		},
	}

	Init()

	dbClient := db.SetupDBConn()
	investorRepoDb := repositories.NewInvestorRepositoryDb(dbClient)
	investorHandler := InvestorHandler{service.NewInvestorService(investorRepoDb)}

	router := chi.NewRouter()
	router.Group(func(r chi.Router) {
		r.Use(auth_middleware.JWTMiddleware)
		r.Get("/investor", investorHandler.viewInvestor)
	})

	for _, test := range tests {
		req := httptest.NewRequest(test.method, test.route, bytes.NewReader(test.requestBody))

		for k, v := range test.requestHeaders {
			req.Header.Add(k, v)
		}

		rr := httptest.NewRecorder()
		router.ServeHTTP(rr, req)

		if test.expectedCode != rr.Code {
			t.Errorf("%s: failed\nStatusCode was incorrect, got: %d, want: %d.", test.description, rr.Code, test.expectedCode)
		}

		expectedBody := []byte(test.responseBody)
		actualBody := rr.Body.Bytes()
		if !bytes.Equal(actualBody, expectedBody) {
			t.Errorf("%s: failed\nResponse body was incorrect, got: %s, want: %s.", test.description, string(actualBody), test.responseBody)
		}
	}
}
