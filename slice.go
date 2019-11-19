package utils

// EntryExists checks if a string exists in a slice of string
func EntryExists(slice []string, entry string) bool {
	for i := 0; i < len(slice); i++ {
		if slice[i] == entry {
			return true
		}
	}
	return false
}
