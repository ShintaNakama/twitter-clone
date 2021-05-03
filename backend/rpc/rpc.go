package rpc

import (
	"fmt"
	"net"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	health "google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"
	grpctrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/google.golang.org/grpc"
)

// Register サービスレジスタ
type Register func(s *grpc.Server)

// Server gRPCサーバー
type Server struct {
	Name string
	Port int
}

// NewServer 新しいServerを生成する
func NewServer(name string, port int) *Server {
	return &Server{
		Name: name,
		Port: port,
	}
}

// Run サーバーを起動する
func (s *Server) Run(register Register, terminator func(server *grpc.Server)) {
	server := s.newServer()
	if register != nil {
		register(server)
	}
	health.RegisterHealthServer(server, &HealthServer{})
	reflection.Register(server)

	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", s.Port))
	if err != nil {
		zap.L().Fatal("Failed to listen", zap.Int("port", s.Port))
	}
	defer listen.Close()
	zap.L().Info("gRPC server is running.")

	if terminator != nil {
		go terminator(server)
	}

	if err := server.Serve(listen); err != nil {
		zap.L().Error("An error occured", zap.Error(err))
	}
}

func (s *Server) newServer() *grpc.Server {
	return grpc.NewServer(
		grpc_middleware.WithUnaryServerChain(
			//grpc_validator.UnaryServerInterceptor(),
			grpc_recovery.UnaryServerInterceptor(),
			grpctrace.UnaryServerInterceptor(grpctrace.WithServiceName(s.Name)),
			grpc_zap.UnaryServerInterceptor(zap.L()),
		),
	)
}
