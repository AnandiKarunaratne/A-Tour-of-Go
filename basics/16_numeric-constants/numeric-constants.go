package main

import "fmt"

// Numeric constants are high-precision values.
// An untyped constant takes the type needed by its context.

const (
	// create huge number by shifting one bit left 100 places
	// ie the binary number that is 1 followed by 100 0s
	Big = 1 << 100

	// shift it again 99 places, so we end up with 1<<1 or 2
	Small = Big >> 99
)

func needInt(x int) int {
	return x*10 + 1
}

func needFloat(x float64) float64 {
	return x * 0.1
}

func main() {
	fmt.Println(needInt(Small))
	// needInt(Big) creates an overflow
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
}
