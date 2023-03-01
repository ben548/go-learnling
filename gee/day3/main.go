package main

import (
	"go-learnling/7days/gee/day2/gee"
	"net/http"
)

func main() {
	e := gee.New()
	e.GET("/", func(c *gee.Context) {
		c.Html(http.StatusOK, "<h1>hello gee</h1>")
	})
	e.GET("/hello", func(c *gee.Context) {
		c.String(http.StatusOK, "hello %s, you're visiting %s\n", c.Query("name"), c.Path)
	})
	e.POST("/login", func(c *gee.Context) {
		c.Json(http.StatusOK, gee.H{
			"username": c.PostForm("username"),
			"passport": c.PostForm("passport"),
		})
	})
	e.Run(":9999")

}
