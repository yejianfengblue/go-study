package string

import (
	"fmt"
	"strings"
	"testing"
)

// nowcoder.com HJ4
func TestStringPadding(t *testing.T) {
	line := "0123456789"
	length := len(line)
	for i := 0; i < length; i += 8 {
		paddingCount := i + 8 - length
		if paddingCount <= 0 {
			fmt.Println(line[i : i+8])
		} else {
			fmt.Println(line[i:length] + strings.Repeat("0", paddingCount))
		}
	}
}

func TestStringPadding2(t *testing.T) {

	line := "01234567"
	originalLength := len(line)
	paddingCount := (8 - originalLength%8) % 8
	fmt.Println(paddingCount)
	newline := line + strings.Repeat("0", paddingCount)
	fmt.Println(newline)
	newLength := len(newline)
	for i := 0; i < newLength; i += 8 {
		fmt.Println(newline[i : i+8])
	}
}
