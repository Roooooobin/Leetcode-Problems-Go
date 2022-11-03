package main

func sumSubarrayMins(arr []int) (res int) {

	const MOD = 1e9 + 7
	n := len(arr)
	left := make([]int, n)
	right := make([]int, n)
	var monoStack []int
	for i, x := range arr {
		for len(monoStack) > 0 && x <= arr[monoStack[len(monoStack)-1]] {
			monoStack = monoStack[:len(monoStack)-1]
		}
		if len(monoStack) == 0 {
			left[i] = i + 1
		} else {
			left[i] = i - monoStack[len(monoStack)-1]
		}
		monoStack = append(monoStack, i)
	}
	monoStack = []int{}
	for i := n - 1; i >= 0; i-- {
		for len(monoStack) > 0 && arr[i] < arr[monoStack[len(monoStack)-1]] {
			monoStack = monoStack[:len(monoStack)-1]
		}
		if len(monoStack) == 0 {
			right[i] = n - i
		} else {
			right[i] = monoStack[len(monoStack)-1] - i
		}
		monoStack = append(monoStack, i)
	}
	for i, x := range arr {
		res = (res + left[i]*right[i]*x) % MOD
	}
	return
}

//
//作者：LeetCode-Solution
//链接：https: //leetcode.cn/problems/sum-of-subarray-minimums/solution/zi-shu-zu-de-zui-xiao-zhi-zhi-he-by-leet-bp3k/
//来源：力扣（LeetCode）
//著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
