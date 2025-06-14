package main

import "fmt"

// func ReverseString(Str string) string {
// 	reversed := ""
// 	for i := len(Str) - 1; i >= 0; i-- {
// 		reversed += string(Str[i])
// 	}
// 	return reversed
// }

func ReverseString(str string) string {
	runes := []rune(str)
	n := len(runes)

	for i := 0; i < n/2; i++ {
		runes[i], runes[n-1-i] = runes[n-1-i], runes[i]
	}
	return string(runes)
}
func main() {
	str := "sithara"
	res := ReverseString(str)
	fmt.Println("original string:", str)
	fmt.Println("after string:", res)
}
