package handlers

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/zuma206/rtml/router"
)

func FromFileRouter(routes router.Router[string]) http.Handler {
	return WithError(func(w http.ResponseWriter, r *http.Request) error {
		result := routes.Get(r.URL.Path)
		if !result.Found {
			return fmt.Errorf("not found: %s", r.URL.Path)
		}
		file, err := os.Open(result.Value)
		if err != nil {
			return err
		}
		_, err = io.Copy(w, file)
		if err != nil {
			return err
		}
		file.Close()
		return nil
	})
}
