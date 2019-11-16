package utils

type StringSlice []string

// EntryExists checks if a string exists in a slice of string
func (s StringSlice) EntryExists(entry string) bool {
	for i := 0; i < len(s); i++ {
		if s[i] == entry {
			return true
		}
	}
	return false
}
