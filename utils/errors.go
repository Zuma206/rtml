package utils

import (
	"fmt"
	"os"
)

func HandleError(err error) {
	if err != nil {
		fmt.Printf("error: %s", err.Error())
		os.Exit(1)
	}
}
