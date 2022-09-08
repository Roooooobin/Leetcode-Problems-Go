package main

func constructArray(n int, k int) []int {

	res := make([]int, 0, n)
	for i := 1; i < n-k; i++ {
		res = append(res, i)
	}
	for i, j := n-k, n; i <= j; i++ {
		res = append(res, i)
		if i != j {
			res = append(res, j)
		}
		j--
	}
	return res
}

/*
- -模拟, 数学
先排n-k-1个数出n-k-2个差为1的数列, 然后交错剩下的k+1排出k-1个不同的差的数列(包括1)
*/
