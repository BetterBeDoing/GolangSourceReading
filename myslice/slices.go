package myslice

type MyComparable interface{ int | string | float64 | bool }

type E MyComparable

// This a function that try to understand the usage of ~
// this function does not use the ~ operator
func EqualWithoutOperator[S []E, E MyComparable](a, b S) bool {
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

func EqualWithOperator[S ~[]E, E MyComparable](a, b S) bool {
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
