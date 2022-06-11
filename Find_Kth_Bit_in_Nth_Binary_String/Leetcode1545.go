package main

func findKthBit(n int, k int) byte {

	if n == 1 {
		return '0'
	}
	length := (1 << n) - 1
	if k == length/2+1 {
		return '1'
	} else if k > length/2+1 {
		c := findKthBit(n-1, length-k+1)
		if c == '0' {
			return '1'
		} else {
			return '0'
		}
	} else {
		return findKthBit(n-1, k)
	}
}
