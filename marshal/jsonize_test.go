package marshal

import (
    "testing"
)

func Test_Jsonize_List(t* testing.T) {
    stringlist := []string{"A", "B", "C"}
    json := Jsonize(stringlist)
    expected := "[\"A\",\"B\",\"C\"]"
    if(string(json) != expected) {
        t.Errorf("Unexpected JSON: %s, expected: %s", string(json), expected)
    }
}

func Test_Jsonize_ListWithMapsEmpty(t* testing.T) {
    input := []map[string]string{}
    json := Jsonize(input)
    expected := "[]"
    if(string(json) != expected) {
        t.Errorf("Unexpected JSON: %s, expected: %s", string(json), expected)
    }
}

func Test_Jsonize_ListWithMaps(t* testing.T) {
    input := []map[string]string{}
    m1 := make(map[string]string)
    m1["key"] = "value"
    m1["anotherkey"] = "value"
    m2 := make(map[string]string)
    m2["key"] = "value"
    input = append(input, m1, m2)
    json := Jsonize(input)
    expected := "[{\"anotherkey\":\"value\",\"key\":\"value\"},{\"key\":\"value\"}]"
    if(string(json) != expected) {
        t.Errorf("Unexpected JSON: %s, expected: %s", string(json), expected)
    }
}
