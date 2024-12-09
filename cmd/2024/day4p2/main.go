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
	defaultFile := "cmd/2024/day4p2/test.input"
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

	sum = findPattern("MAS", &input)

	//fmt.Println(len(matches))
	fmt.Println("Score:")
	fmt.Println(sum)
}

func findPattern(pattern string, data *[][]rune) int {
	found := 0

	for i, row := range *data {
		//fmt.Println(row)
		for j := range row {

			if (*data)[i][j] == rune(pattern[1]) {
				if i == 0 || j == 0 || i == len(*data)-1 || j == len((*data)[i])-1 {
					fmt.Printf(" A Found @%d,%d SKIPPING\n", i, j)
					continue
				}
				fmt.Printf(" A Found @%d,%d \n", i, j)

				// for _, direction := range utils.Directions {
				if dfs(i-1, j-1, pattern, 0, data, *utils.GetDirectionByKey("DR")) && dfs(i-1, j+1, pattern, 0, data, *utils.GetDirectionByKey("DL")) { //start UL match DR
					// Pattern successfully matched
					fmt.Printf("Pattern found starting (%d, %d) and (%d, %d)\n", i-1, j-1, i-1, j+1)
					fmt.Printf("Center character: %c %d %d\n", (*data)[i][j], i, j)
					found++
				}

				// for _, direction := range utils.Directions {
				if dfs(i-1, j-1, pattern, 0, data, *utils.GetDirectionByKey("DR")) && dfs(i+1, j-1, pattern, 0, data, *utils.GetDirectionByKey("UR")) { //start UL match DR
					// Pattern successfully matched
					fmt.Printf("Pattern found starting (%d, %d) and (%d, %d)\n", i-1, j-1, i+1, j-1)
					fmt.Printf("Center character: %c %d %d\n", (*data)[i][j], i, j)
					found++

				}

				if dfs(i+1, j+1, pattern, 0, data, *utils.GetDirectionByKey("UL")) && dfs(i+1, j-1, pattern, 0, data, *utils.GetDirectionByKey("UR")) { //start UL match DR
					// Pattern successfully matched
					fmt.Printf("Pattern found starting (%d, %d) and (%d, %d)\n", i+1, j+1, i+1, j-1)
					fmt.Printf("Center character: %c %d %d\n", (*data)[i][j], i, j)
					found++

				}
				if dfs(i+1, j+1, pattern, 0, data, *utils.GetDirectionByKey("UL")) && dfs(i-1, j+1, pattern, 0, data, *utils.GetDirectionByKey("DL")) { //start UL match DR
					// Pattern successfully matched
					fmt.Printf("Pattern found starting (%d, %d) and (%d, %d)\n", i+1, j+1, i-1, j+1)
					fmt.Printf("Center character: %c %d %d\n", (*data)[i][j], i, j)
					found++

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

	nextX, nextY := x+(direction.Change[0]), y+(direction.Change[1])
	//fmt.Printf("nX: %d, nY:%d\n", nextX, nextY)
	return dfs(nextX, nextY, pattern, idx+1, data, direction)

}
