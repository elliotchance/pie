package pie

import (
	"encoding/json"
	"sort"
)

func (ss Strings) Contains(lookingFor string) bool {
	for _, s := range ss {
		if s == lookingFor {
			return true
		}
	}

	return false
}

func (ss Strings) Select(condition func(string) bool) (ss2 Strings) {
	for _, s := range ss {
		if condition(s) {
			ss2 = append(ss2, s)
		}
	}

	return
}

func (ss Strings) Unselect(condition func(string) bool) (ss2 Strings) {
	for _, s := range ss {
		if !condition(s) {
			ss2 = append(ss2, s)
		}
	}

	return
}

func (ss Strings) Transform(fn func(string) string) (ss2 Strings) {
	if ss == nil {
		return nil
	}

	ss2 = make([]string, len(ss))
	for i, s := range ss {
		ss2[i] = fn(s)
	}

	return
}

func (ss Strings) FirstOr(defaultValue string) string {
	if len(ss) == 0 {
		return defaultValue
	}

	return ss[0]
}

func (ss Strings) LastOr(defaultValue string) string {
	if len(ss) == 0 {
		return defaultValue
	}

	return ss[len(ss)-1]
}

func (ss Strings) First() string {
	return ss.FirstOr("")
}

func (ss Strings) Last() string {
	return ss.LastOr("")
}

func (ss Strings) Len() int {
	return len(ss)
}

func (ss Strings) JSONString() string {
	if ss == nil {
		return "[]"
	}

	data, _ := json.Marshal(ss)

	return string(data)
}

func (ss Strings) Reverse() Strings {

	if len(ss) < 2 {
		return ss
	}

	sorted := make([]string, len(ss))
	for i := 0; i < len(ss); i++ {
		sorted[i] = ss[len(ss)-i-1]
	}

	return sorted
}

func (ss Strings) ToStrings(transform func(string) string) Strings {
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

func (ss Strings) AreSorted() bool {
	return sort.SliceIsSorted(ss, func(i, j int) bool {
		return ss[i] < ss[j]
	})
}

func (ss Strings) Sort() Strings {

	if len(ss) < 2 {
		return ss
	}

	sorted := make([]string, len(ss))
	copy(sorted, ss)
	sort.Slice(sorted, func(i, j int) bool {
		return sorted[i] < sorted[j]
	})

	return sorted
}

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

func (ss Strings) AreUnique() bool {
	return ss.Unique().Len() == ss.Len()
}

func (ss Strings) Unique() Strings {

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

func (ss Strings) Join(glue string) (s string) {
	for i, element := range ss {
		if i > 0 {
			s += glue
		}

		s += string(element)
	}

	return s
}
