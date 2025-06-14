// Selection Sort is a simple sorting algorithm.

// It works like this:

// Find the smallest element in the array.

// Swap it with the first element.

// Then find the next smallest and swap it with the second element.

// Repeat until the array is fully sorted.

package main

import "fmt"

func SelectionSort(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		minIndex := i

		for j := i + 1; j < len(arr); j++ {
			if arr[minIndex] > arr[j] {
				minIndex = j
			}
		}

		arr[i], arr[minIndex] = arr[minIndex], arr[i]
		fmt.Printf("Step %d: %v\n", i+1, arr)
	}
}

func main() {
	arr := []int{2, 4, 21, 5, 3, 12, 1, 5, 2, 5}

	SelectionSort(arr)
	fmt.Println("selection sort:", arr)
}
