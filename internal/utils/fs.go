package utils

import (
	"io"
	"os"
	"path/filepath"

	"github.com/devkcud/mk/internal/mklog"
	"github.com/fatih/color"
)

func CreateDir(path string) {
	if _, err := os.Stat(path); err == nil {
		mklog.Error("Directory already exists:", color.MagentaString(path))
		return
	}

	if os.MkdirAll(path, 0755) != nil {
		mklog.Fatal("Couldn't create directory:", color.MagentaString(path))
	}

	mklog.Log("Created directory:", color.MagentaString(path))
}

func CreateFile(path string) {
	if _, err := os.Stat(path); err == nil {
		mklog.Error("File already exists:", color.MagentaString(path))
		return
	}

	if os.WriteFile(path, []byte{}, 0644) != nil {
		mklog.Fatal("Couldn't create file:", color.MagentaString(path))
	}

	mklog.Log("Created file:", color.MagentaString(path))
}

func CopyFolder(source, destination string) error {
	return filepath.Walk(source, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		destPath := filepath.Join(destination, path[len(source):])

		if info.IsDir() {
			CreateDir(destPath)
		} else {
			sourceFile, err := os.Open(path)
			if err != nil {
				return err
			}
			defer sourceFile.Close()

			CreateFile(destPath)

			destFile, err := os.Create(destPath)
			if err != nil {
				return err
			}
			defer destFile.Close()

			_, err = io.Copy(destFile, sourceFile)
			if err != nil {
				return err
			}
		}

		return nil
	})
}
