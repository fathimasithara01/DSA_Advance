package main

import "fmt"

func SelectionSortDes(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		maxIndex := i

		for j := i + 1; j < len(arr); j++ {
			if arr[j] > arr[maxIndex] {
				maxIndex = j
			}
		}

		arr[i], arr[maxIndex] = arr[maxIndex], arr[i]j
		fmt.Printf("Step %d: %v\n", i+1, arr)
	}
}

func main() {
	arr := []int{2, 4, 21, 5, 3, 12, 1, 5, 2, 5}

	SelectionSortDes(arr)
	fmt.Println("selection sort:", arr)
}
// Metric	Value
// Time (Best)	O(n²)
// Time (Average)	O(n²)
// Time (Worst)	O(n²)
// Space Complexity	O(1)