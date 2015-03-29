package marshal

import (
    "encoding/json"
)

type JsonizeErr struct {
    err error
}

func Jsonize(data interface{}) []byte {
    jsonstr, err := json.Marshal(data)
    if (err != nil) {
        jsonstr, _ = json.Marshal(JsonizeErr{err: err})
    }
    return jsonstr
}
