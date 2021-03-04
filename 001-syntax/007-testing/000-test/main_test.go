package main

import "testing"

func TestSum(t *testing.T) {
	x := sum(10, 9)
	if x != 5 {
		t.Error("Expected", 19, "Got", x)
	}
}
