package main

import "testing"

func TestSumArray(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5}
	var sum int

	for i := 0; i < len(numbers); i++ {
		sum += numbers[i]
	}

	got := SumArray(numbers)

	want := sum

	if got != want {
		t.Errorf("got %d want %d given %v", got, want, numbers)
	}
}
