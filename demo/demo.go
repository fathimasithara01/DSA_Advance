package main

import "fmt"

// func bubbleSort(arr []int) {
// 	for i := 0; i < len(arr)-1; i++ {
// 		for j := 0; j < len(arr)-1; j++ {
// 			if arr[j] > arr[j+1] {
// 				arr[j], arr[j+1] = arr[j+1], arr[j]
// 			}
// 		}
// 	}
// }

// func main() {
// 	arr := []int{3, 6, 2, 8, 1, 4, 6}

// 	bubbleSort(arr)
// 	fmt.Println("result array:", arr)
// }

// func insertionSort(arr []int) {
// 	for i := 1; i < len(arr); i++ {
// 		key := arr[i]
// 		j := i - 1

// 		for j >= 0 && arr[j] > key {
// 			arr[j+1] = arr[j]
// 			j--
// 		}
// 		arr[j+1] = key
// 	}
// }

// func main() {
// 	arr := []int{3, 2, 8, 1, 4, 6}

// 	insertionSort(arr)
// 	fmt.Println("result is:", arr)
// }

// func SelectionSort(arr []int) {
// 	for i := 0; i < len(arr)-1; i++ {
// 		minIndex := i
// 		for j := i + 1; j < len(arr); j++ {
// 			if arr[minIndex] > arr[j] {
// 				minIndex = j
// 			}
// 		}
// 		arr[i], arr[minIndex] = arr[minIndex], arr[i]
// 	}
// }

// func main() {
// 	arr := []int{3, 6, 2, 8, 1, 4, 6}
// 	SelectionSort(arr)
// 	fmt.Println("result is:", arr)
// }

// func ReverseString(str string) string {
// 	reversed := ""
// 	for i := len(str) - 1; i >= 0; i-- {
// 		reversed += string(str[i])
// 	}
// 	return reversed
// }

// func main() {
// 	str := "hello"

// 	res := ReverseString(str)
// 	fmt.Println(res)
// }

// func palindrome(str string) bool {
// 	left := str[0]
// 	right := str[len(str)-1]
// 	n := len(str) - 1

// 	for i := 0; i <= n/2; i++ {
// 		if left == right {
// 			return true
// 		}
// 	}
// 	return false
// }

// func main() {
// 	str := "malayalam"

//		if palindrome(str) {
//			fmt.Println("is palindrome")
//		} else {
//			fmt.Println("is not palindrome")
//		}
//	}
// func countVowelsAndConstant(str string) (int, int) {
// 	strr := strings.ToLower(str)
// 	vowels := []rune("aeiou")

// 	vc := 0
// 	cc := 0

// 	for _, ch := range strr {
// 		if unicode.IsLetter(ch) {
// 			if strings.ContainsRune(string(vowels), ch) {
// 				vc++
// 			} else {
// 				cc++
// 			}
// 		}
// 	}
// 	return vc, cc
// }

// func main() {
// 	str := "golang"

// 	vc, cc := countVowelsAndConstant(str)
// 	fmt.Printf("vowls count:%d, cons count:%d", vc, cc)
// 	fmt.Println("")
// }

// func FreqCounter(str string) {
// 	str = strings.ToLower(str)
// 	text := strings.Fields(str)

// 	freq := make(map[string]int)

// 	for _, count := range text {
// 		freq[count]++
// 	}

// 	for word, count := range freq {
// 		fmt.Printf("%s:%d", word, count)
// 		fmt.Println("")
// 	}
// }

// func main() {
// 	str := "go go golang go "

// 	FreqCounter(str)
// }

type Node struct {
	value int
	left  *Node
	right *Node
}

func Insert(root *Node, data int) *Node {
	if root == nil {
		return &Node{value: data}
	}
	if data < root.value {
		root.left = Insert(root.left, data)
	} else {
		root.right = Insert(root.right, data)
	}
	return root
}

func Search(root *Node, data int) bool {
	if root == nil {
		return false
	}
	if root.value == data {
		return true
	}
	if data < root.value {
		return Search(root.left, data)
	}
	return Search(root.right, data)
}

func FindMin(root *Node) *Node {
	if root == nil {
		return nil
	}
	for root.left != nil {
		root = root.left
	}
	return root
}

func FindMax(root *Node) *Node {
	if root == nil {
		return nil
	}
	for root.right != nil {
		root = root.right
	}
	return root
}

func Delete(root *Node, data int) *Node {
	if root == nil {
		return nil
	}
	if data < root.value {
		root.left = Delete(root.left, data)
	} else if data > root.value {
		root.right = Delete(root.right, data)
	} else {
		if root.left == nil {
			return root.right
		}
		if root.right == nil {
			return root.left
		}
		s := FindMin(root.right)
		root.value = s.value
		root.right = Delete(root.right, s.value)

	}
	return root
}

func Inorder(root *Node) {
	if root == nil {
		return
	}
	Inorder(root.left)
	fmt.Print(root.value, " ")
	Inorder(root.right)
}

func Preorderorder(root *Node) {

	fmt.Print(root.value, " ")
	Inorder(root.left)
	Inorder(root.right)
}
func PostOrder(root *Node) {

	Inorder(root.left)
	Inorder(root.right)
	fmt.Print(root.value, " ")
}
func main() {
	var root *Node
	values := []int{3, 76, 1, 4, 3, 2}

	for _, val := range values {
		root = Insert(root, val)
	}

	fmt.Println("traveersal order wise")
	Inorder(root)
	fmt.Println("")

	PostOrder(root)
	fmt.Println("")

	Preorderorder(root)
	fmt.Println("")

	fmt.Println("search 1", Search(root, 1))

	fmt.Println("min:", FindMin(root).value)
	fmt.Println("max:", FindMax(root).value)

}
