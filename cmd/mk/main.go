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
	if os.Getenv("MK_QUIET") != "" {
		os.Stdout = nil
		os.Stderr = nil
	}

	var dirstack []string

	if len(strings.Join(os.Args[1:], "")) == 0 {
		fmt.Printf("mk (%s) - by: devkcud\n", Version)

		help.ShowHelp()
		return
	}

	for _, name := range os.Args[1:] {
		if name[0] == '.' && strings.TrimLeft(name, ".") == "" {
			dots := strings.Count(name, ".")

			if dots > len(dirstack) {
				dirstack = nil
			} else {
				dirstack = dirstack[:len(dirstack)-dots]
			}

			continue
		}

		dir, file := path.Split(name)

		if dir != "" {
			dirstack = append(dirstack, dir)
			utils.CreateDir(path.Join(dirstack...))
		}

		if file != "" {
			utils.CreateFile(path.Join(append(dirstack, file)...))
		}
	}
}
