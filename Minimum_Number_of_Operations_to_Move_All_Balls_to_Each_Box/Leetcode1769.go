package main

func minOperations(boxes string) []int {

	ones := make([]int, 0)
	for i, c := range boxes {
		if c == '1' {
			ones = append(ones, i)
		}
	}
	n := len(ones)
	res := make([]int, len(boxes))
	if n == 0 {
		return res
	}
	prefixSum := make([]int, n+1)
	for i := range ones {
		prefixSum[i+1] = prefixSum[i] + ones[i]
	}
	binarySearch := func(x int) int {
		if x <= ones[0] {
			return 0
		}
		if x > ones[n-1] {
			return n
		}
		lo, hi := 0, n-1
		for lo <= hi {
			mi := lo + (hi-lo)>>1
			if x > ones[mi] {
				lo = mi + 1
			} else if x < ones[mi] {
				hi = mi - 1
			} else {
				return mi
			}
		}
		return lo
	}
	for idx := range boxes {
		i := binarySearch(idx)
		res[idx] = idx*i - prefixSum[i] + prefixSum[n] - prefixSum[i] - idx*(n-i)
	}
	return res
}

func main() {
	minOperations("110")
}
