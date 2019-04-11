package main

import (
	"sort"
)

// The functions in this file only work for string and numeric slices.

// AreSorted will return true if the slice is already sorted. It is a wrapper
// for sort.SliceTypeAreSorted.
func (ss SliceType) AreSorted() bool {
	return sort.SliceIsSorted(ss, func(i, j int) bool {
		return ss[i] < ss[j]
	})
}

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

	sorted := make([]ElementType, len(ss))
	copy(sorted, ss)
	sort.Slice(sorted, func(i, j int) bool {
		return sorted[i] < sorted[j]
	})

	return sorted
}

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

// AreUnique will return true if the slice contains elements that are all
// different (unique) from each other.
func (ss SliceType) AreUnique() bool {
	return ss.Unique().Len() == ss.Len()
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
