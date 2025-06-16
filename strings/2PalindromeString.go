package main

import "fmt"

// func isPalindrome(str string) bool {
// 	// runes := []rune(str)
// 	n := len(str)

// 	for i := 0; i < n/2; i++ {
// 		if str[i] != str[n-1-i] {
// 			return false
// 		}
// 	}
// 	return true
// }

func isPalindrome(str string) bool {
	left := 0
	right := len(str) - 1

	for i := 0; i < right/2; i++ {
		if str[left] == str[right] && str[right] == str[left] {
			return true
		}
	}
	return false
}
func main() {
	var input string

	fmt.Println("Enter a string:")
	fmt.Scan(&input)

	if isPalindrome(input) {
		fmt.Println("is palindrome:", input)
	} else {
		fmt.Println("is not palindrome")
	}

}
