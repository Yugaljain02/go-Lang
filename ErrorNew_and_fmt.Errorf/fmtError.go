package main

import "fmt"

 func register(age int) error {
   if age <18 {
	 return fmt.Errorf("age %d is not allowed ", age)
   }
   return nil
 }

func main(){
	err := register(14)
	if err!=nil {
		fmt.Println(err)
	}
}