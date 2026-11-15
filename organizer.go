package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// Summary stores the organization results.
type Summary struct {
	Organized int
	Failed    int
	Skipped   int
	Errors    []string
}

// Identify category using extension.
func getCategory(extension string) string {

	categories := map[string]string{
		".jpg":  "Images",
		".jpeg": "Images",
		".png":  "Images",
		".gif":  "Images",
		".pdf":  "Documents",
		".docx": "Documents",
		".txt":  "Documents",
		".mp4":  "Videos",
		".mkv":  "Videos",
		".mp3":  "Audio",
		".wav":  "Audio",
	}

	category, exists := categories[strings.ToLower(extension)]

	if !exists {
		return "Others"
	}

	return category
}

// Validate the source directory.
func validateDirectory(path string) error {

	info, err := os.Stat(path)

	if err != nil {
		return fmt.Errorf("invalid source path: %w", err)
	}

	if !info.IsDir() {
		return fmt.Errorf("source path is not a directory")
	}

	return nil
}

// Organize files inside a directory.
func organizeDirectory(sourceFolder string) (Summary, error) {

	summary := Summary{}

	// Validate input.
	err := validateDirectory(sourceFolder)

	if err != nil {
		return summary, err
	}

	// Read source directory.
	entries, err := os.ReadDir(sourceFolder)

	if err != nil {
		return summary, fmt.Errorf("reading directory: %w", err)
	}

	// Destination is inside the source directory.
	destinationRoot := filepath.Join(sourceFolder, "organized")

	for _, entry := range entries {

		// Skip directories.
		if entry.IsDir() {
			summary.Skipped++
			continue
		}

		fileName := entry.Name()

		sourcePath := filepath.Join(sourceFolder, fileName)

		// Get extension and category.
		extension := filepath.Ext(fileName)
		category := getCategory(extension)

		// Get destination directory.
		categoryFolder := filepath.Join(
			destinationRoot,
			category,
		)

		// Create category directory.
		err := os.MkdirAll(categoryFolder, 0755)

		if err != nil {
			summary.Failed++
			summary.Errors = append(
				summary.Errors,
				fmt.Sprintf("%s: %v", fileName, err),
			)
			continue
		}

		// Build destination path.
		destinationPath := filepath.Join(
			categoryFolder,
			fileName,
		)

		// Prevent overwriting existing files.
		_, err = os.Lstat(destinationPath)

		if err == nil {
			summary.Failed++
			summary.Errors = append(
				summary.Errors,
				fmt.Sprintf("%s: destination already exists", fileName),
			)
			continue
		}

		if !os.IsNotExist(err) {
			summary.Failed++
			summary.Errors = append(
				summary.Errors,
				fmt.Sprintf("%s: checking destination: %v", fileName, err),
			)
			continue
		}

		// Move the file.
		err = os.Rename(sourcePath, destinationPath)

		if err != nil {
			summary.Failed++
			summary.Errors = append(
				summary.Errors,
				fmt.Sprintf("%s: moving file: %v", fileName, err),
			)
			continue
		}

		summary.Organized++
	}

	return summary, nil
}
