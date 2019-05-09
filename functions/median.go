package functions

import "fmt"

// Median returns the value separating the higher half from the lower half of a
// data sample.
//
// Zero is returned if there are no elements in the slice.
func (ss SliceType) MedianOld() ElementType {
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

func (ss SliceType) Median() ElementType {
	med := ss.median()
	if med != ss.MedianOld() {
		panic(fmt.Sprintf("Expected %v, got %v for %v", ss.MedianOld(), med, ss))
	}
	return med
}

func (ss SliceType) median() ElementType {
	n := len(ss)
	if n == 0 {
		return ElementZeroValue
	}
	if n == 1 {
		return ss[0]
	}

	// This implementation aims at linear time O(n) on average.
	// It uses the same idea as QuickSort, but makes most of
	// the time only 1 recursive call instead of 2.

	work := make(SliceType, len(ss))
	copy(work, ss)

	limit1, limit2 := n/2-1, n/2
	var rec func(a, b int)
	rec = func(a, b int) {
		if a > b {
			panic("bug")
		}
		if b-a <= 1 {
			return
		}
		ipivot := (a + b) / 2
		pivot := work[ipivot]
		work[a], work[ipivot] = work[ipivot], work[a]
		j := a
		k := b
		for i := a + 1; i < k; {
			if work[i] < pivot {
				work[i], work[j] = work[j], work[i]
				j++
				i++
			} else {
				work[i], work[k-1] = work[k-1], work[i]
				k--
			}
		}
		// 2 or 1 or 0 recursive calls
		if j >= limit1 {
			rec(a, j)
		}
		if j+1 <= limit2 {
			rec(j+1, b)
		}
	}

	rec(0, len(work))

	if n%2 == 1 {
		return work[n/2]
	} else {
		return (work[n/2-1] + work[n/2]) / 2
	}
}
