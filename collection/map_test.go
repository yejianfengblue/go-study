package collection

import (
	"fmt"
	"testing"
)

// nowcoder.com HJ59
// counter the first letter which occurs only once
func TestList(t *testing.T) {

	input := "qfqqnof"
	// rune : counter
	var counterMap map[rune]int = make(map[rune]int)
	// store the order of each rune in the original string
	var order []rune = make([]rune, len(input))[:0]

	for _, r := range input {
		// if this is the first time the rune occurs, record it in the order
		if counterMap[r] == 0 {
			order = append(order, r)
		}
		// counter++
		counterMap[r]++
	}

	var result rune = -1
	for _, r := range order {
		if counterMap[r] == 1 {
			result = r
			break
		}
	}
	if result > 0 {
		fmt.Printf("%c\n", result)
	} else {
		fmt.Println(-1)
	}

}
