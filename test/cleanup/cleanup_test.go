package main

import (
	"fmt"
	"os"
	"testing"
)

func TestFile(t *testing.T) {

	// Step 1: Create a file
	file, err := os.Create("test.txt")
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println("File created")

	// Step 2: Register cleanup
	t.Cleanup(func() {
		file.Close()
		os.Remove("test.txt")

		fmt.Println("File deleted")
	})

	// Step 3: Use the file
	_, err = file.WriteString("Hello Go")
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println("Test completed")
}
