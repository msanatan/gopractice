package main

import "fmt"

func binarySearch(items []int, target int) int {
	first := 0
	last := len(items) - 1

	for first <= last {
		median := (first + last) / 2

		if items[median] == target {
			return median
		} else if items[median] < target {
			first = median + 1
		} else {
			last = median - 1
		}
	}

	return -1
}

func verify(index int) {
	if index > -1 {
		fmt.Printf("Target was found at position %d\n", index)
	} else {
		fmt.Println("Target was not found")
	}
}

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	result := binarySearch(numbers, 12)
	verify(result)
	result = binarySearch(numbers, 6)
	verify(result)
}
