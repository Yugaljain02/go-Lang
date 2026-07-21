package main

 import (
	"fmt"
    "errors")
func register(age int) error{
	if age<18 {
		return errors.New("age must be greater then 17")
	}
	return nil
}

func main(){
	err := register(13)
	if err !=nil{
		fmt.Println(err) 
	}
}