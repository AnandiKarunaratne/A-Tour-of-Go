package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	switch {
	case x == 0 || math.IsNaN(x) || math.IsInf(x, 1):
		return x
	case x < 0:
		return math.NaN()
	}

	var z float64 = 1
	zPrev := z
	isSame := false
	for !isSame {
		zPrev = z
		z -= (z*z - x) / (2 * z)
		if z == zPrev {
			isSame = true
		}
	}
	return z
}

func main() {
	fmt.Println(Sqrt(0))
}
