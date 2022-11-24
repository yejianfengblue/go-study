package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestName(t *testing.T) {
	ass := assert.New(t)
	//ass.EqualValues(3, LIS([]int{2, 5, 1, 5, 4, 5}))
	ass.EqualValues(4, LIS([]int{268, 90, 179, 129, 204, 224}))
}

func LIS(nums []int) int {
	lis := make([]int, len(nums))
	max := 0
	for i := range nums {
		lis[i] = 1 // the default value, coz at least contains itself
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] && lis[i] < lis[j]+1 {
				lis[i] = lis[j] + 1
			}
		}
		if lis[i] > max {
			max = lis[i]
		}
	}
	return max
}
