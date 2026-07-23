package main
import (
	"fmt"
	"sync"
	"sync/atomic"
)
var counter int64
 //var wg sync.WaitGroup

func worker(id int, wg *sync.WaitGroup){
	defer wg.Done()
	
	atomic.AddInt64(&counter,1)
	
	
}

func main(){

   var wg sync.WaitGroup
      
	   for i:=1;i<=2000;i++{
		 wg.Add(1)
	   go worker(1,&wg)
	   }
	  
	    //go worker(2,&wg)
		 //go worker(3,&wg)
		 wg.Wait()
		  fmt.Println("counter value ", counter)

}