package gee

import (
	"net/http"
	"strings"
)

type route struct {
	roots   map[string]*node
	handles map[string]HandlerFunc
}

func newRoute() *route {
	return &route{
		roots:   make(map[string]*node),
		handles: make(map[string]HandlerFunc),
	}
}

func parsePattern(pattern string) []string {
	vs := strings.Split(pattern, "/")

	parts := make([]string, 0)
	for _, item := range vs {
		if item != "" {
			parts = append(parts, item)
			if string(item[0]) == "*" {
				break
			}
		}
	}
	return parts
}

func (r *route) addRoute(method string, pattern string, handler HandlerFunc) {
	parts := parsePattern(pattern)
	key := method + " - " + pattern
	_, ok := r.roots[method]
	if !ok {
		r.roots[method] = &node{}
	}
	r.roots[method].insert(pattern, parts, 0)
	r.handles[key] = handler
}

func (r *route) getRoute(method string, path string) (*node, map[string]string) {
	//fmt.Println(method, path)
	searchParts := parsePattern(path)
	params := make(map[string]string)
	root, ok := r.roots[method]
	if !ok {
		return nil, nil
	}

	//spew.Dump(searchParts)
	n := root.search(searchParts, 0)

	//spew.Dump(n)

	if n != nil {
		parts := parsePattern(n.pattern)
		for index, part := range parts {
			if string(part[0]) == ":" {
				params[part[1:]] = searchParts[index]
			}
			if string(part[0]) == "*" && len(part) > 1 {
				params[part[1:]] = strings.Join(searchParts[index:], "/")
				break
			}
		}
		return n, params
	}

	return nil, nil
}

func (r *route) handle(c *Context) {
	n, params := r.getRoute(c.Method, c.Path)
	if n != nil {
		c.Params = params
		key := c.Method + " - " + n.pattern
		r.handles[key](c)
	} else {
		c.String(http.StatusNotFound, "404 NOT FOUND %s\n", c.Path)
	}
}
