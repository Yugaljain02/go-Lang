package main

import "fmt"

type AgeError struct{}

func (a AgeError) Error() string {
	return "Age should be greater then 17"
}
func register(age int) error {
	if age < 18 {
		return AgeError{}
	}
	return nil
}
func main() {
	err := register(12)
	if err != nil {
		fmt.Println(err)
	}
}
