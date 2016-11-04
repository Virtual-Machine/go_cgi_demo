package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"net/http/cgi"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	err := cgi.Serve(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		header := w.Header()
		header.Set("Content-Type", "text/plain; charset=utf-8")
		db, err := sql.Open("mysql", "USERNAME:PASSWORD@/DATABASE")
		if err != nil {
			fmt.Fprintln(w, err.Error())
		}
		defer db.Close()

		stmtIns, err := db.Prepare("INSERT INTO table VALUES( ?, ?, ?, ? )")
		if err != nil {
			fmt.Fprintln(w, err.Error())
		}
		defer stmtIns.Close()

		_, err = stmtIns.Exec(PARAM1, PARAM2, PARAM3, PARAM4)
		if err != nil {
			fmt.Fprintln(w, err.Error())
			return
		}
		fmt.Fprintln(w, "Success")
	}))

	if err != nil {
		fmt.Println(err)
	}
}
