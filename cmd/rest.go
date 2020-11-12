package main

import (
    "log"

    "bilalekrem.com/pushnotification-service/internal/config"
    "bilalekrem.com/pushnotification-service/api/rest"
)

func main() {
    log.Println("Starting [rest] server")

    config.InitDefaultConfig()
    restConfig := config.GetAppConfig().GetServersConfig().GetRestConfig()

    server := rest.NewRestServerWithConfig(restConfig)
    server.Start()
}