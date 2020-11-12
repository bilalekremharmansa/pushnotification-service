package main

import (
    "log"

    "bilalekrem.com/pushnotification-service/internal/config"
    "bilalekrem.com/pushnotification-service/api/grpc"
)

func main() {
    log.Println("Starting gRPC server")

    config.InitDefaultConfig()
    gRPCConfig := config.GetAppConfig().GetServersConfig().GetGRPCConfig()

    server := grpc.NewGRPCServerWithConfig(gRPCConfig)
    server.Start()
}