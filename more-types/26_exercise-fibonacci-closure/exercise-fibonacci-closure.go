package main

import "fmt"

// fibonacci is a function that returns a function that returns an int.
func fibonacci() func() int {
	turn := 0
	prev1 := 0
	prev2 := 0
	sum := 0
	return func() int {
		if turn <= 1 {
			sum = 1
		} else {
			sum = prev1 + prev2
		}
		turn++
		prev1 = prev2
		prev2 = sum
		return sum

	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
