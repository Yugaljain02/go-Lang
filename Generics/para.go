package main

import "fmt"

func show[K any, V any](key K, value V) {
	fmt.Println(key, value)
}
func main() {
	show(1, "apple")
}
