package main

import (
	"testing"
	"fmt"
)

func TestAdd(t *testing.T) {
	tests := []struct {
		name     string
		a        int
		b        int
		expected int
	}{
		{"positive number", 10, 20, 30},
		{"negative number", -10, -40, -550},
	{"zero", 0, 0, 0},
		//{"positive number",10,20,30},
	}
	for _, num := range tests {
   t.Run(num.name ,func (t *testing.T){
		result := Add(num.a, num.b)

		if result != num.expected {
			t.Errorf("%s error of %d and %d result is not come expected as %d ", num.name, num.a, num.b, result)
		} else {
			fmt.Println("result is correct")

		}
	})
	}
	}