package main

import (
	"fmt"
	"errors"
)
 func database() error{
     return errors.New("database connection failed")
 }
    func saveuser() error{
		err := database()
		if err != nil{
			return fmt.Errorf("user is not saved %w",err)
		}
		return nil
	}
 func register() error{
	err := saveuser()

	if err != nil {
		return fmt.Errorf("registration failed !! %w",err)
 }
 return nil
 }
func main(){
	err := register()
	if err !=nil{
		fmt.Println(err)
	}
}