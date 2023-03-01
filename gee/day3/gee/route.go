package gee

import "net/http"

type route struct {
	handles map[string]HandlerFunc
}

func newRoute() *route {
	return &route{handles: make(map[string]HandlerFunc)}
}

func (r *route) addRoute(method string, pattern string, handler HandlerFunc) {
	key := method + " - " + pattern
	r.handles[key] = handler
}

func (r *route) handle(c *Context) {
	path := c.Path
	method := c.Method
	if handler, ok := r.handles[method+" - "+path]; ok {
		handler(c)
	} else {
		c.String(http.StatusNotFound, "404 NOT FOUND %s\n", c.Path)
	}
}
