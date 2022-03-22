package pie_test

import (
	"github.com/elliotchance/pie/v2"
	"github.com/stretchr/testify/assert"
	"sort"
	"testing"
)

func TestValues(t *testing.T) {
	assert.Equal(t, []currency(nil), pie.Values(currencies(nil)))

	assert.Equal(t, []currency(nil), pie.Values(currencies{}))

	values := pie.Values(isoCurrencies)
	sort.Slice(values, func(i, j int) bool {
		return values[i].NumericCode < values[j].NumericCode
	})

	assert.Equal(t, []currency{{36, -2}, {840, -2}}, values)
}
