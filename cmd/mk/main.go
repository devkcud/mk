package main

import (
	"flag"
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/devkcud/mk/internal/help"
	"github.com/devkcud/mk/internal/utils"
)

var Version string

func main() {
	showHelp := flag.Bool("help", false, "Show help menu")
	quiet := flag.Bool("quiet", false, "Disable output")
	flag.Parse()

	if *showHelp || flag.NArg() == 0 {
		fmt.Printf("mk (%s) - by: devkcud\n", Version)

		help.ShowHelp()
		return
	}

	if *quiet {
		// os.Stdin = nil
		os.Stdout = nil
		os.Stderr = nil
	}

	var dirstack []string

	for _, name := range flag.Args() {
		if strings.Trim(name, ".") == "" {
			dots := strings.Count(name, ".") - 1

			if dots > len(dirstack) {
				dirstack = nil
			} else {
				dirstack = dirstack[:len(dirstack)-dots]
			}

			continue
		}

		dir, file := path.Split(name)

		if dir != "" {
			separated_dirs := strings.Split(dir, "/")
			dirstack = append(dirstack, separated_dirs[:len(separated_dirs)-1]...)
			utils.CreateDir(path.Join(dirstack...))
		}

		if file != "" {
			utils.CreateFile(path.Join(append(dirstack, file)...))
		}
	}
}
