package middlewares

import (
	"net/http"
)

const (
	accessControlAllowOriginName   = "Access-Control-Allow-Origin"
	accessControlAllowOriginValue  = "*"
	accessControlAllowHeadersValue = "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization"
	accessControlAllowHeadersName  = "Access-Control-Allow-Headers"
	accessControlAllowMethodsName  = "Access-Control-Allow-Methods"
	accessControlAllowMethodsValue = "POST, OPTIONS"
	accessControlMaxAge            = "Access-Control-Max-Age"
	accessControlMaxAgeValue       = "1728000"
)

// CROS Add Access-Control-Allow-Origin to header.
func CROS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set(accessControlAllowOriginName, accessControlAllowOriginValue)
		w.Header().Set(accessControlAllowHeadersName, accessControlAllowHeadersValue)
		w.Header().Set(accessControlAllowMethodsName, accessControlAllowMethodsValue)
		w.Header().Set(accessControlMaxAge, accessControlMaxAgeValue)
		next.ServeHTTP(w, r)
	})
}
