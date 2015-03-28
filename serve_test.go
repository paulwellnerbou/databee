package main

import (
    "testing"
    "net/url"
    "reflect"
)

func Test_SplitPath_withoutQuery(t*testing.T) {
    urlStr := "http://localhost:1234/photo_versions"
    parsedUrl, _ := url.Parse(urlStr)
    params := createParamsFromUrl(parsedUrl)

    expected := Params{tablename:"photo_versions"}
    assertSameParamsObject(t, params, expected)
}

func Test_SplitPath_withoutQueryButMorePathElements(t*testing.T) {
    urlStr := "http://localhost:1234/photos/photo_versions/id"
    parsedUrl, _ := url.Parse(urlStr)
    params := createParamsFromUrl(parsedUrl)

    expected := Params{tablename:"photos"}
    assertSameParamsObject(t, params, expected)
}

func Test_SplitPath_withQuery(t*testing.T) {
    urlStr := "http://localhost:1234/photo_versions?limit=2,3"
    parsedUrl, _ := url.Parse(urlStr)
    params := createParamsFromUrl(parsedUrl)

    expected := Params{tablename:"photo_versions", limit: "2,3"}
    assertSameParamsObject(t, params, expected)
}

func Test_SplitPath_withMultipleSameQueryParams(t*testing.T) {
    urlStr := "http://localhost:1234/photo_versions?limit=2,3&limit=5,6"
    parsedUrl, _ := url.Parse(urlStr)
    params := createParamsFromUrl(parsedUrl)

    expected := Params{tablename:"photo_versions", limit: "2,3"}
    assertSameParamsObject(t, params, expected)
}

func assertSameParamsObject(t*testing.T, actual Params, expected Params) {
    if actual.tablename != expected.tablename {
        t.Errorf("Expected tablename %v but got %v", expected.tablename, actual.tablename)
    }
    if !reflect.DeepEqual(actual, expected) {
        t.Errorf("Expected queryparams %#v but got %#v", expected, actual)
    }
}
