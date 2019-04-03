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
func IntsOnly(ss []int, condition IntConditionFunc) (ss2 []int) {
	for _, s := range ss {
		if condition(s) {
			ss2 = append(ss2, s)
		}
	}

	return
}

// Only is the chained version of IntsOnly.
func (ss Ints) Only(condition IntConditionFunc) (ss2 Ints) {
	return IntsOnly(ss, condition)
}

// IntsWithout works the same as IntsOnly, with a negated condition. That is, it
// will return a new slice only containing the elements that returned false from
// the condition. The returned slice may contain zero elements (nil).
func IntsWithout(ss []int, condition IntConditionFunc) (ss2 []int) {
	for _, s := range ss {
		if !condition(s) {
			ss2 = append(ss2, s)
		}
	}

	return
}

// Without is the chained version of IntsWithout.
func (ss Ints) Without(condition IntConditionFunc) (ss2 Ints) {
	return IntsWithout(ss, condition)
}

// IntsTransform will return a new slice where each element has been
// transformed. The number of element returned will always be the same as the
// input.
func IntsTransform(ss []int, fn IntTransformFunc) (ss2 []int) {
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
func (ss Ints) Transform(fn IntTransformFunc) (ss2 Ints) {
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

// IntsSum is the sum of all of the elements.
func IntsSum(ss []int) (sum int) {
	for _, s := range ss {
		sum += s
	}

	return
}

// Sum is the chained version of IntsSum.
func (ss Ints) Sum() int {
	return IntsSum(ss)
}

// Len returns the number of elements.
func (ss Ints) Len() int {
	return len(ss)
}

// IntsAverage is the average of all of the elements, or zero if there are no
// elements.
func IntsAverage(ss []int) float64 {
	if l := float64(len(ss)); l > 0 {
		return float64(IntsSum(ss)) / l
	}

	return 0
}

// Average is the chained version of IntsAverage.
func (ss Ints) Average() float64 {
	return IntsAverage(ss)
}

// IntsMin is the minimum value, or zero.
func IntsMin(ss []int) (min int) {
	if len(ss) == 0 {
		return
	}

	min = ss[0]
	for _, s := range ss {
		if s < min {
			min = s
		}
	}

	return
}

// Min is the chained version of IntsMin.
func (ss Ints) Min() int {
	return IntsMin(ss)
}

// IntsMax is the maximum value, or zero.
func IntsMax(ss []int) (max int) {
	if len(ss) == 0 {
		return
	}

	max = ss[0]
	for _, s := range ss {
		if s > max {
			max = s
		}
	}

	return
}

// Max is the chained version of IntsMax.
func (ss Ints) Max() int {
	return IntsMax(ss)
}
