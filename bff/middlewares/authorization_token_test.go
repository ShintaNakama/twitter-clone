package middlewares_test

import (
	"context"
	"net/http"
	"testing"

	"github.com/ShintaNakama/twitter-clone/bff/middlewares"

	"github.com/google/go-cmp/cmp"
)

func TestAuthorizationTokenMiddleWare(t *testing.T) {
	reqToken := "testAPItoken"
	req, rec := newTestRequest("GET", "/", "")
	req.Header.Set("Authorization", reqToken)
	middlewares.AuthorizationTokenMiddleWare(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		k, _ := middlewares.GetAuthorizationToken(r.Context())
		if _, err := w.Write([]byte(k)); err != nil {
			t.Error(err)
		}
		w.WriteHeader(http.StatusOK)
	})).ServeHTTP(rec, req)

	if diff := cmp.Diff(rec.Body.String(), reqToken); diff != "" {
		t.Errorf("(-got +want)\n%v", diff)
	}
}

func TestGetAuthorizationToken(t *testing.T) {
	ctx := context.Background()
	token := "testAPItoken"
	ctx = middlewares.SetAuthorizationToken(ctx, token)

	got, err := middlewares.GetAuthorizationToken(ctx)
	if err != nil {
		t.Error(err)
	}
	if diff := cmp.Diff(got, token); diff != "" {
		t.Errorf("(-got +want)\n%v", diff)
	}
}
