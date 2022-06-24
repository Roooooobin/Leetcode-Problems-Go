package main

func longestSubstring(s string, k int) int {

	res := longestSubstringHelper(s, k)
	return res
}

func longestSubstringHelper(s string, k int) int {

	n := len(s)
	if k > n {
		return 0
	}
	res := 0
	hash := make([]int, 26)
	for _, c := range s {
		hash[c-'a']++
	}
	idxFreqSmallerThanK := make([]int, 0)
	for i, c := range s {
		if hash[c-'a'] < k {
			idxFreqSmallerThanK = append(idxFreqSmallerThanK, i)
		}
	}
	//fmt.Println(s)
	if len(idxFreqSmallerThanK) == 0 {
		return n
	} else if len(idxFreqSmallerThanK) == n {
		return 0
	}
	idxFreqSmallerThanK = append(idxFreqSmallerThanK, n)
	//fmt.Println(idxFreqSmallerThanK)
	if checkK(s[0:idxFreqSmallerThanK[0]], k) {
		res = idxFreqSmallerThanK[0]
	} else {
		res = longestSubstringHelper(s[0:idxFreqSmallerThanK[0]], k)
	}
	for i := 0; i < len(idxFreqSmallerThanK)-1; i++ {
		st := idxFreqSmallerThanK[i] + 1
		ed := idxFreqSmallerThanK[i+1]
		if st == ed-1 || ed-st < k {
			continue
		}
		if checkK(s[st:ed], k) {
			res = max(res, ed-st)
		} else {
			res = max(res, longestSubstringHelper(s[st:ed], k))
		}
	}
	return res
}

func checkK(s string, k int) bool {

	hash := make([]int, 26)
	for _, c := range s {
		hash[c-'a']++
	}
	for _, v := range hash {
		if v > 0 && v < k {
			return false
		}
	}
	return true
}

//func max(x, y int) int {
//	if x > y {
//		return x
//	} else {
//		return y
//	}
//}
