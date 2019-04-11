package pie

import (
	"encoding/json"
	"sort"
)

func (ss Float64s) Contains(lookingFor float64) bool {
	for _, s := range ss {
		if s == lookingFor {
			return true
		}
	}

	return false
}

func (ss Float64s) Only(condition func(float64) bool) (ss2 Float64s) {
	for _, s := range ss {
		if condition(s) {
			ss2 = append(ss2, s)
		}
	}

	return
}

func (ss Float64s) Without(condition func(float64) bool) (ss2 Float64s) {
	for _, s := range ss {
		if !condition(s) {
			ss2 = append(ss2, s)
		}
	}

	return
}

func (ss Float64s) Transform(fn func(float64) float64) (ss2 Float64s) {
	if ss == nil {
		return nil
	}

	ss2 = make([]float64, len(ss))
	for i, s := range ss {
		ss2[i] = fn(s)
	}

	return
}

func (ss Float64s) FirstOr(defaultValue float64) float64 {
	if len(ss) == 0 {
		return defaultValue
	}

	return ss[0]
}

func (ss Float64s) LastOr(defaultValue float64) float64 {
	if len(ss) == 0 {
		return defaultValue
	}

	return ss[len(ss)-1]
}

func (ss Float64s) First() float64 {
	return ss.FirstOr(0)
}

func (ss Float64s) Last() float64 {
	return ss.LastOr(0)
}

func (ss Float64s) Len() int {
	return len(ss)
}

func (ss Float64s) JSONString() string {
	if ss == nil {
		return "[]"
	}

	data, _ := json.Marshal(ss)

	return string(data)
}

func (ss Float64s) Reverse() Float64s {

	if len(ss) < 2 {
		return ss
	}

	sorted := make([]float64, len(ss))
	for i := 0; i < len(ss); i++ {
		sorted[i] = ss[len(ss)-i-1]
	}

	return sorted
}

func (ss Float64s) AreSorted() bool {
	return sort.SliceIsSorted(ss, func(i, j int) bool {
		return ss[i] < ss[j]
	})
}

func (ss Float64s) Sort() Float64s {

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

func (ss Float64s) AreUnique() bool {
	return ss.Unique().Len() == ss.Len()
}

func (ss Float64s) Unique() Float64s {

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

func (ss Float64s) Average() float64 {
	if l := float64(len(ss)); l > 0 {
		return float64(ss.Sum()) / float64(l)
	}

	return 0
}

func (ss Float64s) Sum() (sum float64) {
	for _, s := range ss {
		sum += s
	}

	return
}
