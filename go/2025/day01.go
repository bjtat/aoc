//go:build day01

package main

import (
	"aoc/go/utils"
	"fmt"
	"os"
)

func main() {
	inputPath, _ := utils.GetInputPath(2025, 1)
	fileBytes, _ := os.ReadFile(inputPath)
	fmt.Println("Starting day 1")
}
