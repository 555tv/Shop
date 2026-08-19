package middleware

import (
	"net/http"
)

func RequireRole(requiredRole string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			role := r.Context().Value(RoleKey)

			if role == nil {
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}

			userRole, ok := role.(string)

			if !ok {
				http.Error(w, "Invalid role", http.StatusUnauthorized)
				return
			}

			if userRole != requiredRole {
				http.Error(w, "Forbidden", http.StatusForbidden)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}
