package main

import (
	"fmt"
	"sync"
)

var counter int
var mu sync.Mutex

func increment(wg *sync.WaitGroup) {
	defer wg.Done()
	mu.Lock()
	defer mu.Unlock()

	for i := 0; i < 1000; i++ {
		counter++
		if counter == 3 {
			return
		}
	}
}

func main() {

	var wg sync.WaitGroup

	wg.Add(2)

	go increment(&wg)
	go increment(&wg)

	wg.Wait()

	fmt.Println(counter)
}
