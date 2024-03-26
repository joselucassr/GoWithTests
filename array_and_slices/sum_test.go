package main

import "testing"

func TestSum(t *testing.T) {
	t.Run("collection of five numbers", func(t *testing.T) {

		numbers := []int{1, 2, 3, 4, 5}

		got := Sum(numbers)
		expected := 15
		compareResults(t, expected, got, numbers)
	})

	t.Run("collection of any size", func(t *testing.T) {

		numbers := []int{1, 2, 3}

		got := Sum(numbers)
		expected := 6
		compareResults(t, expected, got, numbers)
	})
}

func compareResults(t testing.TB, expected, got int, numbers []int) {
	t.Helper()

	if expected != got {
		t.Errorf("expected %d but got %d, given %v", expected, got, numbers)
	}
}
