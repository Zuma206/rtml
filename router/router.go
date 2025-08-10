package router

type Router[T any] struct{}

func New[T any]() *Router[T] {
	return &Router[T]{}
}

func (router *Router[T]) Get(path string) (T, bool) {
	panic("TODO")
}

func (router *Router[T]) Set(route string, value T) {
	panic("TODO")
}
