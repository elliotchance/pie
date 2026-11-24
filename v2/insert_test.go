package pie_test

import (
	"github.com/elliotchance/pie/v2"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInsert(t *testing.T) {
	assert.Equal(t, []float64(nil), pie.Insert([]float64(nil), 0))
	assert.Equal(t, []float64{2.0, 1.0}, pie.Insert([]float64{1.0}, 0, 2.0))
	assert.Equal(t, []float64{1.0, 2.0}, pie.Insert([]float64{1.0}, 1, 2.0))
	assert.Equal(t, []float64{1.0, 2.0, 3.3}, pie.Insert([]float64{1.0, 3.3}, 1, 2.0))
}

func TestInsertAliasedValues(t *testing.T) {
	tests := []struct {
		name       string
		index      int
		start, end int
		want       []int
	}{
		{"before first", 0, 0, 2, []int{1, 2, 1, 2, 3, 4}},
		{"prefix", 3, 0, 1, []int{1, 2, 3, 1, 4}},
		{"empty prefix", 2, 0, 0, []int{1, 2, 3, 4}},
		{"whole input", 2, 0, 4, []int{1, 2, 1, 2, 3, 4, 3, 4}},
		{"suffix", 1, 2, 4, []int{1, 3, 4, 2, 3, 4}},
		{"middle", 2, 1, 3, []int{1, 2, 2, 3, 3, 4}},
		{"at end", 4, 0, 2, []int{1, 2, 3, 4, 1, 2}},
		{"past end", 7, 0, 2, []int{1, 2, 3, 4, 1, 2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			storage := []int{1, 2, 3, 4, 101, 102, 103, 104}
			before := append([]int(nil), storage...)
			input := storage[:4]
			got := pie.Insert(input, tt.index, input[tt.start:tt.end]...)
			assert.Equal(t, tt.want, got)
			assert.Equal(t, before, storage)

			got[0] = -1
			assert.Equal(t, before, storage)
		})
	}
}

func TestInsertPreservesSeparateInputs(t *testing.T) {
	inputStorage := []int{1, 2, 3, 4, 101, 102, 103, 104}
	valuesStorage := []int{8, 9, 201, 202, 203, 204}
	inputBefore := append([]int(nil), inputStorage...)
	valuesBefore := append([]int(nil), valuesStorage...)

	got := pie.Insert(inputStorage[:4], 1, valuesStorage[:2]...)
	assert.Equal(t, []int{1, 8, 9, 2, 3, 4}, got)
	assert.Equal(t, inputBefore, inputStorage)
	assert.Equal(t, valuesBefore, valuesStorage)

	got[0], got[1] = -1, -2
	assert.Equal(t, inputBefore, inputStorage)
	assert.Equal(t, valuesBefore, valuesStorage)
}

func TestInsertNoValues(t *testing.T) {
	assert.Nil(t, pie.Insert([]int(nil), 0))
	assert.Nil(t, pie.Insert([]int(nil), 5))
	assert.Equal(t, []int{}, pie.Insert([]int{}, 0))
	assert.Equal(t, []int{8, 9}, pie.Insert([]int(nil), 0, 8, 9))
	assert.Equal(t, []int{8, 9}, pie.Insert([]int(nil), 5, 8, 9))
	for _, index := range []int{0, 2, 4, 7} {
		input := []int{1, 2, 3, 4}
		got := pie.Insert(input, index)
		assert.Equal(t, input, got)
		got[0] = -1
		assert.Equal(t, []int{1, 2, 3, 4}, input)
	}
}

func TestOfInsertAliasedValues(t *testing.T) {
	input := []int{1, 2, 3, 4}
	got := pie.Of(input).Insert(3, input[:1]...).Result
	assert.Equal(t, []int{1, 2, 3, 1, 4}, got)
	assert.Equal(t, []int{1, 2, 3, 4}, input)
}
