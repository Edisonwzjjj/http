package gee

import "net/http"

type HandleFunc func(*Context)
type Engine struct {
	router *Router
}

func NewEngine() *Engine {
	return &Engine{
		router: NewRouter(),
	}
}

func (e *Engine) AddNewRouter(method string, pattern string, handler HandleFunc) {
	e.router.addRoute(method, pattern, handler)
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
	c := NewContext(w, r)
	e.router.Handle(c)
}
