package utils

import (
	"fmt"
	"os"
)

func CreateDir(path string) {
	if os.MkdirAll(path, 0755) != nil {
		fmt.Printf("couldn't create folder: %s\n", path)
	}
}

func CreateFile(path string) {
	if _, err := os.Stat(path); err == nil {
		fmt.Printf("path already exists: %s\n", path)
	} else if os.WriteFile(path, []byte{}, 0644) != nil {
		fmt.Printf("couldn't create file: %s\n", path)
	}
}
