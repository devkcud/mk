package main

import (
	"flag"
	"os"
	"path/filepath"
	"strings"

	"github.com/devkcud/mk/internal/help"
	"github.com/devkcud/mk/internal/utils"
)

var Version string

func main() {
	flag.Usage = func() { help.ShowHelp(Version) }

	showHelp := flag.Bool("help", false, "Show help menu")
	quiet := flag.Bool("quiet", false, "Disable output")

	flag.Parse()

	if *showHelp || flag.NArg() == 0 {
		flag.Usage()

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

		dir, file := filepath.Split(name)

		if dir != "" && file != "" {
			temp := append(dirstack, dir)

			utils.CreateDir(filepath.Join(temp...))
			utils.CreateFile(filepath.Join(append(temp, file)...))

			continue
		}

		if dir != "" {
			separated_dirs := strings.Split(dir, "/")
			dirstack = append(dirstack, separated_dirs[:len(separated_dirs)-1]...)
			utils.CreateDir(filepath.Join(dirstack...))
		}

		if file != "" {
			utils.CreateFile(filepath.Join(append(dirstack, file)...))
		}
	}
}
