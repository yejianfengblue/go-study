package main

import (
	"fmt"
)


func main() {

	var a, b, c int

	// when send '1 2 3' to stdio,
	// when send '1 2<enter>3<enter>' to stdio,
	// when send '1 2 3 4<enter>' to stdio,
	// then in all case a = 1, b = 2, c = 3, n = 3, err = nil
	n, err := fmt.Scan(&a, &b, &c)
	fmt.Printf("a = %v\n", a)
	fmt.Printf("b = %v\n", b)
	fmt.Printf("c = %v\n", c)
	fmt.Printf("n = %v\n", n)
	fmt.Printf("err = %#v\n", err)
}
