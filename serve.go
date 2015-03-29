package main

import (
	"fmt"
	"net/http"
	"log"
	"github.com/paulwellnerbou/db-json-server/db"
	"github.com/paulwellnerbou/db-json-server/marshal"
)

const dbname = "/home/paul/src/f-spot/tests/data/f-spot-0.7.0-17.2.db"

func handler(w http.ResponseWriter, r *http.Request) {
	configuredDatabase := db.NewDb(dbname)
	log.Printf("Got request %#v", r.URL.Path)
	params := db.NewParamsFromUrl(r.URL)
	log.Printf("Got params %#v", params)
	data, err := configuredDatabase.SelectAllFrom(params)
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
	} else {
		fmt.Fprint(w, string(marshal.Jsonize(data)))
	}
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":7242", nil)
}
