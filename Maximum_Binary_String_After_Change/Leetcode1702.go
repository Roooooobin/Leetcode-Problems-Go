package main

func maximumBinaryString(binary string) string {

	bytes := make([]byte, len(binary))
	for i := range bytes {
		bytes[i] = '1'
	}
	ones, zeros := 0, 0
	for _, c := range binary {
		if c == '0' {
			zeros++
		} else if zeros == 0 {
			ones++ // 前置1
		}
	}
	// 1往后赶,然后把前面的0换成1,最后把ones+zeros-1的位置设置为0就可以
	if ones < len(binary) {
		bytes[ones+zeros-1] = '0'
	}
	return string(bytes)
}

/*
https://leetcode.com/problems/maximum-binary-string-after-change/discuss/987335/JavaC%2B%2BPython-Solution-with-Explanation
- -贪心
在前面的1不用动, 然后把后面(从第一个0出现之后的)1往最后赶(10->01), 最后把中间的0(00->10)得到11111...0
*/
