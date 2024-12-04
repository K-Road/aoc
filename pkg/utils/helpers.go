package utils

func CountOccurances(values []int) map[int]int {
	counts := make(map[int]int)
	for _, value := range values {
		counts[value]++
	}
	return counts
}
