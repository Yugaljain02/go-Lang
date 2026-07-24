package main

import "fmt"

func main() {
	defer func() {
		r := recover()
		if r != nil {
			fmt.Println("Recovered : ", r)
		}
	}()
	fmt.Println("first")
	panic("something went wrong")
	fmt.Println("end")
}
