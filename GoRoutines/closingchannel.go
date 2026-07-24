package main

import "fmt"

func main() {
	ch := make(chan int, 5)

	ch <- 10
	ch <- 20

	ch <- 30
	ch <- 40
	//  close(ch)
	ch <- 60
	close(ch)
	go func() {
		fmt.Println(<-ch)
		fmt.Println(<-ch)
		fmt.Println(<-ch)
	}()
	for i := range ch {
		fmt.Println(i)
	}
}
