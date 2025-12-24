//go:build day03

package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"aoc/utils"
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
	for xcord := c.x - 1; xcord < c.x+2; xcord++ {
		for ycord := c.y - 1; ycord < c.y+2; ycord++ {
			if !strings.ContainsAny(grid[xcord][ycord], ".0123456789") {
				return true
			}
		}
	}
	return false
}

func detectGear(grid [][]string, c coord) bool {
	for xcord := c.x - 1; xcord < c.x+2; xcord++ {
		for ycord := c.y - 1; ycord < c.y+2; ycord++ {
			if strings.ContainsAny(grid[xcord][ycord], "*") {
				return true
			}
		}
	}
	return false
}

func tagGearsNextToNumber(grid [][]string, c coord, firstIndex, num int) {
	for i := c.x - 1; i <= c.x+1; i++ {
		if strings.ContainsAny(grid[i][firstIndex-1], "*") {
			fmt.Printf("num: %d, [%d, %d]\n", num, i, firstIndex-1)
		}
	}

	for i := firstIndex; i <= c.y; i++ {
		if strings.ContainsAny(grid[c.x-1][i], "*") {
			fmt.Printf("num: %d, [%d, %d]\n", num, c.x-1, i)
		}
		if strings.ContainsAny(grid[c.x+1][i], "*") {
			fmt.Printf("num: %d, [%d, %d]\n", num, c.x+1, i)
		}
	}

	for i := c.x - 1; i <= c.x+1; i++ {
		if strings.ContainsAny(grid[i][c.y+1], "*") {
			fmt.Printf("num: %d, [%d, %d]\n", num, i, c.y+1)
		}
	}
}

func generateNumberFromLastIndexWithGear(grid [][]string, c coord) int {
	numStr := ""
	isValid := false
	firstIndex := 0

	for maxIndex := c.y; maxIndex >= 0; maxIndex-- {
		if strings.ContainsAny(grid[c.x][maxIndex], "0123456789") {
			numStr = grid[c.x][maxIndex] + numStr
			isValid = isValid || detectGear(grid, coord{c.x, maxIndex})
		} else {
			firstIndex = maxIndex + 1
			break
		}
	}

	num, err := strconv.Atoi(numStr)
	if err != nil {
		return -1
	}

	if isValid {
		fmt.Printf("\n\n%s\n", numStr)
		// fmt.Printf("last index: %d, first index: %d\n\n", c.y, firstIndex)
		tagGearsNextToNumber(grid, c, firstIndex, num)
		return num
	} else {
		return 0
	}
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

func checkAroundGear(grid [][]string, c coord) int {
	uniqueSet := make(map[int]bool, 6)

	if strings.ContainsAny(grid[c.x][c.y-1], "0123456789") {
		num := generateNumberFromAnyIndex(grid, coord{c.x, c.y - 1})
		_, ok := uniqueSet[num]
		if !ok {
			uniqueSet[num] = true
		}
	}

	for i := c.y - 1; i <= c.y+1; i++ {
		if strings.ContainsAny(grid[c.x-1][i], "0123456789") {
			num := generateNumberFromAnyIndex(grid, coord{c.x - 1, i})
			_, ok := uniqueSet[num]
			if !ok {
				uniqueSet[num] = true
			}
		}
		if strings.ContainsAny(grid[c.x+1][i], "0123456789") {
			num := generateNumberFromAnyIndex(grid, coord{c.x + 1, i})
			_, ok := uniqueSet[num]
			if !ok {
				uniqueSet[num] = true
			}
		}
	}

	if strings.ContainsAny(grid[c.x][c.y+1], "0123456789") {
		num := generateNumberFromAnyIndex(grid, coord{c.x, c.y + 1})
		_, ok := uniqueSet[num]
		if !ok {
			uniqueSet[num] = true
		}
	}

	valid := 0
	answer := 1
	for num, ok := range uniqueSet {
		if ok {
			valid++
			answer *= num
		}
	}

	if valid == 2 {
		fmt.Printf("gear ratio: %d\n", answer)
		return answer
	}
	return 0
}

func generateNumberFromAnyIndex(grid [][]string, c coord) int {
	numStr := ""
	maxPossibleIndex := len(grid[0])
	firstIndexOfNum := 0

	for someIndex := c.y; someIndex > 0; someIndex-- {
		if strings.ContainsAny(grid[c.x][someIndex], "0123456789") {
			firstIndexOfNum = someIndex
		} else {
			break
		}
	}

	for i := firstIndexOfNum; i < maxPossibleIndex; i++ {
		if strings.ContainsAny(grid[c.x][i], "0123456789") {
			numStr = numStr + grid[c.x][i]
		} else {
			break
		}
	}

	num, err := strconv.Atoi(numStr)
	if err != nil {
		return -1
	}

	return num
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

func checkLineWithGears(grid [][]string, x int) int {
	gridLine := grid[x]
	answer := 0
	for i := 0; i < len(gridLine); i++ {
		if strings.ContainsAny(gridLine[i], "*") {
			fmt.Printf("FOUND GEAR: [%d, %d]\n", x, i)
			answer += checkAroundGear(grid, coord{x, i})
		}
	}
	return answer
}

func main() {
	inputPath, _ := utils.GetInputPath(2023, 3)
	fileBytes, _ := os.ReadFile(inputPath)
	// fileBytes, _ := os.ReadFile("sample.txt")
	grid := generateGrid(fileBytes)

	for _, gridLine := range grid {
		fmt.Println(gridLine)
	}

	// detectSymbol(grid, coord{1, 1})
	// detectSymbol(grid, coord{1, 2})
	// detectSymbol(grid, coord{1, 3})

	answer := 0
	for i, _ := range grid {
		answer += checkLineWithGears(grid, i)
	}
	fmt.Println(answer)

	// generateNumberFromAnyIndex(grid, coord{1, 1})
	// generateNumberFromAnyIndex(grid, coord{1, 2})
	// generateNumberFromAnyIndex(grid, coord{1, 3})

	// generateNumberFromLastIndex(grid, coord{1, 3})
	// generateNumberFromLastIndex(grid, coord{1, 8})

	// tagGearsNextToNumber(grid, coord{1, 3}, 1, 467)
}
