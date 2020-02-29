package utils

import (
	"fmt"
	"testing"
	"time"
)

func TestIsValidDate(t *testing.T) {

	validDates := []string{
		"1993-02-10",
		"2020-01-01",
	}

	invalidDates := []string{
		"2050-01-01",
		"202-01-01",
		"22020-01-01",
		"2020-00-01",
		"2020-1-01",
		"2020-011-01",
		"2020-13-01",
		"2020-00-00",
		"2020-01-1",
		"2020-01-011",
		"2020-01-32",
		"2019-02-29",
		"2020-04-31",
	}

	for _, v := range validDates{
		result, err := IsValidDate(v)
		if result != true {
			t.Errorf("Test Failed!, expected: %v, got: %v", true, result)
		}
		if err != nil {
			t.Errorf("Did not expect an error, but got one")
		}
	}

	for _, v := range invalidDates{
		result, err := IsValidDate(v)
		if result != false {
			t.Errorf("Test Failed!, expected: %v, got: %v", false, result)
		}
		if err == nil {
			t.Errorf("Expected to get an error, but got none")
		}
	}

}

func ExampleIsValidDate() {
	fmt.Println(IsValidDate("2020-01-01"))
	// Output: true <nil>
}

func TestIsValidYear(t *testing.T) {

	validYears := []int{
		1990,
		1993,
		time.Now().Year(),
	}

	invalidYears := []int{
		1889,
		time.Now().Year() + 1,
		20000,
		200,
	}

	for _, v := range validYears{
		result, err := IsValidYear(v)
		if result != true {
			t.Errorf("Test Failed!, expected: %v, got: %v", true, result)
		}
		if err != nil {
			t.Errorf("Did not expect an error, but got one")
		}
	}

	for _, v := range invalidYears{
		result, err := IsValidYear(v)
		if result != false {
			t.Errorf("Test Failed!, expected: %v, got: %v", false, result)
		}
		if err == nil {
			t.Errorf("Expected to get an error, but got none")
		}
	}
}

func ExampleIsValidYear() {
	fmt.Println(IsValidYear(2020))
	// Output: true <nil>
}