package main

func Sum(numbers [5]int) int {
	sum := numbers[0]

	for i := 1; i < len(numbers); i++ {
		sum += numbers[i]
	}

	return sum
}
