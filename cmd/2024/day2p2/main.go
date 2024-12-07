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
	defaultFile := "cmd/2024/day2p2/test.input"

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
		isSafe := false // Assume safe
		if check(row) {
			isSafe = true

		} else {
			//fmt.Println(row)
			for j := 0; j < len(row); j++ {
				tempRow := append([]int{}, row[:j]...)
				tempRow = append(tempRow, row[j+1:]...)
				if check(tempRow) {
					isSafe = true
					break
				}
			}
		}

		if isSafe {
			sum++
			fmt.Println(row, " :SAFE")
		} else {
			//fmt.Println(nums[i])
			fmt.Println(row, ":UNSAFE")
		}
	}

	fmt.Println("Score:")
	fmt.Println(sum)
}

// determineTrend determines the direction (trend) of the entire row: increasing or decreasing
func determineTrend(row []int) string {
	upCount := 0
	downCount := 0
	flat := 0

	// Compare each adjacent pair to determine the overall trend
	for i := 0; i < len(row)-1; i++ {
		if row[i] < row[i+1] {
			upCount++
		} else if row[i] > row[i+1] {
			downCount++
		} else if row[i] == row[i+1] {
			flat++
		}
		if upCount == 2 {
			return "increasing"
		}
		if downCount == 2 {
			return "decreasing"
		}
		if flat == 2 {
			return "flat"
		}
	}

	if upCount > downCount {
		return "increasing"
	} else if downCount > upCount {
		return "decreasing"
	} else {
		return "flat"
	}
}

func findBreakingIdx(row []int, trend string) int {
	n := len(row)
	if n < 2 {
		return -1 // No violation possible if there are less than two elements
	}

	// Check if the first and second elements are equal
	if row[0] == row[1] {
		return 0 // Violation at index 0 if they are equal
	}
	// Handle increasing trend
	if trend == "increasing" {
		for i := 1; i < n; i++ {
			if row[i] < row[i-1] {
				return i // First violation in increasing sequence
			}
		}
	}

	// Handle decreasing trend
	if trend == "decreasing" {
		for i := 1; i < n; i++ {
			if row[i] > row[i-1] {
				return i // First violation in decreasing sequence
			}
		}
	}

	return -1 // No violation found
}

func checkRow(row []int) int {
	lastViolationIdx := -1
	// Check for violations in the sequence
	for j := 0; j < len(row)-1; j++ {
		if row[j] == row[j+1] {
			lastViolationIdx = j
		}

		delta := int(math.Abs(float64(row[j+1] - row[j])))

		if delta > 3 || delta < 1 {
			lastViolationIdx = j + 1
		}
	}

	return lastViolationIdx
}
func check(row []int) bool {
	//lastViolationIdx := -1
	// Check for violations in the sequence
	direction := 0 // 1 for increasing, -1 for decreasing, 0 for undefined
	for j := 0; j < len(row)-1; j++ {
		if row[j] == row[j+1] {
			return false
		}

		if row[j+1] > row[j] { // Increasing
			if direction == -1 { // Previously decreasing
				return false
			}
			direction = 1
		} else if row[j+1] < row[j] { // Decreasing
			if direction == 1 { // Previously increasing
				return false
			}
			direction = -1
		}

		// If row[i] == row[i-1], continue checking (flat segment)

		delta := int(math.Abs(float64(row[j+1] - row[j])))

		if delta > 3 || delta < 1 {
			return false
		}
	}

	return true
}
