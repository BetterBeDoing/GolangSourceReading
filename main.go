package main

import (
	"GolangSourceReading/myslice"
	"fmt"
)

type intX []int

func main() {
	slice0 := []int{1, 2, 3, 4, 5, 6}
	//myslice.CheckTildeOp(slice0)
	slice1 := append(slice0, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20)
	fmt.Println(cap(slice1))
	myslice.Insert(slice0, 1, 4, 5)
}
