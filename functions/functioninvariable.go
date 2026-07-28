package main
import "fmt"

func add(a int, b int) int{
	return a+b
}
func main(){
  myfunction := add
    
    function := myfunction(10,40)
	fmt.Println(function)
}