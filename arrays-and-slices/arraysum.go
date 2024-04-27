package main

func Sum(slice []int) int {
	sum := 0
	for _, num := range slice {
		sum += num
	}
	return sum
}

func SumAll(summedSlices ...[]int) []int {
	var sums []int

	for _, slice := range summedSlices {
		sums = append(sums, Sum(slice))
	}

	return sums
}

func SumAllTails(summedSlices ...[]int) []int {
	var sums []int

	for _, sliceTail := range summedSlices {
		if len(sliceTail) == 0 {
			sums = append(sums, 0)
		} else {
			tail := sliceTail[1:]

			sums = append(sums, Sum(tail))
		}
	}

	return sums
}
