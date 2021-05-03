package main

//go:generate statik -m -src=db/schema -dest=.

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/ShintaNakama/twitter-clone/backend/app/interactor"
	"github.com/ShintaNakama/twitter-clone/backend/db"
	"github.com/ShintaNakama/twitter-clone/backend/rpc"
	_ "github.com/ShintaNakama/twitter-clone/backend/statik"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"pb"
)

const (
	serviceName = "backend"
	port        = 8000
)

func main() {
	logger, _ := zap.NewDevelopment()
	defer func() { _ = logger.Sync() }()

	undo := zap.ReplaceGlobals(logger)
	defer undo()

	dbmap, err := db.GetDbMap()
	if err != nil {
		zap.L().Error("failed to connect to mysql.", zap.Error(err))
		return
	}

	defer db.CloseConnection(dbmap.Db)

	i := interactor.NewInteractor(dbmap)
	c := i.NewApi()

	s := rpc.NewServer(serviceName, port)
	s.Run(func(s *grpc.Server) {
		pb.RegisterTwitterCloneServer(s, c.TwitterCloneServer)
	}, terminate)
}

func terminate(s *grpc.Server) {
	q := make(chan os.Signal, 1)
	signal.Notify(q, syscall.SIGTERM, os.Interrupt)
	zap.L().Info(fmt.Sprintf("signal %d received, so shutting down...", <-q))
	s.GracefulStop()
	zap.L().Info("gRPC server is stopped.")
}
