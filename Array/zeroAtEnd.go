package main

import "fmt"

func main() {
	arr := []int{3, 5, 0, 0, 6, 2, 1}

	pos := 0
	for i := 0; i < len(arr); i++ {
		if arr[i] != 0 {
			arr[pos] = arr[i]
			pos++
		}
	}
	for i := pos; i < len(arr); i++ {
		arr[i] = 0
	}

	fmt.Println(arr)
}
