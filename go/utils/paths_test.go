package utils

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestGetInputPath(t *testing.T) {
	// Test getting regular input path
	path, err := GetInputPath(2025, 1, false)
	if err != nil {
		t.Fatalf("GetInputPath failed: %v", err)
	}

	// Check that path ends with the correct filename
	if !strings.HasSuffix(path, "inputs/2025/day01.txt") {
		t.Errorf("Expected path to end with 'inputs/2025/day01.txt', got: %s", path)
	}

	// Check that the file exists
	if _, err := os.Stat(path); os.IsNotExist(err) {
		t.Errorf("Input file does not exist: %s", path)
	}
}

func TestGetInputPathWithTest(t *testing.T) {
	// Test getting test input path
	path, err := GetInputPath(2025, 1, true)
	if err != nil {
		t.Fatalf("GetInputPath failed: %v", err)
	}

	// Check that path ends with the correct test filename
	if !strings.HasSuffix(path, "inputs/2025/day01_test.txt") {
		t.Errorf("Expected path to end with 'inputs/2025/day01_test.txt', got: %s", path)
	}

	// Check that the test file exists
	if _, err := os.Stat(path); os.IsNotExist(err) {
		t.Errorf("Test input file does not exist: %s", path)
	}

	// Read the test file and verify it has content
	content, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("Failed to read test file: %v", err)
	}

	if len(content) == 0 {
		t.Error("Test file is empty")
	}

	// Verify it contains expected test data
	contentStr := string(content)
	if !strings.Contains(contentStr, "L68") {
		t.Error("Test file doesn't contain expected test data")
	}
}

func TestGetInputPathFormat(t *testing.T) {
	tests := []struct {
		year    int
		day     int
		useTest bool
		want    string
	}{
		{2025, 1, false, "inputs/2025/day01.txt"},
		{2025, 1, true, "inputs/2025/day01_test.txt"},
		{2024, 15, false, "inputs/2024/day15.txt"},
		{2024, 15, true, "inputs/2024/day15_test.txt"},
	}

	for _, tt := range tests {
		path, err := GetInputPath(tt.year, tt.day, tt.useTest)
		if err != nil {
			t.Fatalf("GetInputPath(%d, %d, %v) failed: %v", tt.year, tt.day, tt.useTest, err)
		}

		// Extract just the relative path portion for comparison
		parts := strings.Split(path, string(filepath.Separator))
		var relativePath string
		for i, part := range parts {
			if part == "inputs" && i+2 < len(parts) {
				relativePath = filepath.Join(parts[i:]...)
				break
			}
		}

		expectedPath := filepath.FromSlash(tt.want)
		if relativePath != expectedPath {
			t.Errorf("GetInputPath(%d, %d, %v) = %s, want path ending with %s",
				tt.year, tt.day, tt.useTest, relativePath, expectedPath)
		}
	}
}
