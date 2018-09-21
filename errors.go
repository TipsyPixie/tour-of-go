package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) String() string {
	return "TmpString"
}

func (e ErrNegativeSqrt) Error() string {
    /*
     * Calling fmt.Println() or fmt.Sprint() causes infinite loop here.
     * This is because fmt.print() tries to convert e into string by calling Error().
     * In a similar manner, this infinite loop issue occurs upon calling fmt.Sprint() or fmt.Printf() with %s format.
     */
	return fmt.Sprintf("Cannot Sqrt negative number: %.5f", e)
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}

	return math.Sqrt(x), nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}

