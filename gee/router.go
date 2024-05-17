package gee

import (
	"log"
	"net/http"
)

type Router struct {
	handlers map[string]HandleFunc
}

func NewRouter() *Router {
	return &Router{
		handlers: make(map[string]HandleFunc),
	}
}

func (r *Router) addRoute(method string, pattern string, handler HandleFunc) {
	log.Printf("Route %4s - %s", method, pattern)
	key := method + "-" + pattern
	r.handlers[key] = handler
}

func (r *Router) Handle(c *Context) {
	key := c.Method + "-" + c.Path
	if handler, ok := r.handlers[key]; ok {
		handler(c)
	} else {
		c.String(http.StatusNotFound, "404 NOT FOUND: %s\n", c.Path)
	}
}
