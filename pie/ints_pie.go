package pie

import (
	"encoding/json"
	"sort"
)

func (ss Ints) Contains(lookingFor int) bool {
	for _, s := range ss {
		if s == lookingFor {
			return true
		}
	}

	return false
}

func (ss Ints) Select(condition func(int) bool) (ss2 Ints) {
	for _, s := range ss {
		if condition(s) {
			ss2 = append(ss2, s)
		}
	}

	return
}

func (ss Ints) Unselect(condition func(int) bool) (ss2 Ints) {
	for _, s := range ss {
		if !condition(s) {
			ss2 = append(ss2, s)
		}
	}

	return
}

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

func (ss Ints) FirstOr(defaultValue int) int {
	if len(ss) == 0 {
		return defaultValue
	}

	return ss[0]
}

func (ss Ints) LastOr(defaultValue int) int {
	if len(ss) == 0 {
		return defaultValue
	}

	return ss[len(ss)-1]
}

func (ss Ints) First() int {
	return ss.FirstOr(0)
}

func (ss Ints) Last() int {
	return ss.LastOr(0)
}

func (ss Ints) Len() int {
	return len(ss)
}

func (ss Ints) JSONString() string {
	if ss == nil {
		return "[]"
	}

	data, _ := json.Marshal(ss)

	return string(data)
}

func (ss Ints) Reverse() Ints {

	if len(ss) < 2 {
		return ss
	}

	sorted := make([]int, len(ss))
	for i := 0; i < len(ss); i++ {
		sorted[i] = ss[len(ss)-i-1]
	}

	return sorted
}

func (ss Ints) ToStrings(transform func(int) string) Strings {
	l := len(ss)

	if l == 0 {
		return nil
	}

	result := make(Strings, l)
	for i := 0; i < l; i++ {
		result[i] = transform(ss[i])
	}

	return result
}

func (ss Ints) AreSorted() bool {
	return sort.SliceIsSorted(ss, func(i, j int) bool {
		return ss[i] < ss[j]
	})
}

func (ss Ints) Sort() Ints {

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

func (ss Ints) AreUnique() bool {
	return ss.Unique().Len() == ss.Len()
}

func (ss Ints) Unique() Ints {

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

func (ss Ints) Average() float64 {
	if l := int(len(ss)); l > 0 {
		return float64(ss.Sum()) / float64(l)
	}

	return 0
}

func (ss Ints) Sum() (sum int) {
	for _, s := range ss {
		sum += s
	}

	return
}
