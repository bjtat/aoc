//go:build day01

package main

import (
	"aoc/go/utils"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func GetFile() []string {
	inputPath, err := utils.GetInputPath(2025, 1, false)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error getting input path: %v\n", err)
		os.Exit(1)
	}

	content, err := os.ReadFile(inputPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading file: %v\n", err)
		os.Exit(1)
	}

	file := string(content)
	lines := strings.Split(file, "\n")

	var result []string
	for _, line := range lines {
		trimmed := strings.TrimSpace(line)
		if trimmed != "" {
			result = append(result, trimmed)
		}
	}
	return result
}

func Convert(lines []string) []int {
	var results []int
	for _, line := range lines {
		var result int
		result, err := strconv.Atoi(line[1:])
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error converting to int: %v\n", err)
			continue
		}
		result = result % 100
		if line[0] == 'L' {
			result = -result
		}
		results = append(results, result)
	}
	return results
}

func GetAllExtraTurns(lines []string) int {
	var turns int
	for _, line := range lines {
		var result int
		result, err := strconv.Atoi(line[1:])
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error converting to int: %v\n", err)
			continue
		}
		result = result / 100
		turns += result
	}
	return turns
}

func ComputeResult(lines []string) (int, int) {
	dialPosition := 50
	counter := 0
	dialStartedAtZero := false

	turns := Convert(lines)
	for _, turn := range turns {
		// Apply the turn
		dialPosition += turn

		// Count full rotations (above or below zero)
		counter += dialPosition / 100
		if dialPosition < 0 {
			counter += (-dialPosition) / 100
		}

		// When we cross over to the negative we need to add one as the above floor will not count that
		if dialPosition < 0 && !dialStartedAtZero {
			counter++
		}

		// When we land on exactly 0, we need to increment the counter
		if dialPosition == 0 {
			counter++
		}

		// Reset dial position within bounds (0-99)
		dialPosition = dialPosition % 100
		if dialPosition == 0 {
			dialStartedAtZero = true
		} else {
			dialStartedAtZero = false
		}
		if dialPosition < 0 {
			dialPosition += 100
		}

		fmt.Printf("The dial is now at position %d (counter: %d)\n", dialPosition, counter)
	}
	return counter, 0
}

func main() {
	file := GetFile()
	password, extraTurns := ComputeResult(file)
	turns := GetAllExtraTurns(file)
	totalTurns := password + extraTurns + turns

	fmt.Printf("Original password: %d\n", password)
	fmt.Printf("Extra turns from passing 0: %d\n", extraTurns)
	fmt.Printf("Extra turns from parsing movements: %d\n", turns)
	fmt.Printf("Total turns: %d\n", totalTurns)
}
