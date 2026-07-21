package main

import (
	"errors"
	"fmt"
)

// Custom error type
type ValidationError struct {
	Field string
}

func (v ValidationError) Error() string {
	return "invalid field: " + v.Field
}

func validate() error {
	return ValidationError{
		Field: "Email",
	}
}

func register() error {
	err := validate()
	if err != nil {
		return fmt.Errorf("registration failed: %w", err)
	}
	return nil
}

func main() {
	err := register()

	if err != nil {
		fmt.Println(err)

		var ve ValidationError

		if errors.As(err, &ve) {
			fmt.Println("Validation Error Found")
			fmt.Println("Invalid Field:", ve.Field)
		}
	}
}