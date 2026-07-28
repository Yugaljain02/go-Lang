package main
import "fmt"


func Add(a int, b int) (add ,mul int){
	
	mul = a*b
	add = a+b
	return 
}

func main(){
	
   fmt.Println(Add(10,20))
}