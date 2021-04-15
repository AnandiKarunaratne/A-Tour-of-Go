package main

import "fmt"

// A slice literal is like an array literal without the length.

func main() {
	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)

	// [6]bool{true, false, true, true, false, true}
	// This creates the same array as above,
	// then builds a slice that references it:
	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)

	s := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(s)
}
