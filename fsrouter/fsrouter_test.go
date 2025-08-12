package fsrouter

import (
	"testing"
)

var tests = map[string]string{
	"/about":                      "testdata/about.html",
	"/contact":                    "testdata/contact.html",
	"/":                           "testdata/index.html",
	"/products/game-id-0/preview": "testdata/products/$game/preview.html",
	"/products/340095846/buy":     "testdata/products/$product/buy.html",
	"/products/eoiwijwh/view":     "testdata/products/$product/view.html",
	"/products/all":               "testdata/products/all.html",
	"/users/john/posts":           "testdata/users/$id/posts.html",
	"/users/john/view":            "testdata/users/$id/view.html",
	"/users/aaaaa":                "testdata/users/$id.html",
	"/users/list":                 "testdata/users/list.html",
}

func TestFSRouter(t *testing.T) {
	r := New()
	r.ScanDir("testdata")
	for path, expectedValue := range tests {
		result := r.Get(path)
		if !result.Found {
			t.Errorf("couldn't match path %s to file %s", path, expectedValue)
		} else if result.Value != expectedValue {
			t.Errorf("path %s yielded file %s when %s was expected", path, result.Value, expectedValue)
		}
	}
}
