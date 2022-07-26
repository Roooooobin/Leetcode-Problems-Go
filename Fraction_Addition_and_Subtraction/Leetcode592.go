package main

import (
	"fmt"
	"math/big"
	"strings"
)

func fractionAddition(expression string) string {
	in := strings.NewReader(expression)
	res := big.NewRat(0, 1)
	var a, b int64
	for {
		if n, _ := fmt.Fscanf(in, "%d/%d", &a, &b); n == 0 {
			return res.String()
		}
		res.Add(res, big.NewRat(a, b))
	}
}

//作者：endlesscheng
//链接：https://leetcode.cn/problems/fraction-addition-and-subtraction/solution/go-ku-han-shu-by-endlesscheng-yhwn/
//来源：力扣（LeetCode）
//著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
