package main

import (
	"fmt"
	"net/http"
	"net/http/cgi"

	"github.com/gorilla/securecookie"
)

var hashKey = []byte("b0rg&(G(EvuxiebBXOEvijbxe2jbk2jx")
var blockKey = []byte("@OBFubce0bCBE^&*Dj0beobexojbebAn")
var s = securecookie.New(hashKey, blockKey)

func main() {
	value := map[string]string{
		"foo": "bar",
	}
	err := cgi.Serve(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if encoded, err := s.Encode("test-cookie", value); err == nil {
			cookie := &http.Cookie{
				Name:     "test-cookie",
				Value:    encoded,
				Path:     "/cgi-bin/",
				Secure:   true,
				HttpOnly: true,
			}
			http.SetCookie(w, cookie)
		}
		header := w.Header()
		header.Set("Content-Type", "text/plain; charset=utf-8")
		fmt.Fprintln(w, "It worked!")
	}))

	if err != nil {
		fmt.Println(err)
	}
}
