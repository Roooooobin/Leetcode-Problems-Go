package main

func wordCount(startWords []string, targetWords []string) int {

	calcMask := func(s string) int {
		res := 0
		for _, c := range s {
			res |= 1 << (c - 'a')
		}
		return res
	}
	hash := make(map[int]int)
	seen := make(map[int]bool)
	for _, w := range targetWords {
		mask := calcMask(w)
		v := hash[mask]
		hash[mask] = v + 1
	}
	res := 0
	for _, w := range startWords {
		mask := calcMask(w)
		for k := 0; k < 26; k++ {
			x := 1 << k
			if mask&x == 0 {
				mask ^= x
				if v, ok := hash[mask]; ok && !seen[mask] {
					seen[mask] = true
					res += v
				}
				mask ^= x
			}
		}
	}
	return res
}

/*
- -状态压缩+哈希打表
标准仅含小写字母的状态压缩. 其实从target反向删减一个字母会更好
*/
