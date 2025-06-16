package main

import "fmt"

func main() {
	set := make(map[int]bool)

	set[10] = true
	set[20] = true
	set[30] = true

	set[20] = true

	for key := range set {
		fmt.Println(key)
	}

	if set[20] {
		fmt.Println("set 20 in the set")
	} else {
		fmt.Println("not found")
	}

	delete(set, 30)

	for key := range set {
		fmt.Println(key)
	}
}
