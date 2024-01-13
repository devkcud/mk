package mklog

import (
	"fmt"
	"os"

	"github.com/fatih/color"
)

var LogLevel int = 3

func Log(o ...any) {
	if LogLevel >= 1 {
		custom := color.New(color.FgGreen, color.Bold)
		fmt.Printf("%s     %s ", date(), custom.Sprint("LOG"))
		Print(o...)
	}
}

func Warn(o ...any) {
	if LogLevel >= 2 {
		custom := color.New(color.FgYellow, color.Bold)
		fmt.Printf("%s %s ", date(), custom.Sprint("WARNING"))
		Print(o...)
	}
}

func Error(o ...any) {
	if LogLevel >= 3 {
		custom := color.New(color.FgRed, color.Bold)
		fmt.Printf("%s   %s ", date(), custom.Sprint("ERROR"))
		Print(o...)
	}
}

func Fatal(o ...any) {
	custom := color.New(color.FgWhite, color.Bold)
	fmt.Printf("%s   %s ", date(), custom.Sprint("FATAL"))
	Print(o...)
	os.Exit(1)
}

func Print(o ...any) {
	fmt.Println(o...)
}
