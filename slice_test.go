package utils

import (
	"reflect"
	"testing"
)

func TestEntryExists(t *testing.T) {

	slice := []string{"one", "two", "three"}

	result := EntryExists(slice, "one")

	if result == false {
		t.Errorf("Test Failed!, expected: %v, got: %v", true, result)
	}

	result = EntryExists(slice, "four")

	if result == true {
		t.Errorf("Test Failed!, expected: %v, got: %v", false, result)
	}
}

func TestGetSliceEntryIndex(t *testing.T) {

	slice := []string{"one", "two", "three"}

	result := GetSliceEntryIndex(slice, "one")

	if result != 0 {
		t.Errorf("Test Failed!, expected: %v, got: %v", 0, result)
	}

	result = GetSliceEntryIndex(slice, "doesNotExist")

	if result != -1 {
		t.Errorf("Test Failed!, expected: %v, got: %v", -1, result)
	}
}

func TestRemoveEntryFromSlice(t *testing.T) {
	slice := []string{"one", "two", "three"}
	expectedResult := []string{"one", "three"}

	result := RemoveEntryFromSlice(slice, "two")

	if !reflect.DeepEqual(result, expectedResult){
		t.Errorf("Test Failed!, expected: %v, got: %v", expectedResult, result)
	}

	slice = []string{"one", "two", "three"}
	result = RemoveEntryFromSlice(slice, "four")

	if reflect.DeepEqual(result, expectedResult){
		t.Errorf("Test Failed!, expected: %v, got: %v", slice, result)
	}

	slice = []string{"one", "two", "three", "two"}
	expectedResult = []string{"one", "three", "two"}

	result = RemoveEntryFromSlice(slice, "two")

	if !reflect.DeepEqual(result, expectedResult){
		t.Errorf("Test Failed!, expected: %v, got: %v", expectedResult, result)
	}

	slice = []string{"one", "two",  "two", "three"}
	expectedResult = []string{"one", "two" , "three"}

	result = RemoveEntryFromSlice(slice, "two")

	if !reflect.DeepEqual(result, expectedResult){
		t.Errorf("Test Failed!, expected: %v, got: %v", expectedResult, result)
	}
}

func TestRemoveDuplicateEntries(t *testing.T) {
	slice := []string{"one", "two", "two", "three", "three", "three", "four", "four", "four", "four"}
	expectedResult := []string{"one", "two", "three", "four"}

	result := RemoveDuplicateEntries(slice)

	if !reflect.DeepEqual(result, expectedResult){
		t.Errorf("Test Failed!, expected: %v, got: %v", expectedResult, result)
	}
}

func TestCountDuplicateEntries(t *testing.T) {
	slice := []string{"one", "two", "two", "three", "three", "three", "four", "four", "four", "four"}

	expectedResult := make(map[string]int)
	expectedResult["one"] = 1
	expectedResult["two"] = 2
	expectedResult["three"] = 3
	expectedResult["four"] = 4

	result := CountDuplicateEntries(slice)

	if !reflect.DeepEqual(result, expectedResult){
		t.Errorf("Test Failed!, expected: %v, got: %v", expectedResult, result)
	}
}

func TestDuplicateEntryExists(t *testing.T) {
	slice := []string{"one", "two", "two", "three", "three", "three", "four", "four", "four", "four"}

	result := DuplicateEntryExists(slice)

	if result == false {
		t.Errorf("Test Failed!, expected: %v, got: %v", true, result)
	}

	slice = []string{"one", "two", "three", "four"}

	result = DuplicateEntryExists(slice)

	if result == true {
		t.Errorf("Test Failed!, expected: %v, got: %v", false, result)
	}
}
