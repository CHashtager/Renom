package tests

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/Chashtager/Renom/internal/fileops"
)

func TestRenameFile(t *testing.T) {
	// Create a temporary test directory
	tempDir := t.TempDir()
	oldFile := filepath.Join(tempDir, "test_old.txt")
	newFile := filepath.Join(tempDir, "test_new.txt")

	// Create a test file
	err := os.WriteFile(oldFile, []byte("Sample content"), 0644)
	if err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}

	// Run the rename function
	_, err = fileops.RenameFile(tempDir, "test_old.txt", "old", "new")
	if err != nil {
		t.Fatalf("RenameFile failed: %v", err)
	}

	// Check if the file was renamed
	if _, err := os.Stat(newFile); os.IsNotExist(err) {
		t.Fatalf("Expected file %s, but it does not exist", newFile)
	}
}

func TestReplaceInFile(t *testing.T) {
	// Create a temporary file
	tempDir := t.TempDir()
	testFile := filepath.Join(tempDir, "test.txt")

	// Write initial content
	content := "Hello, old world!"
	err := os.WriteFile(testFile, []byte(content), 0644)
	if err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}

	// Replace "old" with "new"
	err = fileops.ReplaceInFile(testFile, "old", "new")
	if err != nil {
		t.Fatalf("ReplaceInFile failed: %v", err)
	}

	// Read back content
	data, err := os.ReadFile(testFile)
	if err != nil {
		t.Fatalf("Failed to read test file: %v", err)
	}

	// Check if replacement was successful
	expected := "Hello, new world!"
	if string(data) != expected {
		t.Errorf("Expected content %q, but got %q", expected, string(data))
	}
}
