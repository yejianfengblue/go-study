package string

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"log"
	"strings"
	"testing"
	"unicode"
)

func TestMimicWildcard(t *testing.T) {

	ass := assert.New(t)
	ass.True(MimicWildcardTabulationMatch("", ""))
	ass.False(MimicWildcardTabulationMatch("pq", "pppq"))
	ass.True(MimicWildcardTabulationMatch("?*Bc*?", "abcd"))
	ass.False(MimicWildcardTabulationMatch("h*?*a", "h#a"))
	ass.False(MimicWildcardTabulationMatch("p*p*qp**pq*p**p***ppq", "pppppppqppqqppqppppqqqppqppqpqqqppqpqpppqpppqpqqqpqqp"))

	ass.True(MimicWildcardMemoizationMatch("", ""))
	ass.False(MimicWildcardMemoizationMatch("pq", "pppq"))
	ass.True(MimicWildcardMemoizationMatch("?*Bc*?", "abcd"))
	ass.False(MimicWildcardMemoizationMatch("h*?*a", "h#a"))
	ass.False(MimicWildcardMemoizationMatch("p*p*qp**pq*p**p***ppq", "pppppppqppqqppqppppqqqppqppqpqqqppqpqpppqpppqpqqqpqqp"))
}

/*
The complexity comes from the * matching.
The * matches 0, 1, >1 char

	A B * D
	A B C D

	A B *
	A
	not matched

	A B *
	A B
	matched, this * matches 0 char, t[3][2](AB* and AB) = t[2][2] (AB and AB) = t[3-1][2]

	A B *
	A B C
	matched, this * matches 1 char, t[3][3](AB* and ABC) = t[2][2] (AB and AB) = t[3-1][3-1]

	A B *
	A B C D
	matched, this * matches 2 char, t[3][4](AB* and ABCD) = t[3][3] (AB* and ABC) = t[3][4-1]
*/
func MimicWildcardTabulationMatch(expression string, text string) bool {

	// change multiple asterisk to one asterisk
	// change lowercase to uppercase
	var expSb strings.Builder
	starFound := false
	for _, exp := range expression {
		if exp == '*' {
			if !starFound {
				expSb.WriteRune(exp)
			}
			starFound = true
		} else {
			expSb.WriteRune(unicode.ToUpper(exp))
			starFound = false
		}
	}

	exps := []rune(expSb.String())
	expsLen := len(exps)
	texts := []rune(strings.ToUpper(text))
	textsLen := len(texts)
	log.Printf("exp = %c len = %d, texts = %c len = %d", exps, expsLen, texts, textsLen)

	// tabulation[0][0]=true means exps="" matches texts=""
	// tabulation[2][3]=true means exps[0:2] matches texts[0:3]
	// The index 0 room (tabulation[0][?] and tabulation[?][0]) is constructed,
	// to prevent the index 0 check in below switch statement
	tabulation := make([][]bool, expsLen+1)
	for i := 0; i <= expsLen; i++ {
		tabulation[i] = make([]bool, textsLen+1)
	}
	// exps "" matches texts ""
	tabulation[0][0] = true
	// this for loop in unnecessary, because the value is zero-value by default, write it here for better understanding
	for j := 1; j <= textsLen; j++ {
		// exps "" not matches non-empty texts
		tabulation[0][j] = false
	}

	for ei, exp := range exps {
		for ti, text := range texts {
			switch exp {
			case '*':
				tabulation[ei+1][0] = tabulation[ei][0]
				if IsAlphaDigit(text) {
					tabulation[ei+1][ti+1] =
						tabulation[ei][ti+1] || // this * match 0 char
							tabulation[ei][ti] || // this * match 1 char
							tabulation[ei+1][ti] // this * match >1 char
				} else {
					tabulation[ei+1][ti+1] = false
				}
			case '?':
				tabulation[ei+1][0] = false // exp not empty but text empty
				tabulation[ei+1][ti+1] = IsAlphaDigit(text) && tabulation[ei][ti]
			default:
				tabulation[ei+1][0] = false // exp not empty but text empty
				tabulation[ei+1][ti+1] = exp == text && tabulation[ei][ti]
			}
		}
	}

	log.Println(Print2DBoolArray(tabulation))

	return tabulation[expsLen][textsLen]
}

