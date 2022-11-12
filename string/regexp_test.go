package string

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"testing"
)

// nowcoder.com hj17 coordinate move
func TestRegexp(t *testing.T) {

	var input string = "A10;S20;W10;D30;X;A1A;B10A11;;A10;"

	var coordinateRegexp *regexp.Regexp = regexp.MustCompile(`^[ASDW][0-9]{1,2}$`)
	splits := strings.Split(input, ";")
	var x, y int64 = 0, 0
	for i, s := range splits {
		if coordinateRegexp.MatchString(s) {
			distance, _ := strconv.ParseInt(s[1:], 10, 0)
			switch s[0] {
			case 'A':
				x -= distance
			case 'D':
				x += distance
			case 'S':
				y -= distance
			case 'W':
				y += distance
			}
			fmt.Println(i, s, s[0], s[1:], x, y)
		}
	}
	fmt.Printf("%d,%d\n", x, y)

}
