package main

import "fmt"

func QuickSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	pivot := arr[len(arr)-1]
	var left []int
	var right []int

	for _, v := range arr[:len(arr)-1] {
		if v < pivot {
			left = append(left, v)
		} else {
			right = append(right, v)
		}
	}

	return append(append(QuickSort(left), pivot), QuickSort(right)...)
}

func main() {
	arr := []int{4, 2, 3, 7, 1, 3}
	fmt.Println("Originall:", arr)

	sorted := QuickSort(arr)
	fmt.Println("Sorted: ", sorted)
}
