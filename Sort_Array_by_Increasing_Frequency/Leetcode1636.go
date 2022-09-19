package main

import (
	"sort"
	"strings"
)

func frequencySort(nums []int) []int {

	hash := make(map[int]int)
	for _, num := range nums {
		hash[num]++
	}
	sort.Slice(nums, func(i, j int) bool {
		x, y := nums[i], nums[j]
		if hash[x] == hash[y] {
			return x > y
		}
		return hash[x] < hash[y]
	})
	return nums
}

func findDuplicate(paths []string) [][]string {
	mp := make(map[string][]string)
	for _, v := range paths {
		file := strings.Split(v, " ")
		for i := 1; i < len(file); i++ {
			if len(file[i]) > 1 {
				content := strings.Split(file[i], "(")
				txt := content[1][:len(content[1])-1]
				if mp[txt] == nil {
					mp[txt] = make([]string, 0)
				}
				mp[txt] = append(mp[txt], file[0]+"/"+content[0])
			}
		}
	}
	res := make([][]string, 0)
	for _, v := range mp {
		if len(v) > 1 {
			res = append(res, v)
		}
	}
	return res
}
