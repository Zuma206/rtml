package router

import "strings"

type Router[T any] struct {
	children    map[string]*Router[T]
	paramRouter *Router[T]
	index       *SetPair[T]
}

func New[T any]() *Router[T] {
	return &Router[T]{
		children: map[string]*Router[T]{},
	}
}

type GetResult[T any] struct {
	Value T
	Found bool
}

func (router *Router[T]) Get(path string) *GetResult[T] {
	if isIndexRoute(path) {
		if router.index == nil {
			return &GetResult[T]{Found: false}
		}
		return &GetResult[T]{router.index.value, true}
	}
	segment, rest := getSegment(path)
	child, ok := router.children[segment]
	if !ok {
		if router.paramRouter == nil {
			return &GetResult[T]{Found: false}
		}
		return router.paramRouter.Get(rest)
	}
	return child.Get(rest)
}

func getSegment(route string) (string, string) {
	segment, rest, _ := strings.Cut(strings.TrimPrefix(route, "/"), "/")
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

type SetPair[T any] struct {
	route string
	value T
}

func (router *Router[T]) Set(route string, value T) {
	router.set(&SetPair[T]{route, value})
}

func (router *Router[T]) set(setPair *SetPair[T]) {
	if isIndexRoute(setPair.route) {
		router.index = setPair
		return
	}
	segment, rest := getSegment(setPair.route)
	setPair.route = rest
	if _, isParam := strings.CutPrefix(segment, "$"); isParam {
		if router.paramRouter == nil {
			router.paramRouter = New[T]()
		}
		router.paramRouter.set(setPair)
		return
	}
	router.getChild(segment).set(setPair)
}
