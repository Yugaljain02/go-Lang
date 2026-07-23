package main

import(
	"sync"
	"fmt"
)
  func iterate(wg *sync.WaitGroup){

	defer wg.Done()

	for i:=0; i<5; i++{
    
		fmt.Println(i)
	}
	

  }
func main(){

    var wg sync.WaitGroup 

    wg.Add(1)

    go iterate(&wg)

    wg.Wait()

    fmt.Println("Main finished")
}