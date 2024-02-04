package main

import "GolangSourceReading/myslice"

type intX []int
type intY int

func main() {
	slice0 := intX{1, 2, 3}

	myslice.CheckTildeOp(slice0)
	myslice.CheckTildeOp(slice1)
}
