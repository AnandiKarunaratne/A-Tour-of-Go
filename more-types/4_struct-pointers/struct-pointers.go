package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{1, 2}
	// Struct fields can be accessed through a struct pointer.
	p := &v
	p.X = 1e9 // go permits this instead of dereference = (*p).X
	fmt.Println(v)
}
