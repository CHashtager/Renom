package main

import (
	"os"
	"path/filepath"
	"testing"
)

func TestRenameDirectories(t *testing.T) {
	// Create a temporary test directory
	tempDir := t.TempDir()
	oldDir := filepath.Join(tempDir, "old_dir")
	newDir := filepath.Join(tempDir, "new_dir")

	// Create test directory
	err := os.Mkdir(oldDir, 0755)
	if err != nil {
		t.Fatalf("Failed to create test directory: %v", err)
	}

	// Run rename function
	err = RenameDirectories(tempDir, "old", "new")
	if err != nil {
		t.Fatalf("RenameDirectories failed: %v", err)
	}

	// Check if the directory was renamed
	if _, err := os.Stat(newDir); os.IsNotExist(err) {
		t.Fatalf("Expected directory %s, but it does not exist", newDir)
	}
}
