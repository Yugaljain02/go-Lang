package main

import "fmt"

func add(numbers ...int) {
	total := 0
	for _, num := range numbers {
		total += num
	}
	fmt.Println(total)
}
func main() {

	add(3)
	add(34, 32, 45)
	add(23)
}
