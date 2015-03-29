package main

import (
    "fmt"
    "net/http"
    "log"
    "github.com/paulwellnerbou/db-json-server/db"
    "github.com/paulwellnerbou/db-json-server/marshal"
    "github.com/paulwellnerbou/db-json-server/config"
    "os"
    "strings"
)

const defaultConfigFile = "config.json"
const testConfigFile = "config-example.json"

var configuration config.Configuration

func selectHandler(w http.ResponseWriter, r *http.Request, dbConnectionString string) {
    databaseConnection := db.NewDb(dbConnectionString)
    log.Printf("Got request %#v", r.URL.Path)
    params := db.NewParamsFromUrl(r.URL)
    log.Printf("Got params %#v", params)
    data, err := callDatabase(databaseConnection, params)
    renderResult(w, &data, err)
}

func renderResult(w http.ResponseWriter, data*[]map[string]string, err error) {
    w.Header().Set("Content-Type", "application/json")
    if err != nil {
        log.Printf("ERROR: %v", err.Error())
        http.Error(w, err.Error(), http.StatusInternalServerError)
    } else {
        fmt.Fprint(w, string(marshal.Jsonize(data)))
    }
}

func callDatabase(configuredDatabase *db.Db, params *db.Params) ([]map[string]string, error) {
    if (len(params.Tablename) > 0) {
        return configuredDatabase.SelectAllFrom(params)
    } else {
        return configuredDatabase.ShowTables(params)
    }
}

func configHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    fmt.Fprint(w, string(marshal.Jsonize(configuration)))
}

func detectConfigFileToUse(args []string) (string) {
    for _, arg := range args {
        if strings.HasPrefix(arg, "--config=") {
            configFile := strings.Split(arg, "=")[1]
            if !Exists(configFile) {
                log.Printf("FATAL: Expected config file %s not found, Exiting.", configFile)
                os.Exit(1)
            } else {
                return configFile;
            }
        }
    }

    if !Exists(defaultConfigFile) {
        log.Printf("WARNING: Expected config file %s not found, using %s for demonstration and test purposes.", defaultConfigFile, testConfigFile)
        return testConfigFile
    }

    return defaultConfigFile
}

func main() {
    configFile := detectConfigFileToUse(os.Args)
    log.Printf("Loading configuration from %s", configFile)
    configuration.LoadFromFile(configFile)
    log.Printf("Got configuration %#v", configuration)

    http.HandleFunc("/config", configHandler)

    for _, database := range configuration.Databases {
        log.Printf("Registering handler for %s database %s under /%s", database.DbDriver, database.DbConnectionString, database.Alias)
        handlerFunc := func(w http.ResponseWriter, r *http.Request) {
            log.Printf("Handling /%s for url %s", database.Alias, r.URL.Path)
            selectHandler(w, r, database.DbConnectionString)
        }
        http.HandleFunc("/"+database.Alias, handlerFunc)
        http.HandleFunc("/"+database.Alias+"/", handlerFunc)
    }

    log.Printf("Server up and running under port %s. Go to /config to see the actual configuration of databases.", configuration.Port)
    err := http.ListenAndServe(":"+configuration.Port, nil)
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}

// credits to http://stackoverflow.com/a/12527546/4579247
func Exists(name string) bool {
    if _, err := os.Stat(name); err != nil {
        if os.IsNotExist(err) {
            return false
        }
    }
    return true
}
