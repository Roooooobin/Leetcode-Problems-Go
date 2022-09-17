package main

type Node struct {
	ch   [26]int
	flag int
}

var trie []Node

func palindromePairs(words []string) [][]int {
	trie = []Node{{[26]int{}, -1}}
	n := len(words)
	for i := 0; i < n; i++ {
		insert(words[i], i)
	}
	var ret [][]int
	for i := 0; i < n; i++ {
		word := words[i]
		m := len(word)
		for j := 0; j <= m; j++ {
			if isPalindrome(word, j, m-1) {
				leftId := findWord(word, 0, j-1)
				if leftId != -1 && leftId != i {
					ret = append(ret, []int{i, leftId})
				}
			}
			if j != 0 && isPalindrome(word, 0, j-1) {
				rightId := findWord(word, j, m-1)
				if rightId != -1 && rightId != i {
					ret = append(ret, []int{rightId, i})
				}
			}
		}
	}
	return ret
}

func insert(s string, id int) {
	add := 0
	for i := 0; i < len(s); i++ {
		x := int(s[i] - 'a')
		if trie[add].ch[x] == 0 {
			trie = append(trie, Node{[26]int{}, -1})
			trie[add].ch[x] = len(trie) - 1
		}
		add = trie[add].ch[x]
	}
	trie[add].flag = id
}

func findWord(s string, left, right int) int {
	add := 0
	for i := right; i >= left; i-- {
		x := int(s[i] - 'a')
		if trie[add].ch[x] == 0 {
			return -1
		}
		add = trie[add].ch[x]
	}
	return trie[add].flag
}

func isPalindrome(s string, left, right int) bool {
	for i := 0; i < (right-left+1)/2; i++ {
		if s[left+i] != s[right-i] {
			return false
		}
	}
	return true
}

/*
- -字典树
前缀后缀
*/
//作者：LeetCode-Solution
//链接：https://leetcode.cn/problems/palindrome-pairs/solution/hui-wen-dui-by-leetcode-solution/
//来源：力扣（LeetCode）
//著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
