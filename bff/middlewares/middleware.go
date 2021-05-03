package middlewares

import (
	"net/http"
)

// Middleware ミドルウェア.
type Middleware func(hn http.Handler) http.Handler

// Apply ミドルウェアを適用する msに渡された順に適用される.
func Apply(ms ...Middleware) func(h http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		for i := len(ms) - 1; i >= 0; i-- {
			h = ms[i](h)
		}

		return h
	}
}
