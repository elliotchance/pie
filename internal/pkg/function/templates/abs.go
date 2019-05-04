package templates

import (
	"math"
)

// Abs is a function which returns the absolute value of all the
// elements in the slice.
func (ss SliceType) Abs() SliceType {
	for i, val := range ss {
		ss[i] = ElementType(math.Abs(float64(val)))
	}
	return ss
}
