package pie

import (
	"encoding/json"
)

func (ss carPointers) Contains(lookingFor *car) bool {
	for _, s := range ss {
		if s == lookingFor {
			return true
		}
	}

	return false
}

func (ss carPointers) Select(condition func(*car) bool) (ss2 carPointers) {
	for _, s := range ss {
		if condition(s) {
			ss2 = append(ss2, s)
		}
	}

	return
}

func (ss carPointers) Unselect(condition func(*car) bool) (ss2 carPointers) {
	for _, s := range ss {
		if !condition(s) {
			ss2 = append(ss2, s)
		}
	}

	return
}

func (ss carPointers) Transform(fn func(*car) *car) (ss2 carPointers) {
	if ss == nil {
		return nil
	}

	ss2 = make([]*car, len(ss))
	for i, s := range ss {
		ss2[i] = fn(s)
	}

	return
}

func (ss carPointers) FirstOr(defaultValue *car) *car {
	if len(ss) == 0 {
		return defaultValue
	}

	return ss[0]
}

func (ss carPointers) LastOr(defaultValue *car) *car {
	if len(ss) == 0 {
		return defaultValue
	}

	return ss[len(ss)-1]
}

func (ss carPointers) First() *car {
	return ss.FirstOr(&car{})
}

func (ss carPointers) Last() *car {
	return ss.LastOr(&car{})
}

func (ss carPointers) Len() int {
	return len(ss)
}

func (ss carPointers) JSONString() string {
	if ss == nil {
		return "[]"
	}

	data, _ := json.Marshal(ss)

	return string(data)
}

func (ss carPointers) Reverse() carPointers {

	if len(ss) < 2 {
		return ss
	}

	sorted := make([]*car, len(ss))
	for i := 0; i < len(ss); i++ {
		sorted[i] = ss[len(ss)-i-1]
	}

	return sorted
}

func (ss carPointers) ToStrings(transform func(*car) string) Strings {
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
