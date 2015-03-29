package db

import (
    "net/url"
    "strings"
    "path"
)

type Params struct {
    Tablename string
    limit string
}

func NewParamsFromUrl(requestedUrl *url.URL) (*Params) {
    urlParts := strings.Split(requestedUrl.Path[1:], "?")
    cleanedPath := path.Clean(urlParts[0])
    splitted := strings.Split(cleanedPath, "/")
    m, _ := url.ParseQuery(requestedUrl.RawQuery)
    params := new(Params)
    if len(splitted) > 1 {
        params.Tablename = splitted[1]
    }
    params.mapQueryParams(m)
    return params
}

func (params *Params) mapQueryParams(m map[string][]string) {
    if limit, ok := m["limit"]; ok && len(limit) > 0 {
        params.limit = limit[0]
    }
}

func (params *Params) GetSelectSql() (sql string) {
    sql = "SELECT * FROM " + params.Tablename
    if len(params.limit) > 0 {
        sql += " limit "+params.limit
    }
    return
}

func (params *Params) GetShowTablesSql() (sql string) {
    return "SELECT name FROM sqlite_master WHERE type='table'"
}
