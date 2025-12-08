package utils

import (
	"fmt"
	"os/exec"
	"path/filepath"
	"strings"
)

// GetRepoRoot returns the git repository root directory
func GetRepoRoot() (string, error) {
	cmd := exec.Command("git", "rev-parse", "--show-toplevel")
	output, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf("failed to get git root: %w", err)
	}
	return strings.TrimSpace(string(output)), nil
}

// GetInputPath returns the full path to an input file for a given year and day
// If useTest is true, returns the path to the test input file (day##_test.txt)
func GetInputPath(year int, day int, useTest bool) (string, error) {
	root, err := GetRepoRoot()
	if err != nil {
		return "", err
	}

	var filename string
	if useTest {
		filename = fmt.Sprintf("inputs/%d/day%02d_test.txt", year, day)
	} else {
		filename = fmt.Sprintf("inputs/%d/day%02d.txt", year, day)
	}

	return filepath.Join(root, filename), nil
}
