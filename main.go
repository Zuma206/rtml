package main

import (
	"errors"
	"net/http"
	"os"

	"github.com/zuma206/rtml/fsrouter"
	"github.com/zuma206/rtml/handlers"
	"github.com/zuma206/rtml/utils"
)

func main() {
	utils.PrintError(run())
}

func run() error {
	if len(os.Args) != 2 {
		return errors.New("expected exactly one argument")
	}
	dirname := os.Args[1]
	routes := fsrouter.New()
	if err := routes.ScanDir(dirname); err != nil {
		return err
	}
	http.ListenAndServe(":3030", handlers.FromFileRouter(routes))
	return nil
}
