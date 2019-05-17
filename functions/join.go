package functions

import "strings"

// Join returns a string from joining each of the elements.
func (ss SliceType) Join(glue string) (s string) {
	var slice interface{} = []ElementType(ss)

	if y, ok := slice.([]string); ok {
		// The stdlib is efficient for type []string
		return strings.Join(y, glue)
	} else {
		// General case
		parts := make([]string, len(ss))
		for i, element := range ss {
			mightBeString := element
			parts[i] = mightBeString.String()
		}
		return strings.Join(parts, glue)
	}
}
