package main

import "fmt"

type MinHeap struct {
	data []int
}

// Insert an element into the heap
//
//	func (h *MinHeap) Insert(val int) {
//		h.data = append(h.data, val)
//		i := len(h.data) - 1
//		for i > 0 {
//			parent := (i - 1) / 2
//			if h.data[parent] <= h.data[i] {
//				break
//			}
//			h.data[i], h.data[parent] = h.data[parent], h.data[i]
//			i = parent
//		}
//	}
func (h *MinHeap) Insert(val int) {
	h.data = append(h.data, val)
	h.heapifyUp(len(h.data) - 1)
}

// Peek returns the minimum element
func (h *MinHeap) Peek() int {
	if len(h.data) == 0 {
		panic("heap is empty")
	}
	return h.data[0]
}

// ExtractMin removes and returns the minimum element
func (h *MinHeap) ExtractMin() int {
	if len(h.data) == 0 {
		panic("heap is empty")
	}
	min := h.data[0]
	last := h.data[len(h.data)-1]
	h.data = h.data[:len(h.data)-1]
	if len(h.data) > 0 {
		h.data[0] = last
		h.heapifyDown(0)
	}
	return min
}

func (h *MinHeap) Delete(index int) {
	if index < 0 || index >= len(h.data) {
		panic("index out of range")
	}
	h.data[index] = h.data[len(h.data)-1]
	h.data = h.data[:len(h.data)-1]
	h.heapifyDown(index)
	h.heapifyUp(index)
}

func (h *MinHeap) heapifyUp(index int) {
	for index > 0 {
		parent := (index - 1) / 2
		if h.data[parent] <= h.data[index] {
			break
		}
		h.data[parent], h.data[index] = h.data[index], h.data[parent]
		index = parent
	}
}

// heapifyDown restores the heap property from top → down
func (h *MinHeap) heapifyDown(i int) {
	n := len(h.data)
	for {
		left := 2*i + 1
		right := 2*i + 2
		smallest := i

		if left < n && h.data[left] < h.data[smallest] {
			smallest = left
		}
		if right < n && h.data[right] < h.data[smallest] {
			smallest = right
		}
		if smallest == i {
			break
		}
		h.data[i], h.data[smallest] = h.data[smallest], h.data[i]
		i = smallest
	}
}
func main() {
	h := &MinHeap{}

	// Insert values
	h.Insert(20)
	h.Insert(5)
	h.Insert(15)
	h.Insert(2)

	fmt.Println("Min:", h.Peek()) // 2

	// Extract Min
	fmt.Println("Extracted Min:", h.ExtractMin()) // 2

	// Current Min
	fmt.Println("Min now:", h.Peek()) // 5

	// Delete index 1 (value 20)
	h.Delete(1)

	// Min after deletion
	fmt.Println("Min after delete:", h.Peek()) // 5
}
