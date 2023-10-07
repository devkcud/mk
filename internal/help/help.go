package help

import (
	"flag"
	"fmt"
)

var examples []string = []string{"src/main.js", "tests/", "lib/ myclass.go otherclass.go andanotherclass.go", "src/main.cpp Makefile"}

func ShowHelp(version string) {
	fmt.Printf("mk (%s) - by: devkcud\n", version)

	fmt.Printf("Usage: mk [-flags] [..] [folder/] [files...]\n\nFlags:\n")

	flag.PrintDefaults()

	fmt.Println("\nExamples:")

	for _, example := range examples {
		fmt.Printf("    mk %s\n", example)
	}

	fmt.Println("\nIssues/PRs/Help: https://github.com/devkcud/mk")
}
