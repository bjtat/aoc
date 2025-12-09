package utils

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

// getRepoRoot returns the git repository root directory
func getRepoRoot() (string, error) {
	cmd := exec.Command("git", "rev-parse", "--show-toplevel")
	output, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf("failed to get git root: %w", err)
	}
	return strings.TrimSpace(string(output)), nil
}

// getInputPath returns the full path to an input file for a given year and day
// If useTest is true, returns the path to the test input file (day##_test.txt)
func getInputPath(year int, day int, useTest bool) (string, error) {
	root, err := getRepoRoot()
	if err != nil {
		return "", err
	}

	var builder strings.Builder
	builder.WriteString(fmt.Sprintf("inputs/%d/day%02d", year, day))
	if useTest {
		builder.WriteString("_test")
	}
	builder.WriteString(".txt")

	return filepath.Join(root, builder.String()), nil
}

// ReadInputLines reads an input file and returns a slice of trimmed, non-empty lines
func ReadInputLines(year int, day int, useTest bool) ([]string, error) {
	inputPath, err := getInputPath(year, day, useTest)
	if err != nil {
		return nil, fmt.Errorf("failed to get input path: %w", err)
	}

	content, err := os.ReadFile(inputPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %w", err)
	}

	lines := strings.Split(string(content), "\n")

	result := make([]string, 0, len(lines))
	for _, line := range lines {
		trimmed := strings.TrimSpace(line)
		if trimmed != "" {
			result = append(result, trimmed)
		}
	}
	return result, nil
}
