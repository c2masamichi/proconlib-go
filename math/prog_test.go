package math4procon

import (
	"testing"
)

func TestAbs(t *testing.T) {
	parameters := []struct {
		x        int
		expected int
	}{
		{2, 2},
		{-3, 3},
	}

	for i, param := range parameters {
		x, expected := param.x, param.expected
		actual := abs(x)
		if actual != expected {
			t.Errorf("i: %d\nactual: %v\nexpected: %v", i, actual, expected)
		}
	}
}

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

func TestMin(t *testing.T) {
	parameters := []struct {
		a        int
		b        int
		expected int
	}{
		{2, 1, 1},
		{3, 4, 3},
		{5, 5, 5},
	}

	for i, param := range parameters {
		a, b, expected := param.a, param.b, param.expected
		actual := min(a, b)
		if actual != expected {
			t.Errorf("i: %d\nactual: %v\nexpected: %v", i, actual, expected)
		}
	}
}

func TestDivmod(t *testing.T) {
	type result struct {
		quotient  int
		remainder int
	}
	parameters := []struct {
		x        int
		y        int
		expected result
	}{
		{5, 2, result{2, 1}},
		{9, 5, result{1, 4}},
		{4, 5, result{0, 4}},
	}

	for i, param := range parameters {
		x, y, expected := param.x, param.y, param.expected
		q, r := divmod(x, y)
		actual := result{q, r}
		if actual != expected {
			t.Errorf("i: %d\nactual: %v\nexpected: %v", i, actual, expected)
		}
	}
}
