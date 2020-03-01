package utils

import (
	"fmt"
	"regexp"
	"testing"
)

func TestGetRandomPassword(t *testing.T) {

	passwordRegex := "^[A-Za-z0-9]{23}$"
	cRegex, err := regexp.Compile(passwordRegex)
	if err != nil {
		t.Fatal("There was an error compiling the regex")
	}

	result := GetRandomPassword()

	if !cRegex.MatchString(result) {
		t.Errorf("Test Failed!. Expected a 23 bit alphanumeric password but got a %v bit password %v", len(result), result)
	}
}

func ExampleGetRandomPassword() {
	password := GetRandomPassword()
	fmt.Println(password)
}
