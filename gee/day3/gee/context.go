package gee

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type H map[string]interface{}

func NewContext(w http.ResponseWriter, r *http.Request) *Context {
	return &Context{
		Req:    r,
		Writer: w,
		Path:   r.URL.Path,
		Method: r.Method,
	}
}

type Context struct {
	Req        *http.Request
	Writer     http.ResponseWriter
	Path       string
	Method     string
	Params     map[string]string
	StatusCode int
}

func (c *Context) Query(key string) string {
	return c.Req.URL.Query().Get(key)
}

func (c *Context) Param(key string) string {
	v, _ := c.Params[key]
	return v
}

func (c *Context) PostForm(key string) string {
	return c.Req.FormValue(key)
}

func (c *Context) SetHeader(key, value string) {
	c.Writer.Header().Add(key, value)
}

func (c *Context) Status(code int) {
	c.StatusCode = code
	c.Writer.WriteHeader(code)
}

func (c *Context) Json(code int, obj interface{}) {
	c.SetHeader("Content-Type", "application/json")
	encoder := json.NewEncoder(c.Writer)
	c.Status(code)
	if err := encoder.Encode(obj); err != nil {
		http.Error(c.Writer, err.Error(), 500)
	}
}

func (c *Context) String(code int, format string, values ...interface{}) {
	c.SetHeader("Content-Type", "text/plain")
	c.Status(code)
	c.Writer.Write([]byte(fmt.Sprintf(format, values...)))
}

func (c *Context) Data(code int, data []byte) {
	c.Status(code)
	c.Writer.Write(data)
}

func (c *Context) Html(code int, html string) {
	c.SetHeader("Content-Type", "text/html")
	c.Status(code)
	c.Writer.Write([]byte(html))
}
