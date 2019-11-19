package utils

import (
	"testing"
)

func TestStringSlice_EntryExists(t *testing.T) {
	ss := []string{"one", "two", "three"}

	if !EntryExists(ss, "one") {
		t.Errorf("Expected 'true' but got '%t'", EntryExists(ss, "one"))
	}

	if EntryExists(ss, "four") {
		t.Errorf("Expected 'false' but got '%t'", EntryExists(ss, "four"))
	}
}
