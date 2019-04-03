package pie

// Ints is an alias for an int slice.
//
// You can create an Ints directly:
//
//   pie.Ints{1, 2, 3}
//
// Or, cast an existing int slice:
//
//   ss := []int{1, 2, 3}
//   pie.Ints(ss)
//
type Ints []int

// IntsConditionFunc allows ints to be filtered or checked by value.
type IntsConditionFunc func(int) bool

// IntsTransformFunc transforms an int value.
type IntsTransformFunc func(int) int

// IntsContains returns true if the int exists in the slice.
func IntsContains(ss []int, lookingFor int) bool {
	for _, s := range ss {
		if s == lookingFor {
			return true
		}
	}

	return false
}

// Contains is the chained version of IntsContains.
func (ss Ints) Contains(lookingFor int) bool {
	return IntsContains(ss, lookingFor)
}

// IntsOnly will return a new slice containing only the elements that return
// true from the condition. The returned slice may contain zero elements (nil).
//
// IntsWithout works in the opposite way as IntsOnly.
func IntsOnly(ss []int, condition IntsConditionFunc) (ss2 []int) {
	for _, s := range ss {
		if condition(s) {
			ss2 = append(ss2, s)
		}
	}

	return
}

// Only is the chained version of IntsOnly.
func (ss Ints) Only(condition IntsConditionFunc) (ss2 Ints) {
	return IntsOnly(ss, condition)
}

// IntsWithout works the same as IntsOnly, with a negated condition. That is, it
// will return a new slice only containing the elements that returned false from
// the condition. The returned slice may contain zero elements (nil).
func IntsWithout(ss []int, condition IntsConditionFunc) (ss2 []int) {
	for _, s := range ss {
		if !condition(s) {
			ss2 = append(ss2, s)
		}
	}

	return
}

// Without is the chained version of IntsWithout.
func (ss Ints) Without(condition IntsConditionFunc) (ss2 Ints) {
	return IntsWithout(ss, condition)
}

// IntsTransform will return a new slice where each element has been
// transformed. The number of element returned will always be the same as the
// input.
func IntsTransform(ss []int, fn IntsTransformFunc) (ss2 []int) {
	if ss == nil {
		return nil
	}

	ss2 = make([]int, len(ss))
	for i, s := range ss {
		ss2[i] = fn(s)
	}

	return
}

// Transform is the chained version of IntsTransform.
func (ss Ints) Transform(fn IntsTransformFunc) (ss2 Ints) {
	return IntsTransform(ss, fn)
}

// IntsFirstOr returns the first element or a default value if there are no
// elements.
func IntsFirstOr(ss []int, defaultValue int) int {
	if len(ss) == 0 {
		return defaultValue
	}

	return ss[0]
}

// FirstOr is the chained version of IntsFirstOr.
func (ss Ints) FirstOr(defaultValue int) int {
	return IntsFirstOr(ss, defaultValue)
}

// IntsLastOr returns the last element or a default value if there are no
// elements.
func IntsLastOr(ss []int, defaultValue int) int {
	if len(ss) == 0 {
		return defaultValue
	}

	return ss[len(ss)-1]
}

// LastOr is the chained version of IntsLastOr.
func (ss Ints) LastOr(defaultValue int) int {
	return IntsLastOr(ss, defaultValue)
}

// IntsFirst returns the first element, or zero. Also see IntsFirstOr.
func IntsFirst(ss []int) int {
	return IntsFirstOr(ss, 0)
}

// First is the chained version of IntsFirst.
func (ss Ints) First() int {
	return IntsFirst(ss)
}

// IntsLast returns the last element, or zero. Also see IntsLastOr.
func IntsLast(ss []int) int {
	return IntsLastOr(ss, 0)
}

// Last is the chained version of IntsLast.
func (ss Ints) Last() int {
	return IntsLast(ss)
}
