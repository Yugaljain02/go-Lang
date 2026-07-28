package main

import "fmt"

func calculate(a int, b int) (int, int) {
	sum := a + b
	multiply := a * b

	return sum, multiply
}

func main() {
	add, mul := calculate(10, 20)

	fmt.Println("Addition:", add)
	fmt.Println("Multiplication:", mul)
}