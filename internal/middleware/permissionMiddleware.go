package middleware

import (
	"net/http"
)

func PermissionMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user, ok := r.Context().Value(UserKey).(User)
		if !ok {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		if user.Role != "issuer" {
			http.Error(w, "Unauthorized", http.StatusForbidden)
			return
		}
		next.ServeHTTP(w, r)
	})
}
