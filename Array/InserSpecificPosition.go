package main

// func main() {
// 	arr := []int{1, 3, 2, 4, 5}
// 	ind := 2
// 	// value := 6

// 	arr = append(arr[:2], append([]int{6}, arr[ind:]...)...)
// 	fmt.Println(arr)
// }

func main() {
	arr := []int{1, 2, 3, 4, 5}

	ind := 2
	for i := len(arr); i > ind; i-- {
		arr[i]
	}
}
