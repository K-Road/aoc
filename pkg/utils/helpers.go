package utils

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Direction struct {
	Key    string
	Change []int
}

var Directions = []Direction{
	{Key: "UL", Change: []int{-1, -1}}, // Up-left (x-1, y-1)
	{Key: "UR", Change: []int{-1, 1}},  // Up-right (x-1, y+1)
	{Key: "DL", Change: []int{1, -1}},  // Down-left (x+1, y-1)
	{Key: "DR", Change: []int{1, 1}},   // Down-right (x+1, y+1)
	{Key: "U", Change: []int{-1, 0}},   // Up- (x-1, y)
	{Key: "R", Change: []int{0, 1}},    // right (x, y+1)
	{Key: "L", Change: []int{0, -1}},   // left (x, y-1)
	{Key: "D", Change: []int{1, 0}},    // Down (x+1, y)
}

func GetDirectionByKey(key string) *Direction {
	for _, dir := range Directions {
		if dir.Key == key {
			return &dir
		}
	}
	return nil
}

func CountOccurances(values []int) map[int]int {
	counts := make(map[int]int)
	for _, value := range values {
		counts[value]++
	}
	return counts
}

func MatrixImport(file *os.File) [][]int {
	var nums [][]int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		input := scanner.Text() //fmt.Println(scanner.Text())
		strval := strings.Fields(input)
		var row []int
		for _, value := range strval {
			num, err := strconv.Atoi(value)
			if err != nil {
				fmt.Println("Error converting", value, "to int:", err)
				return nil
			}
			row = append(row, num)
		}
		nums = append(nums, row)
	}
	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading file %v\n", err)
	}
	fmt.Println("Input:")
	for _, row := range nums {
		fmt.Println(row)
	}
	return nums
}

func MatrixImportRune(file *os.File) [][]rune {
	var matrix [][]rune

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		input := scanner.Text()
		matrix = append(matrix, []rune(input))
	}
	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading file %v\n", err)
	}
	fmt.Println("Input:")
	for _, row := range matrix {
		for _, r := range row {
			fmt.Printf("%c", r)
		}
		fmt.Printf("\n")
	}
	return matrix
}

func FileImport(file *os.File) string {
	var response string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		input := scanner.Text()
		response += input
	}
	return response
}
