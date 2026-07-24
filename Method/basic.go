package main

import "fmt"

type user struct {
	car_name string
	speed    int
}

// basically we create a method and attach it to reciever  type
func (car user) speedon() {
	fmt.Println(car.car_name)
}
func main() {
	s := user{
		"hyundai", 23,
	}
	s.speedon()

}
