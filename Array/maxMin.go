package main

import "fmt"

func main() {
	arr := []int{3, 45, 2, 4, 1}

	max := arr[0]
	min := arr[0]

	for _, value := range arr {
		if max < value {
			max = value
		}
		if min > value {
			min = value
		}
	}

	fmt.Println(max)
	fmt.Println(min)
}
