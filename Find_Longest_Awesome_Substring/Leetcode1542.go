package main

func longestAwesome(s string) int {

	mp := make([]int, 1<<10)
	for i := range mp {
		mp[i] = -2 // 标志位
	}
	mp[0] = -1
	res := 0
	mask := 0
	for i := 0; i < len(s); i++ {
		digit := s[i] - '0'
		mask ^= 1 << digit
		if mp[mask] != -2 {
			res = max(res, i-mp[mask])
		} else {
			mp[mask] = i
		}
		for j := 0; j < 10; j++ {
			if mp[mask^(1<<j)] != -2 {
				res = max(res, i-mp[mask^(1<<j)])
			}
		}
	}
	return res
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

/*
- -前缀和+位运算
将数字出现次数记录在位上(0-9, 十位bit即可存储), 出现奇数次记为1, 偶数次为0, 利用前缀和思想判断改变一位(或一位都不改变)成偶数, 看看是否能行
比如3242415, 32424这个子串记录为0000001000, 那么改变3(用异或)得到0000000000, 表示该子串可以形成回文(最多一个奇数次)
*/
