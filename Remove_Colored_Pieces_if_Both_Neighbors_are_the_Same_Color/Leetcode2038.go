package main

func winnerOfGame(colors string) bool {

	a, b := 0, 0
	n := len(colors)
	cur := 0
	i := 0
	for i < n {
		cur = 0
		for i < n && colors[i] == 'A' {
			i++
			cur++
		}
		if cur >= 3 {
			a += cur - 2
		}
		i++
	}
	cur, i = 0, 0
	for i < n {
		cur = 0
		for i < n && colors[i] == 'B' {
			i++
			cur++
		}
		if cur >= 3 {
			b += cur - 2
		}
		i++
	}
	return a > b
}
