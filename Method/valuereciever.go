package main
import "fmt"

 type user struct{
	name string
 }

func (u *user)update(){
	u.name = "jainish"
	
}

func main(){
	s :=&user{
		name : "Yugal",
	}
	s.update()
	fmt.Println(s.name)
}