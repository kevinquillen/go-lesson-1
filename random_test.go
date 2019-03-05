package main

import "testing"

func TestRandom(t *testing.T) {
	min := 5
	max := 20
	number := random(min, max)

	if number < min {
		t.Errorf("Number generated was less than the minimum, got: %d. Minimum: %d.", number, min)
	}

	if number > max {
		t.Errorf("Number generated was greater than the maximum, got: %d. Maximum: %d.", number, max)
	}
}
