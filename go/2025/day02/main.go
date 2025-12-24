package main

import (
	"aoc/utils"
	"fmt"
	"os"
)

func main() {
	lines, err := utils.ReadInputLines(2025, 2, false)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading input: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Starting day 2")
	fmt.Printf("Read %d lines\n", len(lines))
}
