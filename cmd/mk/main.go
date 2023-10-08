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
	flagPrompt := flag.Bool("prompt", false, "Prompt every time a folder/file is about to be created")

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

	if *flagHelp || flag.NArg() == 0 {
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

			if *flagPrompt && !utils.YesNoPrompt("Create path "+filepath.Join(append(temp, file)...)+"?", true) {
				continue
			}

			utils.CreateDir(filepath.Join(temp...))
			utils.CreateFile(filepath.Join(append(temp, file)...))

			continue
		}

		if dir != "" {
			separated_dirs := strings.Split(dir, "/")

			if *flagPrompt && !utils.YesNoPrompt("Create folder "+filepath.Join(append(dirstack, separated_dirs[:len(separated_dirs)-1]...)...)+"?", true) {
				continue
			}

			dirstack = append(dirstack, separated_dirs[:len(separated_dirs)-1]...)
			utils.CreateDir(filepath.Join(dirstack...))
		}

		if file != "" {
			if *flagPrompt && !utils.YesNoPrompt("Create file "+filepath.Join(append(dirstack, file)...)+"?", true) {
				continue
			}

			utils.CreateFile(filepath.Join(append(dirstack, file)...))
		}
	}
}
