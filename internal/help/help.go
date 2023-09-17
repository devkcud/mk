package help

import "fmt"

func ShowHelp() {
	fmt.Println(`Usage: mk [-flags] [..] [folder/] [files...]

Flags:
    -help   Show this help menu
    -quiet  Disable output

Examples:
    mk src/main.js
    mk tests/
    mk lib/ myclass.go otherclass.go andanotherclass.go
    mk src/main.cpp Makefile

Issues/PRs/Help: https://github.com/devkcud/mk`)
}
