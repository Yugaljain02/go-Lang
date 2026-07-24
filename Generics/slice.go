package main

import "fmt"

func printslice[k any](arr []k) {
	for _, num := range arr {
		fmt.Println(num)
	}
}
func main() {
	numbers := []int{1, 2, 3, 4, 5}
	numbers = append(numbers, 6)
	names := []string{"Abhinav", "garvit", "vivan", "Yugal", " prankur"}
	printslice(numbers)
	printslice(names)
}
