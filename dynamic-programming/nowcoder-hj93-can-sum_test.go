package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

/*
The problem is find an "x" where 2x = diff + n
diff = |group5sum - group3sum|
n is the sum of non-5 and non-3
*/
func TestCanGroup(t *testing.T) {
	ass := assert.New(t)

	//can, group1, group2 := CanGroup([]int{1, 5, -5, 1})
	//ass.True(can)
	//ass.EqualValues([]int{5, -5, 1}, group1)
	//ass.EqualValues([]int{1}, group2)
	//
	//can, _, _ = CanGroup([]int{3, 5, 8})
	//ass.False(can)
	//
	can, group1, group2 := CanGroup([]int{1, 0, 2, 3, -2})
	ass.True(can)
	fmt.Println(group1)
	fmt.Println(group2)

	//can, group1, group2 := CanGroup([]int{3, -4, -1, 2, -4, -5, 5})
	//ass.True(can)
	//fmt.Println(group1)
	//fmt.Println(group2)
}

func CanGroup(nums []int) (can bool, group5 []int, group3 []int) {

	//group5 := make([]int, len(nums))[:0]
	sum5 := 0
	//group3 := make([]int, len(nums))[:0]
	sum3 := 0
	others := make([]int, len(nums))[:0]
	sum := 0

	for _, num := range nums {
		switch {
		case num == 0:
			// do nothing
		case num%3 == 0:
			group3 = append(group3, num)
			sum3 += num
		case num%5 == 0:
			group5 = append(group5, num)
			sum5 += num
		default:
			others = append(others, num)
		}
		sum += num
	}

	if sum%2 == 0 {
		diff := sum/2 - sum5
		can, used := CanSum(diff, others)
		if can {
			if len(used) > 0 {
				for oi, ui := 0, 0; oi < len(others); oi++ {
					if ui < len(used) && others[oi] == used[ui] {
						group5 = append(group5, others[oi])
						ui++
					} else {
						group3 = append(group3, others[oi])
					}
				}
			} else {
				group3 = append(group3, others...)
			}
			return true, group5, group3
		}
	}
	return false, nil, nil

}

func CanSum(n int, nums []int) (can bool, used []int) {

	if n == 0 {
		return true, []int{}
	}

	var howSum [1001][]int // -500 ~ 0 ~ 500
	var canSum [1001]bool
	canSum[0+500] = true // numValue = index - 500
	for _, num := range nums {
		if num > 0 {
			for i := 1000; i >= 0; i-- {
				if canSum[i] {
					newi := i + num
					if 0 <= newi && newi <= 1000 { // ensure not out of bound
						if !canSum[newi] {
							canSum[newi] = true
							howSum[newi] = append(howSum[newi], howSum[i]...)
							howSum[newi] = append(howSum[newi], num)
						}
					}
					if newi-500 == n {
						return true, howSum[newi]
					}
				}
			}
		} else {
			for i := 0; i <= 1000; i++ {
				if canSum[i] {
					newi := i + num
					if 0 <= newi && newi <= 1000 { // ensure not out of bound
						if !canSum[newi] {
							canSum[newi] = true
							howSum[newi] = append(howSum[newi], howSum[i]...)
							howSum[newi] = append(howSum[newi], num)
						}
					}
					if newi-500 == n {
						return true, howSum[newi]
					}
				}
			}
		}
	}
	return false, nil
}
