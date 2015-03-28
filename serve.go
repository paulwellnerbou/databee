package main

import (
	"fmt"
	"net/http"
	"log"
	"path"
	"strings"
	"net/url"
)

const dbname = "/home/paul/src/f-spot/tests/data/f-spot-0.7.0-17.2.db"

type Params struct {
	tablename string
	limit string
}

func createParamsFromUrl(requestedUrl *url.URL) (Params) {
	urlParts := strings.Split(requestedUrl.Path[1:], "?")
	cleanedPath := path.Clean(urlParts[0])
	splitted := strings.Split(cleanedPath, "/")
	m, _ := url.ParseQuery(requestedUrl.RawQuery)
	params := Params{tablename: splitted[0]}
	params.mapQueryParams(m)
	return params
}

func (params *Params) mapQueryParams(m map[string][]string) {
	if limit, ok := m["limit"]; ok && len(limit) > 0 {
		params.limit = limit[0]
	}
}

func (params *Params) GetSql() (sql string) {
	sql = "SELECT * FROM " + params.tablename
	if len(params.limit) > 0 {
		sql += " limit "+params.limit
	}
	return
}

func handler(w http.ResponseWriter, r *http.Request) {
	db := NewDb(dbname)
	log.Printf("Got request %#v", r.URL.Path)
	params := createParamsFromUrl(r.URL)
	log.Printf("Got params %#v", params)
	data, err := db.SelectAllFrom(params)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
	} else {
		fmt.Fprintf(w, string(Jsonize(data)))
	}
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":7242", nil)
}
