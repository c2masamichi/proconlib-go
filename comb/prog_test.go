package comb

import (
	"testing"
)

func TestPowMod(t *testing.T) {
	parameters := []struct {
		x        int
		y        int
		mod      int
		expected int
	}{
		{2, 0, 101, 1},
		{2, 4, 101, 16},
		{2, 4, 11, 5},
		{2, 5, 101, 32},
	}

	for i, param := range parameters {
		x, y, mod, expected := param.x, param.y, param.mod, param.expected
		actual := powMod(x, y, mod)
		if actual != expected {
			t.Errorf("i: %d\nactual: %v\nexpected: %v", i, actual, expected)
		}
	}
}

func TestCalcComb(t *testing.T) {
	parameters := []struct {
		n        int
		r        int
		mod      int
		expected int
	}{
		{3, 1, 11, 3},
		{5, 2, 11, 10},
		{8, 4, 11, 4},
		{8, 4, 101, 70},
	}

	for i, param := range parameters {
		n, r, mod, expected := param.n, param.r, param.mod, param.expected
		actual := calcComb(n, r, mod)
		if actual != expected {
			t.Errorf("i: %d\nactual: %v\nexpected: %v", i, actual, expected)
		}
	}
}
