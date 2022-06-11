package main

func minimumCardPickup(cards []int) int {

	hash := make(map[int]int)
	res := len(cards) + 1
	for i, card := range cards {
		if v, ok := hash[card]; ok {
			res = min(res, i-v+1)
		}
		hash[card] = i
	}
	if res == len(cards)+1 {
		res = -1
	}
	return res
}

func min(x, y int) int {
	if x < y {
		return x
	} else {
		return y
	}
}
