package router

import "testing"

var staticRoutes = []string{
	"/",
	"/test",
	"/a-path",
	"/a-path/another",
}

func TestStaticRoutes(t *testing.T) {
	// Register all routes
	router := New[int]()
	for index, route := range staticRoutes {
		router.Set(route, index)
	}
	// Test all routes
	for expectedIndex, route := range staticRoutes {
		if index, ok := router.Get(route); !ok || index != expectedIndex {
			t.Errorf("failed to match route %s", route)
		}
	}
}
