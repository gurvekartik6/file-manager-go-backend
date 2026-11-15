package main

import (
	"os"
	"path/filepath"
	"testing"
)

// Test file category classification.
func TestGetCategory(t *testing.T) {

	tests := []struct {
		extension string
		expected  string
	}{
		{".jpg", "Images"},
		{".png", "Images"},
		{".pdf", "Documents"},
		{".mp4", "Videos"},
		{".mp3", "Audio"},
		{".csv", "Others"},
		{".JPG", "Images"},
	}

	for _, test := range tests {

		result := getCategory(test.extension)

		if result != test.expected {
			t.Errorf(
				"getCategory(%q) = %q; expected %q",
				test.extension,
				result,
				test.expected,
			)
		}
	}
}

// Test directory validation.
func TestValidateDirectory(t *testing.T) {

	// Create a temporary directory.
	tempDir := t.TempDir()

	err := validateDirectory(tempDir)

	if err != nil {
		t.Errorf("expected valid directory, got: %v", err)
	}

	// Test a nonexistent directory.
	err = validateDirectory(filepath.Join(tempDir, "missing"))

	if err == nil {
		t.Error("expected an error for nonexistent directory")
	}

	// Test a regular file instead of a directory.
	filePath := filepath.Join(tempDir, "test.txt")

	err = os.WriteFile(filePath, []byte("hello"), 0644)

	if err != nil {
		t.Fatal(err)
	}

	err = validateDirectory(filePath)

	if err == nil {
		t.Error("expected an error when path is a file")
	}
}

// Test actual file organization.
func TestOrganizeDirectory(t *testing.T) {

	// Create temporary source folder.
	source := t.TempDir()

	// Create a sample file.
	filePath := filepath.Join(source, "photo.jpg")

	err := os.WriteFile(filePath, []byte("sample"), 0644)

	if err != nil {
		t.Fatal(err)
	}

	// Run organizer.
	summary, err := organizeDirectory(source)

	if err != nil {
		t.Fatal(err)
	}

	if summary.Organized != 1 {
		t.Errorf(
			"expected 1 organized file, got %d",
			summary.Organized,
		)
	}

	// Verify destination file exists.
	destination := filepath.Join(
		source,
		"organized",
		"Images",
		"photo.jpg",
	)

	_, err = os.Stat(destination)

	if err != nil {
		t.Errorf("expected destination file to exist: %v", err)
	}

	// Verify original file no longer exists.
	_, err = os.Stat(filePath)

	if !os.IsNotExist(err) {
		t.Error("expected original file to be moved")
	}
}
