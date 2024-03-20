package main

import "testing"

func TestSumArray(t *testing.T) {
	numbers := [5]int{1, 2, 3, 4, 5}
	got := SumArray(numbers)
	want := 15

	if got != want {
		t.Errorf("got %d want %d given the array [numbers]int %v ", got, want, numbers)
	}
}
