package functions

import (
	"github.com/elliotchance/pie/pie"
	"strconv"
)

// Float64s transforms each element to a float64.
func (ss SliceType) Float64s() pie.Float64s {
	l := len(ss)

	// Avoid the allocation.
	if l == 0 {
		return nil
	}

	result := make(pie.Float64s, l)
	for i := 0; i < l; i++ {
		mightBeString := ss[i]
		result[i], _ = strconv.ParseFloat(mightBeString.String(), 64)
	}

	return result
}
