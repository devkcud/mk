package help

import "fmt"

func ShowHelp() {
    fmt.Printf("mk (%s) - by: devkcud\n", Version)

	fmt.Println(`Usage: mk [.[folder]] [:<folder>] [[#]filename]

! OUTDATED ! Please refer to the source code until the help is updated

Issues/PRs/Help: https://github.com/devkcud/mk`)
}
