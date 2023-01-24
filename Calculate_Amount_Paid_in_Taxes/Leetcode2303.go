package main

func calculateTax(brackets [][]int, income int) float64 {

	lower := 0
	total := 0
	for _, b := range brackets {
		upper, percent := b[0], b[1]
		tax := (min(income, upper) - lower) * percent
		total += tax
		if income <= upper {
			break
		}
		lower = upper
	}
	return float64(total) / 100
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
