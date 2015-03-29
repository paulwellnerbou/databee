package db

import (
    "testing"
    "net/url"
    "reflect"
)

func Test_SplitPath_withoutQuery(t*testing.T) {
    urlStr := "http://localhost:1234/db/photo_versions"
    parsedUrl, _ := url.Parse(urlStr)
    params := NewParamsFromUrl(parsedUrl)

    expected := Params{Tablename:"photo_versions"}
    assertSameParamsObject(t, params, &expected)
}

func Test_SplitPath_withoutQueryButMorePathElements(t*testing.T) {
    urlStr := "http://localhost:1234/db/photos/photo_versions/id"
    parsedUrl, _ := url.Parse(urlStr)
    params := NewParamsFromUrl(parsedUrl)

    expected := Params{Tablename:"photos"}
    assertSameParamsObject(t, params, &expected)
}

func Test_SplitPath_withQuery(t*testing.T) {
    urlStr := "http://localhost:1234/db/photo_versions?limit=2,3"
    parsedUrl, _ := url.Parse(urlStr)
    params := NewParamsFromUrl(parsedUrl)

    expected := Params{Tablename:"photo_versions", limit: "2,3"}
    assertSameParamsObject(t, params, &expected)
}

func Test_SplitPath_withMultipleSameQueryParams(t*testing.T) {
    urlStr := "http://localhost:1234/db/photo_versions?limit=2,3&limit=5,6"
    parsedUrl, _ := url.Parse(urlStr)
    params := NewParamsFromUrl(parsedUrl)

    expected := Params{Tablename:"photo_versions", limit: "2,3"}
    assertSameParamsObject(t, params, &expected)
}

func assertSameParamsObject(t*testing.T, actual* Params, expected* Params) {
    if actual.Tablename != expected.Tablename {
        t.Errorf("Expected tablename %v but got %v", expected.Tablename, actual.Tablename)
    }
    if !reflect.DeepEqual(actual, expected) {
        t.Errorf("Expected queryparams %#v but got %#v", expected, actual)
    }
}
