package main

import "testing"

func TestSquare1(t *testing.T) {
	rst := square(9)
	if rst != 81 {
		t.Errorf("square(9) sholud be 81 but retruns %d", rst)
	}
}

func TestSquare2(t *testing.T) {
	rst := square(3)
	if rst != 9 {
		t.Errorf("square(9) sholud be 9 but retruns %d", rst)

	}
}
