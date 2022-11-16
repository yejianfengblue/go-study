package string

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRun(t *testing.T) {
	ass := assert.New(t)
	ass.EqualValues(4, LPS("12HHHHA"))
	ass.EqualValues(5, LPS("ABBBA"))
}

func LPS(str string) int {
	var runes []rune = []rune(str)
	max := 0
	if len(runes) < 2 {
		return len(runes)
	} else {
		max = 1
		for i := range runes {
			leftSpace := i - 0
			rightSpace := len(runes) - 1 - i

			// odd length
			tryMoveLength := leftSpace
			if leftSpace > rightSpace {
				tryMoveLength = rightSpace
			}
			for j := 1; j <= tryMoveLength; j++ {
				if runes[i-j] == runes[i+j] {
					fmt.Printf("%c is symmetric\n", runes[i-j:i+j+1])
					if j*2+1 > max {
						max = j*2 + 1
					}
				} else {
					break // not symmetric anymore
				}
			}

			// even length
			if i <= len(runes)-2 && runes[i] == runes[i+1] {
				if max < 2 {
					max = 2
				}
				rightSpace = rightSpace - 1
				tryMoveLength = leftSpace
				if leftSpace > rightSpace {
					tryMoveLength = rightSpace
				}
				for k := 1; k <= tryMoveLength; k++ {
					if runes[i-k] == runes[i+1+k] {
						fmt.Printf("%c is symmetric\n", runes[i-k:i+1+k+1])
						if (k+1)*2 > max {
							max = (k + 1) * 2
						}
					} else {
						break // not symmetric anymore
					}
				}
			}
		}
	}
	return max
}
