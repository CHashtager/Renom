package dirops

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// RenameDirectories renames all directories containing oldStr with newStr, processing from deepest to highest level.
func RenameDirectories(directory, oldStr, newStr string) error {
	// Collect directories in bottom-up order
	var dirs []string

	err := filepath.Walk(directory, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			dirs = append([]string{path}, dirs...) // Prepend to process deepest directories first
		}
		return nil
	})

	if err != nil {
		return fmt.Errorf("error walking directory: %v", err)
	}

	// Rename directories
	for _, dir := range dirs {
		dirName := filepath.Base(dir)
		if strings.Contains(dirName, oldStr) {
			newDirName := strings.Replace(dirName, oldStr, newStr, -1)
			newPath := filepath.Join(filepath.Dir(dir), newDirName)

			err := os.Rename(dir, newPath)
			if err != nil {
				return fmt.Errorf("error renaming directory %s: %v", dir, err)
			}
			fmt.Printf("Renamed directory: %s -> %s\n", dir, newPath)
		}
	}
	return nil
}
