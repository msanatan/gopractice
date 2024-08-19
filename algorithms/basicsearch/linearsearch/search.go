package main

import "fmt"

func linearSearch(items []int, target int) int {
	for i := 0; i < len(items); i++ {
		if items[i] == target {
			return i
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
	result := linearSearch(numbers, 12)
	verify(result)
	result = linearSearch(numbers, 6)
	verify(result)
}
