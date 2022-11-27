package data_structure

import (
	"fmt"
	"testing"
)

type MaxHeap struct {
	array []int
}

func (h *MaxHeap) Len() int {
	return len(h.array)
}

// Push the value to the heap and maintain the heap
func (h *MaxHeap) Push(value int) {
	h.array = append(h.array, value)
	h.HeapifyUp(len(h.array) - 1)
}

func (h *MaxHeap) Pop() int {
	returnValue := h.array[0]

	length := len(h.array) - 1

	// put the last value to the root
	h.array[0] = h.array[length]
	h.array = h.array[:length]

	return returnValue
}

func (h *MaxHeap) HeapifyUp(idx int) {
	// the edge case is, when idx = 0, parentIndex(0) = (0-1)/2 = 0
	// so impossible to be out of bound
	for h.array[parentIndex(idx)] < h.array[idx] {
		h.Swap(parentIndex(idx), idx)
		idx = parentIndex(idx)
	}
}

func (h *MaxHeap) HeapifyDown(idx int) {

	endIndex := len(h.array) - 1
	left, right := leftChildIndex(idx), rightChildIndex(idx)
	comparedChild := 0
	for left <= endIndex {
		if left == endIndex { // when no right child
			comparedChild = left
		} else if h.array[left] > h.array[right] { // when left child value is larger
			comparedChild = left
		} else { // when right child value is larger
			comparedChild = right
		}
		if h.array[idx] < h.array[comparedChild] {
			h.Swap(idx, comparedChild)
			idx = comparedChild
			left, right = leftChildIndex(idx), rightChildIndex(idx)
		} else {
			break
		}
	}
}

func parentIndex(index int) int {
	return (index - 1) / 2
}

func leftChildIndex(index int) int {
	return index*2 + 1
}

func rightChildIndex(index int) int {
	return index*2 + 2
}

func (h *MaxHeap) Swap(idx1, idx2 int) {
	h.array[idx1], h.array[idx2] = h.array[idx2], h.array[idx1]
}

func TestInsert(t *testing.T) {
	var mh MaxHeap
	fmt.Println(mh)

	for _, v := range []int{10, 20, 30} {
		mh.Push(v)
		fmt.Println(mh)
	}
	for mh.Len() > 0 {
		v := mh.Pop()
		fmt.Println(v, mh)
	}
}
