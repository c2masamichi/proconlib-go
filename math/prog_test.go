package math4procon

import (
	"testing"
)

func TestMax(t *testing.T) {
	parameters := []struct {
		a        int
		b        int
		expected int
	}{
		{2, 1, 2},
		{3, 4, 4},
		{5, 5, 5},
	}

	for i, param := range parameters {
		a, b, expected := param.a, param.b, param.expected
		actual := max(a, b)
		if actual != expected {
			t.Errorf("i: %d\nactual: %v\nexpected: %v", i, actual, expected)
		}
	}
}
