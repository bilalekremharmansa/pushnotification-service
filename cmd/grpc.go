package main

import (
    "log"

    "bilalekrem.com/pushnotification-service/internal/config"
    "bilalekrem.com/pushnotification-service/api/grpc"
)

func main() {
    log.Println("Starting gRPC server")

    config.InitConfig("/tmp/config.yaml")

    server := grpc.NewServer()
    server.Start()
}