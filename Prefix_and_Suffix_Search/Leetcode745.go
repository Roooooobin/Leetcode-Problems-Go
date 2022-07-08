package main

type WordFilter struct {
	nodes  [27]*WordFilter
	weight int
}

func Constructor(words []string) WordFilter {
	wf := WordFilter{
		[27]*WordFilter{},
		0,
	}
	for w, word := range words {
		k := word + "{" + word
		for i := 0; i < len(word); i++ {
			wf.insert(k[i:], w)
		}
	}
	return wf
}

func (this *WordFilter) insert(word string, weight int) {
	root := this
	for _, c := range word {
		x := int(c - 'a')
		if root.nodes[x] == nil {
			root.nodes[x] = &WordFilter{
				[27]*WordFilter{},
				weight,
			}
		}
		root = root.nodes[x]
		root.weight = weight
	}
}

func (this *WordFilter) find(word string) int {
	root := this
	for _, c := range word {
		x := int(c - 'a')
		if root.nodes[x] == nil {
			return -1
		}
		root = root.nodes[x]
	}
	return root.weight
}

func (this *WordFilter) F(prefix string, suffix string) int {

	k := suffix + "{" + prefix
	return this.find(k)
}

/*
https://leetcode.cn/problems/prefix-and-suffix-search/solution/go-200can-kao-by-takagi-san/
- -字典树
存word{word的后缀, 用suffix{prefix去查找
*/
