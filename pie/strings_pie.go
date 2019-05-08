package pie

import (
	"context"
	"encoding/json"
	"github.com/elliotchance/pie/pie/util"
	"math/rand"
	"sort"
)

// All will return true if all callbacks return true. It follows the same logic
// as the all() function in Python.
//
// If the list is empty then true is always returned.
func (ss Strings) All(fn func(value string) bool) bool {
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
func (ss Strings) Any(fn func(value string) bool) bool {
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
func (ss Strings) Append(elements ...string) Strings {
	return append(ss, elements...)
}

// AreSorted will return true if the slice is already sorted. It is a wrapper
// for sort.StringsAreSorted.
func (ss Strings) AreSorted() bool {
	return sort.SliceIsSorted(ss, func(i, j int) bool {
		return ss[i] < ss[j]
	})
}

// AreUnique will return true if the slice contains elements that are all
// different (unique) from each other.
func (ss Strings) AreUnique() bool {
	return ss.Unique().Len() == ss.Len()
}

// Bottom will return n elements from bottom
//
// that means that elements is taken from the end of the slice
// for this [1,2,3] slice with n == 2 will be returned [3,2]
// if the slice has less elements then n that'll return all elements
// if n < 0 it'll return empty slice.
func (ss Strings) Bottom(n int) (top Strings) {
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
func (ss Strings) Contains(lookingFor string) bool {
	for _, s := range ss {
		if lookingFor == s {
			return true
		}
	}

	return false
}

// Diff returns the elements that needs to be added or removed from the first
// slice to have the same elements in the second slice.
//
// The order of elements is not taken into consideration, so the slices are
// treated sets that allow duplicate items.
//
// The added and removed returned may be blank respectively, or contain upto as
// many elements that exists in the largest slice.
func (ss Strings) Diff(against Strings) (added, removed Strings) {
	// This is probably not the best way to do it. We do an O(n^2) between the
	// slices to see which items are missing in each direction.

	diffOneWay := func(ss1, ss2raw Strings) (result Strings) {
		ss2 := make(Strings, len(ss2raw))
		copy(ss2, ss2raw)

		for _, s := range ss1 {
			found := false

			for i, element := range ss2 {
				if s == element {
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
func (ss Strings) Each(fn func(string)) Strings {
	for _, s := range ss {
		fn(s)
	}

	return ss
}

// Extend will return a new slice with the slices of elements appended to the
// end.
//
// It is acceptable to provide zero arguments.
func (ss Strings) Extend(slices ...Strings) (ss2 Strings) {
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
func (ss Strings) Filter(condition func(string) bool) (ss2 Strings) {
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
func (ss Strings) FilterNot(condition func(string) bool) (ss2 Strings) {
	for _, s := range ss {
		if !condition(s) {
			ss2 = append(ss2, s)
		}
	}

	return
}

// First returns the first element, or zero. Also see FirstOr().
func (ss Strings) First() string {
	return ss.FirstOr("")
}

// FirstOr returns the first element or a default value if there are no
// elements.
func (ss Strings) FirstOr(defaultValue string) string {
	if len(ss) == 0 {
		return defaultValue
	}

	return ss[0]
}

// Intersect returns items that exist in all lists.
//
// It returns slice without any duplicates.
// If zero slice arguments are provided, then nil is returned.
func (ss Strings) Intersect(slices ...Strings) (ss2 Strings) {
	if slices == nil {
		return nil
	}

	var uniqs = make([]map[string]struct{}, len(slices))
	for i := 0; i < len(slices); i++ {
		m := make(map[string]struct{})
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

// Join returns a string from joining each of the elements.
func (ss Strings) Join(glue string) (s string) {
	for i, element := range ss {
		if i > 0 {
			s += glue
		}

		s += string(element)
	}

	return s
}

// JSONString returns the JSON encoded array as a string.
//
// One important thing to note is that it will treat a nil slice as an empty
// slice to ensure that the JSON value return is always an array.
func (ss Strings) JSONString() string {
	if ss == nil {
		return "[]"
	}

	// An error should not be possible.
	data, _ := json.Marshal(ss)

	return string(data)
}

// Last returns the last element, or zero. Also see LastOr().
func (ss Strings) Last() string {
	return ss.LastOr("")
}

// LastOr returns the last element or a default value if there are no elements.
func (ss Strings) LastOr(defaultValue string) string {
	if len(ss) == 0 {
		return defaultValue
	}

	return ss[len(ss)-1]
}

// Len returns the number of elements.
func (ss Strings) Len() int {
	return len(ss)
}

// Map will return a new slice where each element has been mapped (transformed).
// The number of elements returned will always be the same as the input.
//
// Be careful when using this with slices of pointers. If you modify the input
// value it will affect the original slice. Be sure to return a new allocated
// object or deep copy the existing one.
func (ss Strings) Map(fn func(string) string) (ss2 Strings) {
	if ss == nil {
		return nil
	}

	ss2 = make([]string, len(ss))
	for i, s := range ss {
		ss2[i] = fn(s)
	}

	return
}

// Max is the maximum value, or zero.
func (ss Strings) Max() (max string) {
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

// Min is the minimum value, or zero.
func (ss Strings) Min() (min string) {
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
func (ss Strings) Random(source rand.Source) string {
	n := len(ss)

	// Avoid the extra allocation.
	if n < 1 {
		return ""
	}
	if n < 2 {
		return ss[0]
	}
	rnd := rand.New(source)
	i := rnd.Intn(n)
	return ss[i]
}

// Reduce continually applies the provided function
// over the slice. Reducing the elements to a single value.
//
// Returns a zero value of string if there are no elements in the slice. It will panic if the reducer is nil and the slice has more than one element (required to invoke reduce).
// Otherwise returns result of applying reducer from left to right.
func (ss Strings) Reduce(reducer func(string, string) string) (el string) {
	if len(ss) == 0 {
		return
	}
	el = ss[0]
	for _, s := range ss[1:] {
		el = reducer(el, s)
	}
	return
}

// Reverse returns a new copy of the slice with the elements ordered in reverse.
// This is useful when combined with Sort to get a descending sort order:
//
//   ss.Sort().Reverse()
//
func (ss Strings) Reverse() Strings {
	// Avoid the allocation. If there is one element or less it is already
	// reversed.
	if len(ss) < 2 {
		return ss
	}

	sorted := make([]string, len(ss))
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
func (ss Strings) Send(ctx context.Context, ch chan<- string) Strings {
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

// Sort works similar to sort.Strings(). However, unlike sort.Strings the
// slice returned will be reallocated as to not modify the input slice.
//
// See Reverse() and AreSorted().
func (ss Strings) Sort() Strings {
	// Avoid the allocation. If there is one element or less it is already
	// sorted.
	if len(ss) < 2 {
		return ss
	}

	sorted := make(Strings, len(ss))
	copy(sorted, ss)
	sort.Slice(sorted, func(i, j int) bool {
		return sorted[i] < sorted[j]
	})

	return sorted
}

// SortUsing works similar to sort.Slice. However, unlike sort.Slice the
// slice returned will be reallocated as to not modify the input slice.
func (ss Strings) SortUsing(less func(a, b string) bool) Strings {
	// Avoid the allocation. If there is one element or less it is already
	// sorted.
	if len(ss) < 2 {
		return ss
	}

	sorted := make(Strings, len(ss))
	copy(sorted, ss)
	sort.Slice(sorted, func(i, j int) bool {
		return less(sorted[i], sorted[j])
	})

	return sorted
}

// SortStableUsing works similar to sort.SliceStable. However, unlike sort.SliceStable the
// slice returned will be reallocated as to not modify the input slice.
func (ss Strings) SortStableUsing(less func(a, b string) bool) Strings {
	// Avoid the allocation. If there is one element or less it is already
	// sorted.
	if len(ss) < 2 {
		return ss
	}

	sorted := make(Strings, len(ss))
	copy(sorted, ss)
	sort.SliceStable(sorted, func(i, j int) bool {
		return less(sorted[i], sorted[j])
	})

	return sorted
}

// Shuffle returns shuffled slice by your rand.Source
func (ss Strings) Shuffle(source rand.Source) Strings {
	n := len(ss)

	// Avoid the extra allocation.
	if n < 2 {
		return ss
	}

	// go 1.10+ provides rnd.Shuffle. However, to support older versions we copy
	// the algorithm directly from the go source: src/math/rand/rand.go below,
	// with some adjustments:
	shuffled := make([]string, n)
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
func (ss Strings) Top(n int) (top Strings) {
	for i := 0; i < len(ss) && n > 0; i++ {
		top = append(top, ss[i])
		n--
	}

	return
}

// ToStrings transforms each element to a string.
func (ss Strings) ToStrings(transform func(string) string) Strings {
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
func (ss Strings) Unique() Strings {
	// Avoid the allocation. If there is one element or less it is already
	// unique.
	if len(ss) < 2 {
		return ss
	}

	values := map[string]struct{}{}

	for _, value := range ss {
		values[value] = struct{}{}
	}

	var uniqueValues Strings
	for value := range values {
		uniqueValues = append(uniqueValues, value)
	}

	return uniqueValues
}
