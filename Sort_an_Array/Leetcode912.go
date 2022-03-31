package main

func sortArray(nums []int) []int {

	quicksort(nums, 0, len(nums)-1)
	return nums
}

func quicksort(a []int, l, r int) {

	if l >= r {
		return
	}
	m := partition(a, l, r)
	quicksort(a, l, m-1)
	quicksort(a, m+1, r)
}

func partition(a []int, l, r int) int {

	x := a[l]
	i, j := l, r
	for i < j {
		for i < j && a[j] >= x {
			j--
		}
		if i < j {
			a[i] = a[j]
			i++
		}
		for i < j && a[i] <= x {
			i++
		}
		if i < j {
			a[j] = a[i]
			j--
		}
	}
	a[i] = x
	return i
}

func main() {

	a := []int{1, 4, 7, 5, 6, 3, 2}
	partition(a, 0, len(a)-1)
	for i := range a {
		print(a[i])
	}
}
