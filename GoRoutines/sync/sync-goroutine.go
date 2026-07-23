package main

import(
   "sync"	
	"fmt"
)
       var wg sync.WaitGroup
	   var once sync.Once

func incr(){

	fmt.Println("done 1")

}

  func initialise(id int,wg *sync.WaitGroup){

	defer wg.Done()

	fmt.Printf("worker %d",id)
	fmt.Println()

	once.Do(incr)

  }

func main(){
      

	  wg.Add(3)

	  go initialise(1,&wg)
      go initialise(2,&wg)
	  go initialise(3,&wg)

	  wg.Wait()


}