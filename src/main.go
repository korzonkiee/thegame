package main

import (
	"fmt"
)

func main() {
	fmt.Printf("%d\n", add(2, 2))
}

func add(a int, b int) int {
	return a + b
}
