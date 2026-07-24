package main

import "fmt"

func main() {
	ch := make(chan string)

	go func() {
		ch <- "Pizza"
	}()
	msg := <-ch
	fmt.Println(msg)
}
