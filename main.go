package main

import (
	"fmt"
	"os"
	"path"
	"strings"
)

func main() {
	workingDir := ""

	for _, name := range os.Args[1:] {
		if strings.HasPrefix(name, ":") {
			workingDir = name[1:]

			if os.MkdirAll(workingDir, 0755) != nil {
				fmt.Printf("couldn't create folder: %s\n", workingDir)
			}
			continue
		}

		if strings.HasPrefix(name, "-") {
			if os.MkdirAll(path.Join(workingDir, name[1:]), 0755) != nil {
				fmt.Printf("couldn't create folder: %s\n", path.Join(workingDir, name[1:]))
			}
			continue
		}

		name = path.Join(workingDir, name)

		if _, err := os.Stat(name); err == nil {
			fmt.Printf("file already exists: %s\n", path.Join(workingDir, name[1:]))
		} else if os.WriteFile(name, []byte{}, 0644) != nil {
			fmt.Printf("couldn't create folder: %s\n", path.Join(workingDir, name[1:]))
		}
	}
}