func MimicWildcardMemoizationMatch(expression string, text string) bool {

	// change multiple asterisk to one asterisk
	// change lowercase to uppercase
	var expSb strings.Builder
	starFound := false
	for _, exp := range expression {
		if exp == '*' {
			if !starFound {
				expSb.WriteRune(exp)
			}
			starFound = true
		} else {
			expSb.WriteRune(unicode.ToUpper(exp))
			starFound = false
		}
	}

	exps := []rune(expSb.String())
	expsLen := len(exps)
	texts := []rune(strings.ToUpper(text))
	textsLen := len(texts)
	log.Printf("exp = %c len = %d, texts = %c len = %d", exps, expsLen, texts, textsLen)

	memoization := make([][]int, expsLen)
	for i := 0; i < expsLen; i++ {
		memoization[i] = make([]int, textsLen)
	}

	return MimicWildcardMemoizationMatchRecursively(memoization, exps, expsLen-1, texts, textsLen-1)
}

func MimicWildcardMemoizationMatchRecursively(memoization [][]int, exps []rune, ei int, texts []rune, ti int) bool {

	memo := memoization[ei][ti] // cache the function call return value
	if memo != 0 {
		return memo > 0
	}

	if ei < 0 { // empty expression
		if ti < 0 { // "" and "", matched
			return true
		} else { // "" and non-empty text, mismatched
			return false
		}
	} else { // non-empty expression
		if exps[ei] != '*' && ti < 0 { // non-empty expression and it is not *, and empty text, mismatched
			return false
		}
	}

	exp := exps[ei]
	text := texts[ti]
	matched := false
	switch exp {
	case '*':
		if IsAlphaDigit(text) {
			matched = MimicWildcardMemoizationMatchRecursively(memoization, exps, ei-1, texts, ti) || // match 0 char
				MimicWildcardMemoizationMatchRecursively(memoization, exps, ei-1, texts, ti-1) || // match 1 char
				MimicWildcardMemoizationMatchRecursively(memoization, exps, ei, texts, ti-1) // match 2 char
		} else {
			matched = false
		}
	case '?':
		if IsAlphaDigit(text) {
			matched = MimicWildcardMemoizationMatchRecursively(memoization, exps, ei-1, texts, ti-1)
		} else {
			matched = false
		}
	default:
		if exp == text {
			matched = MimicWildcardMemoizationMatchRecursively(memoization, exps, ei-1, texts, ti-1)
		} else {
			matched = false
		}
	}

	if matched {
		memoization[ei][ti] = 1
	} else {
		memoization[ei][ti] = -1
	}

	return matched
}

func IsAlphaDigit(r rune) bool {
	return ('0' <= r && r <= '9') || ('A' <= r && r <= 'Z')
}

func Print2DBoolArray(array [][]bool) string {

	var sb strings.Builder
	sb.WriteString("\n   ")
	for i := 0; i < len(array[0]); i++ {
		sb.WriteString(fmt.Sprintf("%2d ", i))
	}
	sb.WriteString("\n")
	for i, tabi := range array {
		sb.WriteString(fmt.Sprintf("%2d%v \n", i, BoolArrayToString(tabi)))
	}

	return sb.String()
}

func BoolToRune(b bool) rune {
	if b {
		return '1'
	} else {
		return '0'
	}
}

func BoolArrayToString(bools []bool) string {
	var sb strings.Builder
	for i := 0; i < len(bools); i++ {
		sb.WriteRune(' ')
		sb.WriteRune(' ')
		sb.WriteRune(BoolToRune(bools[i]))
	}
	return sb.String()
}
