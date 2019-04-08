package pie_test

import (
	"github.com/elliotchance/pie"
	"github.com/stretchr/testify/assert"
	"testing"
)

func assertImmutableStrings(t *testing.T, ss *pie.Strings) func() {
	before := (*ss).JSONString()

	return func() {
		after := (*ss).JSONString()
		assert.Equal(t, before, after)
		assert.True(t, before == after)
	}
}

func assertImmutableFloat64s(t *testing.T, ss *pie.Float64s) func() {
	before := (*ss).JSONString()

	return func() {
		after := (*ss).JSONString()
		assert.Equal(t, before, after)
		assert.True(t, before == after)
	}
}

func assertImmutableInts(t *testing.T, ss *pie.Ints) func() {
	before := (*ss).JSONString()

	return func() {
		after := (*ss).JSONString()
		assert.Equal(t, before, after)
		assert.True(t, before == after)
	}
}
