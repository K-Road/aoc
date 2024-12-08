package main

import (
	"flag"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"

	"github.com/K-Road/aoc/pkg/utils"
)

func main() {
	//default for testing
	//TODO
	//defaultFile := "cmd/2024/day3p2/test.input"
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
	matchesidx := re.FindAllStringSubmatchIndex(input, -1)

	dontPattern := `don't\(\)`
	doPattern := `do\(\)`

	re = regexp.MustCompile(dontPattern)
	dontidx := re.FindAllStringSubmatchIndex(input, -1)
	re = regexp.MustCompile(doPattern)
	doidx := re.FindAllStringSubmatchIndex(input, -1)

	type KeyValue struct {
		Key  int
		List []interface{}
	}

	inputs := []KeyValue{}

	for i := 0; i < len(matchesidx); i++ {
		inputs = append(inputs, KeyValue{
			Key:  matchesidx[i][0],
			List: []interface{}{matches[i]},
		})
	}
	for i := 0; i < len(doidx); i++ {
		inputs = append(inputs, KeyValue{
			Key:  doidx[i][0],
			List: []interface{}{true},
		})
	}
	for i := 0; i < len(dontidx); i++ {
		inputs = append(inputs, KeyValue{
			Key:  dontidx[i][0],
			List: []interface{}{false},
		})
	}

	// Sort the inputs slice by Key field using a custom sort function
	sort.Slice(inputs, func(i, j int) bool {
		return inputs[i].Key < inputs[j].Key // Compare by Key field
	})
	enabled := true
	// Print the results
	for _, kv := range inputs {
		//fmt.Printf("Key: %d, List: %v\n", kv.Key, kv.List)

		if enabled {
			if kv.List[0] == true {
				continue
			}
			if kv.List[0] == false {
				enabled = false
				//	fmt.Printf("DISABLING:DISABLING:DISABLING:DISABLING: SCORE @: %d\n", sum)
				continue
			}
			mulMatch := kv.List[0].([]string)
			if len(mulMatch) > 2 {

				num1, err1 := strconv.Atoi(mulMatch[1])
				num2, err2 := strconv.Atoi(mulMatch[2])

				if err1 != nil || err2 != nil {
					fmt.Println("Error converting ints, err1, err2")
					continue
				}
				//fmt.Printf("SUMMING: %v %v\n", num1, num2)
				sum += (num1 * num2)
			}
		} else {
			if kv.List[0] == true {
				enabled = true
				//fmt.Printf("ENABLING:ENABLING:ENABLING:ENABLING:SCORE @: %d\n", sum)
				continue
			}
			if kv.List[0] == false {
				enabled = false
				continue
			}
		}
	}
	//fmt.Println(len(matches))
	fmt.Println("Score:")
	fmt.Println(sum)
}
