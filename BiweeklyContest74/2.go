package main

func maximumSubsequenceCount(text string, pattern string) int64 {

	a := make([]int, 0)
	b := make([]int, 0)
	c := make([]int, 0)
	a = append(a, -1)
	for i := range text {
		if text[i] == pattern[0] {
			a = append(a, i)
			c = append(c, i)
		} else if text[i] == pattern[1] {
			b = append(b, i)
		}
	}
	if pattern[0] == pattern[1] {
		length := int64(len(a))
		return length * (length - 1) / 2
	}

	res1 := int64(0)
	res2 := int64(0)
	if len(b) != 0 {
		for _, num := range a {
			res1 += int64(len(b) - binaryFind(b, num))
		}
	}

	b = append(b, len(text))
	for _, num := range c {
		res2 += int64(len(b) - binaryFind(b, num))
	}
	if res1 > res2 {
		return res1
	} else {
		return res2
	}

}

func binaryFind(a []int, x int) int {
	if x < a[0] {
		return 0
	} else if x > a[len(a)-1] {
		return len(a)
	}
	lo, hi := 0, len(a)-1
	for lo <= hi {
		mid := lo + (hi-lo)>>1
		if x > a[mid] {
			lo = mid + 1
		} else {
			hi = mid - 1
		}
	}
	return lo
}

func main() {
	print(maximumSubsequenceCount("aaaaa", "aa"))
}
