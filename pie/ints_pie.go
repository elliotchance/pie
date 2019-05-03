package pie

import (
	"encoding/json"
	"github.com/elliotchance/pie/pie/util"
	"math"
	"math/rand"
	"sort"
)

// Abs is a function which returns the absolute value of all the
// elements in the slice.
func (ss Ints) Abs() Ints {
	for i, val := range ss {
		ss[i] = int(math.Abs(float64(val)))
	}
	return ss
}

// All will return true if all callbacks return true. It follows the same logic
// as the all() function in Python.
//
// If the list is empty then true is always returned.
func (ss Ints) All(fn func(value int) bool) bool {
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
func (ss Ints) Any(fn func(value int) bool) bool {
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
func (ss Ints) Append(elements ...int) Ints {
	return append(ss, elements...)
}

// AreSorted will return true if the slice is already sorted. It is a wrapper
// for sort.IntsAreSorted.
func (ss Ints) AreSorted() bool {
	return sort.SliceIsSorted(ss, func(i, j int) bool {
		return ss[i] < ss[j]
	})
}

// AreUnique will return true if the slice contains elements that are all
// different (unique) from each other.
func (ss Ints) AreUnique() bool {
	return ss.Unique().Len() == ss.Len()
}

// Average is the average of all of the elements, or zero if there are no
// elements.
func (ss Ints) Average() float64 {
	if l := int(len(ss)); l > 0 {
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
func (ss Ints) Bottom(n int) (top Ints) {
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
func (ss Ints) Contains(lookingFor int) bool {
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
func (ss Ints) Each(fn func(int)) Ints {
	for _, s := range ss {
		fn(s)
	}

	return ss
}

// Extend will return a new slice with the slices of elements appended to the
// end.
//
// It is acceptable to provide zero arguments.
func (ss Ints) Extend(slices ...Ints) (ss2 Ints) {
	ss2 = ss

	for _, slice := range slices {
		ss2 = ss2.Append(slice...)
	}

	return ss2
}

// First returns the first element, or zero. Also see FirstOr().
func (ss Ints) First() int {
	return ss.FirstOr(0)
}

// FirstOr returns the first element or a default value if there are no
// elements.
func (ss Ints) FirstOr(defaultValue int) int {
	if len(ss) == 0 {
		return defaultValue
	}

	return ss[0]
}

// JSONString returns the JSON encoded array as a string.
//
// One important thing to note is that it will treat a nil slice as an empty
// slice to ensure that the JSON value return is always an array.
func (ss Ints) JSONString() string {
	if ss == nil {
		return "[]"
	}

	// An error should not be possible.
	data, _ := json.Marshal(ss)

	return string(data)
}

// Last returns the last element, or zero. Also see LastOr().
func (ss Ints) Last() int {
	return ss.LastOr(0)
}

// LastOr returns the last element or a default value if there are no elements.
func (ss Ints) LastOr(defaultValue int) int {
	if len(ss) == 0 {
		return defaultValue
	}

	return ss[len(ss)-1]
}

// Len returns the number of elements.
func (ss Ints) Len() int {
	return len(ss)
}

// Max is the maximum value, or zero.
func (ss Ints) Max() (max int) {
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
func (ss Ints) Median() int {
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
func (ss Ints) Min() (min int) {
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

// Reverse returns a new copy of the slice with the elements ordered in reverse.
// This is useful when combined with Sort to get a descending sort order:
//
//   ss.Sort().Reverse()
//
func (ss Ints) Reverse() Ints {
	// Avoid the allocation. If there is one element or less it is already
	// reversed.
	if len(ss) < 2 {
		return ss
	}

	sorted := make([]int, len(ss))
	for i := 0; i < len(ss); i++ {
		sorted[i] = ss[len(ss)-i-1]
	}

	return sorted
}

// Select will return a new slice containing only the elements that return
// true from the condition. The returned slice may contain zero elements (nil).
//
// Unselect works in the opposite way as Select.
func (ss Ints) Select(condition func(int) bool) (ss2 Ints) {
	for _, s := range ss {
		if condition(s) {
			ss2 = append(ss2, s)
		}
	}

	return
}

// Sort works similar to sort.Ints(). However, unlike sort.Ints the
// slice returned will be reallocated as to not modify the input slice.
//
// See Reverse() and AreSorted().
func (ss Ints) Sort() Ints {
	// Avoid the allocation. If there is one element or less it is already
	// sorted.
	if len(ss) < 2 {
		return ss
	}

	sorted := make([]int, len(ss))
	copy(sorted, ss)
	sort.Slice(sorted, func(i, j int) bool {
		return sorted[i] < sorted[j]
	})

	return sorted
}

// Sum is the sum of all of the elements.
func (ss Ints) Sum() (sum int) {
	for _, s := range ss {
		sum += s
	}

	return
}

// Shuffle returns shuffled slice by your rand.Source
func (ss Ints) Shuffle(source rand.Source) Ints {
	n := len(ss)

	// Avoid the extra allocation.
	if n < 2 {
		return ss
	}

	// go 1.10+ provides rnd.Shuffle. However, to support older versions we copy
	// the algorithm directly from the go source: src/math/rand/rand.go below,
	// with some adjustments:
	shuffled := make([]int, n)
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
func (ss Ints) Top(n int) (top Ints) {
	for i := 0; i < len(ss) && n > 0; i++ {
		top = append(top, ss[i])
		n--
	}

	return
}

// ToStrings transforms each element to a string.
func (ss Ints) ToStrings(transform func(int) string) Strings {
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

// Transform will return a new slice where each element has been transformed.
// The number of element returned will always be the same as the input.
//
// Be careful when using this with slices of pointers. If you modify the input
// value it will affect the original slice. Be sure to return a new allocated
// object or deep copy the existing one.
func (ss Ints) Transform(fn func(int) int) (ss2 Ints) {
	if ss == nil {
		return nil
	}

	ss2 = make([]int, len(ss))
	for i, s := range ss {
		ss2[i] = fn(s)
	}

	return
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
func (ss Ints) Unique() Ints {
	// Avoid the allocation. If there is one element or less it is already
	// unique.
	if len(ss) < 2 {
		return ss
	}

	values := map[int]struct{}{}

	for _, value := range ss {
		values[value] = struct{}{}
	}

	var uniqueValues Ints
	for value := range values {
		uniqueValues = append(uniqueValues, value)
	}

	return uniqueValues
}

// Unselect works the same as Select, with a negated condition. That is, it will
// return a new slice only containing the elements that returned false from the
// condition. The returned slice may contain zero elements (nil).
func (ss Ints) Unselect(condition func(int) bool) (ss2 Ints) {
	for _, s := range ss {
		if !condition(s) {
			ss2 = append(ss2, s)
		}
	}

	return
}
