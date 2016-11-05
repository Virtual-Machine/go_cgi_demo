package main

import (
	"fmt"
	"net/http"
	"net/http/cgi"
)

func main() {
	body := `<h2>Hello</h2><p style='color: red;'>This is a CGI scripted response</p>`

	err := cgi.Serve(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		header := w.Header()
		header.Set("Content-Type", "text/html; charset=UTF-8")
		fmt.Fprintln(w, body)
	}))

	if err != nil {
		fmt.Println(err)
	}
}
