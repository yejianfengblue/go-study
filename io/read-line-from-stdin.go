package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	b := scanner.Scan()

	// when scan stdin, Scan() always return True, if there is no error, because there is no END
	fmt.Printf("b = %v\n", b)
	if b {
		var line string = scanner.Text()
		fmt.Printf("Input line = %v\n", line)
	}
}
