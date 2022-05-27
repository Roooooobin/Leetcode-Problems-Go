package main

import "fmt"

type Number interface {
	int | int64 | float64
}

func SumNumbers[K comparable, V Number](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}

type Addable interface {
	int | int64 | float64 | string
}

func SumTwoNumbers[T Addable](a, b T) T {
	return a + b
}

func main() {
	fmt.Println(SumTwoNumbers(1, 2))
	fmt.Println(SumTwoNumbers(1.4, 2.2))
	fmt.Println(SumTwoNumbers("1.4", "2.2"))
}
