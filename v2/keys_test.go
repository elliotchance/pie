package pie_test

import (
	"github.com/elliotchance/pie/v2"
	"github.com/stretchr/testify/assert"
	"sort"
	"testing"
)

type currency struct {
	NumericCode, Exponent int
}

type currencies map[string]currency

var isoCurrencies = currencies{
	"AUD": {36, -2},
	"USD": {840, -2},
}

func TestKeys(t *testing.T) {
	assert.Equal(t, []string(nil), pie.Keys(currencies(nil)))

	assert.Equal(t, []string(nil), pie.Keys(currencies{}))

	keys := pie.Keys(isoCurrencies)
	sort.Strings(keys)
	assert.Equal(t, []string{"AUD", "USD"}, keys)
}
