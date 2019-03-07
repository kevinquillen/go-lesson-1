package main

import "testing"

func TestRandom(t *testing.T) {
	min := 5
	max := 20
	number, _ := randomRange(min, max)

	if number < min {
		t.Errorf("Number generated was less than the minimum, got: %d. Minimum: %d.", number, min)
	}

	if number > max {
		t.Errorf("Number generated was greater than the maximum, got: %d. Maximum: %d.", number, max)
	}

	// String testing is probably asinine, how do you do a better 'exception expected' style test?
	_, err := randomRange(50, 10)

	if err == nil {
		t.Errorf("Maximum passed is greater than the minimum, an error should be returned, got nil.")
	}
}
