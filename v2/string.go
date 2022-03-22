package pie

import (
	"fmt"
	"golang.org/x/exp/constraints"
)

// String transforms a value into a string. Nil values will be treated as empty
// strings.
//
// If the element type implements fmt.Stringer it will be used. Otherwise it
// will fallback to the result of:
//
//   fmt.Sprintf("%v")
//
func String[T constraints.Ordered](s T) string {
	return fmt.Sprintf("%v", s)
}
