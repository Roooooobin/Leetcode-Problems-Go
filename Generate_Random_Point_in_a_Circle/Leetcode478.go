package main

import "math/rand"

type Solution struct {
	r, xCenter, yCenter float64
}

func Constructor(radius float64, x_center float64, y_center float64) Solution {
	return Solution{radius, x_center, y_center}
}

func (this *Solution) RandPoint() []float64 {

	for {
		// [-1, 1)
		x := rand.Float64()*2 - 1
		y := rand.Float64()*2 - 1
		if x*x+y*y < 1 {
			return []float64{this.xCenter + this.r*x, this.yCenter + this.r*y}
		}
	}
}
