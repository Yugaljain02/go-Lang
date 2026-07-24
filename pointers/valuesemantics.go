package main

import "fmt"

func Update(age int) {
	age = 30
}

func main() {
	age := 22

	Update(age)

	fmt.Println(age)
}
