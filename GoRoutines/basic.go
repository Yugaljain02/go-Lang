package main

import (
	"fmt"
	"time"
)

func A() {
	fmt.Println("A Start")
	time.Sleep(2 * time.Second)
	fmt.Println("A End")
}
func B(){
	fmt.Println("B started")
    time.Sleep(2 * time.Second)
	fmt.Println("B is finished")
}

func main() {
	fmt.Println("1")
	go A()
	fmt.Println("2")
	B()
//time.Sleep(1 * time.Second)
}