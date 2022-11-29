package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var s [9][9]int // sudoku board

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	for r := 0; r < 9; r++ {
		scanner.Scan()
		line := scanner.Text()
		for c := 0; c < 9; c++ {
			s[r][c], _ = strconv.Atoi(line[c*2 : c*2+1])
		}
	}

	if try(0, 0) {
		for r := 0; r < 9; r++ {
			for c := 0; c < 9; c++ {
				fmt.Print(s[r][c], " ")
			}
			fmt.Println()
		}
	} else {
		fmt.Println("Not resolved")
	}
}

func try(r, c int) bool {
	if r == 9 || c == 9 {
		return true
	}
	nextr, nextc := r+(c+1)/9, (c+1)%9
	if s[r][c] == 0 {
		for num := 1; num <= 9; num++ {
			if IsThisNumNotExistInRow(r, num) && IsThisNumNotExistInColumn(c, num) && IsThisNumNotExistIn3x3(r, c, num) {
				s[r][c] = num
				if try(nextr, nextc) {
					return true
				} else {
					continue
				}
			}
		}
		s[r][c] = 0
		return false
	} else {
		return try(nextr, nextc)
	}
}

func IsThisNumNotExistInRow(r int, num int) bool {
	for c := 0; c < 9; c++ {
		if s[r][c] == num {
			return false
		}
	}
	return true
}

func IsThisNumNotExistInColumn(c int, num int) bool {
	for r := 0; r < 9; r++ {
		if s[r][c] == num {
			return false
		}
	}
	return true
}

func IsThisNumNotExistIn3x3(r, c int, num int) bool {
	for ri := 0; ri < 3; ri++ {
		for ci := 0; ci < 3; ci++ {
			if s[r/3*3+ri][c/3*3+ci] == num {
				return false
			}
		}
	}
	return true
}
