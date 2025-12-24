package main

import (
	"aoc/utils"
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
		dialPosition += turn

		// Count full rotations in the positive direction
		counter += dialPosition / 100

		// If we pass zero in the negative direction, count it
		// We check if the dial was already at zero when going negative because we don't
		// want to double count "clicking to zero" (aka from landing on the following case in the prior turn)
		if dialPosition < 0 && !dialStartedAtZero {
			counter++
		}

		// When we land on exactly 0, we need to increment the counter for clicking to it
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

		// Counting for part 1 where we count the time we land on zero
		if dialPosition == 0 {
			endOnZero++
		}

		fmt.Printf("The dial is rotated %d to point to %d\n", turn, dialPosition)
	}
	return counter, endOnZero
}

func main() {
	file, err := utils.ReadInputLines(2025, 1, false)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading 2025 Day 1 file: %v\n", err)
		os.Exit(1)
	}

	password, endOnZero := ComputeResult(file)
	extraTurns := GetAllExtraTurns(file)
	totalTurns := password + extraTurns

	fmt.Printf("Original password: %d\n", endOnZero)
	fmt.Printf("Total turns: %d\n", totalTurns)
}
