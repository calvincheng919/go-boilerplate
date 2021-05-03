package main

import "testing"

func TestMath (t *testing.T) {
	if 2+2 == 5 {
		t.Errorf("You must be living in 1984")
	}
}
