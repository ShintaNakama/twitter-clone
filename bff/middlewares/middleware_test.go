package middlewares_test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/ShintaNakama/twitter-clone/bff/middlewares"
)

func TestApply(t *testing.T) {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if _, err := w.Write([]byte("h")); err != nil {
			t.Error(err)
		}
		w.WriteHeader(http.StatusOK)
	})
	ms := []middlewares.Middleware{
		func(next http.Handler) http.Handler {
			return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				if _, err := w.Write([]byte("m1")); err != nil {
					t.Error(err)
				}
				next.ServeHTTP(w, r)
			})
		},
		func(next http.Handler) http.Handler {
			return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				if _, err := w.Write([]byte("m2")); err != nil {
					t.Error(err)
				}
				next.ServeHTTP(w, r)
			})
		},
	}

	req, rec := newTestRequest("GET", "/test", "")
	middlewares.Apply(ms...)(handler).ServeHTTP(rec, req)

	if rec.Body.String() != "m1m2h" {
		t.Errorf("want: m1m2h, got: %s", rec.Body.String())
	}
}

func newTestRequest(method string, url string, body string) (*http.Request, *httptest.ResponseRecorder) {
	req := httptest.NewRequest(method, url, strings.NewReader(body))
	rec := httptest.NewRecorder()

	return req, rec
}
