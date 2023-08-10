package main

import (
	"fmt"
	"os"
	"path"
)

func createDir(path string) {
	if os.MkdirAll(path, 0755) != nil {
		fmt.Printf("couldn't create folder: %s\n", path)
	}
}

func main() {
	var dirs []string

	if len(os.Args[1:]) == 0 {
		// TODO: Create a help menu
		help := "Usage: [+<folder>] [-] [filename]"

		fmt.Println(help)
		return
	}

	for _, name := range os.Args[1:] {
		switch []rune(name)[0] {
		case '+':
			dirs = append(dirs, name[1:])
			createDir(path.Join(dirs...))

			break
		case '-':
			if len(dirs) > 0 {
				dirs = dirs[:len(dirs)-1]
			}
			break
		default:
			curPath := path.Join(append(dirs, name)...)

			if _, err := os.Stat(curPath); err == nil {
				fmt.Printf("path already exists: %s\n", curPath)
			} else if os.WriteFile(curPath, []byte{}, 0644) != nil {
				fmt.Printf("couldn't create file: %s\n", curPath)
			}
			break
		}
	}
}
