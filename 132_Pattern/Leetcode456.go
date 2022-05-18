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
