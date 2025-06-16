package main

import "fmt"

type maxHeap struct {
	arr []int
}

func (h *maxHeap) Insert(value int) {
	h.arr = append(h.arr, value)
	i := len(h.arr) - 1
	for i > 0 {
		parent := (i - 1) / 2
		if h.arr[parent] >= h.arr[i] {
			break
		}

		h.arr[i], h.arr[parent] = h.arr[parent], h.arr[i]
		i = parent
	}
}

func (h *maxHeap) Peek() int {
	if len(h.arr) == 0 {
		panic("heap is empty")
	}
	return h.arr[0]
}

func (h *maxHeap) ExtractMax() int {
	if len(h.arr) == 0 {
		panic("heap is empty")
	}
	max := h.arr[0]
	last := h.arr[len(h.arr)-1]
	h.arr = h.arr[:len(h.arr)-1]
	if len(h.arr) > 0 {
		h.arr[0] = last
		h.heapifyDown(0)
	}
	return max
}

func (h *maxHeap) heapifyDown(i int) {
	n := len(h.arr)
	for {
		left := 2*i + 1
		right := 2*i + 2
		largest := i

		if left < n && h.arr[left] > h.arr[largest] {
			largest = left
		}
		if right < n && h.arr[right] > h.arr[largest] {
			largest = right
		}
		if largest == i {
			break
		}
		h.arr[i], h.arr[largest] = h.arr[largest], h.arr[i]
		i = largest
	}
}

func main() {
	h := &maxHeap{}
	h.Insert(5)
	h.Insert(10)
	h.Insert(3)
	h.Insert(15)

	fmt.Println("Max:", h.Peek())               // 15
	fmt.Println("Extract Max:", h.ExtractMax()) // 15
	fmt.Println("Now Max:", h.Peek())           // 10
}
