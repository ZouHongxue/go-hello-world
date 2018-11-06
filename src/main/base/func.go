package base

import "fmt"

//multi-return
func vals() (int, int) {
	return 3, 7
}

//multi-arg
func sum(nums ...int) {
	var total int = 0
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

func main() {
	a, b := vals()
	fmt.Println(a)
	fmt.Println(b)

	_, c := vals()
	fmt.Println(c)

	sum(1, 2, 3, 4, 5)
}
