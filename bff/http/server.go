package http

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"go.uber.org/zap"
)

const limitTime = 60 * time.Second

// Server HTTP server.
type Server struct {
	name   string
	port   int
	server *http.Server
}

// NewServer returns HTTPServer.
func NewServer(name string, port int, handler http.Handler) *Server {
	sv := &Server{
		name: name,
		port: port,
	}

	sv.server = &http.Server{
		Addr:              fmt.Sprintf(":%d", sv.port),
		Handler:           handler,
		ReadHeaderTimeout: limitTime,
	}

	return sv
}

// Run start server.
func (sv *Server) Run() {
	go func() {
		if err := sv.server.ListenAndServe(); err != nil {
			zap.L().Error("failed to ListenAndAServe", zap.Error(err))
		}
	}()
}

func (sv *Server) Terminate() {
	ctx, cancel := context.WithTimeout(context.Background(), limitTime)
	defer cancel()

	if err := sv.server.Shutdown(ctx); err != nil {
		zap.L().Error("failed to gracefully shutdown", zap.Error(err))
		sv.server.Close()
		zap.L().Info("server closed")
	}

	zap.L().Info("server shutdown")
}
