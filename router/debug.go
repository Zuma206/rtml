package router

import "fmt"

func PrintSegmentRouter[T any](router *SegmentRouter[T]) {
	printSegmentRouterAtDepth(router, 0)
}

func printSegmentRouterAtDepth[T any](router *SegmentRouter[T], depth int) {
	if router.index != nil {
		printMargin(depth)
		println("termination")
	}
	for key, child := range router.children {
		printMargin(depth)
		fmt.Printf("\"%s\"\n", key)
		printSegmentRouterAtDepth(child, depth+1)
	}
	if router.paramRouter != nil {
		printMargin(depth)
		println("catch-all param")
		printSegmentRouterAtDepth(router.paramRouter, depth+1)
	}
}

func printMargin(depth int) {
	for range depth {
		print("\t")
	}
}
