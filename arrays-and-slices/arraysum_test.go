package main

import "testing"

func TestSumArray(t *testing.T) {
	t.Run("collection of exactly 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		var sum int

		for _, num := range numbers {
			sum += num
		}

		got := SumArray(numbers)

		want := sum

		if got != want {
			t.Errorf("got %d want %d given the array [numbers]int %v ", got, want, numbers)
		}
	})
	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		var sum int

		for i := 0; i < len(numbers); i++ {
			sum += numbers[i]
		}

		got := SumArray(numbers)

		want := sum

		if got != want {
			t.Errorf("got %d want %d given %v", got, want, numbers)
		}
	})
}
