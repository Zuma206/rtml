package main

import (
	"github.com/zuma206/rtml/utils"
)

func main() {
	utils.HandleError(run())
}

func run() error {
	println("Hello, World!")
	return nil
}
