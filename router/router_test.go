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
		if result := router.Get(route); !result.Found || result.Value != strconv.Itoa(expectedIndex) {
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
		result := router.Get(route)
		if result.Found {
			t.Errorf("router unexpectedly matched %s", route)
		}
	}
}

func TestParamRoutes(t *testing.T) {
	router := newTestRouter()
	for expectedIndex, routeInfo := range paramRoutes {
		for _, test := range routeInfo.tests {
			if result := router.Get(test); !result.Found || result.Value != expectedIndex {
				t.Errorf("failed to match %s to param route %s", test, routeInfo.route)
			}
		}
	}
}

var paramRouteValues = []struct {
	route  string
	path   string
	values map[string]string
}{
	{"/profile/$profile/followers/$follower", "/profile/user-0/followers/user-1", map[string]string{
		"profile":  "user-0",
		"follower": "user-1",
	}},
	{"/profile/$admin/privileges", "/profile/main-man/privileges", map[string]string{
		"admin": "main-man",
	}},
}

func TestParamRouteValues(t *testing.T) {
	router := New[int]()
	for _, test := range paramRouteValues {
		router.Set(test.route, 0)
	}
	for _, test := range paramRouteValues {
		result := router.Get(test.path)
		if !result.Found {
			t.Errorf("failed to match %s to %s\n", test.path, test.route)
			continue
		}
		for key, expectedValue := range result.Params {
			value, ok := result.Params[key]
			if !ok {
				t.Errorf("missing param %s from route %s using path %s\n", key, test.route, test.path)
				continue
			}
			if value != expectedValue {
				t.Errorf("expected param %s to be %s for path %s on route %s but got %s instead\n", key, expectedValue, test.path, test.route, value)
				continue
			}
		}
	}
}
