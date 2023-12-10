package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fileBytes, _ := os.ReadFile("sample.txt")
	linesArray := strings.Split(string(fileBytes), "\n")
	fmt.Println(linesArray)
}
