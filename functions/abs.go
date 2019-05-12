package functions

import (
	"math"
)

// Abs is a function which returns the absolute value of all the
// elements in the slice.
func (ss SliceType) Abs() SliceType {
	result := make(SliceType, len(ss))
	for i, val := range ss {
		result[i] = ElementType(math.Abs(float64(val)))
	}
	return result
}
