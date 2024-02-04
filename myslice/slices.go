package myslice

import "fmt"

type MyComparable interface{ int | string | float64 | bool }

// Try to understand the ~ operator
func Equal[S ~[]E, E MyComparable](a, b S) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func CheckTildeOp[S ~[]E, E int](a S) {
	for _, v := range a {
		fmt.Println(v)
	}
}
