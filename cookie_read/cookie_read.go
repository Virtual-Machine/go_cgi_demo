package main

import (
	"fmt"
	"net/http"
	"net/http/cgi"

	"github.com/gorilla/securecookie"
)

var hashKey = []byte("b0rg&(G(EvuxiebBXOEvijbxe2jbk2jx")
var blockKey = []byte("@OBFubce0bCBE^&*Dj0beobexojbebAn")
var s2 = securecookie.New(hashKey, blockKey)

func main() {
	err := cgi.Serve(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		header := w.Header()
		header.Set("Content-Type", "text/plain; charset=utf-8")
		if cookie, err := r.Cookie("test-cookie"); err == nil {
			value := make(map[string]string)
			if err = s2.Decode("test-cookie", cookie.Value, &value); err == nil {
				fmt.Fprintf(w, "The value of foo is %q", value["foo"])
				return
			}
			fmt.Fprintf(w, "Unable to decode cookie value")
			return
		}
		fmt.Fprintf(w, "No cookie present")
	}))

	if err != nil {
		fmt.Println(err)
	}
}
