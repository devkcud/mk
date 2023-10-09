package utils

import (
	"os"

	"github.com/devkcud/mk/internal/mklog"
	"github.com/fatih/color"
)

func CreateDir(path string) {
	if os.MkdirAll(path, 0755) != nil {
		mklog.Error("Couldn't create directory:", color.MagentaString(path))
	}

	mklog.Log("Created directory:", color.MagentaString(path))
}

func CreateFile(path string) {
	if _, err := os.Stat(path); err == nil {
		mklog.Error("Path already exists:", color.MagentaString(path))
	} else if os.WriteFile(path, []byte{}, 0644) != nil {
		mklog.Error("Couldn't create file:", color.MagentaString(path))
	}

	mklog.Log("Created file:", color.MagentaString(path))
}
