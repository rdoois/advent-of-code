package utils

import (
	"fmt"
	"os"
)

func ReadInput(filename string) string {
	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Printf("error reading file '%s'\n%s\n", filename, err)
		return ""
	}
	return string(data)
}
