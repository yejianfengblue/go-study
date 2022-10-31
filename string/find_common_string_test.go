package string

import (
	"fmt"
	"testing"
)

// nowcoder.com HJ65
func TestFindCommonString(t *testing.T) {

	//str1 := "abcdefghijklmnop"
	str1 := "bcsxxxopqxx"
	str2 := "abcsafjklmnopqrstuvw"

	rs1 := []rune(str1)
	rs2 := []rune(str2)
	if len(rs1) > len(rs2) {
		rs1, rs2 = rs2, rs1
	}

	commonLength := 0
	var commonSlice []rune
	for rs1i := len(rs1) - 1; rs1i >= 0; rs1i-- {
		fmt.Printf("rs1i = %d, rs1[i] = %c\n", rs1i, rs1[rs1i])
		for rs2i := len(rs2) - 1; rs2i >= 0; rs2i-- {
			fmt.Printf("rs2i = %d, rs2[i] = %c\n", rs2i, rs2[rs2i])
			if rs1[rs1i] == rs2[rs2i] {
				fmt.Printf("i matched: rs1[rs1i] = %c, rs2[rs2i] = %c\n", rs1[rs1i], rs2[rs2i])
				rs1j := rs1i - 1
				rs2j := rs2i - 1

				for ; rs1j >= 0 && rs2j >= 0; rs1j, rs2j = rs1j-1, rs2j-1 {
					if rs1[rs1j] != rs2[rs2j] {
						// +1 to find the last matched index
						rs1j++
						rs2j++
						break
					}
				}
				// if index goes before 0, treat it as mismatch, so +1 to find the last matched index
				if rs1j < 0 || rs2j < 0 {
					rs1j++
					rs2j++
				}
				fmt.Printf("j mismatched: rs1i = %d, rs1j = %d, rs2i = %d, rs2j = %d, length = %d, slice = %c\n",
					rs1i, rs1j, rs2i, rs2j, rs1i-rs1j, rs1[rs1j:rs1i+1])
				if commonLength <= rs1i-rs1j+1 {
					commonLength = rs1i - rs1j + 1
					commonSlice = rs1[rs1j : rs1i+1]
					fmt.Printf("commonLength = %d\n", commonLength)
					fmt.Printf("commonSlice = %c\n", commonSlice)
				}
			}
		}
	}
	fmt.Println("========================")
	fmt.Printf("commonLength = %d\n", commonLength)
	fmt.Printf("commonSlice = %c\n", commonSlice)

}
