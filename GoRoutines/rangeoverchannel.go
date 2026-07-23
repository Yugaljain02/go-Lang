package main
import "fmt"

func main(){
	ch := make(chan string)

	go func(){
		ch <- "Pizza"
		ch <- "Pizza"
		ch <- "Pizza"
		close(ch)
		ch<-"Pasta"
	 }()
	 for msg := range ch{
        
	fmt.Println(msg)
}
}