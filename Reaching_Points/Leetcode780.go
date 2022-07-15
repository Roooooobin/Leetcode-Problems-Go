package main

//https://leetcode-cn.com/problems/reaching-points/solution/dao-da-zhong-dian-by-leetcode-solution-77fo/
func reachingPoints(sx int, sy int, tx int, ty int) bool {

	for tx > sx && ty > sy && tx != ty {
		if tx > ty {
			tx %= ty
		} else {
			ty %= tx
		}
	}
	if tx == sx && ty == sy {
		return true
	} else if tx == sx {
		return ty > sy && (ty-sy)%tx == 0
	} else if ty == sy {
		return tx > sx && (tx-sx)%ty == 0
	}
	return false
}

/*
- -反向
*/
