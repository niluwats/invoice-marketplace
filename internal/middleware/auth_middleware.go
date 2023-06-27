package middleware

import (
	"context"
	"net/http"

	"github.com/niluwats/invoice-marketplace/internal/auth"
)

type User struct {
	InvestorId string
	Role       string
}

type contextKey int

const (
	UserKey contextKey = iota
)

func JWTMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := auth.ExtractTokenFromHeader(r)

		if tokenString == "" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		id, role, err := auth.ValidateTokenAndGetClaims(tokenString)
		if err != nil {
			http.Error(w, "Invalid JWT ", http.StatusUnauthorized)
			return
		}

		user := User{
			InvestorId: id,
			Role:       role,
		}
		ctx := context.WithValue(r.Context(), UserKey, user)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
