package handlers

import (
	"fmt"
	"io/fs"
	"net/http"
	"os"
	"sync"

	"github.com/zuma206/rtml/router"
	"github.com/zuma206/rtml/runtime"
)

func FromFileRouter(routes router.Router[string]) http.Handler {
	rt := struct {
		*runtime.Runtime
		*sync.Mutex
	}{runtime.New(), &sync.Mutex{}}
	return WithError(func(w http.ResponseWriter, r *http.Request) error {
		file, err := lookupFile(routes, r.URL.Path)
		if err != nil {
			return err
		}
		defer file.Close()
		rt.Lock()
		defer rt.Unlock()
		rt.Output = w
		if err := rt.RunCode(file); err != nil {
			return err
		}
		return nil
	})
}

func lookupFile(routes router.Router[string], path string) (fs.File, error) {
	result := routes.Get(path)
	if !result.Found {
		return nil, fmt.Errorf("not found: %s", path)
	}
	file, err := os.Open(result.Value)
	if err != nil {
		return nil, err
	}
	return file, nil
}
