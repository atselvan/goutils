package utils

import (
	"testing"
)

func TestStringSlice_EntryExists(t *testing.T) {
	ss := StringSlice{"one", "two", "three"}

	if !ss.EntryExists("one") {
		t.Errorf("Expected 'true' but got '%t'", ss.EntryExists("one"))
	}

	if ss.EntryExists("four") {
		t.Errorf("Expected 'false' but got '%t'", ss.EntryExists("one"))
	}
}
