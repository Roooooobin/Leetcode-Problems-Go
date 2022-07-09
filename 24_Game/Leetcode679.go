package main

const (
	TARGET   = 24
	EPSILON  = 1e-6
	ADD      = 0
	MULTIPLY = 1
	SUBTRACT = 2
	DIVIDE   = 3
)

func judgePoint24(nums []int) bool {
	list := make([]float64, 0)
	for _, num := range nums {
		list = append(list, float64(num))
	}
	return solve(list)
}

func solve(list []float64) bool {

	n := len(list)
	if n == 0 {
		return false
	}
	if n == 1 {
		return abs(list[0]-TARGET) < EPSILON
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i != j {
				list2 := make([]float64, 0)
				for k := 0; k < n; k++ {
					if k != i && k != j {
						list2 = append(list2, list[k])
					}
				}
				for k := 0; k < 4; k++ {
					if k < 2 && i < j {
						continue
					}
					switch k {
					case ADD:
						list2 = append(list2, list[i]+list[j])
					case MULTIPLY:
						list2 = append(list2, list[i]*list[j])
					case SUBTRACT:
						list2 = append(list2, list[i]-list[j])
					case DIVIDE:
						if abs(list[j]) < EPSILON {
							continue
						} else {
							list2 = append(list2, list[i]/list[j])
						}
					}
					if solve(list2) {
						return true
					}
					list2 = list2[:len(list2)-1]
				}
			}
		}
	}
	return false
}

func abs(x float64) float64 {
	if x < 0 {
		return -x
	}
	return x
}

/*
//作者：LeetCode-Solution
//链接：https://leetcode.cn/problems/24-game/solution/24-dian-you-xi-by-leetcode-solution/
//来源：力扣（LeetCode）
//著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
- -回溯
先在四个数字中有序选出2个, 加减乘除中选择一个操作, 取代选择的两个数字之后继续
剩下三个同样选2个, ..., 以此类推
*/
