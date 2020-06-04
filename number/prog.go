package number

import "sort"

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func maxList(x []int) int {
	maxv := x[0]
	for i := 1; i < len(x); i++ {
		maxv = max(maxv, x[i])
	}
	return maxv
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func minList(x []int) int {
	minv := x[0]
	for i := 1; i < len(x); i++ {
		minv = min(minv, x[i])
	}
	return minv
}

func divmod(x, y int) (int, int) {
	return x / y, x % y
}

func gcd(x, y int) int {
	if x < y {
		x, y = y, x
	}
	for y > 0 {
		r := x % y
		x = y
		y = r
	}
	return x
}

func gcdList(x []int) int {
	result := x[0]
	for i := 1; i < len(x); i++ {
		result = gcd(result, x[i])
	}
	return result
}

func lcm(x, y int) int {
	return x * y / gcd(x, y)
}

func lcmList(x []int) int {
	result := x[0]
	for i := 1; i < len(x); i++ {
		result = lcm(result, x[i])
	}
	return result
}

func isPrime(x int) bool {
	if x < 2 {
		return false
	}
	if x == 2 {
		return true
	}
	if x%2 == 0 {
		return false
	}
	for i := 3; i*i <= x; i += 2 {
		if x%i == 0 {
			return false
		}
	}
	return true
}

func countFactors(x, mod int) int {
	cnt := 0
	r := 0
	for r == 0 {
		x, r = divmod(x, mod)
		if r == 0 {
			cnt++
		}
	}
	return cnt
}

func makeDivisors(x int) []int {
	if x == 1 {
		return []int{1}
	}

	divisor := []int{1, x}
	for i := 2; i*i <= x; i++ {
		q, r := divmod(x, i)
		if r == 0 {
			divisor = append(divisor, i)
			if q != i {
				divisor = append(divisor, q)
			}
		}
	}
	sort.Ints(divisor)
	return divisor
}
