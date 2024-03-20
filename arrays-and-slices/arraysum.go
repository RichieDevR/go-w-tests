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

	sumOfNumbersInArray := make([]int, numberOfArraysToSum)

	for i, numbersInArray := range arraysToSum {
		sumOfNumbersInArray[i] = SumArray(numbersInArray)
	}

	return sumOfNumbersInArray
}
