// Make sure the array is sorted before using binary search.
// Time complexity: O(log n)
// Returns the index if found, or -1 if not found.

//1. Start with the middle element of the array.
//2. Compare the target with the middle element:
// If target == middle → Found
// If target < middle → Search in the left half
// If target > middle → Search in the right half
//3. Repeat the steps on the selected half until the element is found or search space becomes empty.

package main

import (
	"fmt"
	"sort"
)

func binarySearch(arr []int, target int) int {
	left := arr[0]
	right := len(arr) - 1

	for left < right {
		mid := left + (right-left)/2

		if arr[mid] == target {
			return mid
		} else if arr[mid] < target { //If target < middle →// Target is in the right half
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}

func main() {
	arr := []int{324, 5, 32, 43, 5, 3, 21, 3, 2}

	sort.Ints(arr)
	fmt.Println("sorted aray:", arr)

	var target int
	fmt.Println("Enter target")
	fmt.Scan(&target)

	index := binarySearch(arr, target)

	if index != -1 {
		fmt.Printf("Elemnt %d at index %d", target, index)
	} else {
		fmt.Println("Not found")
	}

}
