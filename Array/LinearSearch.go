// Linear search checks each element one by one.
// Time complexity: O(n) — good for small arrays or unsorted data.
// Returns -1 if not found, or the index where the element is found.

package main

import "fmt"

func LinearSearch(arr []int, target int) int {
	for i, value := range arr {
		if value == target {
			return i
		}
	}
	return -1
}

func main() {
	arr := []int{3, 5, 1, 235, 64, 32, 4, 234}
	var target int

	fmt.Print("enter the target:")
	fmt.Scan(&target)

	index := LinearSearch(arr, target)

	if index != -1 {
		fmt.Printf("Element %d found at index %d", target, index)
	} else {
		fmt.Printf("Element %d is not found", target)
	}
	fmt.Println("")
}
