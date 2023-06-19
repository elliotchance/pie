package pie

import (
	"context"
	"math/rand"
)

// Of encapsulates a slice to be used in multiple chained operations.
func Of[T any](ss []T) OfSlice[T] {
	return OfSlice[T]{ss}
}

// OfSlice provides the proxy methods that operate on slices. If the last method
// in the chain does not return a single value, you can access the Result to get
// final slice.
type OfSlice[T any] struct {
	Result []T
}

// All will return true if all callbacks return true. It follows the same logic
// as the all() function in Python.
//
// If the list is empty then true is always returned.
func (o OfSlice[T]) All(fn func(value T) bool) bool {
	return All(o.Result, fn)
}

// Any will return true if any callbacks return true. It follows the same logic
// as the any() function in Python.
//
// If the list is empty then false is always returned.
func (o OfSlice[T]) Any(fn func(value T) bool) bool {
	return Any(o.Result, fn)
}

// Bottom will return n elements from bottom
//
// that means that elements is taken from the end of the slice
// for this [1,2,3] slice with n == 2 will be returned [3,2]
// if the slice has less elements then n that'll return all elements
// if n < 0 it'll return empty slice.
func (o OfSlice[T]) Bottom(n int) OfSlice[T] {
	return OfSlice[T]{Bottom(o.Result, n)}
}

// DropTop will return the rest slice after dropping the top n elements
// if the slice has less elements then n that'll return empty slice
// if n < 0 it'll return empty slice.
func (o OfSlice[T]) DropTop(n int) OfSlice[T] {
	return OfSlice[T]{DropTop(o.Result, n)}
}

// Each is more condensed version of Transform that allows an action to happen
// on each elements and pass the original slice on.
//
//	pie.Each(cars, func (car *Car) {
//	    fmt.Printf("Car color is: %s\n", car.Color)
//	})
//
// Pie will not ensure immutability on items passed in so they can be
// manipulated, if you choose to do it this way, for example:
//
//	// Set all car colors to Red.
//	pie.Each(cars, func (car *Car) {
//	    car.Color = "Red"
//	})
func (o OfSlice[T]) Each(fn func(T)) OfSlice[T] {
	return OfSlice[T]{Each(o.Result, fn)}
}

// Filter will return a new slice containing only the elements that return
// true from the condition. The returned slice may contain zero elements (nil).
//
// FilterNot works in the opposite way of Filter.
func (o OfSlice[T]) Filter(condition func(T) bool) OfSlice[T] {
	return OfSlice[T]{Filter(o.Result, condition)}
}

// FilterNot works the same as Filter, with a negated condition. That is, it will
// return a new slice only containing the elements that returned false from the
// condition. The returned slice may contain zero elements (nil).
func (o OfSlice[T]) FilterNot(condition func(T) bool) OfSlice[T] {
	return OfSlice[T]{FilterNot(o.Result, condition)}
}

// FindFirstUsing will return the index of the first element when the callback
// returns true or -1 if no element is found.
// It follows the same logic as the findIndex() function in Javascript.
//
// If the list is empty then -1 is always returned.
func (o OfSlice[T]) FindFirstUsing(fn func(value T) bool) int {
	return FindFirstUsing(o.Result, fn)
}

// First returns the first element or a zero value if there are no elements.
func (o OfSlice[T]) First() T {
	return First(o.Result)
}

// FirstOr returns the first element or a default value if there are no
// elements.
func (o OfSlice[T]) FirstOr(defaultValue T) T {
	return FirstOr(o.Result, defaultValue)
}

// Insert a value at an index.
func (o OfSlice[T]) Insert(index int, values ...T) OfSlice[T] {
	return OfSlice[T]{Insert(o.Result, index, values...)}
}

// Last returns the last element or a zero value if there are no elements.
func (o OfSlice[T]) Last() T {
	return Last(o.Result)
}

