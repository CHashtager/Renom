package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Chashtager/Renom/internal/dirops"
	"github.com/Chashtager/Renom/internal/fileops"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: renom <old_string> <new_string>")
		os.Exit(1)
	}

	oldStr, newStr := os.Args[1], os.Args[2]
	currentDir, err := os.Getwd()
	if err != nil {
		log.Fatalf("Error getting current directory: %v", err)
	}

	// Rename directories first
	err = dirops.RenameDirectories(currentDir, oldStr, newStr)
	if err != nil {
		log.Fatalf("Error renaming directories: %v", err)
	}

	// Rename files and replace content
	err = fileops.ProcessFiles(currentDir, oldStr, newStr)
	if err != nil {
		log.Fatalf("Error processing files: %v", err)
	}

	fmt.Println("Processing complete.")
}
