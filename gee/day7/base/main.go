package main

import (
	"fmt"
	"net/http"
)

func main() {
	fileServer := http.StripPrefix("/asset/", http.FileServer(http.Dir("../static")))
	http.Handle("/", fileServer)
	err := http.ListenAndServe(":8088", nil)
	if err != nil {
		fmt.Println(err)
	}
}
