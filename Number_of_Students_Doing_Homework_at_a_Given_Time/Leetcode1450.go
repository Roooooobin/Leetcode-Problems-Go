package main

func busyStudent(startTime []int, endTime []int, queryTime int) int {

	res := 0
	for i, st := range startTime {
		if st <= queryTime && queryTime <= endTime[i] {
			res++
		}
	}
	return res
}
