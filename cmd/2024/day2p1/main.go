package main

import (
	"flag"
	"fmt"
	"math"
	"os"

	"github.com/K-Road/aoc/pkg/utils"
)

func main() {
	//default for testing
	//TODO
	defaultFile := "cmd/2024/day2p1/test.input"

	filepath := flag.String("e", defaultFile, "Path to the file to read")
	flag.Parse()
	sum := 0

	file, err := os.Open(*filepath)
	if err != nil {
		fmt.Println("Error reading file", err)
		return
	}
	defer file.Close()

	nums := utils.MatrixImport(file)
	fmt.Println("")

	for i := range nums {
		row := nums[i]

		var direction bool
		if row[0] > row[1] {
			direction = true //upward
		} else {
			direction = false //downward
		}

		isSafe := true // assume safe
		// for j := range len(row) - 1 {
		// 	if row[j] == row[j+1] {
		// 		isSafe = false
		// 		break // not safe move to next row
		// 	}
		// 	if row[j] > row[j+1] && !direction {
		// 		isSafe = false
		// 		break //not safe direction change
		// 	} else if row[j] < row[j+1] && direction {
		// 		isSafe = false
		// 		break
		// 	}

		// 	delta := int(math.Abs(float64(row[j+1] - row[j])))

		// 	if delta > 3 || delta < 1 {
		// 		isSafe = false
		// 		break //not safe delta not 1-3
		// 	}
		// }
		for j := range len(row) - 1 {
			if row[j] == row[j+1] {
				isSafe = false
				break // not safe move to next row
			}
			if row[j] > row[j+1] && !direction {
				isSafe = false
				break //not safe direction change
			} else if row[j] < row[j+1] && direction {
				isSafe = false
				break
			}

			delta := int(math.Abs(float64(row[j+1] - row[j])))

			if delta > 3 || delta < 1 {
				isSafe = false
				break //not safe delta not 1-3
			}
		}

		if isSafe {
			sum++
			fmt.Println(row, " :SAFE")
		}
	}

	score := 0
	score = sum

	fmt.Println("Score:")
	fmt.Println(score)
}
