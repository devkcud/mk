package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/devkcud/mk/internal/help"
	"github.com/devkcud/mk/internal/mklog"
	"github.com/devkcud/mk/internal/utils"
	"github.com/fatih/color"
	flag "github.com/spf13/pflag"
)

var Version string

func main() {
	flag.Usage = func() { help.ShowHelp(Version) }

	flagHelp := flag.BoolP("help", "h", false, "Show help menu")
	flagVersion := flag.BoolP("version", "V", false, "Show version")
	flagQuiet := flag.BoolP("quiet", "q", false, "Disable output")
	flagPrompt := flag.Bool("prompt", false, "Ask to create file/directory")
	flagNoColor := flag.Bool("nocolor", false, "Disable colors in output (logging)")
    flagLogLevel := flag.IntP("loglevel", "l", 0, "Set log level (0: none, 1: info, 2: warn, 3: error, 4: all)")

	flag.Parse()

	mklog.LogLevel = *flagLogLevel

	if *flagNoColor {
		os.Setenv("NO_COLOR", "1")
	}

	if *flagQuiet {
		if *flagPrompt {
			mklog.Warn("--prompt and --quiet are incompatible; skipping --prompt")
		}

		os.Stdin = nil
		os.Stdout = nil
		os.Stderr = nil
	}

	if *flagVersion {
		mklog.Print("mk", Version)
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

			quantSay := "directory"

			if dots > 1 {
				quantSay = "directories"
			}

			mklog.Log("Going", color.MagentaString(fmt.Sprint(dots)), quantSay, "back (dirstack)")

			continue
		}

		dir, file := filepath.Split(name)

		if dir != "" && file != "" {
			temp := append(dirstack, dir)

			if *flagPrompt && !utils.YesNoPrompt("Create path "+filepath.Join(append(temp, file)...)+"?", true) {
				mklog.Log("Skipping path", color.MagentaString(filepath.Join(append(temp, file)...)))
				continue
			}

			utils.CreateDir(filepath.Join(temp...))
			utils.CreateFile(filepath.Join(append(temp, file)...))

			continue
		}

		if dir != "" {
			separated_dirs := strings.Split(dir, "/")

			if *flagPrompt && !utils.YesNoPrompt("Create directory "+filepath.Join(append(dirstack, separated_dirs[:len(separated_dirs)-1]...)...)+"?", true) {
				mklog.Log("Skipping directory", color.MagentaString(filepath.Join(append(dirstack, separated_dirs[:len(separated_dirs)-1]...)...)))
				mklog.Warn(color.MagentaString(filepath.Join(append(dirstack, separated_dirs[:len(separated_dirs)-1]...)...)), "not added to the dirstack")
				continue
			}

			dirstack = append(dirstack, separated_dirs[:len(separated_dirs)-1]...)
			utils.CreateDir(filepath.Join(dirstack...))
		}

		if file != "" {
			if *flagPrompt && !utils.YesNoPrompt("Create file "+filepath.Join(append(dirstack, file)...)+"?", true) {
				mklog.Log("Skipping file", color.MagentaString(filepath.Join(append(dirstack, file)...)))
				continue
			}

			utils.CreateFile(filepath.Join(append(dirstack, file)...))
		}
	}
}
