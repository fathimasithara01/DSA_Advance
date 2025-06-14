package main

import (
	"fmt"
	"strings"
	"unicode"
)

func countVowelsAndConstant(str string) (int, int) {
	s := strings.ToLower(str)
	vowels := []rune("aeiou")

	vowelsCount := 0
	constantsCount := 0

	for _, ch := range s {
		if unicode.IsLetter(ch) {
			if strings.ContainsRune(string(vowels), ch) {
				vowelsCount++
			} else {
				constantsCount++
			}
		}
	}
	return vowelsCount, constantsCount
}

func main() {
	var input string

	fmt.Println("Enter a string:")
	fmt.Scan(&input)

	v, c := countVowelsAndConstant(input)
	fmt.Printf("vowels count:%d, constant count:%d", v, c)
	fmt.Println("")
}
