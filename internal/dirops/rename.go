package dirops

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// RenameDirectories renames all directories containing oldStr with newStr.
func RenameDirectories(directory, oldStr, newStr string) error {
	return filepath.Walk(directory, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() && strings.Contains(info.Name(), oldStr) {
			newDirname := strings.Replace(info.Name(), oldStr, newStr, -1)
			newPath := filepath.Join(filepath.Dir(path), newDirname)

			err := os.Rename(path, newPath)
			if err != nil {
				return fmt.Errorf("error renaming directory %s: %v", path, err)
			}

			fmt.Printf("Renamed directory: %s -> %s\n", info.Name(), newDirname)
		}
		return nil
	})
}
