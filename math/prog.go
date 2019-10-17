package math4procon

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
	for i := 0; i < len(x); i++ {
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
	for i := 0; i < len(x); i++ {
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
	ret := x[0]
	for i := 1; i < len(x); i++ {
		ret = gcd(ret, x[i])
	}
	return ret
}

func lcm(x, y int) int {
	return x * y / gcd(x, y)
}
