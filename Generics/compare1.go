package main

import "fmt"

func show[T comparable](arr []T, target T) bool {

	for _, num := range arr {
		if target == num {
			return true
		}
	}
	return false

}
func main() {
	numbers := []int{1, 2, 3, 4, 5}
	fmt.Println(show(numbers, 22))
}
