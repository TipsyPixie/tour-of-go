package main

import (
	"golang.org/x/tour/pic"
	"math"
)

func Pic(dx, dy int) [][]uint8 {
	table := make([][]uint8, dy)

	for row := 0; row < len(table); row++ {
		table[row] = make([]uint8, dx)
		for column := 0; column < len(table[row]); column++ {
			squareSum := math.Pow(float64(row), 2) + math.Pow(float64(column - 128), 2)
			table[row][column] = 255 - uint8(math.Sqrt(squareSum))
		}
	}

	return table
}

func main() {
	pic.Show(Pic)
}

