package main

import (
	"errors"
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return math.NaN(), errors.New(ErrNegativeSqrt(x).Error())
	}
	result := x
	for i := 0; i < 10; i++ {
		result = result - ((result*result - x) / (2 * result))
	}
	return result, error(nil)
}

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}
