package main

import (
	"fmt"
	"os"
	"os/exec"
	"path"
	"strings"
)

func createDir(path string) {
	if os.MkdirAll(path, 0755) != nil {
		fmt.Printf("couldn't create folder: %s\n", path)
	}
}

func runCommand(path string, command string) {
	run := exec.Command("sh", "-c", command)

	if len(path) != 0 {
		run = exec.Command("sh", "-c", fmt.Sprintf("cd %s && %s", path, command))
	}

	if run.Run() != nil {
		fmt.Printf("couldn't not run: %s\n", strings.Join(run.Args, " "))
	}
}

func main() {
	var dirs []string

	if len(os.Args[1:]) == 0 {
		fmt.Println(`mk: Simply make files/folders

Usage: mk [+<folder>] [-] [filename]

+<folder> : Create a new directory and add it to the directory stack.
-         : Remove the last added directory from the directory stack.
filename  : Create an empty file with the specified name.

Issues/PRs/Help: https://github.com/devkcud/mk`)
		return
	}

	for _, name := range os.Args[1:] {
		switch []rune(name)[0] {
		case '+':
			dirs = append(dirs, name[1:])
			createDir(path.Join(dirs...))
			break
		case '-':
			count := strings.Count(name, "-")

			if count > len(dirs) {
				dirs = nil
				break
			}

			dirs = dirs[:len(dirs)-count]
			break
		case '%':
			curPath := path.Join(dirs...)
			runCommand(curPath, name[1:])
			break
		default:
			if strings.HasPrefix(name, "#") { // So you can use "+" or "-" in file names
				name = name[1:]
			}

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
