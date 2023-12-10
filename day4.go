package main

import (
	"fmt"
	"os"
	"strings"
)

func parseNumbers(line string) ([]string, []string) {
	fmt.Println(line)
	splitLine := strings.Split(line, "|")
	winningNumberStr := strings.Split(splitLine[0], ":")
	winningNumberArr := strings.Split(winningNumberStr[1], " ")

	actualNumberArr := strings.Split(splitLine[1], " ")

	fmt.Println(winningNumberArr, actualNumberArr)

	return nil, nil
}

func main() {
	fileBytes, _ := os.ReadFile("sample.txt")
	trimmedFile := strings.TrimSpace(string(fileBytes))
	linesArray := strings.Split(trimmedFile, "\n")

	for _, line := range linesArray {
		parseNumbers(line)
	}
}
