package fileops

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// RenameFile renames a file if it contains oldStr.
func RenameFile(directory, filename, oldStr, newStr string) (string, error) {
	if strings.Contains(filename, oldStr) {
		newFilename := strings.Replace(filename, oldStr, newStr, -1)
		oldPath := filepath.Join(directory, filename)
		newPath := filepath.Join(directory, newFilename)

		err := os.Rename(oldPath, newPath)
		if err != nil {
			return filename, fmt.Errorf("error renaming file %s: %v", filename, err)
		}

		fmt.Printf("Renamed file: %s -> %s\n", filename, newFilename)
		return newFilename, nil
	}
	return filename, nil
}
