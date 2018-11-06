package object_oriented

import "fmt"

type person struct {
	name string
	age  int
}

type rect struct {
	width, height int
}

func (r *rect) area() int {
	return r.height * r.width
}

func (r rect) perim() int {
	return 2*r.width + 2*r.height
}

func main() {
	fmt.Println(person{"Bob", 20})

	fmt.Println(person{name: "Alice", age: 20})

	fmt.Println(person{age: 20})

	fmt.Println(&person{name: "Ann", age: 30})

	s := person{name: "Sean", age: 50}

	sp := &s

	fmt.Println(sp.age)
	sp.age = 51
	fmt.Println(sp.age)

	r := rect{width: 20, height: 10}
	fmt.Println("area:", r.area())
	fmt.Println("perim:", r.perim())

	rp := &r
	fmt.Println("area:", rp.area())
	fmt.Println("perim:", rp.perim())

}
