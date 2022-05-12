package main

func oneEditAway(first string, second string) bool {

	m, n := len(first), len(second)
	if m-n > 1 || n-m > 1 {
		return false
	}
	lenDiff := m - n
	diff := 1
	for i, j := 0, 0; i < m && j < n; i++ {
		if first[i] != second[j] {
			if lenDiff == 1 {
				j--
			} else if lenDiff == -1 {
				i--
			}
			diff--
		}
		if diff < 0 {
			return false
		}
		j++
	}
	return true
}
