package db

import (
    "database/sql"
    _ "github.com/mattn/go-sqlite3"
    "log"
)

type Db struct {
    dbstr string
}

func NewDb(dbstr string) (db *Db) {
    db = new(Db)
    db.dbstr = dbstr
    return db
}

func (db *Db) SelectAllFrom(params* Params) ([]map[string]string, error) {
    sqldb, err := connect(db.dbstr)

    if err != nil {
        return nil, err
    }

    defer sqldb.Close()
    sqlstmt := params.GetSql()
    return db.querySql(sqldb, sqlstmt);
}

func (db *Db) querySql(sqldb *sql.DB, sqlstmt string) ([]map[string]string, error) {
    log.Printf("Executing sql statement %s", sqlstmt)
    rows, err := sqldb.Query(sqlstmt)
    if err != nil {
        return nil, err
    }

    defer rows.Close()

    tabledata := []map[string]string{}
    columns, _ := rows.Columns()
    rawResult := make([][]byte, len(columns))

    // credits to http://stackoverflow.com/questions/14477941/read-select-columns-into-string-in-go/14500756?sgp=2#14500756
    dest := make([]interface{}, len(columns)) // A temporary interface{} slice
    for i, _ := range rawResult {
        dest[i] = &rawResult[i] // Put pointers to each string in the interface slice
    }

    for rows.Next() {
        rowdata := make(map[string]string)
        rows.Scan(dest...)

        for i, raw := range rawResult {
            if raw == nil {
                rowdata[columns[i]] = "\\N"
            } else {
                rowdata[columns[i]] = string(raw)
            }
        }
        tabledata = append(tabledata, rowdata)
    }

    return tabledata, nil
}

func connect(dbfile string) (db *sql.DB, err error) {
    db, err = sql.Open("sqlite3", dbfile)
    if err != nil {
        log.Fatal(err)
    }

    return db, err
}
