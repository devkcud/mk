package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	for _, name := range os.Args[1:] {
		if strings.HasPrefix(name, "-") {
			if os.MkdirAll(name[1:], 0755) != nil {
				fmt.Printf("couldn't create folder: %s\n", name[1:])
			}
			continue
		}

		if _, err := os.Stat(name); err == nil {
			fmt.Printf("file already exists: %s\n", name)
		} else if os.WriteFile(name, []byte{}, 0644) != nil {
			fmt.Printf("couldn't create folder: %s\n", name)
		}
	}
}