// LastOr returns the last element or a default value if there are no elements.
func (o OfSlice[T]) LastOr(defaultValue T) T {
	return LastOr(o.Result, defaultValue)
}

// Map will return a new slice where each element has been mapped (transformed).
// The number of elements returned will always be the same as the input.
//
// Be careful when using this with slices of pointers. If you modify the input
// value it will affect the original slice. Be sure to return a new allocated
// object or deep copy the existing one.
func (o OfSlice[T]) Map(fn func(T) T) OfSlice[T] {
	return OfSlice[T]{Map(o.Result, fn)}
}

// Reverse returns a new copy of the slice with the elements ordered in reverse.
// This is useful when combined with Sort to get a descending sort order:
//
//	ss.Sort().Reverse()
func (o OfSlice[T]) Reverse() OfSlice[T] {
	return OfSlice[T]{Reverse(o.Result)}
}

// Send sends elements to channel
// in normal act it sends all elements but if func canceled it can be less
//
// it locks execution of gorutine
// it doesn't close channel after work
// returns sent elements if len(this) != len(old) considered func was canceled
func (o OfSlice[T]) Send(ctx context.Context, ch chan<- T) OfSlice[T] {
	return OfSlice[T]{Send(ctx, o.Result, ch)}
}

// SequenceUsing generates slice in range using creator function
//
// There are 3 variations to generate:
//  1. [0, n).
//  2. [min, max).
//  3. [min, max) with step.
//
// if len(params) == 1 considered that will be returned slice between 0 and n,
// where n is the first param, [0, n).
// if len(params) == 2 considered that will be returned slice between min and max,
// where min is the first param, max is the second, [min, max).
// if len(params) > 2 considered that will be returned slice between min and max with step,
// where min is the first param, max is the second, step is the third one, [min, max) with step,
// others params will be ignored
func (o OfSlice[T]) SequenceUsing(creator func(int) T, params ...int) OfSlice[T] {
	return OfSlice[T]{SequenceUsing(o.Result, creator, params...)}
}

// Shuffle returns a new shuffled slice by your rand.Source. The original slice
// is not modified.
func (o OfSlice[T]) Shuffle(source rand.Source) OfSlice[T] {
	return OfSlice[T]{Shuffle(o.Result, source)}
}

// SortUsing works similar to sort.Slice. However, unlike sort.Slice the
// slice returned will be reallocated as to not modify the input slice.
func (o OfSlice[T]) SortUsing(less func(a, b T) bool) OfSlice[T] {
	return OfSlice[T]{SortUsing(o.Result, less)}
}

// StringsUsing transforms each element to a string.
func (o OfSlice[T]) StringsUsing(transform func(T) string) []string {
	return StringsUsing(o.Result, transform)
}

// SubSlice will return the subSlice from start to end(excluded)
//
// Condition 1: If start < 0 or end < 0, nil is returned.
// Condition 2: If start >= end, nil is returned.
// Condition 3: Return all elements that exist in the range provided,
// if start or end is out of bounds, zero items will be placed.
func (o OfSlice[T]) SubSlice(start int, end int) OfSlice[T] {
	return OfSlice[T]{SubSlice(o.Result, start, end)}
}

// Top will return n elements from head of the slice
// if the slice has less elements then n that'll return all elements
// if n < 0 it'll return empty slice.
func (o OfSlice[T]) Top(n int) OfSlice[T] {
	return OfSlice[T]{Top(o.Result, n)}
}

// Unshift adds one or more elements to the beginning of the slice
// and returns the new slice.
func (o OfSlice[T]) Unshift(elements ...T) OfSlice[T] {
	return OfSlice[T]{Unshift(o.Result, elements...)}
}

// Removes element at index in idx from input slice, returns resulting slice.
// If an index in idx out of bounds, skip it.
func (o OfSlice[T]) Delete(idx ...int) OfSlice[T] {
	return OfSlice[T]{Delete(o.Result, idx...)}
}
