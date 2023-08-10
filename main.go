package main

import (
	"fmt"
	"os"
)

func createDir(path string) {
	if os.MkdirAll(path, 0755) != nil {
		fmt.Printf("couldn't create folder: %s\n", path)
	}
}

func main() {
	for _, name := range os.Args[1:] {
		switch []rune(name)[0] {
		case ':':
			createDir(name[1:])
			break
		default:
			if _, err := os.Stat(name); err == nil {
				fmt.Printf("file already exists: %s\n", name[1:])
			} else if os.WriteFile(name, []byte{}, 0644) != nil {
				fmt.Printf("couldn't create folder: %s\n", name[1:])
			}
			break
		}
	}
}
