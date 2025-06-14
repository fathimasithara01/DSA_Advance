package main

import "fmt"

func bubbleSort(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

func main() {
	arr := []int{5, 3, 2, 3, 1, 34, 2, 55}
	fmt.Println("Original array", arr)

	bubbleSort(arr)
	fmt.Println("Sorted array", arr)
}
