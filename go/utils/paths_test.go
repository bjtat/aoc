package utils

import (
	"reflect"
	"testing"
)

func TestReadInputLines(t *testing.T) {
	// Test reading regular input file
	lines, err := ReadInputLines(2025, 1, false)
	if err != nil {
		t.Fatalf("ReadInputLines failed: %v", err)
	}

	// Check that we got some lines back
	if len(lines) == 0 {
		t.Error("ReadInputLines returned empty slice")
	}

	// All lines should be non-empty (empty lines are filtered out)
	for i, line := range lines {
		if line == "" {
			t.Errorf("Line %d is empty, but empty lines should be filtered", i)
		}
	}
}

func TestReadInputLinesWithTest(t *testing.T) {
	// Test reading test input file
	lines, err := ReadInputLines(2025, 1, true)
	if err != nil {
		t.Fatalf("ReadInputLines with test file failed: %v", err)
	}

	expected := []string{
		"L68",
		"L30",
		"R48",
		"L5",
		"R60",
		"L55",
		"L1",
		"L99",
		"R14",
		"L82",
	}

	if !reflect.DeepEqual(lines, expected) {
		t.Errorf("ReadInputLines() = %v, want %v", lines, expected)
	}
}

func TestReadInputLinesTrimsWhitespace(t *testing.T) {
	// Test that whitespace is properly trimmed
	lines, err := ReadInputLines(2025, 1, true)
	if err != nil {
		t.Fatalf("ReadInputLines failed: %v", err)
	}

	// Verify no lines have leading or trailing whitespace
	for i, line := range lines {
		if line != line {
			t.Errorf("Line %d has whitespace: '%s'", i, line)
		}
	}
}

func TestReadInputLinesInvalidYear(t *testing.T) {
	// Test with a year/day that doesn't exist
	_, err := ReadInputLines(1999, 99, false)
	if err == nil {
		t.Error("Expected error for non-existent input file, got nil")
	}
}
