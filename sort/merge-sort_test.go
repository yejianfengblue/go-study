package sort

import (
	"fmt"
	"testing"
)

type StuScore struct {
	name  string
	score int
}

func (s StuScore) LessEquals(o StuScore) bool {
	return s.score <= o.score
}

func TestMergeSort(t *testing.T) {

	var n int
	fmt.Scan(&n)
	stus := make([]StuScore, n)
	var isIncreasing int
	fmt.Scan(&isIncreasing)

	var name string
	var score int
	_, err := fmt.Scan(&name, &score)
	for err != nil {
		stus = append(stus, StuScore{name, score})
		_, err = fmt.Scan(&name, &score)
	}
	for _, stu := range stus {
		fmt.Println(stu)
	}
	sort(stus, 0, n)
	for _, stu := range stus {
		fmt.Println(stu)
	}
}

func merge(array []StuScore, startIdx, midIdx, endIdx int) {

	leftSize := midIdx - startIdx + 1 // include mid
	rightSize := endIdx - midIdx

	left := make([]StuScore, leftSize)
	right := make([]StuScore, rightSize)

	// copy data to temp arrays
	copy(left, array[startIdx:midIdx+1])
	copy(right, array[midIdx+1:endIdx])

	li, ri := 0, 0
	ai := startIdx
	for li < leftSize && ri < rightSize {
		if left[li].LessEquals(right[ri]) {
			array[ai] = left[li]
			li++
		} else {
			array[ai] = right[ri]
			ri++
		}
		ai++
	}
	// copy the remaining elems of left
	for li < leftSize {
		array[ai] = left[li]
		li++
		ai++
	}
	for ri < rightSize {
		array[ai] = right[ri]
		ri++
		ai++
	}
}

func sort(array []StuScore, startIdx, endIdx int) {
	if startIdx < endIdx {
		midIdx := startIdx + (endIdx-startIdx)/2
		sort(array, startIdx, midIdx)
		sort(array, midIdx+1, endIdx)
		merge(array, startIdx, midIdx, endIdx)
	}
}
