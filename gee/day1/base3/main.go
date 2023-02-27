package main

import (
	"7days/gee/day1/base3/gee"
	"fmt"
	"net/http"
)

func main()  {
	e := gee.New()
	e.GET("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "the request path is %s", r.URL.Path)
	})
	e.POST("/hello", func(w http.ResponseWriter, r *http.Request) {
		for k, v := range r.Header {
			fmt.Fprintf(w, "key is %s, value is %s <br/>", k, v)
		}
	})
	e.Run(":9999")

}
