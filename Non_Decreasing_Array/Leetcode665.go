package main

func checkPossibility(nums []int) bool {

	cnt := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] < nums[i-1] {
			if i > 1 && nums[i-2] > nums[i] {
				nums[i] = nums[i-1]
			}
			cnt++
		}
		if cnt > 1 {
			return false
		}
	}
	return true
}

/*
逻辑题
找到第一组降序对nums[i-1], nums[i], nums[i] < nums[i-1], 如果存在nums[i-2]且nums[i-2]>nums[i],那么nums[i]至少得=nums[i-1]
才能保持升序, 修改后如果还出现了降序对, 那么无法再维持整体升序了
*/
