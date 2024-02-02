package main

import (
	"GolangSourceReading/myslice"
	"fmt"
)

type intX int

func main() {

	a := []int{1, 2, 3}
	c := []intX(1, 2, 3)
	b := []string{"a", "b", "c"}
	flag := []bool{}
	flag = append(flag, myslice.EqualWithoutOperator(a, a))
	flag = append(flag, myslice.EqualWithoutOperator(a, b))
	flag = append(flag, myslice.EqualWithOperator(b, b))
	flag = append(flag, myslice.EqualWithOperator(a, c))
	fmt.Println(flag)
}
