package main

import (
	"reflect"
	"testing"
)

func TestSumArray(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5}
	var sum int

	for i := 0; i < len(numbers); i++ {
		sum += numbers[i]
	}

	got := SumArray(numbers)

	want := sum

	if got != want {
		t.Errorf("I got %d but I wanted %d when I'm given %v", got, want, numbers)
	}
}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2}, []int{0, 9}, []int{9, 1})

	want := []int{3, 9, 10}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("I got %v but I want %v", got, want)
	}
}
