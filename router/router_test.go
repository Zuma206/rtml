package router

import "testing"

var staticRoutes = []string{
	"/",
	"/test",
	"/a-path",
	"/a-path/another",
}

func newTestRouter() *Router[int] {
	// Register all routes
	router := New[int]()
	for index, route := range staticRoutes {
		router.Set(route, index)
	}
	return router
}

func TestStaticRoutes(t *testing.T) {
	router := newTestRouter()
	// Test all routes
	for expectedIndex, route := range staticRoutes {
		if index, ok := router.Get(route); !ok || index != expectedIndex {
			t.Errorf("failed to match route %s", route)
		}
	}
}

var notFoundRoutes = []string{
	"/my-page",
	"/another-page",
	"/this/page/doesnt/exist",
	"/test/not",
	"/a-path/never/a/path",
}

func TestRouterFailure(t *testing.T) {
	router := newTestRouter()
	for _, route := range notFoundRoutes {
		_, ok := router.Get(route)
		if ok {
			t.Errorf("router unexpectedly matched %s", route)
		}
	}
}
