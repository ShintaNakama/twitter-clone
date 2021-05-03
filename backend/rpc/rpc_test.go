package rpc_test

import (
	"testing"

	"github.com/ShintaNakama/twitter-clone/backend/rpc"
	"google.golang.org/grpc"
)

func TestNewServer(t *testing.T) {
	got := rpc.NewServer("name", 3000)
	if got.Name != "name" {
		t.Errorf("name want: %s, got: %s", "name", got.Name)
	}
	if got.Port != 3000 {
		t.Errorf("port port: %d, got: %d", 3000, got.Port)
	}
}

func TestRun(t *testing.T) {
	s := rpc.NewServer("name", 9001)
	s.Run(nil, terminate)
	t.Log("without any errors.")
}

func terminate(s *grpc.Server) {
	s.GracefulStop()
}
