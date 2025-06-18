package main

import (
	"fmt"
	"strings"
)

func main() {
	text := "go is simple and go is powerful and GO is fast"

	text = strings.ToLower(text)
	words := strings.Fields(text)

	frequency := make(map[string]int)

	for _, word := range words {
		frequency[word]++
	}

	fmt.Println("word frequency:")
	for word, count := range frequency {
		fmt.Printf("%s:%d\n", word, count)
	}

}
