package main

import "fmt"

func print(v any) {
	switch value := v.(type) {

	case int:
		fmt.Printf("integer %d\n", value)
	case bool:
		fmt.Printf("boolean %t\n", value)
	case float64:
		fmt.Printf("decimal %f\n", value)
	case string:
		fmt.Printf("string %s\n", value)
	default:
		fmt.Printf("undefined type\n")
	}
}

func main() {
	print(6)
	print(2.22)
	print(true)
	print("char")
}
