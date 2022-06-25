package scalar

import (
	"math"
)

// EqualWithinAbs returns true when a and b have an absolute difference
// not greater than tol
func EqualWithinAbs(a, b, tol float64) bool {
	return a == b || math.Abs(a-b) <= tol
}
