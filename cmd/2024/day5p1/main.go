package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/K-Road/aoc/pkg/utils"
)

func main() {
	//default for testing
	//TODO
	//defaultFile := "cmd/2024/day5p1/test.input"
	//defaultFile := "test.input"

	//filepath := flag.String("e", defaultFile, "Path to the file to read")
	flag.Parse()
	sum := 0
	args := os.Args[1:]
	var input []string
	if len(args) == 0 {
		fmt.Println("No args")
		input = utils.TestInput()
	} else if len(args) == 3 {
		input = utils.Input(2024, 5, 1)
	} else {
		fmt.Println("Usage: ./main.go 2024 5 1")
		log.Fatal()
	}

	for _, line := range input {
		fmt.Println(line)
	}

	joinedInput := strings.Join(input, "\n")

	sections := strings.Split(joinedInput, "\n\n")

	rulesInput := strings.Split(sections[0], "\n")
	updatesInput := strings.Split(sections[1], "\n")

	//var rules [][]int
	fmt.Println("Section 1: Rules")
	rules := CreateMatrix(rulesInput, "|")
	for _, rule := range rules {
		fmt.Println(rule)
	}

	fmt.Println("Section 2: Updates")
	updates := CreateMatrix(updatesInput, ",")
	for _, update := range updates {
		fmt.Println(update)
	}

	fmt.Println("Score:")
	fmt.Println(sum)
}

func CreateMatrix(input []string, rowsep string) [][]int {
	var nums [][]int

	for _, line := range input {
		strval := strings.Join(strings.Split(line, rowsep), " ")
		// fmt.Println(strval)
		var row []int
		for _, value := range strings.Fields(strval) {
			num, err := strconv.Atoi(value)
			if err != nil {
				fmt.Println("Error converting", value, "to int:", err)
				return nil
			}
			row = append(row, num)
		}
		nums = append(nums, row)
	}
	// fmt.Println("Input:")
	// for _, row := range nums {
	// 	fmt.Println(row)
	// }
	return nums
}
