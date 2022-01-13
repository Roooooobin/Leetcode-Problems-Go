package main

func canThreePartsEqualSum(arr []int) bool {

	sum := 0
	for _, x := range arr {
		sum += x
	}
	if sum%3 != 0 {
		return false
	}
	targetSum := sum / 3
	sum = 0
	flag1, flag2 := false, false
	for i, x := range arr {
		sum += x
		if !flag1 && sum == targetSum {
			flag1 = true
		} else if flag1 && sum == 2*targetSum && i != len(arr)-1 {
			flag2 = true
		}
	}
	return flag1 && flag2
}
