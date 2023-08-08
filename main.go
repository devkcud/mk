package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	args := os.Args[1:]

	for _, name := range args {
		if strings.HasPrefix(name, "-") {
			if err := os.MkdirAll(name[1:], 0755); err != nil {
				fmt.Printf("couldn't create folder: %s\n", name[1:])
			}
			continue
		}

		if _, err := os.ReadFile(name); err == nil {
			fmt.Printf("file already exists: %s\n", name)
			continue
		}

		if err := os.WriteFile(name, []byte{}, 0644); err != nil {
			fmt.Printf("couldn't create folder: %s\n", name)
		}
	}

}
