package main

import (
	"fmt"
	"net/http"
)

var dbname = "/home/paul/src/f-spot/tests/data/f-spot-0.7.0-17.2.db"

func handler(w http.ResponseWriter, r *http.Request) {
	db := NewDb(dbname)
	tablename := r.URL.Path[1:]
	fmt.Fprintf(w, string(Jsonize(db.SelectAllFrom(tablename))))
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":7242", nil)
}
