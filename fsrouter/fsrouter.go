package fsrouter

import (
	"os"
	"path"

	"github.com/zuma206/rtml/router"
)

type FSRouter struct {
	*router.Router[string]
}

func New() *FSRouter {
	return &FSRouter{router.New[string]()}
}

func (router *FSRouter) ScanDir(dirPath string) error {
	return router.scanDir(dirPath, "")
}

func (router *FSRouter) scanDir(dirPath string, prefix string) error {
	entries, err := os.ReadDir(dirPath)
	if err != nil {
		return err
	}
	for _, entry := range entries {
		itemPath := path.Join(dirPath, entry.Name())
		route := path.Join(prefix, entry.Name())
		if entry.IsDir() {
			if err := router.scanDir(itemPath, route); err != nil {
				return err
			}
			continue
		}
		router.Set(route, itemPath)
	}
	return nil
}
