// Code generated by go generate; DO NOT EDIT.
package main

var pieTemplates = map[string]string{
	"Abs": `package functions

import (
	"math"
)

// Abs is a function which returns the absolute value of all the
// elements in the slice.
func (ss SliceType) Abs() SliceType {
	for i, val := range ss {
		ss[i] = ElementType(math.Abs(float64(val)))
	}
	return ss
}
`,
	"All": `package functions

// All will return true if all callbacks return true. It follows the same logic
// as the all() function in Python.
//
// If the list is empty then true is always returned.
func (ss SliceType) All(fn func(value ElementType) bool) bool {
	for _, value := range ss {
		if !fn(value) {
			return false
		}
	}

	return true
}
`,
	"Any": `package functions

// Any will return true if any callbacks return true. It follows the same logic
// as the any() function in Python.
//
// If the list is empty then false is always returned.
func (ss SliceType) Any(fn func(value ElementType) bool) bool {
	for _, value := range ss {
		if fn(value) {
			return true
		}
	}

	return false
}
`,
	"Append": `package functions

// Append will return a new slice with the elements appended to the end.
//
// It is acceptable to provide zero arguments.
func (ss SliceType) Append(elements ...ElementType) SliceType {
	// Copy ss, to make sure no memory is overlapping between input and
	// output. See issue #97.
	result := append(SliceType{}, ss...)

	result = append(result, elements...)
	return result
}
`,
	"AreSorted": `package functions

import (
	"sort"
)

// AreSorted will return true if the slice is already sorted. It is a wrapper
// for sort.SliceTypeAreSorted.
func (ss SliceType) AreSorted() bool {
	return sort.SliceIsSorted(ss, func(i, j int) bool {
		return ss[i] < ss[j]
	})
}
`,
	"AreUnique": `package functions

// AreUnique will return true if the slice contains elements that are all
// different (unique) from each other.
func (ss SliceType) AreUnique() bool {
	return ss.Unique().Len() == ss.Len()
}
`,
	"Average": `package functions

// Average is the average of all of the elements, or zero if there are no
// elements.
func (ss SliceType) Average() float64 {
	if l := ElementType(len(ss)); l > 0 {
		return float64(ss.Sum()) / float64(l)
	}

	return 0
}
`,
	"Bottom": `package functions

// Bottom will return n elements from bottom
//
// that means that elements is taken from the end of the slice
// for this [1,2,3] slice with n == 2 will be returned [3,2]
// if the slice has less elements then n that'll return all elements
// if n < 0 it'll return empty slice.
func (ss SliceType) Bottom(n int) (top SliceType) {
	var lastIndex = len(ss) - 1
	for i := lastIndex; i > -1 && n > 0; i-- {
		top = append(top, ss[i])
		n--
	}

	return
}
`,
	"Contains": `package functions

// Contains returns true if the element exists in the slice.
//
// When using slices of pointers it will only compare by address, not value.
func (ss SliceType) Contains(lookingFor ElementType) bool {
	for _, s := range ss {
		if lookingFor.Equals(s) {
			return true
		}
	}

	return false
}
`,
	"Diff": `package functions

// Diff returns the elements that needs to be added or removed from the first
// slice to have the same elements in the second slice.
//
// The order of elements is not taken into consideration, so the slices are
// treated sets that allow duplicate items.
//
// The added and removed returned may be blank respectively, or contain upto as
// many elements that exists in the largest slice.
func (ss SliceType) Diff(against SliceType) (added, removed SliceType) {
	// This is probably not the best way to do it. We do an O(n^2) between the
	// slices to see which items are missing in each direction.

	diffOneWay := func(ss1, ss2raw SliceType) (result SliceType) {
		ss2 := make(SliceType, len(ss2raw))
		copy(ss2, ss2raw)

		for _, s := range ss1 {
			found := false

			for i, element := range ss2 {
				if s.Equals(element) {
					ss2 = append(ss2[:i], ss2[i+1:]...)
					found = true
				}
			}

			if !found {
				result = append(result, s)
			}
		}

		return
	}

	removed = diffOneWay(ss, against)
	added = diffOneWay(against, ss)

	return
}
`,
	"Each": `package functions

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
func (ss SliceType) Each(fn func(ElementType)) SliceType {
	for _, s := range ss {
		fn(s)
	}

	return ss
}
`,
	"Extend": `package functions

// Extend will return a new slice with the slices of elements appended to the
// end.
//
// It is acceptable to provide zero arguments.
func (ss SliceType) Extend(slices ...SliceType) (ss2 SliceType) {
	ss2 = ss

	for _, slice := range slices {
		ss2 = ss2.Append(slice...)
	}

	return ss2
}
`,
	"Filter": `package functions

// Filter will return a new slice containing only the elements that return
// true from the condition. The returned slice may contain zero elements (nil).
//
// FilterNot works in the opposite way of Filter.
func (ss SliceType) Filter(condition func(ElementType) bool) (ss2 SliceType) {
	for _, s := range ss {
		if condition(s) {
			ss2 = append(ss2, s)
		}
	}
	return
}
`,
	"FilterNot": `package functions

// FilterNot works the same as Filter, with a negated condition. That is, it will
// return a new slice only containing the elements that returned false from the
// condition. The returned slice may contain zero elements (nil).
func (ss SliceType) FilterNot(condition func(ElementType) bool) (ss2 SliceType) {
	for _, s := range ss {
		if !condition(s) {
			ss2 = append(ss2, s)
		}
	}

	return
}
`,
	"FindFirst": `package functions

// FindFirst will return the index of the first element when the callback returns true or -1 if no element is find.
// It follows the same logic as the findIndex() function in Javascript.
//
// If the list is empty then -1 is always returned.
func (ss SliceType) FindFirst(fn func(value ElementType) bool) int {
	for idx, value := range ss {
		if fn(value) {
			return idx
		}
	}

	return -1
}
`,
	"First": `package functions

// First returns the first element, or zero. Also see FirstOr().
func (ss SliceType) First() ElementType {
	return ss.FirstOr(ElementZeroValue)
}
`,
	"FirstOr": `package functions

// FirstOr returns the first element or a default value if there are no
// elements.
func (ss SliceType) FirstOr(defaultValue ElementType) ElementType {
	if len(ss) == 0 {
		return defaultValue
	}

	return ss[0]
}
`,
	"Float64s": `package functions

import (
	"github.com/elliotchance/pie/pie"
	"strconv"
)

// Float64s transforms each element to a float64.
func (ss SliceType) Float64s() pie.Float64s {
	l := len(ss)

	// Avoid the allocation.
	if l == 0 {
		return nil
	}

	result := make(pie.Float64s, l)
	for i := 0; i < l; i++ {
		mightBeString := ss[i]
		result[i], _ = strconv.ParseFloat(mightBeString.String(), 64)
	}

	return result
}
`,
	"Intersect": `package functions

// Intersect returns items that exist in all lists.
//
// It returns slice without any duplicates.
// If zero slice arguments are provided, then nil is returned.
func (ss SliceType) Intersect(slices ...SliceType) (ss2 SliceType) {
	if slices == nil {
		return nil
	}

	var uniqs = make([]map[ElementType]struct{}, len(slices))
	for i := 0; i < len(slices); i++ {
		m := make(map[ElementType]struct{})
		for _, el := range slices[i] {
			m[el] = struct{}{}
		}
		uniqs[i] = m
	}

	var containsInAll = false
	for _, el := range ss.Unique() {
		for _, u := range uniqs {
			if _, exists := u[el]; !exists {
				containsInAll = false
				break
			}
			containsInAll = true
		}
		if containsInAll {
			ss2 = append(ss2, el)
		}
	}

	return
}
`,
	"Ints": `package functions

import (
	"github.com/elliotchance/pie/pie"
	"strconv"
)

// Ints transforms each element to an integer.
func (ss SliceType) Ints() pie.Ints {
	l := len(ss)

	// Avoid the allocation.
	if l == 0 {
		return nil
	}

	result := make(pie.Ints, l)
	for i := 0; i < l; i++ {
		mightBeString := ss[i]
		f, _ := strconv.ParseFloat(mightBeString.String(), 64)
		result[i] = int(f)
	}

	return result
}
`,
	"JSONString": `package functions

import (
	"encoding/json"
)

// JSONString returns the JSON encoded array as a string.
//
// One important thing to note is that it will treat a nil slice as an empty
// slice to ensure that the JSON value return is always an array.
func (ss SliceType) JSONString() string {
	if ss == nil {
		return "[]"
	}

	// An error should not be possible.
	data, _ := json.Marshal(ss)

	return string(data)
}
`,
	"Join": `package functions

// Join returns a string from joining each of the elements.
func (ss StringSliceType) Join(glue string) (s string) {
	for i, element := range ss {
		if i > 0 {
			s += glue
		}

		s += string(element)
	}

	return s
}
`,
	"Keys": `package functions

// Keys returns the keys in the map. All of the items will be unique.
//
// Due to Go's randomization of iterating maps the order is not deterministic.
func (m MapType) Keys() KeySliceType {
	// Avoid allocation
	l := len(m)
	if l == 0 {
		return nil
	}

	i := 0
	keys := make(KeySliceType, len(m))
	for key := range m {
		keys[i] = key
		i++
	}

	return keys
}
`,
	"Last": `package functions

// Last returns the last element, or zero. Also see LastOr().
func (ss SliceType) Last() ElementType {
	return ss.LastOr(ElementZeroValue)
}
`,
	"LastOr": `package functions

// LastOr returns the last element or a default value if there are no elements.
func (ss SliceType) LastOr(defaultValue ElementType) ElementType {
	if len(ss) == 0 {
		return defaultValue
	}

	return ss[len(ss)-1]
}
`,
	"Len": `package functions

// Len returns the number of elements.
func (ss SliceType) Len() int {
	return len(ss)
}
`,
	"Map": `package functions

// Map will return a new slice where each element has been mapped (transformed).
// The number of elements returned will always be the same as the input.
//
// Be careful when using this with slices of pointers. If you modify the input
// value it will affect the original slice. Be sure to return a new allocated
// object or deep copy the existing one.
func (ss SliceType) Map(fn func(ElementType) ElementType) (ss2 SliceType) {
	if ss == nil {
		return nil
	}

	ss2 = make([]ElementType, len(ss))
	for i, s := range ss {
		ss2[i] = fn(s)
	}

	return
}
`,
	"Max": `package functions

// Max is the maximum value, or zero.
func (ss SliceType) Max() (max ElementType) {
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
`,
	"Median": `package functions

// Median returns the value separating the higher half from the lower half of a
// data sample.
//
// Zero is returned if there are no elements in the slice.
func (ss SliceType) Median() ElementType {
	l := len(ss)

	switch {
	case l == 0:
		return ElementZeroValue

	case l == 1:
		return ss[0]
	}

	sorted := ss.Sort()

	if l%2 != 0 {
		return sorted[l/2]
	}

	return (sorted[l/2-1] + sorted[l/2]) / 2
}
`,
	"Min": `package functions

// Min is the minimum value, or zero.
func (ss SliceType) Min() (min ElementType) {
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
`,
	"Product": `package functions

// Product is the product of all of the elements.
func (ss SliceType) Product() (product ElementType) {
	if len(ss) == 0 {
		return
	}
	product = ss[0]
	for _, s := range ss[1:] {
		product *= s
	}

	return
}
`,
	"Random": `package functions

import (
	"math/rand"
)

// Random returns a random element by your rand.Source, or zero
func (ss SliceType) Random(source rand.Source) ElementType {
	n := len(ss)

	// Avoid the extra allocation.
	if n < 1 {
		return ElementZeroValue
	}
	if n < 2 {
		return ss[0]
	}
	rnd := rand.New(source)
	i := rnd.Intn(n)
	return ss[i]
}
`,
	"Reduce": `package functions

// Reduce continually applies the provided function
// over the slice. Reducing the elements to a single value.
//
// Returns a zero value of ElementType if there are no elements in the slice. It will panic if the reducer is nil and the slice has more than one element (required to invoke reduce).
// Otherwise returns result of applying reducer from left to right.
func (ss SliceType) Reduce(reducer func(ElementType, ElementType) ElementType) (el ElementType) {
	if len(ss) == 0 {
		return
	}
	el = ss[0]
	for _, s := range ss[1:] {
		el = reducer(el, s)
	}
	return
}
`,
	"Reverse": `package functions

// Reverse returns a new copy of the slice with the elements ordered in reverse.
// This is useful when combined with Sort to get a descending sort order:
//
//   ss.Sort().Reverse()
//
func (ss SliceType) Reverse() SliceType {
	// Avoid the allocation. If there is one element or less it is already
	// reversed.
	if len(ss) < 2 {
		return ss
	}

	sorted := make([]ElementType, len(ss))
	for i := 0; i < len(ss); i++ {
		sorted[i] = ss[len(ss)-i-1]
	}

	return sorted
}
`,
	"Send": `package functions

import (
	"context"
)

// Send sends elements to channel
// in normal act it sends all elements but if func canceled it can be less
//
// it locks execution of gorutine
// it doesn't close channel after work
// returns sended elements if len(this) != len(old) considered func was canceled
func (ss SliceType) Send(ctx context.Context, ch chan<- ElementType) SliceType {
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
`,
	"Sequence": `package functions

import "github.com/elliotchance/pie/pie/util"

// Sequence generates all numbers in range or returns nil if params invalid
//
// There are 3 variations to generate:
// 		1. [0, n).
//		2. [min, max).
//		3. [min, max) with step.
//
// if len(params) == 1 considered that will be returned slice between 0 and n,
// where n is the first param, [0, n).
// if len(params) == 2 considered that will be returned slice between min and max,
// where min is the first param, max is the second, [min, max).
// if len(params) > 2 considered that will be returned slice between min and max with step,
// where min is the first param, max is the second, step is the third one, [min, max) with step,
// others params will be ignored
func (ss SliceType) Sequence(params ...int) SliceType {
	var seq = func(min, max, step int) (seq SliceType) {
		lenght := int(util.Round(float64(max-min) / float64(step)))
		if lenght < 1 {
			return
		}

		seq = make(SliceType, lenght)
		for i := 0; i < lenght; min += step {
			seq[i] = ElementType(min)
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
`,
	"Shuffle": `package functions

import (
	"github.com/elliotchance/pie/pie/util"
	"math/rand"
)

// Shuffle returns shuffled slice by your rand.Source
func (ss SliceType) Shuffle(source rand.Source) SliceType {
	n := len(ss)

	// Avoid the extra allocation.
	if n < 2 {
		return ss
	}

	// go 1.10+ provides rnd.Shuffle. However, to support older versions we copy
	// the algorithm directly from the go source: src/math/rand/rand.go below,
	// with some adjustments:
	shuffled := make([]ElementType, n)
	copy(shuffled, ss)

	rnd := rand.New(source)

	util.Shuffle(rnd, n, func(i, j int) {
		shuffled[i], shuffled[j] = shuffled[j], shuffled[i]
	})

	return shuffled
}
`,
	"Sort": `package functions

import (
	"sort"
)

// Sort works similar to sort.SliceType(). However, unlike sort.SliceType the
// slice returned will be reallocated as to not modify the input slice.
//
// See Reverse() and AreSorted().
func (ss SliceType) Sort() SliceType {
	// Avoid the allocation. If there is one element or less it is already
	// sorted.
	if len(ss) < 2 {
		return ss
	}

	sorted := make(SliceType, len(ss))
	copy(sorted, ss)
	sort.Slice(sorted, func(i, j int) bool {
		return sorted[i] < sorted[j]
	})

	return sorted
}
`,
	"SortStableUsing": `package functions

import (
	"sort"
)

// SortStableUsing works similar to sort.SliceStable. However, unlike sort.SliceStable the
// slice returned will be reallocated as to not modify the input slice.
func (ss SliceType) SortStableUsing(less func(a, b ElementType) bool) SliceType {
	// Avoid the allocation. If there is one element or less it is already
	// sorted.
	if len(ss) < 2 {
		return ss
	}

	sorted := make(SliceType, len(ss))
	copy(sorted, ss)
	sort.SliceStable(sorted, func(i, j int) bool {
		return less(sorted[i], sorted[j])
	})

	return sorted
}
`,
	"SortUsing": `package functions

import (
	"sort"
)

// SortUsing works similar to sort.Slice. However, unlike sort.Slice the
// slice returned will be reallocated as to not modify the input slice.
func (ss SliceType) SortUsing(less func(a, b ElementType) bool) SliceType {
	// Avoid the allocation. If there is one element or less it is already
	// sorted.
	if len(ss) < 2 {
		return ss
	}

	sorted := make(SliceType, len(ss))
	copy(sorted, ss)
	sort.Slice(sorted, func(i, j int) bool {
		return less(sorted[i], sorted[j])
	})

	return sorted
}
`,
	"Strings": `package functions

import (
	"github.com/elliotchance/pie/pie"
)

// Strings transforms each element to a string.
//
// If the element type implements fmt.Stringer it will be used. Otherwise it
// will fallback to the result of:
//
//   fmt.Sprintf("%v")
//
func (ss SliceType) Strings() pie.Strings {
	l := len(ss)

	// Avoid the allocation.
	if l == 0 {
		return nil
	}

	result := make(pie.Strings, l)
	for i := 0; i < l; i++ {
		mightBeString := ss[i]
		result[i] = mightBeString.String()
	}

	return result
}
`,
	"Sum": `package functions

// Sum is the sum of all of the elements.
func (ss SliceType) Sum() (sum ElementType) {
	for _, s := range ss {
		sum += s
	}

	return
}
`,
	"ToStrings": `package functions

import (
	"github.com/elliotchance/pie/pie"
)

// ToStrings transforms each element to a string.
func (ss SliceType) ToStrings(transform func(ElementType) string) pie.Strings {
	l := len(ss)

	// Avoid the allocation.
	if l == 0 {
		return nil
	}

	result := make(pie.Strings, l)
	for i := 0; i < l; i++ {
		result[i] = transform(ss[i])
	}

	return result
}
`,
	"Top": `package functions

// Top will return n elements from head of the slice
// if the slice has less elements then n that'll return all elements
// if n < 0 it'll return empty slice.
func (ss SliceType) Top(n int) (top SliceType) {
	for i := 0; i < len(ss) && n > 0; i++ {
		top = append(top, ss[i])
		n--
	}

	return
}
`,
	"Unique": `package functions

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
func (ss SliceType) Unique() SliceType {
	// Avoid the allocation. If there is one element or less it is already
	// unique.
	if len(ss) < 2 {
		return ss
	}

	values := map[ElementType]struct{}{}

	for _, value := range ss {
		values[value] = struct{}{}
	}

	var uniqueValues SliceType
	for value := range values {
		uniqueValues = append(uniqueValues, value)
	}

	return uniqueValues
}
`,
	"Values": `package functions

// Values returns the values in the map.
//
// Due to Go's randomization of iterating maps the order is not deterministic.
func (m MapType) Values() []ElementType {
	// Avoid allocation
	l := len(m)
	if l == 0 {
		return nil
	}

	i := 0
	keys := make([]ElementType, len(m))
	for _, value := range m {
		keys[i] = value
		i++
	}

	return keys
}
`,
}
