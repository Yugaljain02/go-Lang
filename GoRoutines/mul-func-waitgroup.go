package main

import (
	"fmt"
	"sync"
)
func worker(id int,wg *sync.WaitGroup){
	defer wg.Done()
	
		fmt.Println("worker")
	
}
func main(){
	var wg sync.WaitGroup

	wg.Add(3) // 3 goroutine
	go worker(1, &wg)
	//go worker(2, &wg)
	go worker(3, &wg)
	wg.Wait()
	fmt.Println("main finished")
}