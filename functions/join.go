package functions

// Join returns a string from joining each of the elements.
func (ss StringSliceType) Join(glue string) (s string) {
	// Ideally we would call directly strings.Join, however the code generation
	// makes it non-trivial to have the correct parameter type everywhere.

	// We use the same idea, which is to make only 1 string allocation for the result.
	if len(ss) == 0 {
		return ""
	}
	targetLength := 0
	for _, element := range ss {
		targetLength += len(element)
	}
	nglue := len(glue)
	targetLength += (len(ss) - 1) * nglue

	buffer := make([]byte, targetLength)
	b := buffer[:]
	for i, element := range ss {
		if i > 0 {
			copy(b[:nglue], glue)
			b = b[nglue:]
		}
		n := len(element)
		copy(b[:n], element)
		b = b[n:]
	}

	return string(buffer)
}
