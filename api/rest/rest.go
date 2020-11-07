package rest

import (
    "fmt"
    "log"
    "net/http"
    "strconv"

    "github.com/gorilla/mux"

    "bilalekrem.com/pushnotification-service/internal/config"

    "bilalekrem.com/pushnotification-service/api/rest/router"

    "bilalekrem.com/pushnotification-service/api/rest/mgmt"
    "bilalekrem.com/pushnotification-service/api/rest/push"
)

var (
    DEFAULT_HOST = ""
    DEFAULT_PORT = 8080
)

type Server struct {
    muxRouter *mux.Router
    host string
    port int
}

func NewRestServer() *Server {
    return NewRestServerWithAddress(DEFAULT_HOST, DEFAULT_PORT)
}

func NewRestServerWithAddress(host string, port int) *Server {
    muxRouter := mux.NewRouter().StrictSlash(true)
    muxRouter.Use(restMiddleware)

    server := Server{muxRouter: muxRouter, host: host, port: port}

    server.init()

    return &server
}

func NewRestServerWithConfig() *Server {
    cfg := config.GetConfig()
    serverConfig := cfg.ServerConfig

    port, err := strconv.Atoi(serverConfig.Port)
    if err != nil {
        log.Fatalf("config -- server port is not decimal: [%s]", serverConfig.Port)
    }

    return NewRestServerWithAddress(serverConfig.Host, port)
}

func (server *Server) init() {
    server.registerRoutes(mgmt.Routes())

    r := push.NewRouter()
    server.registerRoutes(r.Routes())
}

func (server *Server) Start() {
    addr := fmt.Sprintf("%s:%d", server.host, server.port)

    log.Fatal(http.ListenAndServe(addr, server.muxRouter))
}

func (server *Server) registerRoutes(routes []router.Route) {
    for _, route := range routes {
        log.Printf("Initializing route -> [%s] %s", route.Method(), route.Path())
        server.
            muxRouter.
            Methods(route.Method()).
            Path(route.Path()).
            HandlerFunc(route.Handler())
    }
}

func restMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        w.Header().Add("Content-Type", "application/json; charset=UTF-8")
        next.ServeHTTP(w, r)
    })
}