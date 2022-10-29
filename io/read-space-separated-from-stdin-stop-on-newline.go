package main

import (
	"fmt"
)


func main() {

	var a, b, c int

	// when send '1234<enter>' to stdio,
	// then a = 1234, n = 1, err = nil
	n, err := fmt.Scan(&a)
	fmt.Printf("a = %v\n", a)
	fmt.Printf("n = %v\n", n)
	fmt.Printf("err = %#v\n", err)

	// when send '1 2 3' to stdio,
	// then a = 1, b = 2, c = 3, n = 3, err = nil

	// when send '1 2<enter>' to stdio,
	// then a = 1, b = 2, c = 0 (default value), n = 2, err = unexpected newline

	// when send '1 2 3 4<enter>' to stdio,
	// then a = 1, b = 2, c = 3, n = 3, err = expected newline
	n, err = fmt.Scanln(&a, &b, &c)
	fmt.Printf("a = %v\n", a)
	fmt.Printf("b = %v\n", b)
	fmt.Printf("c = %v\n", c)
	fmt.Printf("n = %v\n", n)
	fmt.Printf("err = %#v\n", err)
}
