package main

import (
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/devkcud/mk/internal/help"
	"github.com/devkcud/mk/internal/utils"
)

var Version string

func main() {
	fmt.Printf("mk %s\n", Version)

	var dirs []string

	if len(strings.Join(os.Args[1:], "")) == 0 {
		help.ShowHelp(nil)
		return
	}

	for _, name := range os.Args[1:] {
		switch []rune(name)[0] {
		case '+':
			dirs = append(dirs, name[1:])
			utils.Mkdir(path.Join(dirs...))
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
			utils.ExecCommand(path.Join(dirs...), name[1:])
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

	fmt.Println("Done")
}