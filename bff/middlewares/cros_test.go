package middlewares_test

import (
	"net/http"
	"testing"

	"github.com/ShintaNakama/twitter-clone/bff/middlewares"
)

func TestCROS(t *testing.T) {
	req, rec := newTestRequest("GET", "/", "")
	middlewares.CROS(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if _, err := w.Write([]byte("test")); err != nil {
			t.Error(err)
		}
		w.WriteHeader(http.StatusOK)
	})).ServeHTTP(rec, req)

	cros := rec.Header().Get("Access-Control-Allow-Origin")
	if cros != "*" {
		t.Errorf("want: %s, got: %s", "*", cros)
	}
}
