package main

import "fmt"

func insertionSort(arr []int) {
	n := len(arr)

	for i := 1; i < n; i++ {
		// key := arr[i]
		j := i - 1

		fmt.Printf("\nStep %d: inserting %d\n", i, arr[i])

		for j >= 0 && arr[j] > arr[i] {
			arr[(j)+1] = arr[j]
			j--
		}

		arr[j+1] = arr[i]

		fmt.Println("Current array:", arr)
	}
}

func main() {
	arr := []int{2, 5, 4, 6, 1, 7}

	fmt.Println("Original array:", arr)

	insertionSort(arr)

	fmt.Println("\nFinal sorted array:", arr)
}

//  Time Complexity:
// Case	Time
// Best Case	O(n)
// Average	O(n²)
// Worst Case	O(n²)
// Space	O(1)
