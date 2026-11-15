package main

import (
	"fmt"
	"os"
)

func main() {

	fmt.Println("================================")
	fmt.Println("       GO FILE ORGANIZER")
	fmt.Println("================================")

	// Check command-line arguments.
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run . <directory-path>")
		return
	}

	sourceFolder := os.Args[1]

	fmt.Println("Source folder:", sourceFolder)
	fmt.Println("Starting organization...")

	// Execute organizer.
	summary, err := organizeDirectory(sourceFolder)

	if err != nil {
		fmt.Println("Application error:", err)
		return
	}

	// Display individual errors.
	for _, message := range summary.Errors {
		fmt.Println("WARNING:", message)
	}

	// Display final report.
	fmt.Println("\n================================")
	fmt.Println("       ORGANIZATION REPORT")
	fmt.Println("================================")

	fmt.Println("Organized:", summary.Organized)
	fmt.Println("Failed:", summary.Failed)
	fmt.Println("Skipped:", summary.Skipped)

	fmt.Println("================================")
	fmt.Println("       PROCESS COMPLETED")
	fmt.Println("================================")
}
