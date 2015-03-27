package main

import (
    "encoding/json"
)

func Jsonize(data interface{}) []byte {
    jsonstr, _ := json.Marshal(data)
    return jsonstr
}
