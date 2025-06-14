package main

import "fmt"

func swapAlternate(arr []int) []int {
    for i := 0; i < len(arr)-1; i += 2 {
        arr[i], arr[i+1] = arr[i+1], arr[i]
    }
    return arr
}
func main() {
	arr := []int{2, 5, 21, 3, 2, 3, 34, 2}

	res := swapAlternate(arr)
	fmt.Println(res)
}
