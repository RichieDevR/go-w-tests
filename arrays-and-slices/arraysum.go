package main

func SumArray(array [5]int) int {
	sum := 0
	for _, num := range array {
		sum += num
	}
	return sum
}
