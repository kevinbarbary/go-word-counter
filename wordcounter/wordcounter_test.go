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

func TestCounts1(t *testing.T) {
	NewWordStore()
	dictionary.add("test")
	ans := dictionary.counts()
	if ans["test"] != 1 {
		t.Errorf("counts = %d, want 1", ans["test"])
	}
}

func TestCounts2(t *testing.T) {
	NewWordStore()
	dictionary.add("test")
	dictionary.add("test")
	ans := dictionary.counts()
	if ans["test"] != 2 {
		t.Errorf("counts = %d, want 2", ans["test"])
	}
}

func TestCounts3(t *testing.T) {
	NewWordStore()
	dictionary.add("test")
	dictionary.add("thing")
	dictionary.add("thing")
	dictionary.add("test")
	dictionary.add("test")
	ans := dictionary.counts()
	if ans["test"] != 3 {
		t.Errorf("counts = %d, want 3", ans["test"])
	}
}
