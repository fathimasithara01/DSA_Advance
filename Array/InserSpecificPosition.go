package main

import "fmt"

// func main() {
// 	arr := []int{1, 3, 2, 4, 5}
// 	ind := 2
// 	// value := 6

// 	arr = append(arr[:2], append([]int{6}, arr[ind:]...)...)
// 	fmt.Println(arr)
// }

func main() {
	arr := make([]int, 6)

	arr[0] = 1
	arr[1] = 1
	arr[2] = 2
	arr[3] = 3
	arr[4] = 4
	arr[5] = 5

	value := 6
	ind := 2
	size := 5
	for i := size; i > ind; i-- {
		arr[i] = arr[i-1]
	}

	arr[ind] = value
	size++

	for i := 0; i < size; i++ {
		fmt.Print(arr[i], " ")
	}

	fmt.Println()
}
