package utils

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

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

func FileImport(file *os.File) string {
	var response string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		input := scanner.Text()
		response += input
	}
	return response
}
