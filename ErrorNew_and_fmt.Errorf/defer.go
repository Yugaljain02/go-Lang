package main

import "fmt"

func main() {
	fmt.Println("first")
	defer fmt.Println("second")
	fmt.Println("third")
	defer fmt.Println("fourth")
	defer fmt.Println("fifth")
	defer fmt.Println("six")

}
