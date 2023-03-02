package gee

import (
	"log"
	"net/http"
	"strings"
)

type HandlerFunc func(*Context)

type RouteGroup struct {
	prefix      string
	middlewares []HandlerFunc
	parent      *RouteGroup
	engine      *Engine
}

type Engine struct {
	route *route
	*RouteGroup
	groups []*RouteGroup
}

func New() *Engine {
	engine := &Engine{route: newRoute()}
	engine.RouteGroup = &RouteGroup{
		engine: engine,
	}
	engine.groups = append(engine.groups, engine.RouteGroup)
	return engine
}

func (g *RouteGroup) Group(prefix string) *RouteGroup {
	engine := g.engine
	newGroup := &RouteGroup{
		prefix: g.prefix + prefix,
		parent: g,
		engine: engine,
	}
	engine.groups = append(engine.groups, newGroup)
	return newGroup
}

func (g *RouteGroup) Use(handlers ...HandlerFunc) {
	g.middlewares = append(g.middlewares, handlers...)
}

func (g *RouteGroup) addRoute(method string, comp string, handler HandlerFunc) {
	pattern := g.prefix + comp
	log.Printf("Route %4s - %s", method, pattern)
	g.engine.route.addRoute(method, pattern, handler)
}

func (g *RouteGroup) GET(pattern string, handler HandlerFunc) {
	g.addRoute("GET", pattern, handler)
}

func (g *RouteGroup) POST(pattern string, handler HandlerFunc) {
	g.addRoute("POST", pattern, handler)
}

func (e *Engine) GET(pattern string, handler HandlerFunc) {
	e.route.addRoute("GET", pattern, handler)
}

func Default() *Engine {
	e := New()
	e.Use(Logger(), Recover())
	return e
}

func (e *Engine) POST(pattern string, handler HandlerFunc) {
	e.route.addRoute("POST", pattern, handler)
}

func (e *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var handlers []HandlerFunc
	c := NewContext(w, r)
	for _, g := range e.groups {
		if strings.HasPrefix(c.Req.URL.Path, g.prefix) {
			handlers = append(handlers, g.middlewares...)
		}
	}
	c.handlers = handlers
	e.route.handle(c)
}

func (e *Engine) Run(addr string) {
	log.Fatal(http.ListenAndServe(addr, e))
}
