package utils

import (
	"fmt"
	"os"
	"path/filepath"
)

func BuildTree(root string, indent string) {
	base := filepath.Base(root)
	fmt.Printf("%s%s/\n", indent, base)

	fileInfos, err := os.ReadDir(root)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, fileInfo := range fileInfos {
		if fileInfo.IsDir() {
			BuildTree(filepath.Join(root, fileInfo.Name()), indent+"  ")
		} else {
			fmt.Printf("%s  %s\n", indent, fileInfo.Name())
		}
	}
}
