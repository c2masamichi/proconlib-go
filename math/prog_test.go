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
		x        int
		y        int
		expected int
	}{
		{2, 1, 2},
		{3, 4, 4},
		{5, 5, 5},
	}

	for i, param := range parameters {
		x, y, expected := param.x, param.y, param.expected
		actual := max(x, y)
		if actual != expected {
			t.Errorf("i: %d\nactual: %v\nexpected: %v", i, actual, expected)
		}
	}
}

func TestMaxList(t *testing.T) {
	parameters := []struct {
		x        []int
		expected int
	}{
		{[]int{2}, 2},
		{[]int{3, 4, 1, 2}, 4},
	}

	for i, param := range parameters {
		x, expected := param.x, param.expected
		actual := maxList(x)
		if actual != expected {
			t.Errorf("i: %d\nactual: %v\nexpected: %v", i, actual, expected)
		}
	}
}

func TestMin(t *testing.T) {
	parameters := []struct {
		x        int
		y        int
		expected int
	}{
		{2, 1, 1},
		{3, 4, 3},
		{5, 5, 5},
	}

	for i, param := range parameters {
		x, y, expected := param.x, param.y, param.expected
		actual := min(x, y)
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

func TestGcd(t *testing.T) {
	parameters := []struct {
		x        int
		y        int
		expected int
	}{
		{8, 2, 2},
		{17, 13, 1},
		{147, 105, 21},
	}

	for i, param := range parameters {
		x, y, expected := param.x, param.y, param.expected
		actual := gcd(x, y)
		if actual != expected {
			t.Errorf("i: %d\nactual: %v\nexpected: %v", i, actual, expected)
		}
	}
}

func TestLcm(t *testing.T) {
	parameters := []struct {
		x        int
		y        int
		expected int
	}{
		{2, 4, 4},
		{2, 3, 6},
		{16, 24, 48},
	}

	for i, param := range parameters {
		x, y, expected := param.x, param.y, param.expected
		actual := lcm(x, y)
		if actual != expected {
			t.Errorf("i: %d\nactual: %v\nexpected: %v", i, actual, expected)
		}
	}
}
