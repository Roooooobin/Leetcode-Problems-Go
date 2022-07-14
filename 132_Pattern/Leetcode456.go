package main

func find132pattern(nums []int) bool {

	// 枚举i，单调栈存j，维护最大k
	stack := make([]int, 0)
	maxK := -1000000007
	n := len(nums)
	for i := n - 1; i >= 0; i-- {
		if nums[i] < maxK {
			return true
		}
		for len(stack) > 0 && nums[i] > stack[len(stack)-1] {
			maxK = max(maxK, stack[len(stack)-1])
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, nums[i])
	}
	return false
}

func max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}

func main() {
	print(find132pattern([]int{-2, 1, 1, -2, 1, 1}))
}

/*
https://leetcode.cn/problems/132-pattern/solution/xiang-xin-ke-xue-xi-lie-xiang-jie-wei-he-95gt/
- -单调栈
枚举i，单调栈存j，维护最大k
维护单调递减栈, 保证有一个(j,k)对且j>k, 因为k是从栈中弹出来的, 如果k大于当前的i, 表示找到了132
*/
