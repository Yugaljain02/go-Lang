package main

import (
	"GoProject/Visibility/calculator"
	"fmt"
	// "github.com/gin-gonic/gin"
)

func main() {
	a := 12
	b := 20
	fmt.Println(calculator.Add(a, b))
	fmt.Println(calculator.Sub(a, b))
}
