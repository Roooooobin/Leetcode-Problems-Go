package main

func longestPrefix(s string) string {

	var l, r int64
	var p int64 = 1
	MOD := int64(1e9 + 7)
	k := 0
	for i := 0; i < len(s)-1; i++ {
		l = (l*128 + int64(s[i])) % MOD
		r = (r + p*int64(s[len(s)-i-1])) % MOD
		if l == r {
			k = i + 1
		}
		p = p * int64(128) % MOD
	}
	return s[:k]
}

/*
- -滚动哈希
分别计算前缀和后缀的哈希值然后比较
*/
