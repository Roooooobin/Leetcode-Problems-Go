package main

// space: O(n)
func duplicateZerosSpaceN(arr []int) {

	n := len(arr)
	a := make([]int, 0)
	for i := range arr {
		if arr[i] == 0 {
			a = append(a, 0)
			a = append(a, 0)
		} else {
			a = append(a, arr[i])
		}
		if len(a) >= n {
			break
		}
	}
	for i := range arr {
		arr[i] = a[i]
	}
}

// space: O(1)
func duplicateZerosBetter(arr []int) {

	n := len(arr)
	count := 0
	for i := range arr {
		if arr[i] == 0 {
			count++
		}
	}
	for i := n - 1; i >= 0; i-- {
		if arr[i] == 0 {
			count--
		}
		if i+count < n {
			arr[i+count] = arr[i]
		}
		if arr[i] == 0 && i+count+1 < n {
			arr[i+count+1] = 0
		}
	}
}
