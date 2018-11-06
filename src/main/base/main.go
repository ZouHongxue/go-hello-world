package base

import (
	"config"
	"fmt"
)

func main() {
	config.LoadConfig()
	fmt.Println("Hello, World")

	i := 1

	for i <= 3 {
		fmt.Println(i)
		i++
	}

	for j := 1; j <= 9; j++ {
		fmt.Println(j)
	}

	for {
		fmt.Println("loop")
		if a := "loop"; a == "loop" {
			break
		}
	}

	//数组
	var a [5]int
	fmt.Println("emp", a)

	a[4] = 100
	fmt.Println("set:", a)
	fmt.Println("get:", a[4])

	fmt.Println("len:", len(a))

	b := [5]int{1, 2, 3, 4, 5}

	fmt.Println("dcl:", b)

	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)

	//slice
	s := make([]string, 3)
	fmt.Println("emp2:", s)

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])
	fmt.Println("len:", len(s))

	fmt.Println("len:", len(s))

	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd:", s)

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("copy:", c)

	l := s[2:5]
	fmt.Println("sl1:", l)

	l = s[:5]
	fmt.Println("sl2:", l)

	t := []string{"g", "h", "i"}
	fmt.Println("dcl", t)

	//map
	m := make(map[string]int)
	m["k1"] = 7
	m["k2"] = 13
	fmt.Println("map", m)

	v1 := m["k1"]
	fmt.Println("v1:", v1)

	fmt.Println("len", len(m))

	delete(m, "k2")
	fmt.Println("map:", m)

	_, prs := m["k2"]
	fmt.Println("prs", prs)

	n := map[string]int{"foo": 1, "bar": 2}

	fmt.Println("map:", n)

	nums := []int{2, 3, 4}
	sum := 0

	//index, value
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)

	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}
}
