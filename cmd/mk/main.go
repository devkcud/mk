package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/devkcud/mk/internal/help"
	"github.com/devkcud/mk/internal/utils"
	flag "github.com/spf13/pflag"
)

var Version string

func main() {
	flag.Usage = func() { help.ShowHelp(Version) }

	flagHelp := flag.BoolP("help", "h", false, "Show help menu")
	flagVersion := flag.BoolP("version", "V", false, "Show version")
	flagQuiet := flag.BoolP("quiet", "q", false, "Disable output")

	flag.Parse()

	if *flagQuiet {
		// os.Stdin = nil
		os.Stdout = nil
		os.Stderr = nil
	}

	if *flagVersion {
		fmt.Printf("mk %s\n", Version)
		return
	}

	if *flagHelp {
		flag.Usage()
		return
	}

	if flag.NArg() == 0 {
		flag.Usage()
		return
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
