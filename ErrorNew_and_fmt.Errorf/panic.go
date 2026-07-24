package main

import "fmt"

func main() {
	fmt.Println("first")
	panic("something went wrong")
	fmt.Println("second")
}
