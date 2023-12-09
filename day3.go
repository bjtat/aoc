package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// we wants a map where k= gear coord and v = the numbers that are valid next to it
// then we look at any key where it has exactly two values to it
var gearMap = make(map[int][]coord)

type coord struct {
	x int
	y int
}

type number struct {
	numInt    int
	coordList []coord
}

func generateGrid(fileBytes []byte) [][]string {
	fileArray := strings.Split(string(fileBytes), "\n")

	// We only add 1 because there's a \n on the last line which creates an empty element
	heightLen := len(fileArray) + 1
	grid := make([][]string, heightLen)

	// Generate top and bottom padding by repeating "." then converting into an array
	widthLen := len(fileArray[0]) + 2
	paddingArray := strings.Split(strings.Repeat(".", widthLen), "")
	grid[0] = paddingArray
	grid[heightLen-1] = paddingArray

	for i, line := range fileArray {
		if len(line) > 0 {
			line = fmt.Sprintf(".%s.", line) // Add BOL and EOL padding
			lineArray := strings.Split(line, "")
			grid[i+1] = lineArray
		}
	}

	return grid
}

func detectSymbol(grid [][]string, c coord) bool {
	// fmt.Printf("checking %s\n", grid[c.x][c.y])
	for xcord := c.x - 1; xcord < c.x+2; xcord++ {
		for ycord := c.y - 1; ycord < c.y+2; ycord++ {
			// fmt.Printf("[%s] ", grid[xcord][ycord])
			if !strings.ContainsAny(grid[xcord][ycord], ".0123456789") {
				// fmt.Println("found one!\n")
				return true
			}
		}
		// fmt.Println()
	}
	return false
}

func generateNumberFromLastIndex(grid [][]string, c coord) int {
	numStr := ""
	isValid := false

	for maxIndex := c.y; maxIndex > 0; maxIndex-- {
		if strings.ContainsAny(grid[c.x][maxIndex], "0123456789") {
			numStr = grid[c.x][maxIndex] + numStr
			isValid = isValid || detectSymbol(grid, coord{c.x, maxIndex})
		} else {
			break
		}
	}

	fmt.Printf("%s: %v\n", numStr, isValid)

	num, err := strconv.Atoi(numStr)
	if err != nil {
		return -1
	}

	if isValid {
		return num
	} else {
		return 0
	}
}

func checkLine(grid [][]string, x int) int {
	gridLine := grid[x]
	num := 0
	for i := 0; i < len(gridLine); i++ {
		if strings.ContainsAny(gridLine[i], "01234567899") && !strings.ContainsAny(gridLine[i+1], "0123456789") {
			num += generateNumberFromLastIndex(grid, coord{x, i})
		}
	}
	return num
}

func main() {
	fileBytes, _ := os.ReadFile("aoc-day-3.txt")
	grid := generateGrid(fileBytes)

	for _, gridLine := range grid {
		fmt.Println(gridLine)
	}

	// detectSymbol(grid, coord{1, 1})
	// detectSymbol(grid, coord{1, 2})
	// detectSymbol(grid, coord{1, 3})

	answer := 0
	for i, _ := range grid {
		answer += checkLine(grid, i)
	}
	fmt.Print(answer)

	// generateNumberFromLastIndex(grid, coord{1, 3})
	// generateNumberFromLastIndex(grid, coord{1, 8})
}