package comb

func powMod(x, y, mod int) int {
	if y == 0 {
		return 1
	}

	result := 1
	for i := 0; i < y; i++ {
		result *= x
		result %= mod
	}
	return result
}

func calcComb(n, r, mod int) int {
	result := 1
	if r > n-r {
		r = n - r
	}
	for i := n; i > n-r; i-- {
		result *= i
		result %= mod
	}
	for i := 2; i <= r; i++ {
		result *= powMod(i, mod-2, mod)
		result %= mod
	}
	return result
}
