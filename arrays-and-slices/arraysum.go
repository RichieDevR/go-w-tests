package main

func Sum(array []int) int {
	sum := 0
	for _, num := range array {
		sum += num
	}
	return sum
}

func SumAll(arraysToSum ...[]int) []int {
	var sums []int

	for _, arrays := range arraysToSum {
		sums = append(sums, Sum(arrays))
	}

	return sums
}
