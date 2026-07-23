package main
import (
    "sync"	
	"fmt"
)
 var once sync.Once
func initialise(){
	fmt.Println("hello")
}
func main(){

	once.Do(initialise)
	once.Do(initialise)
    once.Do(initialise)

}