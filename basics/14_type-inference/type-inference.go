package main

import "fmt"

// When declaring a variable without specifying an explicit type (either by using the := syntax or var = expression syntax),
// the variable's type is inferred from the value on the right hand side.

func main() {
	v := 42 // change me!
	fmt.Printf("vis of type %T\n", v)
}
