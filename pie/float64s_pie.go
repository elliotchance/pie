package pie

import (
	"context"
	"encoding/json"
	"github.com/elliotchance/pie/pie/util"
	"math"
	"math/rand"
	"sort"
)

// Abs is a function which returns the absolute value of all the
// elements in the slice.
func (ss Float64s) Abs() Float64s {
	for i, val := range ss {
		ss[i] = float64(math.Abs(float64(val)))
	}
	return ss
}

// All will return true if all callbacks return true. It follows the same logic
// as the all() function in Python.
//
// If the list is empty then true is always returned.
func (ss Float64s) All(fn func(value float64) bool) bool {
	for _, value := range ss {
		if !fn(value) {
			return false
		}
	}

	return true
}

// Any will return true if any callbacks return true. It follows the same logic
// as the any() function in Python.
//
// If the list is empty then false is always returned.
func (ss Float64s) Any(fn func(value float64) bool) bool {
	for _, value := range ss {
		if fn(value) {
			return true
		}
	}

	return false
}

// Append will return a new slice with the elements appended to the end. It is a
// wrapper for the internal append(). It is offered as a function so that it can
// more easily chained.
//
// It is acceptable to provide zero arguments.
func (ss Float64s) Append(elements ...float64) Float64s {
	return append(ss, elements...)
}

// AreSorted will return true if the slice is already sorted. It is a wrapper
// for sort.Float64sAreSorted.
func (ss Float64s) AreSorted() bool {
	return sort.SliceIsSorted(ss, func(i, j int) bool {
		return ss[i] < ss[j]
	})
}

// AreUnique will return true if the slice contains elements that are all
// different (unique) from each other.
func (ss Float64s) AreUnique() bool {
	return ss.Unique().Len() == ss.Len()
}

// Average is the average of all of the elements, or zero if there are no
// elements.
func (ss Float64s) Average() float64 {
	if l := float64(len(ss)); l > 0 {
		return float64(ss.Sum()) / float64(l)
	}

	return 0
}

// Bottom will return n elements from bottom
//
// that means that elements is taken from the end of the slice
// for this [1,2,3] slice with n == 2 will be returned [3,2]
// if the slice has less elements then n that'll return all elements
// if n < 0 it'll return empty slice.
func (ss Float64s) Bottom(n int) (top Float64s) {
	var lastIndex = len(ss) - 1
	for i := lastIndex; i > -1 && n > 0; i-- {
		top = append(top, ss[i])
		n--
	}

	return
}

// Contains returns true if the element exists in the slice.
//
// When using slices of pointers it will only compare by address, not value.
func (ss Float64s) Contains(lookingFor float64) bool {
	for _, s := range ss {
		if s == lookingFor {
			return true
		}
	}

	return false
}

// Each is more condensed version of Transform that allows an action to happen
// on each elements and pass the original slice on.
//
//   cars.Each(func (car *Car) {
//       fmt.Printf("Car color is: %s\n", car.Color)
//   })
//
// Pie will not ensure immutability on items passed in so they can be
// manipulated, if you choose to do it this way, for example:
//
//   // Set all car colors to Red.
//   cars.Each(func (car *Car) {
//       car.Color = "Red"
//   })
//
func (ss Float64s) Each(fn func(float64)) Float64s {
	for _, s := range ss {
		fn(s)
	}

	return ss
}

// Extend will return a new slice with the slices of elements appended to the
// end.
//
// It is acceptable to provide zero arguments.
func (ss Float64s) Extend(slices ...Float64s) (ss2 Float64s) {
	ss2 = ss

	for _, slice := range slices {
		ss2 = ss2.Append(slice...)
	}

	return ss2
}

// Filter will return a new slice containing only the elements that return
// true from the condition. The returned slice may contain zero elements (nil).
//
// FilterNot works in the opposite way of Filter.
func (ss Float64s) Filter(condition func(float64) bool) (ss2 Float64s) {
	for _, s := range ss {
		if condition(s) {
			ss2 = append(ss2, s)
		}
	}
	return
}

// FilterNot works the same as Filter, with a negated condition. That is, it will
// return a new slice only containing the elements that returned false from the
// condition. The returned slice may contain zero elements (nil).
func (ss Float64s) FilterNot(condition func(float64) bool) (ss2 Float64s) {
	for _, s := range ss {
		if !condition(s) {
			ss2 = append(ss2, s)
		}
	}

	return
}

// First returns the first element, or zero. Also see FirstOr().
func (ss Float64s) First() float64 {
	return ss.FirstOr(0)
}

