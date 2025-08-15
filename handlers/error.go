package handlers

import (
	"io"
	"net/http"
	"os"

	"github.com/zuma206/rtml/utils"
)

func WithError(handler func(w http.ResponseWriter, r *http.Request) error) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		err := handler(w, r)
		if err == nil {
			return
		}
		utils.FprintError(io.MultiWriter(w, os.Stderr), err)
	})
}
