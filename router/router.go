package router

import "strings"

type Router[T any] interface {
	Get(path string) *GetResult[T]
	Set(route string, value T)
}

type SegmentRouter[T any] struct {
	children    map[string]*SegmentRouter[T]
	paramRouter *SegmentRouter[T]
	index       *SetPair[T]
}

func New[T any]() *SegmentRouter[T] {
	return &SegmentRouter[T]{
		children: map[string]*SegmentRouter[T]{},
	}
}

type GetResult[T any] struct {
	Value  T
	Found  bool
	Params map[string]string
}

type GetLookup[T any] struct {
	path   string
	params []string
}

func (router *SegmentRouter[T]) Get(path string) *GetResult[T] {
	return router.get(&GetLookup[T]{path, []string{}})
}

func newGetResultSuccess[T any](value T, paramNames []string, paramValues []string) *GetResult[T] {
	params := map[string]string{}
	for i, name := range paramNames {
		params[name] = paramValues[i]
	}
	return &GetResult[T]{value, true, params}
}

func (router *SegmentRouter[T]) get(lookup *GetLookup[T]) *GetResult[T] {
	if isIndexRoute(lookup.path) {
		if router.index == nil {
			return &GetResult[T]{Found: false}
		}
		return newGetResultSuccess(router.index.value, router.index.paramNames, lookup.params)
	}
	segment, rest := getSegment(lookup.path)
	child, ok := router.children[segment]
	lookup.path = rest
	if !ok {
		if router.paramRouter == nil {
			return &GetResult[T]{Found: false}
		}
		lookup.params = append(lookup.params, segment)
		return router.paramRouter.get(lookup)
	}
	return child.get(lookup)
}

func getSegment(route string) (string, string) {
	segment, rest, _ := strings.Cut(strings.TrimPrefix(route, "/"), "/")
	return segment, rest
}

func isIndexRoute(path string) bool {
	return path == "/" || path == ""
}

func (router *SegmentRouter[T]) getChild(segment string) *SegmentRouter[T] {
	child, ok := router.children[segment]
	if !ok {
		router.children[segment] = New[T]()
		child = router.children[segment]
	}
	return child
}

type SetPair[T any] struct {
	route      string
	value      T
	paramNames []string
}

func (router *SegmentRouter[T]) Set(route string, value T) {
	router.set(&SetPair[T]{route, value, []string{}})
}

func (router *SegmentRouter[T]) set(setPair *SetPair[T]) {
	if isIndexRoute(setPair.route) {
		router.index = setPair
		return
	}
	segment, rest := getSegment(setPair.route)
	setPair.route = rest
	if paramName, isParam := strings.CutPrefix(segment, "$"); isParam {
		if router.paramRouter == nil {
			router.paramRouter = New[T]()
		}
		setPair.paramNames = append(setPair.paramNames, paramName)
		router.paramRouter.set(setPair)
		return
	}
	router.getChild(segment).set(setPair)
}
