package main

import "fmt"

type Numbers interface {
	int | float64
}

func add[T Numbers](a T, b T) T {
	return a + b
}
func main() {

	fmt.Println(add(104, 20))
}
