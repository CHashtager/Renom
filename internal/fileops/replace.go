package fileops

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// ReplaceInFile replaces all occurrences of oldStr with newStr in the given file.
func ReplaceInFile(filePath, oldStr, newStr string) error {
	content, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("error reading file %s: %v", filePath, err)
	}

	updatedContent := strings.ReplaceAll(string(content), oldStr, newStr)

	err = os.WriteFile(filePath, []byte(updatedContent), 0644)
	if err != nil {
		return fmt.Errorf("error writing to file %s: %v", filePath, err)
	}

	return nil
}

// ProcessFiles walks through all files and processes them.
func ProcessFiles(directory, oldStr, newStr string) error {
	return filepath.Walk(directory, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() {
			// Rename the file if needed
			updatedFilename, err := RenameFile(filepath.Dir(path), info.Name(), oldStr, newStr)
			if err != nil {
				return err
			}

			// Replace text inside the file
			err = ReplaceInFile(filepath.Join(filepath.Dir(path), updatedFilename), oldStr, newStr)
			if err != nil {
				return err
			}
		}
		return nil
	})
}
