package main

import (
	"fmt"
	"log"
	"net/http"
)

type gee struct {
	
}

func (g *gee) ServeHTTP(w http.ResponseWriter, r *http.Request)  {
	switch r.URL.Path {
	case "/":
		fmt.Fprintf(w, "the request path is %s", r.URL.Path)
	case "/hello":
		for k, v := range r.Header {
			fmt.Fprintf(w, "key is %s, value is %s <br/>", k, v)
		}
	default:
		fmt.Fprintf(w, "404 not found")
	}
}

func main()  {
	g := new(gee)
	log.Fatal(http.ListenAndServe(":9999", g))
}