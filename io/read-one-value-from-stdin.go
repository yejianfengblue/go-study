package main

import (
	"fmt"
)


func main() {

	var a int

	// when send '1234<enter>' to stdio,
	// then a = 1234, n = 1, err = nil
	n, err := fmt.Scan(&a)
	fmt.Printf("a = %v\n", a)
	fmt.Printf("n = %v\n", n)
	fmt.Printf("err = %#v\n", err)
}
