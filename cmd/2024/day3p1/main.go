package main

import (
	"flag"
	"fmt"
	"os"
	"regexp"
	"strconv"

	"github.com/K-Road/aoc/pkg/utils"
)

func main() {
	//default for testing
	//TODO
	//defaultFile := "cmd/2024/day3p1/test.input"
	defaultFile := "test.input"

	filepath := flag.String("e", defaultFile, "Path to the file to read")
	flag.Parse()
	sum := 0

	file, err := os.Open(*filepath)
	if err != nil {
		fmt.Println("Error reading file", err)
		return
	}
	defer file.Close()

	input := utils.FileImport(file)
	fmt.Println("")

	fmt.Println(input)
	mulPattern := `mul\((\d+),(\d+)\)`
	re := regexp.MustCompile(mulPattern)
	matches := re.FindAllStringSubmatch(input, -1)

	for _, match := range matches {
		num1, err1 := strconv.Atoi(match[1])
		num2, err2 := strconv.Atoi(match[2])

		if err1 != nil || err2 != nil {
			fmt.Println("Error converting ints, err1, err2")
			continue
		}
		sum += (num1 * num2)
	}

	fmt.Println("Score:")
	fmt.Println(sum)
}

// func mulPattern() string {
// 	return `mul\((\d+),(\d+)\)`
// }
