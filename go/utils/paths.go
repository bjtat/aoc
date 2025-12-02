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
func GetInputPath(year int, day int) (string, error) {
	root, err := GetRepoRoot()
	if err != nil {
		return "", err
	}
	return filepath.Join(root, fmt.Sprintf("inputs/%d/day%02d.txt", year, day)), nil
}
