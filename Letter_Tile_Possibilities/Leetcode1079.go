package main

// too bad
//func numTilePossibilities(tiles string) int {
//
//	resSet := make(map[string]struct{})
//	for k := 1; k <= len(tiles); k++ {
//		bit := 0
//		dfs(tiles, "", k, resSet, bit)
//	}
//	//for s := range resSet {
//	//	fmt.Println(s)
//	//}
//	return len(resSet) - 1
//}
//
//func dfs(s, cur string, k int, set map[string]struct{}, bit int) {
//
//	if len(cur) > k {
//		return
//	}
//	set[cur] = struct{}{}
//	for i := 0; i < len(s); i++ {
//		if bit&(1<<i) != 0 {
//			continue
//		}
//		bit ^= 1 << i
//		dfs(s, fmt.Sprintf("%s%s", cur, string(rune(int(s[i])))), k, set, bit)
//		bit ^= 1 << i
//	}
//}

func numTilePossibilities(tiles string) int {

	arr := make([]int, 26)
	for _, tile := range tiles {
		arr[tile-'A']++
	}
	var dfs func(a []int) int
	dfs = func(a []int) int {
		sum := 0
		for i := 0; i < 26; i++ {
			if arr[i] == 0 {
				continue
			}
			sum++
			arr[i]--
			sum += dfs(arr)
			arr[i]++
		}
		return sum
	}
	return dfs(arr)
}

func main() {

	res := numTilePossibilities("AAB")
	print(res)
}
