package main

func nthMagicalNumber(n int, a int, b int) int {

	MOD := int64(1e9 + 7)
	lcs := a / gcd(a, b) * b
	lo, hi := int64(0), int64(1e15)
	for lo <= hi {
		mi := lo + (hi-lo)/2
		cur := mi/int64(a) + mi/int64(b) - mi/int64(lcs)
		if cur < int64(n) {
			lo = mi + 1
		} else {
			hi = mi - 1
		}
	}
	return int(lo % MOD)
}

func gcd(x, y int) int {
	if y == 0 {
		return x
	}
	return gcd(y, x%y)
}
