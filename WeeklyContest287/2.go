package main

import "sort"

func findWinners(matches [][]int) [][]int {

	lossCount := make(map[int]int)
	playersSet := make(map[int]struct{})
	for _, match := range matches {
		if _, ok := lossCount[match[1]]; !ok {
			lossCount[match[1]] = 1
		} else {
			lossCount[match[1]]++
		}
		if _, ok := playersSet[match[0]]; !ok {
			playersSet[match[0]] = struct{}{}
		}
		if _, ok := playersSet[match[1]]; !ok {
			playersSet[match[1]] = struct{}{}
		}
	}
	players := make([]int, 0)
	for x := range playersSet {
		players = append(players, x)
	}
	sort.Ints(players)
	answer1 := make([]int, 0)
	answer2 := make([]int, 0)
	for _, player := range players {
		if _, ok := lossCount[player]; !ok {
			answer1 = append(answer1, player)
		} else {
			if lossCount[player] == 1 {
				answer2 = append(answer2, player)
			}
		}
	}
	return [][]int{answer1, answer2}
}
