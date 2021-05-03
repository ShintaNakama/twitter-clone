package middlewares

import (
	"context"
	"fmt"
	"net/http"

	"go.uber.org/zap"
)

type key int

const (
	authorizationTokenContextKey key = iota
)

// GetAuthorizationToken get Authorization token
func GetAuthorizationToken(ctx context.Context) (string, error) {
	v := ctx.Value(authorizationTokenContextKey)
	authorizationToken, ok := v.(string)
	if !ok || authorizationToken == "" {
		return "", fmt.Errorf("authorizationToken not found")
	}
	return authorizationToken, nil
}

// SetAuthorizationToken set Authorization token to context
func SetAuthorizationToken(parent context.Context, value string) context.Context {
	return context.WithValue(parent, authorizationTokenContextKey, value)
}

// AuthorizationTokenMiddleWare get API token from `Authorization` header.
func AuthorizationTokenMiddleWare(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		zap.L().Debug("token", zap.String("token", token))
		ctx := SetAuthorizationToken(r.Context(), token)
		r = r.WithContext(ctx)
		h.ServeHTTP(w, r)
	})
}
