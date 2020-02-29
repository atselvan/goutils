package utils

import (
	"log"
)

// EntryExists checks if a string exists in a slice of strings
func EntryExists(slice []string, entry string) bool {
	for i := 0; i < len(slice); i++ {
		if slice[i] == entry {
			return true
		}
	}
	return false
}

// GetSliceEntryIndex returns the index of an entry in the slice of strings
func GetSliceEntryIndex(slice []string, entry string) int {
	for i := 0; i < len(slice); i++ {
		if slice[i] == entry {
			return i
		}
	}
	return -1
}

// RemoveEntryFromSlice removes a string entry from a slice of strings
func RemoveEntryFromSlice(slice []string, entry string) []string {
	i := GetSliceEntryIndex(slice, entry)
	if i == -1 {
		log.Printf(entryDoesNotExistMsg, entry)
		return slice
	}
	return append(slice[:i], slice[i+1:]...)
}

// RemoveDuplicates removes duplicate entries in a slice of strings and returns a slice with unique entries
func RemoveDuplicates(stringSlice []string) []string {
	keys := make(map[string]bool)
	var result []string
	for _, entry := range stringSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			result = append(result, entry)
		}
	}
	return result
}

// CountDuplicates counts the number of entries in a slice
// If a duplicate entry exists the count of that entry is incremented
// The function returns a map of the duplicate frequencies
func CountDuplicates(list []string) map[string]int {

	duplicateCount := make(map[string]int)

	for _, item := range list {
		// check if the item/element exist in the duplicateCount map
		if _, exist := duplicateCount[item]; exist {
			duplicateCount[item] += 1 // increase counter by 1 if already in the map
		} else {
			duplicateCount[item] = 1 // else start counting from 1
		}
	}
	return duplicateCount
}

// DuplicateExists checks if a slice has duplicate entries
// If yes the function returns true else the function returns false
func DuplicateExists(stringSlice []string) bool {

	duplicateCount := CountDuplicates(stringSlice)

	for _, k := range duplicateCount {
		if k > 1 {
			return true
		}
	}

	return false
}
