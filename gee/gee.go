package gee

import "net/http"

type HandleFunc func(w http.ResponseWriter, r *http.Request)
type Engine struct {
	router map[string]HandleFunc
}

func NewEngine() *Engine {
	return &Engine{
		router: make(map[string]HandleFunc),
	}
}

func (e *Engine) AddNewRouter(method string, pattern string, handler HandleFunc) {
	key := method + "-" + pattern
	e.router[key] = handler
}
func (e *Engine) GET(pattern string, handler HandleFunc) {
	e.AddNewRouter("GET", pattern, handler)
}

func (e *Engine) POST(pattern string, handler HandleFunc) {
	e.AddNewRouter("POST", pattern, handler)
}
func (e *Engine) Run(addr string) (err error) {
	return http.ListenAndServe(addr, e)
}

func (e *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	key := r.Method + "-" + r.URL.Path
	if handler, ok := e.router[key]; ok {
		handler(w, r)
	} else {
		http.NotFound(w, r)
	}
}
