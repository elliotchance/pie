package pie

// Float64s is an alias for an float64 slice.
//
// You can create an Float64s directly:
//
//   pie.Float64s{1, 2, 3}
//
// Or, cast an existing float64 slice:
//
//   ss := []float64{1, 2, 3}
//   pie.Float64s(ss)
//
type Float64s []float64

// Float64sContains returns true if the float64 exists in the slice.
func Float64sContains(ss []float64, lookingFor float64) bool {
	for _, s := range ss {
		if s == lookingFor {
			return true
		}
	}

	return false
}

// Contains is the chained version of Float64sContains.
func (ss Float64s) Contains(lookingFor float64) bool {
	return Float64sContains(ss, lookingFor)
}

// Float64sOnly will return a new slice containing only the elements that return
// true from the condition. The returned slice may contain zero elements (nil).
//
// Float64sWithout works in the opposite way as Float64sOnly.
func Float64sOnly(ss []float64, condition Float64ConditionFunc) (ss2 []float64) {
	for _, s := range ss {
		if condition(s) {
			ss2 = append(ss2, s)
		}
	}

	return
}

// Only is the chained version of Float64sOnly.
func (ss Float64s) Only(condition Float64ConditionFunc) (ss2 Float64s) {
	return Float64sOnly(ss, condition)
}

// Float64sWithout works the same as Float64sOnly, with a negated condition. That is, it
// will return a new slice only containing the elements that returned false from
// the condition. The returned slice may contain zero elements (nil).
func Float64sWithout(ss []float64, condition Float64ConditionFunc) (ss2 []float64) {
	for _, s := range ss {
		if !condition(s) {
			ss2 = append(ss2, s)
		}
	}

	return
}

// Without is the chained version of Float64sWithout.
func (ss Float64s) Without(condition Float64ConditionFunc) (ss2 Float64s) {
	return Float64sWithout(ss, condition)
}

// Float64sTransform will return a new slice where each element has been
// transformed. The number of element returned will always be the same as the
// input.
func Float64sTransform(ss []float64, fn Float64TransformFunc) (ss2 []float64) {
	if ss == nil {
		return nil
	}

	ss2 = make([]float64, len(ss))
	for i, s := range ss {
		ss2[i] = fn(s)
	}

	return
}

// Transform is the chained version of Float64sTransform.
func (ss Float64s) Transform(fn Float64TransformFunc) (ss2 Float64s) {
	return Float64sTransform(ss, fn)
}

// Float64sFirstOr returns the first element or a default value if there are no
// elements.
func Float64sFirstOr(ss []float64, defaultValue float64) float64 {
	if len(ss) == 0 {
		return defaultValue
	}

	return ss[0]
}

// FirstOr is the chained version of Float64sFirstOr.
func (ss Float64s) FirstOr(defaultValue float64) float64 {
	return Float64sFirstOr(ss, defaultValue)
}

// Float64sLastOr returns the last element or a default value if there are no
// elements.
func Float64sLastOr(ss []float64, defaultValue float64) float64 {
	if len(ss) == 0 {
		return defaultValue
	}

	return ss[len(ss)-1]
}

// LastOr is the chained version of Float64sLastOr.
func (ss Float64s) LastOr(defaultValue float64) float64 {
	return Float64sLastOr(ss, defaultValue)
}

// Float64sFirst returns the first element, or zero. Also see Float64sFirstOr.
func Float64sFirst(ss []float64) float64 {
	return Float64sFirstOr(ss, 0)
}

// First is the chained version of Float64sFirst.
func (ss Float64s) First() float64 {
	return Float64sFirst(ss)
}

// Float64sLast returns the last element, or zero. Also see Float64sLastOr.
func Float64sLast(ss []float64) float64 {
	return Float64sLastOr(ss, 0)
}

// Last is the chained version of Float64sLast.
func (ss Float64s) Last() float64 {
	return Float64sLast(ss)
}

// Float64sSum is the sum of all of the elements.
func Float64sSum(ss []float64) (sum float64) {
	for _, s := range ss {
		sum += s
	}

	return
}

// Sum is the chained version of Float64sSum.
func (ss Float64s) Sum() float64 {
	return Float64sSum(ss)
}

// Len returns the number of elements.
func (ss Float64s) Len() int {
	return len(ss)
}

// Float64sAverage is the average of all of the elements, or zero if there are no
// elements.
func Float64sAverage(ss []float64) float64 {
	if l := float64(len(ss)); l > 0 {
		return Float64sSum(ss) / l
	}

	return 0
}

// Average is the chained version of Float64sAverage.
func (ss Float64s) Average() float64 {
	return Float64sAverage(ss)
}

// Float64sMin is the minimum value, or zero.
func Float64sMin(ss []float64) (min float64) {
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

// Min is the chained version of Float64sMin.
func (ss Float64s) Min() float64 {
	return Float64sMin(ss)
}

// Float64sMax is the maximum value, or zero.
func Float64sMax(ss []float64) (max float64) {
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

// Max is the chained version of Float64sMax.
func (ss Float64s) Max() float64 {
	return Float64sMax(ss)
}
