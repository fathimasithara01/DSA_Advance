package main

import "fmt"

type Node struct {
	value string
	Left  *Node
	Right *Node
}

func Preorder(root *Node) {
	if root == nil {
		return
	}

	fmt.Print(root.value, " ")
	Preorder(root.Left)
	Preorder(root.Right)
}

func Inorder(root *Node) {
	if root == nil {
		return
	}

	Inorder(root.Left)
	fmt.Print(root.value + " ")
	Inorder(root.Right)
}

func Postorder(root *Node) {
	if root == nil {
		return
	}

	Postorder(root.Left)
	Postorder(root.Right)
	fmt.Print(root.value, "")
}

func LevelOrder(root *Node) {
	if root == nil {
		return
	}

	queue := []*Node{root}

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		fmt.Print(node.value, " ")
		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
	}
}

func Height(root *Node) int {
	if root == nil {
		return 0
	}

	left := Height(root.Left)
	right := Height(root.Left)

	if left > right {
		return left + 1
	}

	return right + 1
}

func CountNodes(root *Node) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 1
	}
	return 1 + CountNodes(root.Left) + CountNodes(root.Right)
}

func CountLeave(root *Node) int {
	if root == nil {
		return 0
	}

	if root.Left == nil && root.Right == nil {
		return 1
	}

	return CountLeave(root.Left) + CountLeave(root.Right)
}

func Mirror(root *Node) *Node {
	if root == nil {
		return nil
	}

	root.Left, root.Right = Mirror(root.Right), Mirror(root.Left)
	return root
}

func BuildTree() *Node {
	root := &Node{value: "A"}
	root.Left = &Node{value: "B"}
	root.Right = &Node{value: "C"}
	root.Left.Left = &Node{value: "D"}
	root.Left.Right = &Node{value: "E"}
	root.Right.Right = &Node{value: "F"}

	return root
}

func main() {
	root := BuildTree()

	fmt.Println("DFS Traversals:")
	fmt.Print("Preorder: ")
	Preorder(root)
	fmt.Println()

	fmt.Print("Inorder:  ")
	Inorder(root)
	fmt.Println("")

	fmt.Print("Postorder: ")
	Postorder(root)
	fmt.Println("")

	fmt.Println("Height of the tree:", Height(root))
	fmt.Println("Total nodes", CountNodes(root))
	fmt.Println("leaf nodes", CountLeave(root))

	fmt.Println("Mirrored Tree Order")
	Mirror(root)
	Inorder(root)
	fmt.Println("")
}
