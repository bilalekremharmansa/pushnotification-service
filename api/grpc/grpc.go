package grpc

import (
	"log"
	"fmt"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"bilalekrem.com/pushnotification-service/internal/config"
	"bilalekrem.com/pushnotification-service/api/grpc/pushservice"
	pb "bilalekrem.com/pushnotification-service/proto/pushservice"
)

type grpcServer struct {
    server *grpc.Server
    enabled bool
    host string
    port int
}

func NewGRPCServerWithConfig(cfg config.GRPCConfig) *grpcServer {
    return NewGRPCServerWithAddress(cfg.GetEnabled(), cfg.GetHost(), cfg.GetPort())
}

func NewGRPCServerWithAddress(enabled bool, host string, port int) *grpcServer {
    gRpcServer := grpc.NewServer()

    return &grpcServer{server: gRpcServer, enabled: enabled, host: host, port:port}
}

func (g *grpcServer) init() {
    reflection.Register(g.server)

    pb.RegisterPushNotificationServiceServer(g.server, pushservice.NewService())
}

func (g *grpcServer) Start() {
    if !g.enabled {
        log.Println("[gRPC] server is disabled, not starting...")
        return
    }

    g.init()

    address := fmt.Sprintf("%s:%d", g.host, g.port)

    lis, err := net.Listen("tcp", address)
	if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }

    log.Printf("[gRPC] server is about to listen on: [%s]", address)

    if err := g.server.Serve(lis); err != nil {
        log.Fatalf("[gRPC] failed to serve: %v", err)
    }
}