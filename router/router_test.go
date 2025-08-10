package router

import (
	"strconv"
	"testing"
)

var staticRoutes = []string{
	"/",
	"/test",
	"/a-path",
	"/a-path/another",
	"/a-path/thing",
}

var paramRoutes = map[string]struct {
	route string
	tests []string
}{
	"friends-id": {"/friends/$id", []string{
		"/friends/friend-0",
		"/friends/f4",
		"/friends/a-person",
	}},
	"view-blog": {"/blogs/$key/view", []string{
		"/blogs/blog-0/view",
		"/blogs/my-post/view",
		"/blogs/an-article/view",
		"/blogs/example/view",
	}},
	"buy-it-now": {"/products/item/$pid/buy", []string{
		"/products/item/product_0/buy",
		"/products/item/3294536/buy",
		"/products/item/some-item/buy",
	}},
}

func newTestRouter() *Router[string] {
	// Register all routes
	router := New[string]()
	for index, route := range staticRoutes {
		router.Set(route, strconv.Itoa(index))
	}
	// Add param routes
	for index, routeInfo := range paramRoutes {
		router.Set(routeInfo.route, index)
	}
	return router
}

func TestStaticRoutes(t *testing.T) {
	router := newTestRouter()
	// Test all routes
	for expectedIndex, route := range staticRoutes {
		if index, ok := router.Get(route); !ok || index != strconv.Itoa(expectedIndex) {
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

func TestParamRoutes(t *testing.T) {
	router := newTestRouter()
	for expectedIndex, routeInfo := range paramRoutes {
		for _, test := range routeInfo.tests {
			if index, ok := router.Get(test); !ok || index != expectedIndex {
				t.Errorf("failed to match %s to param route %s", test, routeInfo.route)
			}
		}
	}
}
