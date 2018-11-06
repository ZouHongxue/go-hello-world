package object_oriented

import (
	"fmt"
	"math"
)

//接口
type geometry interface {
	area() float64
	perim() float64
}

//类
type square struct {
	width, height float64
}

type circle struct {
	radius float64
}

//实现
func (s square) area() float64 {
	return s.height * s.height
}

func (s square) perim() float64 {
	return 2 * (s.width + s.height)
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func main() {
	s := square{width: 3, height: 4}
	c := circle{radius: 5}

	measure(s)
	measure(c)
}
