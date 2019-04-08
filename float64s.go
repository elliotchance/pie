package pie

import "encoding/json"

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

// Contains returns true if the float64 exists in the slice.
func (ss Float64s) Contains(lookingFor float64) bool {
	for _, s := range ss {
		if s == lookingFor {
			return true
		}
	}

	return false
}

// Only will return a new slice containing only the elements that return
// true from the condition. The returned slice may contain zero elements (nil).
//
// Float64sWithout works in the opposite way as Float64sOnly.
func (ss Float64s) Only(condition Float64ConditionFunc) (ss2 Float64s) {
	for _, s := range ss {
		if condition(s) {
			ss2 = append(ss2, s)
		}
	}

	return
}

// Without works the same as Float64sOnly, with a negated condition. That is, it
// will return a new slice only containing the elements that returned false from
// the condition. The returned slice may contain zero elements (nil).
func (ss Float64s) Without(condition Float64ConditionFunc) (ss2 Float64s) {
	for _, s := range ss {
		if !condition(s) {
			ss2 = append(ss2, s)
		}
	}

	return
}

// Transform will return a new slice where each element has been transformed.
// The number of element returned will always be the same as the input.
func (ss Float64s) Transform(fn Float64TransformFunc) (ss2 Float64s) {
	if ss == nil {
		return nil
	}

	ss2 = make([]float64, len(ss))
	for i, s := range ss {
		ss2[i] = fn(s)
	}

	return
}

// FirstOr returns the first element or a default value if there are no
// elements.
func (ss Float64s) FirstOr(defaultValue float64) float64 {
	if len(ss) == 0 {
		return defaultValue
	}

	return ss[0]
}

// LastOr returns the last element or a default value if there are no elements.
func (ss Float64s) LastOr(defaultValue float64) float64 {
	if len(ss) == 0 {
		return defaultValue
	}

	return ss[len(ss)-1]
}

// First returns the first element, or zero. Also see FirstOr().
func (ss Float64s) First() float64 {
	return ss.FirstOr(0)
}

// Last returns the last element, or zero. Also see LastOr().
func (ss Float64s) Last() float64 {
	return ss.LastOr(0)
}

// Sum is the sum of all of the elements.
func (ss Float64s) Sum() (sum float64) {
	for _, s := range ss {
		sum += s
	}

	return
}

// Len returns the number of elements.
func (ss Float64s) Len() int {
	return len(ss)
}

// Average is the average of all of the elements, or zero if there are no
// elements.
func (ss Float64s) Average() float64 {
	if l := float64(len(ss)); l > 0 {
		return ss.Sum() / l
	}

	return 0
}

// Min is the minimum value, or zero.
func (ss Float64s) Min() (min float64) {
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

// Max is the maximum value, or zero.
func (ss Float64s) Max() (max float64) {
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

// JSONString returns the JSON encoded array as a string.
//
// One important thing to note is that it will treat a nil slice as an empty
// slice to ensure that the JSON value return is always an array.
func (ss Float64s) JSONString() string {
	if ss == nil {
		return "[]"
	}

	// An error should not be possible.
	data, _ := json.Marshal(ss)

	return string(data)
}
