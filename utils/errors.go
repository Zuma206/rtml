package utils

import (
	"fmt"
	"io"
	"os"
)

func FprintError(w io.Writer, err error) {
	fmt.Fprintf(w, "error: %s\n", err.Error())
}

func PrintError(err error) {
	FprintError(os.Stderr, err)
}
