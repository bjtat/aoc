package main

import (
	"aoc/go/utils"
	"fmt"
	"os"
	"strconv"
)

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
	endOnZero := 0
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
			endOnZero++
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
	}
	return counter, endOnZero
}

func main() {
	file, err := utils.ReadInputLines(2025, 1, true)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading 2025 Day 1 file: %v\n", err)
		os.Exit(1)
	}

	password, endOnZero := ComputeResult(file)
	extraTurns := GetAllExtraTurns(file)
	totalTurns := password + endOnZero + extraTurns

	fmt.Printf("Original password: %d\n", endOnZero)
	fmt.Printf("Total turns: %d\n", totalTurns)
}
