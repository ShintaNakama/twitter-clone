package client

import (
	"fmt"
	"time"

	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/backoff"
	"google.golang.org/grpc/balancer/roundrobin"
	"google.golang.org/grpc/resolver"
	grpctrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/google.golang.org/grpc"
)

const defaultGrpcMaxDelay = 1 * time.Second

func gRPCConn(name string, port int) (*grpc.ClientConn, error) {
	//ns := "default"

	//dns := fmt.Sprintf("%s.%s.svc.cluster.local:%d", name, ns, port)
	dns := fmt.Sprintf("localhost:%d", port)
	zap.L().Info("dial DNS info", zap.String("service", name), zap.String("dns", dns))

	// gRPCコネクション失敗時、1秒待機してリトライする.
	bc := backoff.DefaultConfig
	bc.MaxDelay = defaultGrpcMaxDelay
	connParams := grpc.ConnectParams{Backoff: bc}

	resolver.SetDefaultScheme("dns")

	cc, err := grpc.Dial(
		dns,
		grpc.WithInsecure(),
		// client LB RoundRobinを設定
		grpc.WithDefaultServiceConfig(fmt.Sprintf(`{"LoadBalancingPolicy": "%s"}`, roundrobin.Name)),
		// gRPC Clientのoptionを設定
		grpc.WithConnectParams(connParams),
		// datadogのtraceを設定
		grpc.WithUnaryInterceptor(grpctrace.UnaryClientInterceptor()),
		// error messageをラップしないようにする
		grpc.FailOnNonTempDialError(true),
		// 接続が確立するまでブロックする
		//grpc.WithBlock(),
	)
	if err != nil {
		return nil, err
	}

	return cc, nil
}
