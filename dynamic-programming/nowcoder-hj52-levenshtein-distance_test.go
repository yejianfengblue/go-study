package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLevenshteinDistance(t *testing.T) {

	ass := assert.New(t)
	ass.EqualValues(1, LevenshteinDistance("abcdefg", "abcdef"))
}

func LevenshteinDistance(s1, s2 string) int {

	rs1 := []rune(s1)
	rs2 := []rune(s2)
	len1 := len(rs1)
	len2 := len(rs2)
	tab := make([][]int, len1+1)
	for i := 0; i <= len1; i++ {
		tab[i] = make([]int, len2+1)
	}
	tab[0][0] = 0
	for i1 := 1; i1 <= len1; i1++ {
		tab[i1][0] = tab[i1-1][0] + 1
	}
	for i2 := 1; i2 <= len2; i2++ {
		tab[0][i2] = tab[0][i2-1] + 1
	}
	for i1, r1 := range rs1 {
		for i2, r2 := range rs2 {
			if r1 == r2 {
				tab[i1+1][i2+1] = tab[i1][i2]
			} else {
				tab[i1+1][i2+1] = minOf3(tab[i1][i2], tab[i1][i2+1], tab[i1+1][i2]) + 1
			}
		}
	}
	return tab[len1][len2]
}

func minOf3(a, b, c int) int {
	minab := a
	if b < a {
		minab = b
	}
	minabc := minab
	if c < minab {
		minabc = c
	}
	return minabc
}
