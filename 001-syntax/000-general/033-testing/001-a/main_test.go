package main

import "testing"

func TestAdd(t *testing.T) {
	result := Add(1,3)
	if result != 4{
		t.Errorf("Expected: 4, got: %d", result)
	}
}