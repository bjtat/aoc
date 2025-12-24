package main

import (
	"aoc/utils"
	"fmt"
	"os"
	"strings"
)

type Range struct {
	Start int
	End   int
}

func Parse(line string) []Range {
	var ranges []Range

	for part := range strings.SplitSeq(line, ",") {
		numbers := strings.Split(part, "-")
		if len(numbers) == 2 {
			start := strings.TrimSpace(numbers[0])
			end := strings.TrimSpace(numbers[1])

			var rangeObj Range
			fmt.Sscanf(start, "%d", &rangeObj.Start)
			fmt.Sscanf(end, "%d", &rangeObj.End)

			ranges = append(ranges, rangeObj)
		}
	}

	return ranges
}

func main() {
	lines, err := utils.ReadInputLines(2025, 2, true)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading input: %v\n", err)
		os.Exit(1)
	}

	productRanges := Parse(lines[0])

	for _, r := range productRanges {
		fmt.Println(r)
	}

	fmt.Println("Starting day 2")
}
