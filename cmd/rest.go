package main

import (
    "log"

    "bilalekrem.com/pushnotification-service/api/rest"
)

func main() {
    log.Println("Starting server")

    server := rest.NewRestServer()
    server.Start()
}