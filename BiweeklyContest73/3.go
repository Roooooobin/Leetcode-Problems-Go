package main

import "sort"

func getAncestors(n int, edges [][]int) [][]int {

	inDegrees := make([]int, n)
	child := make([][]int, n)
	for i := range child {
		child[i] = make([]int, 0)
	}
	for _, edge := range edges {
		inDegrees[edge[1]]++
		child[edge[0]] = append(child[edge[0]], edge[1])
	}
	res := make([][]int, n)
	for i := range res {
		res[i] = make([]int, 0)
	}
	queue := make([]int, 0)
	for i, v := range inDegrees {
		if v == 0 {
			queue = append(queue, i)
		}
	}

	hashMap := make([][]int, n)
	for i := range hashMap {
		hashMap[i] = make([]int, n)
	}
	for len(queue) > 0 {
		anc := queue[0]
		queue = queue[1:]
		for _, v := range child[anc] {
			if hashMap[v][anc] == 0 {
				res[v] = append(res[v], anc)
				hashMap[v][anc]++
			}
			for _, x := range res[anc] {
				if hashMap[v][x] == 0 {
					res[v] = append(res[v], x)
				}
				hashMap[v][x]++
			}
			inDegrees[v]--
			if inDegrees[v] == 0 {
				queue = append(queue, v)
			}
		}
	}
	for i, l := range res {
		sort.Ints(l)
		res[i] = l
	}
	return res
}
