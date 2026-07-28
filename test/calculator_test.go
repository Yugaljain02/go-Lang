package main
import (
	"testing"
   "fmt"
)
func TestAdd(t *testing.T){
	result := Add(10,30)

	if result !=40{
		t.Errorf(" addition should be %d but we found 60\n",result)
	}
	fmt.Println("checking compleleted")
}
func TestMultiply(t *testing.T){
	result := Multiply(10,30)

	if result !=300{

		t.Fatalf("multiply should be %d but result comes 3400\n", result)

	}
	fmt.Println("checking compleleted")
}
