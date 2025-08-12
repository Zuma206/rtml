package router

import "fmt"

func Print[T any](router *Router[T]) {
	printDepth(router, 0)
}

func printDepth[T any](router *Router[T], depth int) {
	if router.index != nil {
		printMargin(depth)
		println("termination")
	}
	for key, child := range router.children {
		printMargin(depth)
		fmt.Printf("\"%s\"\n", key)
		printDepth(child, depth+1)
	}
	if router.paramRouter != nil {
		printMargin(depth)
		println("catch-all param")
		printDepth(router.paramRouter, depth+1)
	}
}

func printMargin(depth int) {
	for range depth {
		print("\t")
	}
}
