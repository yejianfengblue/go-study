package string

import (
	"fmt"
	"sort"
	"strings"
	"testing"
	"unicode"
)

// sort.Strings(x []string)
func TestSortStrings(t *testing.T) {

	n := 3
	var array []string = []string{"magic", "hydra", "apple"}
	sort.Strings(array)
	for i := 0; i < n; i++ {
		fmt.Println(array[i])
	}
}

// sort.SliceStable(slice any, less func)
func TestRuneInsensitive(t *testing.T) {

	input := "AaPpPle"
	runeSlice := []rune(input)
	sort.SliceStable(runeSlice, func(i, j int) bool {
		return unicode.ToUpper(runeSlice[i]) < unicode.ToUpper(runeSlice[j])
	})
	fmt.Printf(string(runeSlice))
	fmt.Println()
}

func TestSortReverse(t *testing.T) {

	strSlice := []string{"banana", "apple", "cherry", "orange"}
	sort.Sort(
		sort.Reverse(
			sort.StringSlice(strSlice)))
	fmt.Println(strSlice)
}

/*
nowcoder.com
HJ26
描述
编写一个程序，将输入字符串中的字符按如下规则排序。

规则 1 ：英文字母从 A 到 Z 排列，不区分大小写。

如，输入： Type 输出： epTy

规则 2 ：同一个英文字母的大小写同时存在时，按照输入顺序排列。

如，输入： BabA 输出： aABb

规则 3 ：非英文字母的其它字符保持原来的位置。

如，输入： By?e 输出： Be?y

数据范围：输入的字符串长度满足 1 \le n \le 1000 \1≤n≤1000

A Famous Saying: Much Ado About Nothing (2012/8).
A aaAAbc dFgghh: iimM nNn oooos Sttuuuy (2012/8).
*/
func TestHJ26(t *testing.T) {

	input := "A Famous Saying: Much Ado About Nothing (2012/8)."

	// store all the letters
	var sb strings.Builder
	for _, r := range input {
		if unicode.IsLetter(r) {
			sb.WriteRune(r)
		}
	}
	// sort letters
	letters := []rune(sb.String())
	sort.SliceStable(letters, func(i, j int) bool {
		return unicode.ToUpper(letters[i]) < unicode.ToUpper(letters[j])
	})

	// reset the string builder for reuse
	sb.Reset()

	// re-iterate the original input string
	// when see letter, write sorted letter
	// when see non-letter, write it as is
	runeCount := 0
	for _, r := range input {
		if unicode.IsLetter(r) {
			sb.WriteRune(letters[runeCount])
			runeCount++
		} else {
			sb.WriteRune(r)
		}
	}

	fmt.Println(sb.String())
}
