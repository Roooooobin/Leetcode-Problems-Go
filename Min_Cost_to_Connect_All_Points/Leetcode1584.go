package main

import "sort"

type Node struct {
	x, y, edge int
}

func minCostConnectPoints(points [][]int) int {

	n := len(points)
	par := make([]int, n)
	for i := range par {
		par[i] = i
	}
	find := func(x int) int {
		r := x
		for r != par[r] {
			r = par[r]
		}
		for x != r {
			tmp := par[x]
			par[x] = r
			x = tmp
		}
		return r
	}
	union := func(x, y int) {
		x = find(x)
		y = find(y)
		if x != y {
			par[x] = y
		}
	}

	nodes := make([]Node, n*(n-1))
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			edge := Abs(points[i][0]-points[j][0]) + Abs(points[i][1]-points[j][1])
			nodes = append(nodes, Node{
				x:    i,
				y:    j,
				edge: edge,
			})
			nodes = append(nodes, Node{
				x:    j,
				y:    i,
				edge: edge,
			})
		}
	}
	sort.Slice(nodes, func(i, j int) bool {
		return nodes[i].edge < nodes[j].edge
	})
	res := 0
	left := n - 1
	for i := 0; i < len(nodes); i++ {
		x := nodes[i].x
		y := nodes[i].y
		if find(x) != find(y) {
			union(x, y)
			res += nodes[i].edge
			left--
		}
		if left == 0 {
			return res
		}
	}
	return res
}

func Abs(x int) int {
	if x < 0 {
		x = -x
	}
	return x
}
