package main

import (
	"fmt"
	"time"
)

func main() {
	pizza := make(chan string, 3)
	burger := make(chan string, 3)
	pizza <- "pizza"
	burger <- "burger"
	for i := 0; i < 2; i++ {
		select {
		case msg := <-pizza:
			fmt.Println(msg)
		case msg := <-burger:
			fmt.Println(msg)
			time.Sleep(2 * time.Second)
		default:
			fmt.Println("No food is ready")
		}
	}
	// go func(){

	// }()
}
