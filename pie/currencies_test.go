package pie

import (
	"github.com/stretchr/testify/assert"
	"sort"
	"testing"
)

func TestCurrencies_Keys(t *testing.T) {
	assert.Equal(t, []string(nil), currencies(nil).Keys())

	assert.Equal(t, []string(nil), currencies{}.Keys())

	keys := isoCurrencies.Keys()
	sort.Strings(keys)
	assert.Equal(t, []string{"AUD", "USD"}, keys)
}

func TestCurrencies_Values(t *testing.T) {
	assert.Equal(t, []currency(nil), currencies(nil).Values())

	assert.Equal(t, []currency(nil), currencies{}.Values())

	values := isoCurrencies.Values()
	sort.Slice(values, func(i, j int) bool {
		return values[i].NumericCode < values[j].NumericCode
	})

	assert.Equal(t, []currency{{36, -2}, {840, -2}}, values)
}
