package main
import "fmt"

func main(){
	ch :=make(chan int,3)
go func(){
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}()
	ch <- 10
	ch <- 20
	ch <- 30
	ch <- 40
	ch <- 50
	fmt.Println(<-ch)
}