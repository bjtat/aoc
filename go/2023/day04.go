//go:build day04

package main

import (
	"fmt"
	"math"
	"os"
	"strings"

	"aoc/utils"
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

func addCards(cards *[250]int, set map[string]bool, actualNumbers []string, cardNum int) {
	points := 0
	for _, num := range actualNumbers {
		_, seen := set[num]
		if seen {
			points += 1
		}
	}

	numCopies := cards[cardNum]
	for i := 1; i <= points; i++ {
		fmt.Printf("appending %d copies to %d\n", numCopies, cardNum+i)
		cards[cardNum+i] += numCopies
	}
}

func main() {
	cards := [250]int{}
	fmt.Println(cards)

	inputPath, _ := utils.GetInputPath(2023, 4)
	fileBytes, _ := os.ReadFile(inputPath)
	trimmedFile := strings.TrimSpace(string(fileBytes))
	linesArray := strings.Split(trimmedFile, "\n")

	for i, line := range linesArray {
		cards[i+1] += 1
		winning, actual := parseNumbers(line)
		fmt.Println(winning, actual)
		seenNumberSet := generateSet(winning)
		fmt.Println(seenNumberSet)

		addCards(&cards, seenNumberSet, actual, i+1)
	}

	fmt.Println(cards)

	answer := 0
	for _, copies := range cards {
		answer += copies
	}
	fmt.Println(answer)
}
