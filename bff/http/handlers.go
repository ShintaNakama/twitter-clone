package http

import (
	"context"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/apollotracing"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/ShintaNakama/twitter-clone/bff/generated"
	"github.com/ShintaNakama/twitter-clone/bff/middlewares"
	"github.com/ShintaNakama/twitter-clone/bff/resolvers"
	"github.com/vektah/gqlparser/v2/gqlerror"
	"go.uber.org/zap"
	httptrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/net/http"
)

// RegisterHandlers リクエストハンドラーを設定する.
func RegisterHandlers(serviceName string, r *resolvers.Resolver) {
	// /graphql
	config := generated.Config{Resolvers: r}
	schema := generated.NewExecutableSchema(config)
	server := handler.NewDefaultServer(schema)
	server.Use(apollotracing.Tracer{})
	server.SetRecoverFunc(func(ctx context.Context, err interface{}) error {
		zap.L().Error("Recoverd panic", zap.Error(err.(error)))
		return &gqlerror.Error{
			Message: "Internal Server Error",
			Extensions: map[string]interface{}{
				"codes": 500,
			},
		}
	})

	gqlHandler := httptrace.WrapHandler(
		middlewares.Apply(
			middlewares.CROS,
			middlewares.AuthorizationTokenMiddleWare,
		)(server),
		serviceName,
		"/graphql",
	)

	http.Handle("/graphql", gqlHandler)
	http.Handle("/health", httptrace.WrapHandler(healthcheck(), serviceName, "/health"))

	http.Handle("/playground", httptrace.WrapHandler(playground.Handler("hinata-app API playground", "/graphql"), serviceName, "/graphql"))
}

// health check handler.
func healthcheck() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		if _, err := w.Write([]byte(`{"status": "OK"}`)); err != nil {
			zap.L().Error("failed to write response", zap.Error(err))
		}
	}
}
