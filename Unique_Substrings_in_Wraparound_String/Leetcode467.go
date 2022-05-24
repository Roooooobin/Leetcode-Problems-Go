package main

// WA
//func findSubstringInWraproundString(p string) int {
//
//	n := len(p)
//	res := 0
//	i := 0
//	seen := make(map[string]struct{})
//	for i < n {
//		j := i + 1
//		cur := 0
//		var builder strings.Builder
//		for ; j < n; j++ {
//			if (p[j]-p[j-1]+26)%26 == 1 {
//				builder.WriteString(string(p[j]))
//				s := builder.String()
//				if _, ok := seen[s]; !ok {
//					cur++
//					seen[s] = struct{}{}
//					res += cur
//				}
//			} else {
//				break
//			}
//		}
//		i = j
//	}
//	seen = make(map[string]struct{})
//	for j := 0; j < n; j++ {
//		s := string(p[j])
//		if _, ok := seen[s]; !ok {
//			res++
//			seen[s] = struct{}{}
//		}
//	}
//	return res
//}

func findSubstringInWraproundString(p string) int {

	// count the max length of each letter as the end
	count := make([]int, 26)
	n := len(p)
	i := 0
	for i < n {
		count[p[i]-'a'] = max(count[p[i]-'a'], 1)
		j := i + 1
		for ; j < n; j++ {
			if (p[j]-p[j-1]+26)%26 == 1 {
				count[p[j]-'a'] = max(count[p[j]-'a'], j-i+1)
			} else {
				break
			}
		}
		i = j
	}
	res := 0
	for _, v := range count {
		res += v
	}
	return res
}

func max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}

func main() {
	print(findSubstringInWraproundString("zabc"))
}
