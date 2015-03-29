package main

import (
    "testing"
)

func Test_Exists_exists(t*testing.T) {
    f := "serve.go"
    if !Exists(f) {
        t.Errorf("Expected file %s to exist.")
    }
}

func Test_Exists_notexists(t*testing.T) {
    f := "inexistantfile"
    if Exists(f) {
        t.Errorf("Not expected file %s to exist.")
    }
}

func Test_detectConfigFileToUse(t*testing.T) {
	args := []string{"./..", "--config=config-example.json"}
    expected := "config-example.json"
    actual := detectConfigFileToUse(args)

    if actual != expected {
        t.Errorf("Expected %s but got %s", expected, actual)
    }
}
