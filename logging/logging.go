package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/cgi"
	"os"
	"path/filepath"
)

func main() {
	err := cgi.Serve(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fileP, err := filepath.Abs("../../tmp/error_log")

		if err != nil {
			log.Fatalf("error generating file path: %v", err)
		}

		f, err := os.OpenFile(fileP, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

		if err != nil {
			log.Fatalf("error opening file: %v", err)
		}
		defer f.Close()

		log.SetOutput(f)
		log.Println("This is a test log entry")

		header := w.Header()
		header.Set("Content-Type", "text/plain; charset=utf-8")
		fmt.Fprintln(w, "It worked!")
	}))

	if err != nil {
		fmt.Println(err)
	}
}
