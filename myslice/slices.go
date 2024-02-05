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

// Have a question about the Insert
func Insert[S ~[]E, E any](s S, i int, v ...E) S {
	m := len(v)
	n := len(s)

	// Use append rather than make so that we bump the size of
	// the slice up to the next storage class.
	// This is what Grow does but we don't call Grow because
	// that might copy the values twice.
	// because the capacity of v + s is bigger than the capacity of s
	// so the append will create a new slice
	s2 := append(s[:i], make(S, n+m-i)...)
	fmt.Println(s2)
	copy(s2[i:], v)
	fmt.Println(s2)
	copy(s2[i+m:], s[i:])
	fmt.Println(s2)
	return s2
}
