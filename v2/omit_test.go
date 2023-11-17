package pie_test

import (
	"github.com/elliotchance/pie/v2"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOmit(t *testing.T) {
	assert.Equal(t,
		map[string]int{"b": 2},
		pie.Omit([]string{"a"}, map[string]int{"a": 1, "b": 2}),
	)

	assert.Equal(t,
		map[int]string{1: "one", 3: "three"},
		pie.Omit([]int{2}, map[int]string{1: "one", 2: "two", 3: "three"}),
	)

	assert.Equal(t,
		map[string]string{},
		pie.Omit([]string{"x", "y", "z"}, map[string]string{}),
	)

	assert.Equal(t,
		map[string]int{"a": 1, "b": 2, "c": 3},
		pie.Omit([]string{}, map[string]int{"a": 1, "b": 2, "c": 3}),
	)

	assert.Equal(t,
		map[string]int{"a": 1, "b": 2, "c": 3},
		pie.Omit([]string{"x", "y", "z"}, map[string]int{"a": 1, "b": 2, "c": 3}),
	)
}
