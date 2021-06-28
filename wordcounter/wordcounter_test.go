package wordcounter

import (
	"testing"
)

func TestIsEmptyWhenEmpty(t *testing.T) {
	NewWordStore()
	if !dictionary.isEmpty() {
		t.Errorf("Word store is empty but isEmpty returned false")
	}
}

func TestIsEmptyWhenNotEmpty(t *testing.T) {
	NewWordStore()
	dictionary.add("test")
	if dictionary.isEmpty() {
		t.Errorf("Word store is not empty but isEmpty returned true")
	}
}
