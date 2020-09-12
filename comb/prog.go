package comb

func powMod(x, y, mod int) int {
	if y == 0 {
		return 1
	}

	if y%2 == 0 {
		d := powMod(x, y/2, mod)
		return (d * d) % mod
	} else {
		return (x * powMod(x, y-1, mod)) % mod
	}
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
