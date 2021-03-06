package number

import (
	"reflect"
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

func TestMinList(t *testing.T) {
	parameters := []struct {
		x        []int
		expected int
	}{
		{[]int{2}, 2},
		{[]int{3, 4, 1, 2}, 1},
	}

	for i, param := range parameters {
		x, expected := param.x, param.expected
		actual := minList(x)
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

func TestGcdList(t *testing.T) {
	parameters := []struct {
		x        []int
		expected int
	}{
		{[]int{3}, 3},
		{[]int{8, 2}, 2},
		{[]int{6, 15, 9}, 3},
		{[]int{17, 13, 7}, 1},
	}

	for i, param := range parameters {
		x, expected := param.x, param.expected
		actual := gcdList(x)
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

func TestLcmList(t *testing.T) {
	parameters := []struct {
		x        []int
		expected int
	}{
		{[]int{3}, 3},
		{[]int{2, 4}, 4},
		{[]int{2, 3, 5}, 30},
		{[]int{10, 6, 8}, 120},
	}

	for i, param := range parameters {
		x, expected := param.x, param.expected
		actual := lcmList(x)
		if actual != expected {
			t.Errorf("i: %d\nactual: %v\nexpected: %v", i, actual, expected)
		}
	}
}

func TestIsPrime(t *testing.T) {
	parameters := []struct {
		x        int
		expected bool
	}{
		{1, false},
		{2, true},
		{4, false},
		{5, true},
		{71, true},
		{91, false},
	}

	for i, param := range parameters {
		x, expected := param.x, param.expected
		actual := isPrime(x)
		if actual != expected {
			t.Errorf("i: %d\nactual: %v\nexpected: %v", i, actual, expected)
		}
	}
}

func TestCountFactors(t *testing.T) {
	parameters := []struct {
		x        int
		mod      int
		expected int
	}{
		{3, 3, 1},
		{7, 2, 0},
		{16, 2, 4},
		{48, 2, 4},
	}

	for i, param := range parameters {
		x, mod, expected := param.x, param.mod, param.expected
		actual := countFactors(x, mod)
		if actual != expected {
			t.Errorf("i: %d\nactual: %v\nexpected: %v", i, actual, expected)
		}
	}
}
func TestMakeDivisors(t *testing.T) {
	parameters := []struct {
		x        int
		expected []int
	}{
		{1, []int{1}},
		{3, []int{1, 3}},
		{6, []int{1, 2, 3, 6}},
		{16, []int{1, 2, 4, 8, 16}},
		{25, []int{1, 5, 25}},
	}

	for i, param := range parameters {
		x, expected := param.x, param.expected
		actual := makeDivisors(x)
		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("i: %d\nactual: %v\nexpected: %v", i, actual, expected)
		}
	}
}
