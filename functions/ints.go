package functions

import (
	"github.com/elliotchance/pie/pie"
	"strconv"
)

// Ints transforms each element to an integer.
func (ss SliceType) Ints() pie.Ints {
	l := len(ss)

	// Avoid the allocation.
	if l == 0 {
		return nil
	}

	result := make(pie.Ints, l)
	for i := 0; i < l; i++ {
		mightBeString := ss[i]
		f, _ := strconv.ParseFloat(mightBeString.String(), 64)
		result[i] = int(f)
	}

	return result
}
