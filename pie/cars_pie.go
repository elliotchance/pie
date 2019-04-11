package pie

import (
	"encoding/json"
)

func (ss cars) Contains(lookingFor car) bool {
	for _, s := range ss {
		if s == lookingFor {
			return true
		}
	}

	return false
}

func (ss cars) Only(condition func(car) bool) (ss2 cars) {
	for _, s := range ss {
		if condition(s) {
			ss2 = append(ss2, s)
		}
	}

	return
}

func (ss cars) Without(condition func(car) bool) (ss2 cars) {
	for _, s := range ss {
		if !condition(s) {
			ss2 = append(ss2, s)
		}
	}

	return
}

func (ss cars) Transform(fn func(car) car) (ss2 cars) {
	if ss == nil {
		return nil
	}

	ss2 = make([]car, len(ss))
	for i, s := range ss {
		ss2[i] = fn(s)
	}

	return
}

func (ss cars) FirstOr(defaultValue car) car {
	if len(ss) == 0 {
		return defaultValue
	}

	return ss[0]
}

func (ss cars) LastOr(defaultValue car) car {
	if len(ss) == 0 {
		return defaultValue
	}

	return ss[len(ss)-1]
}

func (ss cars) First() car {
	return ss.FirstOr(car{})
}

func (ss cars) Last() car {
	return ss.LastOr(car{})
}

func (ss cars) Len() int {
	return len(ss)
}

func (ss cars) JSONString() string {
	if ss == nil {
		return "[]"
	}

	data, _ := json.Marshal(ss)

	return string(data)
}

func (ss cars) Reverse() cars {

	if len(ss) < 2 {
		return ss
	}

	sorted := make([]car, len(ss))
	for i := 0; i < len(ss); i++ {
		sorted[i] = ss[len(ss)-i-1]
	}

	return sorted
}
