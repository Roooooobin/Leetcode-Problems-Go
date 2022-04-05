package main

func minimizedMaximum(n int, quantities []int) int {

	lo, hi := 1, 200000
	for lo <= hi {
		mid := lo + (hi-lo)>>1
		cnt := 0
		for _, q := range quantities {
			cnt += q/mid + 1
			if q%mid == 0 {
				cnt--
			}
			if cnt > n {
				break
			}
		}
		if cnt <= n {
			hi = mid - 1
		} else {
			lo = mid + 1
		}
	}
	return lo
}
