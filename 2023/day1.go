package main

import (
	"fmt"
	"os"
	"strings"
)

const (
	ZERO_UNICODE byte = 48
	NINE_UNICODE byte = 57
)

func isDigit(b byte) bool {
	return ZERO_UNICODE <= b && b <= NINE_UNICODE
}

func parseFirstDigit(str string) byte {
	for i := 0; i < len(str); i++ {
		if isDigit(str[i]) {
			return str[i] - ZERO_UNICODE
		}
	}
	return 0
}

func parseLastDigit(str string) byte {
	for i := len(str) - 1; i > -1; i-- {
		if isDigit(str[i]) {
			return str[i] - ZERO_UNICODE
		}
	}
	return 0
}

func convertToNumber(firstDigit byte, secondDigit byte) int {
	return int(firstDigit*10 + secondDigit)
}

func main() {
	fileBytes, _ := os.ReadFile("aoc-day-1.txt")
	cleanstring := string(fileBytes)
	fileString := string(fileBytes)
	fileString = strings.ReplaceAll(fileString, "one", "o1ne")
	fileString = strings.ReplaceAll(fileString, "two", "t2wo")
	fileString = strings.ReplaceAll(fileString, "three", "th3ree")
	fileString = strings.ReplaceAll(fileString, "four", "fo4ur")
	fileString = strings.ReplaceAll(fileString, "five", "fi5ve")
	fileString = strings.ReplaceAll(fileString, "six", "s6ix")
	fileString = strings.ReplaceAll(fileString, "seven", "se7ven")
	fileString = strings.ReplaceAll(fileString, "eight", "ei8ght")
	fileString = strings.ReplaceAll(fileString, "nine", "ni9ne")
	linesArray := strings.Split(fileString, "\n")
	cleanlinesArray := strings.Split(cleanstring, "\n")

	for _, line := range linesArray {
		fmt.Println(line)
	}

	answer := 0
	for i, line := range linesArray {
		number := convertToNumber(parseFirstDigit(line), parseLastDigit(line))
		fmt.Printf("%v : %v\n", cleanlinesArray[i], number)
		if number > 0 {
			answer += number
		}
	}
	fmt.Println(answer)
}
