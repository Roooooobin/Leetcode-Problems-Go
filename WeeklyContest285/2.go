package main

func countCollisions(directions string) int {

	res := 0
	i := 0
	n := len(directions)
	for i < n && directions[i] == 'L' {
		i++
	}
	if i == n {
		return res
	}
	top := directions[i]
	numR := 0
	if directions[i] == 'R' {
		numR++
	}
	for i = i + 1; i < n; i++ {
		if directions[i] == 'L' {
			if top == 'S' {
				res++
			} else if top == 'R' {
				res += 2 + numR - 1
				top = 'S'
				numR = 0
			}
		} else if directions[i] == 'S' {
			if top == 'R' {
				res += numR
				top = 'S'
				numR = 0
			}
		} else if directions[i] == 'R' {
			top = 'R'
			numR++
		}
	}
	return res
}
