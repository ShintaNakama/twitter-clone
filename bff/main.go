package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/ShintaNakama/twitter-clone/bff/client"
	"github.com/ShintaNakama/twitter-clone/bff/http"
	"github.com/ShintaNakama/twitter-clone/bff/resolvers"
	"go.uber.org/zap"
)

const (
	serviceName = "bff"
	port        = 8080
)

func main() {
	logger, _ := zap.NewDevelopment()
	defer func() { _ = logger.Sync() }()

	undo := zap.ReplaceGlobals(logger)
	defer undo()

	c, clientClose, err := client.NewBackendClient("backend", 8000)
	if err != nil {
		zap.L().Fatal("fatal conn backend", zap.Any("err", err))
	}

	r := resolvers.NewResolver(c)
	http.RegisterHandlers(serviceName, r)
	sv := http.NewServer(serviceName, port, nil)

	sv.Run()

	closeService(sv, clientClose)
}

func closeService(sv *http.Server, closeClient func() error) {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, os.Interrupt)

	zap.L().Info(fmt.Sprintf("signal %d received, then shutdown...", <-quit))

	if err := closeClient(); err != nil {
		zap.L().Error("error close grpc client", zap.Any("err", err))
	}

	sv.Terminate()
}
