package gee

import (
	"fmt"
	"net/http"
)

type HandlerFunc func(http.ResponseWriter, *http.Request)

type Engine struct {
	routes map[string]HandlerFunc
}

func New() *Engine {
	return &Engine{routes:make(map[string]HandlerFunc)}
}

func (e *Engine) addRoute(method string, pattern string, handler HandlerFunc)  {
	key := method + " - "+ pattern
	e.routes[key] = handler
}

func (e *Engine) GET (pattern string, handler HandlerFunc)  {
	e.addRoute("GET", pattern, handler)
}

func (e *Engine) POST (pattern string, handler HandlerFunc)  {
	e.addRoute("POST", pattern, handler)
}

func (e *Engine) ServeHTTP (w http.ResponseWriter, r *http.Request)  {
	path := r.URL.Path
	method := r.Method
	if handler, ok := e.routes[method + " - "+ path]; ok {
		handler(w, r)
	} else {
		fmt.Fprintf(w, "404 not found")
	}
}

func (e *Engine) Run (addr string)  {
	http.ListenAndServe(addr, e)
}