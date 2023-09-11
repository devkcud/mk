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
		switch []rune(name)[0] {
		case ':':
			utils.CreateDir(path.Join(append(dirstack, name[1:])...))
		case '.':
			dots := strings.Count(name, ".")

			if strings.Replace(name, ".", "", -1) != "" {
				dirstack = append(dirstack, name[1:])
				utils.CreateDir(path.Join(dirstack...))
				break
			}

			if dots > len(dirstack) {
				dirstack = nil
			} else {
				dirstack = dirstack[:len(dirstack)-(dots)]
			}
		// case '%':
		// TODO: Redo commands again
		// utils.ExecCommand(path.Join(dirstack...), name[1:])
		default:
			utils.CreateFile(path.Join(append(dirstack, strings.Replace(name, "#", "", 1))...))
		}
	}
}
