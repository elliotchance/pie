package pie

import (
	testify_stats "github.com/elliotchance/testify-stats"
	"github.com/elliotchance/testify-stats/assert"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	os.Exit(testify_stats.Run(m))
}

func assertImmutableStrings(t *testing.T, ss *Strings) func() {
	before := (*ss).JSONString()

	return func() {
		after := (*ss).JSONString()
		assert.Equal(t, before, after)
		assert.True(t, before == after)
	}
}

func assertImmutableInts(t *testing.T, ss *Ints) func() {
	before := (*ss).JSONString()

	return func() {
		after := (*ss).JSONString()
		assert.Equal(t, before, after)
		assert.True(t, before == after)
	}
}

func assertImmutableFloat64s(t *testing.T, ss *Float64s) func() {
	before := (*ss).JSONString()

	return func() {
		after := (*ss).JSONString()
		assert.Equal(t, before, after)
		assert.True(t, before == after)
	}
}

func assertImmutableCars(t *testing.T, ss *cars) func() {
	before := (*ss).JSONString()

	return func() {
		after := (*ss).JSONString()
		assert.Equal(t, before, after)
		assert.True(t, before == after)
	}
}

func assertImmutableCarPointers(t *testing.T, ss *carPointers) func() {
	before := (*ss).JSONString()

	return func() {
		after := (*ss).JSONString()
		assert.Equal(t, before, after)
		assert.True(t, before == after)
	}
}
