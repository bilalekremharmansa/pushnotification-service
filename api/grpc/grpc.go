package grpc

import (
	"log"
	"fmt"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"bilalekrem.com/pushnotification-service/api/grpc/pushservice"
	pb "bilalekrem.com/pushnotification-service/proto/pushservice"
)

var (
    DEFAULT_HOST = ""
    DEFAULT_PORT = 8080
)

type grpcServer struct {
    server *grpc.Server
    host string
    port int
}

func NewServer() *grpcServer {
    gRpcServer := grpc.NewServer()

    s := &grpcServer{server: gRpcServer, host: DEFAULT_HOST, port:DEFAULT_PORT}

    s.init()

    return s
}

func (g *grpcServer) init() {
    reflection.Register(g.server)

    pb.RegisterPushNotificationServiceServer(g.server, pushservice.NewService())
}

func (g *grpcServer) Start() {
    address := fmt.Sprintf("%s:%d", g.host, g.port)

    lis, err := net.Listen("tcp", address)
	if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }

    log.Printf("server is about to listen on: [%s]", address)

    if err := g.server.Serve(lis); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }
}