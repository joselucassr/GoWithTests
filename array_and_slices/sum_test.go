package main

import (
	"slices"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("collection of five numbers", func(t *testing.T) {

		numbers := []int{1, 2, 3, 4, 5}

		got := Sum(numbers)
		expected := 15
		compareSum(t, expected, got, numbers)
	})

	t.Run("collection of any size", func(t *testing.T) {

		numbers := []int{1, 2, 3}

		got := Sum(numbers)
		expected := 6
		compareSum(t, expected, got, numbers)
	})
}

func compareSum(t testing.TB, expected, got int, numbers []int) {
	t.Helper()

	if expected != got {
		t.Errorf("expected %d but got %d, given %v", expected, got, numbers)
	}
}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2, 3}, []int{1, 2, 3, 4, 5})
	expected := []int{6, 15}
	compareSlices(t, expected, got)
}

func TestSumAllTails(t *testing.T) {
	t.Run("make sum of some slices", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{0, 9})
		expected := []int{2, 9}
		compareSlices(t, expected, got)
	})

	t.Run("safely sum empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{3, 4, 5})
		expected := []int{0, 9}
		compareSlices(t, expected, got)
	})
}

func compareSlices(t testing.TB, expected, got []int) {
	t.Helper()

	if !slices.Equal(expected, got) {
		t.Errorf("expected %v but got %v", expected, got)
	}
}
