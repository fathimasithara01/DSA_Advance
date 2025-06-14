package main

import "fmt"

func insertionSortDescending(arr []int) {
	n := len(arr)

	for i := 1; i < n; i++ {
		key := arr[i]
		j := i - 1

		fmt.Printf("\nStep %d: inserting %d\n", i, key)

		for j >= 0 && arr[j] < key {
			arr[j+1] = arr[j]
			j--
		}

		arr[j+1] = key

		fmt.Println("Current array:", arr)
	}
}

func main() {
	arr := []int{2, 5, 4, 6, 1, 7}

	fmt.Println("Original array:", arr)

	insertionSortDescending(arr)

	fmt.Println("\nFinal sorted array (descending):", arr)
}
