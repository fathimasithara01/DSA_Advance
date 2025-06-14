package main

import "fmt"

func main() {
	str := "hello golang"

	fmt.Println(str)
	fmt.Println(len(str))
	fmt.Println(str[0])
	fmt.Println(string(str[0]))

	for i := 0; i < len(str); i++ {
		fmt.Printf("%c ", str[i])
	}
	fmt.Println()
}
