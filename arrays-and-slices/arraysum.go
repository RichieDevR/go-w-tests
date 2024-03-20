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

func SumAllTails(arraysToSum ...[]int) []int {
	var sums []int

	for _, arrayTail := range arraysToSum {
		if len(arrayTail) == 0 {
			sums = append(sums, 0)
		} else {

			tail := arrayTail[1:]

			sums = append(sums, Sum(tail))
		}
	}

	return sums
}
