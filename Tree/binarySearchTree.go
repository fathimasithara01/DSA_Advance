package main

import "fmt"

// Traversal	Time	Space (avg)	Space (worst)
// Inorder	O(n)	O(log n)	O(n)
// Preorder	O(n)	O(log n)	O(n)
// Postorder	O(n)	O(log n)	O(n)

type Nod struct {
	val   int
	Left  *Nod
	Right *Nod
}

func Insert(root *Nod, data int) *Nod {
	if root == nil {
		return &Nod{val: data}
	}
	if data < root.val {
		root.Left = Insert(root.Left, data)
	} else {
		root.Right = Insert(root.Right, data)
	}
	return root
}

func Search(root *Nod, data int) bool {
	if root == nil {
		return false
	}
	if root.val == data {
		return true
	}
	if data < root.val {
		return Search(root.Left, data)
	}
	return Search(root.Right, data)
}

func FindMin(root *Nod) *Nod {
	if root == nil {
		return nil
	}
	for root.Left != nil {
		root = root.Left
	}
	return root
}

func FindMax(root *Nod) *Nod {
	if root == nil {
		return nil
	}
	for root.Right != nil {
		root = root.Right
	}
	return root
}

func Delete(root *Nod, data int) *Nod {
	if root == nil {
		return nil
	}
	if data < root.val {
		root.Left = Delete(root.Left, data)
	} else if data > root.val {
		root.Right = Delete(root.Right, data)
	} else {
		if root.Left == nil {
			return root.Right
		}
		if root.Right == nil {
			return root.Left
		}

		successor := FindMin(root.Right)
		root.val = successor.val
		root.Right = Delete(root.Right, successor.val)
	}
	return root
}

func Inorderr(root *Nod) {
	if root == nil {
		return
	}
	Inorderr(root.Left)
	fmt.Print(root.val, " ")
	Inorderr(root.Right)
}

func Preorderr(root *Nod) {
	if root == nil {
		return
	}
	fmt.Print(root.val, " ")
	Preorderr(root.Left)
	Preorderr(root.Right)
}

func Postorderr(root *Node) {
	if root == nil {
		return
	}

	Postorderr(root.Left)
	Postorderr(root.Right)
	fmt.Print(root.value, " ")
}

func main() {
	var root *Nod

	values := []int{3, 5, 2, 1, 23, 54, 5, 2, 1, 7, 3}
	for _, v := range values {
		root = Insert(root, v)
	}

	fmt.Println("Inorder Traversl (sorted):")
	Inorderr(root)
	fmt.Println("")

	fmt.Println("Search 6", Search(root, 6))
	fmt.Println("Search 15", Search(root, 15))

	fmt.Println("min", FindMin(root).val)
	fmt.Println("max", FindMax(root).val)
}
