package mklog

import (
	"fmt"
	"os"
	"strings"

	"github.com/fatih/color"
)

func Log(o ...string) {
	output := date()

	custom := color.New(color.FgGreen, color.Bold)
	output += custom.Sprint("    INFO ") + strings.Join(o, " ")

	fmt.Println(output)
}

func Warn(o ...string) {
	output := date()

	custom := color.New(color.FgYellow, color.Bold)
	output += custom.Sprint(" WARNING ") + strings.Join(o, " ")

	fmt.Println(output)
}

func Error(o ...string) {
	output := date()

	custom := color.New(color.FgRed, color.Bold)
	output += custom.Sprint("   ERROR ") + strings.Join(o, " ")

	fmt.Println(output)
}

func Fatal(o ...string) {
	output := date()

	custom := color.New(color.FgWhite, color.BgRed, color.Bold)
	output += "   " + custom.Sprint("FATAL") + " " + strings.Join(o, " ")

	fmt.Println(output)
	os.Exit(1)
}

func Print(o ...string) {
	fmt.Println(strings.Join(o, " "))
}
