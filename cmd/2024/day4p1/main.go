package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/K-Road/aoc/pkg/utils"
)

func main() {
	//default for testing
	//TODO
	defaultFile := "cmd/2024/day4p1/test.input"
	//defaultFile := "test.input"

	filepath := flag.String("e", defaultFile, "Path to the file to read")
	flag.Parse()
	sum := 0

	file, err := os.Open(*filepath)
	if err != nil {
		fmt.Println("Error reading file", err)
		return
	}
	defer file.Close()

	input := utils.MatrixImportRune(file)
	fmt.Println("")
	//fmt.Println(input)

	sum = findPattern("XMAS", &input)

	//fmt.Println(len(matches))
	fmt.Println("Score:")
	fmt.Println(sum)
}

func findPattern(pattern string, data *[][]rune) int {
	found := 0

	for i, row := range *data {
		//fmt.Println(row)
		for j := range row {

			if (*data)[i][j] == rune(pattern[0]) {
				//fmt.Printf(" X Found @%d,%d \n", i, j)

				for _, direction := range utils.Directions {
					if dfs(i, j, pattern, 0, data, direction) {
						fmt.Printf("Pattern found starting (%d, %d) %v\n", i, j, direction.Key)
						found++
					}
				}
			}
		}
	}
	return found
}

func dfs(x, y int, pattern string, idx int, data *[][]rune, direction utils.Direction) bool {
	if idx == len(pattern) {
		return true
	}

	if x < 0 || x >= len(*data) || y < 0 || y >= len((*data)[x]) {
		return false
	}

	if (*data)[x][y] != rune(pattern[idx]) {
		return false
	}

	nextX, nextY := x+direction.Change[0], y+direction.Change[1]

	return dfs(nextX, nextY, pattern, idx+1, data, direction)

}
