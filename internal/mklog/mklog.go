package mklog

import (
	"fmt"
	"os"

	"github.com/fatih/color"
	"golang.org/x/exp/slices"
)

var LogLevel []string = []string{"error"}

func Log(o ...any) {
	if !slices.Contains(LogLevel, "info") && !slices.Contains(LogLevel, "all") {
		return
	}

	custom := color.New(color.FgGreen, color.Bold)
	fmt.Printf("%s     %s ", date(), custom.Sprint("LOG"))
	Print(o...)
}

func Warn(o ...any) {
	if !slices.Contains(LogLevel, "warn") && !slices.Contains(LogLevel, "all") {
		return
	}

	custom := color.New(color.FgYellow, color.Bold)
	fmt.Printf("%s %s ", date(), custom.Sprint("WARNING"))
	Print(o...)
}

func Error(o ...any) {
	if !slices.Contains(LogLevel, "error") && !slices.Contains(LogLevel, "all") {
		return
	}

	custom := color.New(color.FgRed, color.Bold)
	fmt.Printf("%s   %s ", date(), custom.Sprint("ERROR"))
	Print(o...)
}

func Fatal(o ...any) {
	if !slices.Contains(LogLevel, "error") && !slices.Contains(LogLevel, "all") {
		return
	}

	custom := color.New(color.FgWhite, color.Bold)
	fmt.Printf("%s   %s ", date(), custom.Sprint("FATAL"))
	Print(o...)
	os.Exit(1)
}

func Print(o ...any) {
	fmt.Println(o...)
}
