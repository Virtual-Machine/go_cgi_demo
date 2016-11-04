package main

import (
	"fmt"
	"net/http"
	"net/http/cgi"
)

func main() {
	err := cgi.Serve(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		header := w.Header()
		header.Set("Content-Type", "text/plain; charset=utf-8")
		fmt.Fprintln(w, "It worked!")
	}))

	if err != nil {
		fmt.Println(err)
	}
}
