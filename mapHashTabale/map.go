package main

import "fmt"

func main() {
	fruits := make(map[string]int)

	fruits["apple"] = 5
	fruits["banana"] = 10
	fruits["mango"] = 21

	fmt.Println(fruits, " ", fruits["mango"])

	value, exists := fruits["grapes"]
	if exists {
		fmt.Println("Grapes", value)
	} else {
		fmt.Println("not found grapes")
	}

	fruits["mangp"] = 23

	delete(fruits, "apple")

	for fruit, count := range fruits {
		fmt.Printf("fruit:%s and count:%d", fruit, count)
		fmt.Println("")
	}
}
