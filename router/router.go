package router

import "strings"

type Router[T any] struct {
	children map[string]*Router[T]
	index    *T
}

func New[T any]() *Router[T] {
	return &Router[T]{
		children: map[string]*Router[T]{},
	}
}

func (router *Router[T]) Get(path string) (T, bool) {
	panic("TODO")
}

func getSegment(route string) (string, string) {
	segment := strings.TrimLeftFunc(route, func(r rune) bool { return r != '/' })
	rest := strings.TrimPrefix(route, segment)
	return segment, rest
}

func isIndexRoute(path string) bool {
	return path == "/" || path == ""
}

func (router *Router[T]) getChild(segment string) *Router[T] {
	child, ok := router.children[segment]
	if !ok {
		router.children[segment] = New[T]()
		child = router.children[segment]
	}
	return child
}

func (router *Router[T]) Set(route string, value T) {
	if isIndexRoute(route) {
		router.index = &value
		return
	}
	segment, rest := getSegment(route)
	router.getChild(segment).Set(rest, value)
}
