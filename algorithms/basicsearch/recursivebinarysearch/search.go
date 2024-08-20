package main

import "fmt"

func recursiveBinarySearch(items []int, target int) bool {
	if len(items) == 0 {
		return false
	}

	midpoint := len(items) / 2

	if items[midpoint] == target {
		return true
	}

	if items[midpoint] < target {
		return recursiveBinarySearch(items[midpoint+1:], target)
	} else {
		return recursiveBinarySearch(items[:midpoint], target)
	}
}

func verify(result bool) {
	fmt.Printf("Target found: %v\n", result)
}

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	result := recursiveBinarySearch(numbers, 12)
	verify(result)
	result = recursiveBinarySearch(numbers, 6)
	verify(result)
}
