package middleware

import (
	"net/http"
	"sync"
	"time"
)

type client struct {
	count       int
	lastRequest time.Time
}

var (
	clients = make(map[string]*client)
	mu      sync.Mutex
)

const (
	maxRequests = 10
	timeWindow  = time.Second
)

func ThrottleMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		ip := r.RemoteAddr

		mu.Lock()

		now := time.Now()

		c, exists := clients[ip]

		if !exists {
			clients[ip] = &client{
				count:       1,
				lastRequest: now,
			}

			mu.Unlock()

			next.ServeHTTP(w, r)
			return
		}

		if now.Sub(c.lastRequest) >= timeWindow {
			c.count = 1
			c.lastRequest = now

			mu.Unlock()

			next.ServeHTTP(w, r)
			return
		}

		c.count++

		if c.count > maxRequests {
			mu.Unlock()

			http.Error(
				w,
				"Too many requests",
				http.StatusTooManyRequests,
			)
			return
		}

		c.lastRequest = now

		mu.Unlock()

		next.ServeHTTP(w, r)
	})
}
