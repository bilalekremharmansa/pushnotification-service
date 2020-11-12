package rest

import (
    "fmt"
    "log"
    "net/http"

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
    enabled bool
    host string
    port int
}

func NewRestServerWithAddress(enabled bool, host string, port int) *Server {
    muxRouter := mux.NewRouter().StrictSlash(true)
    muxRouter.Use(restMiddleware)

    server := Server{muxRouter: muxRouter, enabled: enabled, host: host, port: port}

    return &server
}

func NewRestServerWithConfig(cfg config.RestConfig) *Server {
    return NewRestServerWithAddress(cfg.GetEnabled(), cfg.GetHost(), cfg.GetPort())
}

func (server *Server) init() {
    server.registerRoutes(mgmt.Routes())

    r := push.NewRouter()
    server.registerRoutes(r.Routes())
}

func (server *Server) Start() {
    if !server.enabled {
        log.Println("[rest] server is disabled, not starting...")
        return
    }

    server.init()

    addr := fmt.Sprintf("%s:%d", server.host, server.port)

    log.Printf("[rest] server is about to listen on: [%s]", addr)
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