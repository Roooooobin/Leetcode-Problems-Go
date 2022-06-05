package main

func waysToBuyPensPencils(total int, cost1 int, cost2 int) int64 {

	res := int64(0)
	max1 := total / cost1
	for i := 0; i <= max1; i++ {
		left := total - i*cost1
		res += int64(left/cost2) + 1
	}
	return res
}
