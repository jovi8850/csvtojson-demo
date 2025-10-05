package main

import "testing"

// Test helper function atoi
func TestAtoi(t *testing.T) {
	result := atoi("42")
	if result != 42 {
		t.Errorf("Expected 42, got %d", result)
	}
}

// Test helper function atof
func TestAtof(t *testing.T) {
	result := atof("3.14")
	if result != 3.14 {
		t.Errorf("Expected 3.14, got %f", result)
	}
}
