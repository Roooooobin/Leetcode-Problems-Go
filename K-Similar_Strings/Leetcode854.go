package main

func kSimilarity(s1 string, s2 string) (res int) {

	type pair struct {
		s string
		i int
	}
	queue := []pair{{s1, 0}}
	seen := map[string]bool{s1: true}
	n := len(s1)
	for ; ; res++ {
		length := len(queue)
		for l := 0; l < length; l++ {
			front := queue[0]
			queue = queue[1:]
			s, i := front.s, front.i
			if s == s2 {
				return
			}
			for i < n && s[i] == s2[i] {
				i++
			}
			ss := []byte(s)
			for j := i + 1; j < n; j++ {
				if s[j] == s2[i] && s[j] != s2[j] {
					ss[i], ss[j] = ss[j], ss[i]
					if t := string(ss); !seen[t] {
						seen[t] = true
						queue = append(queue, pair{t, i + 1})
					}
					ss[i], ss[j] = ss[j], ss[i]
				}
			}
		}
	}
}

/*
- -BFS+剪枝
当枚举的位置j与目标相等时不交换(ss[j] == s2[j]), 因为到后面总会换的, 避免重复
*/
