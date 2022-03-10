package main

func countHighestScoreNodes(parents []int) int {

	n := len(parents)
	outDegree := make([]int, n)
	for _, parent := range parents {
		if parent == -1 {
			continue
		}
		outDegree[parent]++
	}
	queue := make([]int, 0)
	for i := 0; i < n; i++ {
		if outDegree[i] == 0 {
			queue = append(queue, i)
		}
	}
	childrenCount := make([][]int, n)
	for i := 0; i < n; i++ {
		childrenCount[i] = make([]int, 2)
	}
	for len(queue) > 0 {
		front := queue[0]
		queue = queue[1:]
		parent := parents[front]
		if parent == -1 {
			continue
		}
		outDegree[parent]--
		childrenCount[parent][outDegree[parent]] = childrenCount[front][0] + childrenCount[front][1] + 1
		if outDegree[parent] == 0 {
			queue = append(queue, parent)
		}
	}
	//for i, v := range childrenCount {
	//	fmt.Printf("%d %d\n", i, v)
	//}

	max := 0
	val := make([]int, n)
	for i := 0; i < n; i++ {
		left, right, par := 1, 1, 1
		if childrenCount[i][0] != 0 {
			left = childrenCount[i][0]
		}
		if childrenCount[i][1] != 0 {
			right = childrenCount[i][1]
		}
		parent := parents[i]
		if parent != -1 {
			par = n - childrenCount[i][0] - childrenCount[i][1] - 1
		}
		v := left * right * par
		val[i] = v
		if v > max {
			max = v
		}
	}
	res := 0
	for _, v := range val {
		if v == max {
			res++
		}
	}
	return res
}

func main() {

	print(countHighestScoreNodes([]int{-1, 2, 0, 2, 0}))
	print(countHighestScoreNodes([]int{-1, 2, 0}))
}
