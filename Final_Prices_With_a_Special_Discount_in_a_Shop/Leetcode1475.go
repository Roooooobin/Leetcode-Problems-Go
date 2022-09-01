package main

func finalPrices(prices []int) []int {

	stack := make([]int, 0)
	res := make([]int, len(prices))
	copy(res, prices)
	for i, p := range prices {
		for len(stack) != 0 && prices[stack[len(stack)-1]] >= p {
			top := stack[len(stack)-1]
			res[top] = prices[top] - p
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}
	return res
}
