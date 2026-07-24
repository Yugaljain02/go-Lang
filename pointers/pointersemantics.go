package main

import "fmt"

func update(age *int) {
	*age = 30
}

func main() {
	age := 22
	update(&age)
	fmt.Println(age)
}
