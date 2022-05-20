package main

func maxCompatibilitySum(students [][]int, mentors [][]int) int {

	res := 0
	m := len(students)
	indexes := make([]int, m)
	for i := 0; i < m; i++ {
		indexes[i] = i
	}
	permutations := getPermutations(indexes)
	for _, permutation := range permutations {
		cur := 0
		idx := 0
		for i := range permutation {
			for j, stu := range students[permutation[i]] {
				if stu == mentors[idx][j] {
					cur++
				}
			}
			idx++
		}
		if cur > res {
			res = cur
		}
	}
	return res
}

func getPermutations(arr []int) [][]int {
	var helper func([]int, int)
	res := [][]int{}

	helper = func(arr []int, n int) {
		if n == 1 {
			tmp := make([]int, len(arr))
			copy(tmp, arr)
			res = append(res, tmp)
		} else {
			for i := 0; i < n; i++ {
				helper(arr, n-1)
				if n%2 == 1 {
					tmp := arr[i]
					arr[i] = arr[n-1]
					arr[n-1] = tmp
				} else {
					tmp := arr[0]
					arr[0] = arr[n-1]
					arr[n-1] = tmp
				}
			}
		}
	}
	helper(arr, len(arr))
	return res
}

func main() {
	maxCompatibilitySum([][]int{{1, 1, 0}, {1, 0, 1}, {0, 0, 1}}, [][]int{{1, 0, 0}, {0, 0, 1}, {1, 1, 0}})
}
