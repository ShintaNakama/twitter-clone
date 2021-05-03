package rpc

import (
	"context"

	health "google.golang.org/grpc/health/grpc_health_v1"
)

// HealthServer ヘルスチェックサーバー
type HealthServer struct {
	health.HealthServer
}

// Check ヘルスチェックを実行する
func (h *HealthServer) Check(ctx context.Context, _ *health.HealthCheckRequest) (*health.HealthCheckResponse, error) {
	status := health.HealthCheckResponse_NOT_SERVING
	return &health.HealthCheckResponse{
		Status: status,
	}, nil
}
