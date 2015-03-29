package config

import (
    "os"
    "encoding/json"
)

type DatabaseConfig struct {
    Alias string                `json:"alias"`
    DbDriver string             `json:"dbDriver"`
    DbConnectionString string   `json:"dbConnectionString"`
}

type Configuration struct {
    Port string                 `json:"port"`
    Databases []DatabaseConfig  `json:"databases"`
}

// Credits to https://github.com/etsy/Hound/blob/master/config/config.go
func (c *Configuration) LoadFromFile(filename string) error {
    f, err := os.Open(filename)
    if err != nil {
        return err
    }
    defer f.Close()
    if err := json.NewDecoder(f).Decode(c); err != nil {
        return err
    }

    return nil
}
