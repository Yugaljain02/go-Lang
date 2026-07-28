package main

import (
	"fmt"
)

func show[a comparable](b a, c a) bool {
	return c == b
}

func main() {
	fmt.Println(show(1034, 10))
	fmt.Println(show("augst", "august"))
}
