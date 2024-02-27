package utils

import (
	"fmt"
	"os"
)

func TreatNilObjectError(obj error, errorMessage string) {
	if obj != nil {
		fmt.Print(errorMessage)
		os.Exit(1)
	}
}
