package functions

import (
	"github.com/elliotchance/pie/pie"
)

// StringsUsing transforms each element to a string.
func (ss SliceType) StringsUsing(transform func(ElementType) string) pie.Strings {
	l := len(ss)

	// Avoid the allocation.
	if l == 0 {
		return nil
	}

	result := make(pie.Strings, l)
	for i := 0; i < l; i++ {
		result[i] = transform(ss[i])
	}

	return result
}
