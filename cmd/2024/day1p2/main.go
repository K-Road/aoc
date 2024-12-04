package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/K-Road/aoc/pkg/utils"
)

func main() {
	//default for testing
	//TODO
	defaultFile := "cmd/2024/day1p2/test.input"

	filepath := flag.String("e", defaultFile, "Path to the file to read")
	flag.Parse()

	file, err := os.Open(*filepath)
	if err != nil {
		fmt.Println("Error reading file", err)
		return
	}
	defer file.Close()

	var left []int
	var right []int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		input := scanner.Text() //fmt.Println(scanner.Text())
		strval := strings.Fields(input)

		l, err := strconv.Atoi(strval[0])
		if err != nil {
			fmt.Println("Error converting", strval[0], "to int:", err)
			return
		}
		r, err := strconv.Atoi(strval[1])
		if err != nil {
			fmt.Println("Error converting", strval[1], "to int:", err)
			return
		}

		left = append(left, l)
		right = append(right, r)

	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading file %v\n", err)
	}

	// fmt.Println(left)
	// fmt.Println(right)

	//sort.Ints(left)
	//sort.Ints(right)

	counts := utils.CountOccurances(right)

	// fmt.Println(left)
	// fmt.Println(right)

	sum := 0
	for i := range left {
		fmt.Println(counts[left[i]])
		sum += left[i] * counts[left[i]]
	}

	score := 0
	score = sum
	fmt.Println(score)
}
