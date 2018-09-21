package main

import (
	"fmt"
	"math"
)

const initialValue = 1

func Sqrt(x float64) float64 {
	z := float64(initialValue)

	for i := 0; i < 10000; i++ {
		z -= (math.Pow(z, 2) - x) / (2 * z)
	}

	return z
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println("Diff:", math.Abs(math.Sqrt(2) - Sqrt(2)))
}

