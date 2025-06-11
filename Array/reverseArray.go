package main

import "fmt"

func reverseArray(arr []int) []int {
	left := 0
	right := len(arr) - 1

	for left < right {
		arr[left], arr[right] = arr[right], arr[left]
		left++
		right--
	}
	return arr
}

func main() {
	arr := []int{32, 4, 2, 1, 2}

	res := reverseArray(arr)
	fmt.Println(res)
}
