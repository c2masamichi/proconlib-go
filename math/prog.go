package number

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
