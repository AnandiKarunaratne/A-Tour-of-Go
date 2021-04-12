package main

import (
	"fmt"
	"math"
)

// In Go, a name is exported if it begins with a capital letter.
// For example, Pi is exported from the math package.
// pi does not start with a capital letter, so it is not exported.
// When importing a package, you can refer only to its exported names.
// Any "unexported" names are not accessible from outside the package.

func main() {
	fmt.Println(math.Pi)
}
