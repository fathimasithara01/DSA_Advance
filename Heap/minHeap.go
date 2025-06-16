package main

import "fmt"

type minHeap struct {
	data []int
}

func (h *minHeap) Insert(value int) {
	h.data = append(h.data, value)
	h.heapifyUp(len(h.data) - 1)
}

func (h *minHeap) Peek() int {
	if len(h.data) == 0 {
		panic("Heap is empty")
	}
	return h.data[0]
}

// remove and return the smallest element)
func (h *minHeap) ExtractMin() int {
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

func BuildHeap(arr []int) *minHeap {
	h := &minHeap{data: arr}
	for i := len(h.data)/2 - 1; i >= 0; i-- {
		h.heapifyDown(i)
	}
	return h
}

// heapifyUp Used after inserting a new element at the end of the heap.
func (h *minHeap) heapifyUp(i int) {
	for i > 0 {
		parent := (i - 1) / 2
		if h.data[parent] <= h.data[i] {
			break
		}

		h.data[i], h.data[parent] = h.data[parent], h.data[i]
		i = parent
	}
}

// Used after removing the root
func (h *minHeap) heapifyDown(i int) {
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
	h := &minHeap{}

	h.Insert(10)
	h.Insert(5)
	h.Insert(3)
	h.Insert(2)
	h.Insert(8)

	fmt.Println("Min:", h.Peek())
	fmt.Println("Eexracted:", h.ExtractMin())
	fmt.Println("New min:", h.Peek())

	arr := []int{7, 9, 4, 1, 6}
	newHeap := BuildHeap(arr)
	fmt.Println("Min from built heap:", newHeap.Peek())
}
