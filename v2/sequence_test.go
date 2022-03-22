package pie_test

import (
	"github.com/elliotchance/pie/v2"
	"github.com/stretchr/testify/assert"
	"testing"
)

var sequenceAndSequenceUsingTests = []struct {
	ss       []float64
	params   []int
	expected []float64
}{
	// n
	{
		nil,
		nil,
		nil,
	},
	{
		nil,
		[]int{-1},
		nil,
	},
	{
		nil,
		[]int{0},
		nil,
	},
	{
		nil,
		[]int{3},
		[]float64{0, 1, 2},
	},
	{
		[]float64{},
		[]int{3},
		[]float64{0, 1, 2},
	},
	// range
	{
		nil,
		[]int{2, 2},
		nil,
	},
	{
		[]float64{},
		[]int{3, 2},
		nil,
	},
	{
		nil,
		[]int{0, 3},
		[]float64{0, 1, 2},
	},
	{
		[]float64{},
		[]int{3, 6},
		[]float64{3, 4, 5},
	},
	{
		[]float64{},
		[]int{-5, 0},
		[]float64{-5, -4, -3, -2, -1},
	},
	{
		[]float64{},
		[]int{-5, -10},
		nil,
	},
	// range with step
	{
		nil,
		[]int{3, 3, 1},
		nil,
	},
	{
		[]float64{},
		[]int{3, 6, 2},
		[]float64{3, 5},
	},
	{
		[]float64{},
		[]int{3, 7, 2},
		[]float64{3, 5},
	},
	{
		[]float64{},
		[]int{-10, -6, 1},
		[]float64{-10, -9, -8, -7},
	},
	{
		[]float64{},
		[]int{-6, -10, -1},
		[]float64{-6, -7, -8, -9},
	},
	{
		[]float64{},
		[]int{-6, -10, 1},
		nil,
	},
}

func TestSequence(t *testing.T) {
	for _, test := range sequenceAndSequenceUsingTests {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, test.expected, pie.Sequence(test.ss, test.params...))
		})
	}
}
