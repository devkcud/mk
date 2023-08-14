package utils

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func Mkdir(path string) {
	if os.MkdirAll(path, 0755) != nil {
		fmt.Printf("couldn't create folder: %s\n", path)
	}
}

func ExecCommand(path string, command string) {
	cmd := exec.Command("sh", "-c", command)
	cmd.Dir = path

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if cmd.Run() != nil {
		fmt.Printf("couldn't not run: %s\n", strings.Join(cmd.Args, " "))
	}
}
