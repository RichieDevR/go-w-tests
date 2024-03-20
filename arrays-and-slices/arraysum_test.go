package main

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5}
	var sum int

	for i := 0; i < len(numbers); i++ {
		sum += numbers[i]
	}

	got := Sum(numbers)

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

func TestSumAllTails(t *testing.T) {
	t.Run("make array of sum for tail of slices", func(t *testing.T) {
		got := SumAllTails([]int{1, 2, 3}, []int{0, 9, 3})

		want := []int{5, 12}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("I got %v but I want %v", got, want)
		}
	})
	t.Run("safely sum empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{3, 9, 3})

		want := []int{0, 12}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("I got %v but I want %v", got, want)
		}
	})
}
