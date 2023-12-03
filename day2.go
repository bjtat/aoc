package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func parseGames(gamesList []string, gameId, r, g, b int) bool {
	for _, game := range gamesList {
		values := strings.Split(game, ", ")
		fmt.Printf("Game %d : %#v\n", gameId, values)
		if !validateValue(values, r, g, b) {
			return false
		}
	}
	return true
}

func validateValue(values []string, r, g, b int) bool {
	for _, numberAndColor := range values {
		tuple := strings.Split(numberAndColor, " ")
		parsedValue, _ := strconv.Atoi(tuple[0])
		fmt.Printf("number: %v, color: %v\n", parsedValue, tuple[1])
		switch tuple[1] {
		case "red":
			if parsedValue > r {
				fmt.Println("FALSE GAME")
				return false
			}
		case "green":
			if parsedValue > g {
				fmt.Println("FALSE GAME")
				return false
			}
		case "blue":
			if parsedValue > b {
				fmt.Println("FALSE GAME")
				return false
			}
		default:
		}
	}
	return true
}

func parseGames2(gamesList string, gameId int) int {
	gamesList = strings.ReplaceAll(gamesList, ";", ",")
	gamesListParsed := strings.Split(gamesList, ", ")
	fmt.Printf("%#v\n", gamesListParsed)

	return validateValue2(gamesListParsed)
}

func validateValue2(values []string) int {
	maxRed := 0
	maxGreen := 0
	maxBlue := 0

	for _, numberAndColor := range values {
		tuple := strings.Split(numberAndColor, " ")
		parsedValue, _ := strconv.Atoi(tuple[0])
		// fmt.Printf("number: %v, color: %v\n", parsedValue, tuple[1])
		switch tuple[1] {
		case "red":
			if parsedValue > maxRed {
				maxRed = parsedValue
			}
		case "green":
			if parsedValue > maxGreen {
				maxGreen = parsedValue
			}
		case "blue":
			if parsedValue > maxBlue {
				maxBlue = parsedValue
			}
		default:
		}
	}

	fmt.Printf("minRed: %d, minGreen: %d, minBlue: %d -- power: %d\n\n\n", maxRed, maxGreen, maxBlue, maxRed*maxGreen*maxBlue)
	return maxRed * maxGreen * maxBlue
}

func main() {
	fileBytes, _ := os.ReadFile("aoc-day-2.txt")
	linesArray := strings.Split(string(fileBytes), "\n")

	// for i, line := range linesArray {
	// 	game := strings.Split(line, ": ")
	// 	if len(game) > 1 {
	// 		results := strings.Split(game[1], "; ")
	// 		if parseGames(results, i, 12, 13, 14) {
	// 			fmt.Println(i + 1)
	// 			answer += i + 1
	// 		}
	// 	}
	// }
	answer := 0
	for i, line := range linesArray {
		fmt.Println(line)
		game := strings.Split(line, ": ")
		if len(game) > 1 {
			answer += parseGames2(game[1], i)
		}
	}

	fmt.Println(answer)
}
