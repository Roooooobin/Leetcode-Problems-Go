package main

import (
	"math"
	"math/rand"
	"time"
)

func randInt(a, b int) int {
	return a + rand.Intn(b-a)
}

// quick power
func pow(x, n, mod int) int {
	res := 1
	for ; n > 0; n >>= 1 {
		if n&1 > 0 {
			res = res * x % mod
		}
		x = x * x % mod
	}
	return res
}

func check(arr []byte, m, a1, a2, mod1, mod2 int) int {
	aL1, aL2 := pow(a1, m, mod1), pow(a2, m, mod2)
	h1, h2 := 0, 0
	for _, c := range arr[:m] {
		h1 = (h1*a1 + int(c)) % mod1
		h2 = (h2*a2 + int(c)) % mod2
	}
	seen := map[[2]int]bool{{h1, h2}: true}
	for start := 1; start <= len(arr)-m; start++ {
		h1 = (h1*a1 - int(arr[start-1])*aL1 + int(arr[start+m-1])) % mod1
		if h1 < 0 {
			h1 += mod1
		}
		h2 = (h2*a2 - int(arr[start-1])*aL2 + int(arr[start+m-1])) % mod2
		if h2 < 0 {
			h2 += mod2
		}
		if seen[[2]int{h1, h2}] {
			return start
		}
		seen[[2]int{h1, h2}] = true
	}
	return -1
}

func longestDupSubstring(s string) string {

	rand.Seed(time.Now().UnixNano())
	a1, a2 := randInt(26, 100), randInt(26, 100)
	mod1, mod2 := randInt(1e9+7, math.MaxInt32), randInt(1e9+7, math.MaxInt32)
	arr := []byte(s)
	for i := range arr {
		arr[i] -= 'a'
	}
	lo, hi := 1, len(s)-1
	length, start := 0, -1
	for lo <= hi {
		mid := lo + ((hi - lo) >> 1)
		idx := check(arr, mid, a1, a2, mod1, mod2)
		if idx != -1 {
			lo = mid + 1
			length = mid
			start = idx
		} else {
			hi = mid - 1
		}
	}
	if start == -1 {
		return ""
	}
	return s[start : start+length]
}
