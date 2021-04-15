package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	pic := make([][]uint8, dx)
	for i := 0; i < dx; i++ {
		for j := 0; j < dy; j++ {
			pic[i] = append(pic[i], uint8(i*j))
		}
	}
	return pic
}

func main() {
	pic.Show(Pic)
}
