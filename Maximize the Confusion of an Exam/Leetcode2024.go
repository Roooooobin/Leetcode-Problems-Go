package main

func maxConsecutiveAnswers(answerKey string, k int) int {

	// 计算窗口中为c的字符个数，向右扩展，如果超过k，左边界向右收直到满足条件
	slidingWindow := func(c byte, k int) int {
		res := 0
		n := len(answerKey)
		cnt := 0
		for l, r := 0, 0; r < n; r++ {
			if answerKey[r] == c {
				cnt++
			}
			for cnt > k {
				if answerKey[l] == c {
					cnt--
				}
				l++
			}
			res = max(res, r-l+1)
		}
		return res
	}
	return max(slidingWindow('T', k), slidingWindow('F', k))
}

func max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}

func main() {
	print(maxConsecutiveAnswers("TFFT", 1))
}
