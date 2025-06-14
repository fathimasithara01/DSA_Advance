package main

import "fmt"

func arrFront(arr []int) []int {
	pos := len(arr) - 1

	for i := pos; i >= 0; i-- {
		if arr[i] != 0 {
			arr[pos] = arr[i]
			pos--
		}
	}
	for i := 0; i <= pos; i++ {
		arr[i] = 0
	}
	return arr
}

func main() {
	arr := []int{3, 52, 0, 5, 0, 2, 5}

	res := arrFront(arr)
	fmt.Println(res)
}
