package main

func subarraysWithKDistinct(nums []int, k int) int {

	res := 0
	hash := make(map[int]int)
	n := len(nums)
	for r, l := 0, 0; r < n; r++ {
		v, ok := hash[nums[r]]
		if !ok {
			v = 0
		}
		hash[nums[r]] = v + 1
		for l < r && len(hash) > k {
			if hash[nums[l]] == 1 {
				delete(hash, nums[l])
			} else {
				hash[nums[l]]--
			}
			l++
		}
		if len(hash) == k {
			i := l
			for i <= r {
				if hash[nums[i]] > 1 {
					res++
					hash[nums[i]]--
					i++
				} else {
					res++
					break
				}
			}
			for m := l; m < i; m++ {
				hash[nums[m]]++
			}
		}
	}
	return res
}

/*
better
func subarraysWithKDistinct(nums []int, k int) (ans int) {
    n := len(nums)
    num1 := make([]int, n+1)
    num2 := make([]int, n+1)
    var tot1, tot2, left1, left2 int
    for _, v := range nums {
        if num1[v] == 0 {
            tot1++
        }
        num1[v]++
        if num2[v] == 0 {
            tot2++
        }
        num2[v]++
        for tot1 > k {
            num1[nums[left1]]--
            if num1[nums[left1]] == 0 {
                tot1--
            }
            left1++
        }
        for tot2 > k-1 {
            num2[nums[left2]]--
            if num2[nums[left2]] == 0 {
                tot2--
            }
            left2++
        }
        ans += left2 - left1
    }
    return ans
}

作者：LeetCode-Solution
链接：https://leetcode.cn/problems/subarrays-with-k-different-integers/solution/k-ge-bu-tong-zheng-shu-de-zi-shu-zu-by-l-9ylo/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/
