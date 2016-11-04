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

		rows, err := db.Query("SELECT * FROM table")
		if err != nil {
			fmt.Fprintln(w, err.Error())
		}

		columns, err := rows.Columns()
		if err != nil {
			fmt.Fprintln(w, err.Error())
		}

		values := make([]sql.RawBytes, len(columns))
		scanArgs := make([]interface{}, len(values))
		for i := range values {
			scanArgs[i] = &values[i]
		}

		for rows.Next() {
			err = rows.Scan(scanArgs...)
			if err != nil {
				fmt.Fprintln(w, err.Error())
			}
			var value string
			for i, col := range values {
				if col == nil {
					value = "NULL"
				} else {
					value = string(col)
				}
				fmt.Fprintln(w, columns[i], ": ", value)
			}
			fmt.Fprintln(w, "-----------------------------------")
		}
		if err = rows.Err(); err != nil {
			fmt.Fprintln(w, err.Error())
		}
	}))

	if err != nil {
		fmt.Println(err)
	}
}
