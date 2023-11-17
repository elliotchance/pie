package pie_test

import (
	"github.com/elliotchance/pie/v2"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPick(t *testing.T) {
	assert.Equal(t,
		map[string]string{"a": "1", "c": "3"},
		pie.Pick([]string{"a", "c"}, map[string]string{"a": "1", "b": "2", "c": "3"}),
	)

	assert.Equal(t,
		map[int]string{2: "two", 3: "three"},
		pie.Pick([]int{2, 3}, map[int]string{1: "one", 2: "two", 3: "three"}),
	)

	assert.Equal(t,
		map[string]string{},
		pie.Pick([]string{"x", "y", "z"}, map[string]string{}),
	)

	assert.Equal(t,
		map[string]int{},
		pie.Pick([]string{}, map[string]int{"a": 1, "b": 2, "c": 3}),
	)

	assert.Equal(t,
		map[string]int{},
		pie.Pick([]string{"x", "y", "z"}, map[string]int{"a": 1, "b": 2, "c": 3}),
	)
}
