package utils

import "fmt"

func ExampleError_NewError() {
	err := Error{
		Message: "Some Error",
		Detail:  "",
	}
	fmt.Println(err.NewError())
	// Output: Some Error
}

func ExampleError_NewErrorWithDetails() {
	err := Error{
		Message: "Some Error",
		Detail:  "Error Details",
	}
	fmt.Println(err.NewError())
	// Output: Some Error : Error Details
}
