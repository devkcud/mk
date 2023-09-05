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

	var dirs []string

	if len(strings.Join(os.Args[1:], "")) == 0 {
		fmt.Printf("mk (%s) - by: devkcud\n", Version)

		help.ShowHelp()
		return
	}

	for _, name := range os.Args[1:] {
		switch []rune(name)[0] {
		case ':':
			utils.CreateDir(path.Join(append(dirs, name[1:])...))
		case '.':
			dots := strings.Count(name, ".")

			if strings.Replace(name, ".", "", -1) != "" {
				dirs = append(dirs, name[1:])
				utils.CreateDir(path.Join(dirs...))
				break
			}

			if dots > len(dirs) {
				dirs = nil
			} else {
				dirs = dirs[:len(dirs)-(dots)]
			}
        // TODO: Redo commands again
		// case '%':
		// 	utils.ExecCommand(path.Join(dirs...), name[1:])
		default:
			utils.CreateFile(path.Join(append(dirs, strings.Replace(name, "#", "", 1))...))
		}
	}
}
