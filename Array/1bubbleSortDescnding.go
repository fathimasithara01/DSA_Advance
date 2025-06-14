package main

import "fmt"

func bubbleSortDes(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-1; j++ {
			if arr[j] < arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

func main() {
	arr := []int{34, 52, 1, 3, 5, 3, 12, 43}
	fmt.Println("Original arry:", arr)

	bubbleSortDes(arr)
	fmt.Println("Sorted Array:", arr)
}

// ️ Time Complexity:
// Case	Time
// Best	O(n)
// Average	O(n²)
// Worst	O(n²)
// Space	O(1)
