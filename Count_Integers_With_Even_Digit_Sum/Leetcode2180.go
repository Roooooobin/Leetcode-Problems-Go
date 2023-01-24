package main

func countEven(num int) (res int) {

	for i := 2; i <= num; i++ {
		sum := 0
		x := i
		for x > 0 {
			sum += x % 10
			x /= 10
		}
		if sum&1 == 0 {
			res++
		}
	}
	return
}
