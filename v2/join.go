package pie

import (
	"golang.org/x/exp/constraints"
	"strings"
)

// Join returns a string from joining each of the elements.
func Join[T constraints.Ordered](ss []T, glue string) (s string) {
	parts := make([]string, len(ss))
	for i, element := range ss {
		parts[i] = String(element)
	}

	return strings.Join(parts, glue)
}
