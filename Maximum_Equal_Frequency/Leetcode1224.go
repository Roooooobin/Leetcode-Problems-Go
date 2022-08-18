package main

func maxEqualFreq(nums []int) int {

	frequencyMap := make(map[int]int)
	frequencyCountMap := make(map[int]int)
	res := 1
	maxFreq := 0
	for i, num := range nums {
		v := frequencyMap[num]
		if v > 0 {
			frequencyCountMap[v]--
		}
		frequencyMap[num] = v + 1
		maxFreq = max(maxFreq, v+1)
		fq := frequencyCountMap[v+1]
		frequencyCountMap[v+1] = fq + 1
		if maxFreq == 1 ||
			frequencyCountMap[maxFreq]*maxFreq+frequencyCountMap[maxFreq-1]*(maxFreq-1) == i+1 &&
				frequencyCountMap[maxFreq] == 1 ||
			frequencyCountMap[maxFreq]*maxFreq == i && frequencyCountMap[1] == 1 {
			res = i + 1
		}

	}
	return res
}

func max(x, y int) int {

	if x > y {
		return x
	}
	return y
}

/*
- -哈希表+逻辑
记录每个num的频率以及频率的count
*/
