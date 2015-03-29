package config

import (
    "testing"
    "os"
    "encoding/json"
)

func Test_LoadFromFile(t*testing.T) {

    wd, _ := os.Getwd()
    t.Logf("Working directory: %s", wd)

    c := new(Configuration)
    err := c.LoadFromFile("testdata/config.json")

    if err != nil {
        t.Errorf("Got error %s reading configuration file", err)
        return
    }

    expectedPort := "7242"
    expectedDatabaseCount := 1
    expectedDatabase := DatabaseConfig{Alias:"db", DbConnectionString:"db/testdata/f-spot-test.db", DbDriver: "sqlite3"}

    t.Logf("Got configuration: %#v", c)
    if (c.Port != expectedPort) {
        t.Errorf("Expected port %v but got %v", expectedPort, c.Port)
    }
    if (len(c.Databases) != expectedDatabaseCount) {
        t.Errorf("Expected database count %v but got %v", expectedDatabaseCount, len(c.Databases))
        return
    }
    if (c.Databases[0] != expectedDatabase) {
        t.Errorf("Expected database %v but got %v", expectedDatabase, c.Databases[0])
    }
}

func Test_GenerateExampleConfig(t*testing.T) {
    dbConfig := DatabaseConfig{Alias:"db", DbConnectionString:"db/testdata/f-spot-test.db", DbDriver: "sqlite3"}
    c := Configuration{Port:"9999", Databases: []DatabaseConfig{dbConfig}}
    t.Logf("Got configuration: %#v", c)
    json, err := json.Marshal(c)
    if (err != nil) {
        t.Errorf("Unable to generate example json: %v", err)
    }
    os.Stdout.Write(json)
}
