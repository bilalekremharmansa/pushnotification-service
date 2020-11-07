package router

// impl ref: https://github.com/moby/moby/blob/master/api/server/router/local.go

import (
    "net/http"
)

type Router interface {
	Routes() []Route
}

type Route interface {
	Method() string
	Path() string
	Handler() http.HandlerFunc
}

type route struct {
	method  string
	path    string
	handler http.HandlerFunc
}

func (r route) Path() string {
    return r.path
}

func (r route) Method() string {
    return r.method
}

func (r route) Handler() http.HandlerFunc {
    return r.handler
}

func NewRoute(method, path string, handler http.HandlerFunc) Route {
	return &route{method, path, handler}
}

func NewGetRoute(path string, handler http.HandlerFunc) Route {
	return NewRoute(http.MethodGet, path, handler)
}

func NewPostRoute(path string, handler http.HandlerFunc) Route {
	return NewRoute(http.MethodPost, path, handler)
}