// FirstOr returns the first element or a default value if there are no
// elements.
func (ss Float64s) FirstOr(defaultValue float64) float64 {
	if len(ss) == 0 {
		return defaultValue
	}

	return ss[0]
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

// Last returns the last element, or zero. Also see LastOr().
func (ss Float64s) Last() float64 {
	return ss.LastOr(0)
}

// LastOr returns the last element or a default value if there are no elements.
func (ss Float64s) LastOr(defaultValue float64) float64 {
	if len(ss) == 0 {
		return defaultValue
	}

	return ss[len(ss)-1]
}

// Len returns the number of elements.
func (ss Float64s) Len() int {
	return len(ss)
}

// Map will return a new slice where each element has been mapped (transformed).
// The number of elements returned will always be the same as the input.
//
// Be careful when using this with slices of pointers. If you modify the input
// value it will affect the original slice. Be sure to return a new allocated
// object or deep copy the existing one.
func (ss Float64s) Map(fn func(float64) float64) (ss2 Float64s) {
	if ss == nil {
		return nil
	}

	ss2 = make([]float64, len(ss))
	for i, s := range ss {
		ss2[i] = fn(s)
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

// Median returns the value separating the higher half from the lower half of a
// data sample.
//
// Zero is returned if there are no elements in the slice.
func (ss Float64s) Median() float64 {
	l := len(ss)

	switch {
	case l == 0:
		return 0

	case l == 1:
		return ss[0]
	}

	sorted := ss.Sort()

	if l%2 != 0 {
		return sorted[l/2]
	}

	return (sorted[l/2-1] + sorted[l/2]) / 2
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

// Random returns a random element by your rand.Source, or zero
func (ss Float64s) Random(source rand.Source) float64 {
	n := len(ss)

	// Avoid the extra allocation.
	if n < 1 {
		return 0
	}
	if n < 2 {
		return ss[0]
	}
	rnd := rand.New(source)
	i := rnd.Intn(n)
	return ss[i]
}

// Reverse returns a new copy of the slice with the elements ordered in reverse.
// This is useful when combined with Sort to get a descending sort order:
//
//   ss.Sort().Reverse()
//
func (ss Float64s) Reverse() Float64s {
	// Avoid the allocation. If there is one element or less it is already
	// reversed.
	if len(ss) < 2 {
		return ss
	}

	sorted := make([]float64, len(ss))
	for i := 0; i < len(ss); i++ {
		sorted[i] = ss[len(ss)-i-1]
	}

	return sorted
}

// Send sends elements to channel
// in normal act it sends all elements but if func canceled it can be less
//
// it locks execution of gorutine
// it doesn't close channel after work
// returns sended elements if len(this) != len(old) considered func was canceled
func (ss Float64s) Send(ctx context.Context, ch chan<- float64) Float64s {
	for i, s := range ss {
		select {
		case <-ctx.Done():
			return ss[:i]
		default:
			ch <- s
		}
	}

	return ss
}

// Sequence returns a sequencent elements.
//
// There are 3 variations to generate:
// 		1. [0, n).
//		2. [min, max).
//		3. [min, max) with step.
//
// if lenght of params == 1 considered that will returned slice between 0 and n,
// where n is the first param, [0, n).
// if lenght of params == 2 considered that will returned slice between min and max,
// where min is the first param, max is the second, [min, max).
// if lenght of params == 3 considered that will returned slice between min and max with step,
// where min is the first param, max is the second, step is the third one, [min, max) with step.
//
// if len(params) > 3, uses 3 approach, with min, max and step others params will be ignored
func (ss Float64s) Sequence(params ...int) Float64s {
	var seq = func(min, max, step int) (seq Float64s) {
		lenght := int(math.Round(float64(max-min) / float64(step)))
		if lenght < 1 {
			return
		}

		seq = make(Float64s, lenght)
		for i := 0; i < lenght; min += step {
			seq[i] = float64(min)
			i++
		}

		return seq
	}

	if len(params) > 2 {
		return seq(params[0], params[1], params[2])
	} else if len(params) > 1 {
		return seq(params[0], params[1], 1)
	} else {
		return seq(0, params[0], 1)
	}
}

// Sort works similar to sort.Float64s(). However, unlike sort.Float64s the
// slice returned will be reallocated as to not modify the input slice.
//
// See Reverse() and AreSorted().
func (ss Float64s) Sort() Float64s {
	// Avoid the allocation. If there is one element or less it is already
	// sorted.
	if len(ss) < 2 {
		return ss
	}

	sorted := make([]float64, len(ss))
	copy(sorted, ss)
	sort.Slice(sorted, func(i, j int) bool {
		return sorted[i] < sorted[j]
	})

	return sorted
}

// Sum is the sum of all of the elements.
func (ss Float64s) Sum() (sum float64) {
	for _, s := range ss {
		sum += s
	}

	return
}

// Shuffle returns shuffled slice by your rand.Source
func (ss Float64s) Shuffle(source rand.Source) Float64s {
	n := len(ss)

	// Avoid the extra allocation.
	if n < 2 {
		return ss
	}

	// go 1.10+ provides rnd.Shuffle. However, to support older versions we copy
	// the algorithm directly from the go source: src/math/rand/rand.go below,
	// with some adjustments:
	shuffled := make([]float64, n)
	copy(shuffled, ss)

	rnd := rand.New(source)

	util.Shuffle(rnd, n, func(i, j int) {
		shuffled[i], shuffled[j] = shuffled[j], shuffled[i]
	})

	return shuffled
}

// Top will return n elements from head of the slice
// if the slice has less elements then n that'll return all elements
// if n < 0 it'll return empty slice.
func (ss Float64s) Top(n int) (top Float64s) {
	for i := 0; i < len(ss) && n > 0; i++ {
		top = append(top, ss[i])
		n--
	}

	return
}

// ToStrings transforms each element to a string.
func (ss Float64s) ToStrings(transform func(float64) string) Strings {
	l := len(ss)

	// Avoid the allocation.
	if l == 0 {
		return nil
	}

	result := make(Strings, l)
	for i := 0; i < l; i++ {
		result[i] = transform(ss[i])
	}

	return result
}

// Unique returns a new slice with all of the unique values.
//
// The items will be returned in a randomized order, even with the same input.
//
// The number of items returned may be the same as the input or less. It will
// never return zero items unless then input slice has zero items.
//
// A slice with zero elements is considered to be unique.
//
// See AreUnique().
func (ss Float64s) Unique() Float64s {
	// Avoid the allocation. If there is one element or less it is already
	// unique.
	if len(ss) < 2 {
		return ss
	}

	values := map[float64]struct{}{}

	for _, value := range ss {
		values[value] = struct{}{}
	}

	var uniqueValues Float64s
	for value := range values {
		uniqueValues = append(uniqueValues, value)
	}

	return uniqueValues
}
