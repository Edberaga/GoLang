package main

import "fmt"

type triangle struct {
	base float64
	height float64
}

func(t triangle) area() float64 {
	return 0.5 * t.base * t.height
}

func(t triangle) increaseSize(n float64) {
	t.base += n
	t.height += n
	fmt.Printf("Size has been increased for %f.2\nTriangle size now: \nbase: %f.2\nheight: %f.2\n", n, t.base, t.height)
}

func main() {
	triangle1 := triangle{10, 11}
	area := triangle1.area()
	fmt.Println("Area: ", area)

	triangle1.increaseSize(2)
	fmt.Println("triangle now: ", triangle1.base, triangle1.height)
}