package main

func SumArray(array [5]int) int {
	var sum int
	for i := 0; i < 5; i++ {
		sum += array[i]
	}
	return sum
}
