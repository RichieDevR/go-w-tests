package main

func SumArray(array []int) int {
	sum := 0
	for _, num := range array {
		sum += num
	}
	return sum
}

func SumAll(arraysToSum ...[]int) []int {
	numberOfArraysToSum := len(arraysToSum)

	sumOfNumbers := make([]int, numberOfArraysToSum)

	for i, numbersInArray := range arraysToSum {
		sumOfNumbers[i] = SumArray(numbersInArray)
	}

	return sumOfNumbers
}
