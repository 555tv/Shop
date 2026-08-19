package middleware

import (
	"context"
	"net/http"
	"strings"

	"gate/internal/auth"
	"gate/internal/config"

	"github.com/golang-jwt/jwt/v5"
)

type contextKey string

const (
	UserIDKey    contextKey = "userId"
	FirstNameKey contextKey = "firstName"
	LastNameKey  contextKey = "lastName"
	EmailKey     contextKey = "email"
	RoleKey      contextKey = "role"
)

func JWTMiddleware(cfg config.Config) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			authHeader := r.Header.Get("Authorization")

			if authHeader == "" {
				http.Error(
					w,
					"Authorization header required",
					http.StatusUnauthorized,
				)
				return
			}

			parts := strings.SplitN(authHeader, " ", 2)

			if len(parts) != 2 || parts[0] != "Bearer" {
				http.Error(
					w,
					"Invalid Authorization header",
					http.StatusUnauthorized,
				)
				return
			}

			tokenString := parts[1]

			token, err := auth.ValidateToken(
				tokenString,
				cfg.JWTSecret,
			)

			if err != nil || !token.Valid {
				http.Error(
					w,
					"Invalid or expired token",
					http.StatusUnauthorized,
				)
				return
			}

			claims, ok := token.Claims.(jwt.MapClaims)
			if !ok {
				http.Error(
					w,
					"Invalid token claims",
					http.StatusUnauthorized,
				)
				return
			}

			ctx := context.WithValue(
				r.Context(),
				UserIDKey,
				claims["userId"],
			)

			ctx = context.WithValue(
				ctx,
				FirstNameKey,
				claims["firstName"],
			)

			ctx = context.WithValue(
				ctx,
				LastNameKey,
				claims["lastName"],
			)

			ctx = context.WithValue(
				ctx,
				EmailKey,
				claims["email"],
			)

			ctx = context.WithValue(
				ctx,
				RoleKey,
				claims["role"],
			)

			next.ServeHTTP(
				w,
				r.WithContext(ctx),
			)
		})
	}
}
