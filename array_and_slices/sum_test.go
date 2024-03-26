package main

import "testing"

func TestSum(t *testing.T) {
	numbers := [5]int{1, 2, 3, 4, 5}

	got := Sum(numbers)
	expected := 15

	if expected != got {
		t.Errorf("expected %d but got %d, given %v", expected, got, numbers)
	}
}