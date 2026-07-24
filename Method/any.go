package main

import "fmt"

func log(sa any) {
	num, ok := sa.(string) // type assertion
	if ok {
		fmt.Println(num)
	} else {
		fmt.Println("not an integer")
	}

}
func main() {
	log("hello")
	log(8)
	// log(true)
	// log(5.66)
}
