package help

import "fmt"

func ShowHelp() {
	fmt.Println(`Usage: mk [+<folder>] [-] [%'<command>'] [filename]

+<folder>    : Create a new directory and add it to the directory stack.
-            : Remove the last added directory from the directory stack.
%'<command>' : Runs a command in the last folder in the directory stack.
filename     : Create an empty file with the specified name.

Issues/PRs/Help: https://github.com/devkcud/mk`)
}
