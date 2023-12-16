package main

import (
	"fmt"
	"math"
	"os"
	"strings"
)

func parseNumbers(line string) ([]string, []string) {
	splitLine := strings.Split(line, "|")
	winningNumberStr := strings.Split(splitLine[0], ":")
	winningNumberArr := strings.Split(winningNumberStr[1], " ")
	actualNumberArr := strings.Split(splitLine[1], " ")

	return winningNumberArr, actualNumberArr
}

func generateSet(winningNumbers []string) map[string]bool {
	uniqueSet := make(map[string]bool, 10)

	for _, numStr := range winningNumbers {
		_, ok := uniqueSet[numStr]
		if !ok && len(numStr) > 0 {
			uniqueSet[numStr] = true
		}
	}
	return uniqueSet
}

func calculatePoints(set map[string]bool, actualNumbers []string) int {
	points := -1
	for _, num := range actualNumbers {
		_, seen := set[num]
		if seen {
			points += 1
		}
	}

	if points > -1 {
		return int(math.Pow(2, float64(points)))
	}
	return 0
}

func main() {
	fileBytes, _ := os.ReadFile("aoc-day-4.txt")
	trimmedFile := strings.TrimSpace(string(fileBytes))
	linesArray := strings.Split(trimmedFile, "\n")

	answer := 0
	for _, line := range linesArray {
		winning, actual := parseNumbers(line)
		fmt.Println(winning, actual)
		seenNumberSet := generateSet(winning)
		fmt.Println(seenNumberSet)
		answer += calculatePoints(seenNumberSet, actual)
	}

	fmt.Println(answer)
}
