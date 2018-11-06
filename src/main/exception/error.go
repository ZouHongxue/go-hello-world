package main

import (
	"errors"
	"fmt"
)

func func1(arg int) (int, error) {
	if arg == 42 {
		return -1, errors.New("can't work with 42")
	}

	return arg + 3, nil
}

type argError struct {
	arg  int
	prob string
}

func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.prob)
}

func func2(arg int) (int, error) {
	if arg == 42 {
		return -1, &argError{arg, "can't work with it"}
	}
	return arg + 3, nil
}

func main() {
	for _, i := range []int{7, 42} {
		if r, e := func1(i); e != nil {
			fmt.Println("func1 failed:", e)
		} else {
			fmt.Println("func1 failed:", r)
		}
	}

	for _, i := range []int{7, 42} {
		if r, e := func2(i); e != nil {
			fmt.Println("func2 failed:", e)
		} else {
			fmt.Println("fun2 worked", r)
		}
	}

	_, e := func2(42)
	if ae, ok := e.(*argError); ok {
		fmt.Println(ae.arg)
		fmt.Println(ae.prob)
	}
}
