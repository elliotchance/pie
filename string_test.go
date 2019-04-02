package pie_test

import (
	"github.com/elliotchance/pie"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPrefix(t *testing.T) {
	for _, test := range []struct {
		s        string
		prefix   string
		expected bool
	}{
		{"", "a", false},
		{"", "", true},
		{"a", "a", true},
		{"ab", "a", true},
		{"a", "b", false},
		{"a", "A", false},
		{"ab", "b", false},
	} {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, test.expected, pie.Prefix(test.prefix)(test.s))
		})
	}
}

func TestSuffix(t *testing.T) {
	for _, test := range []struct {
		s        string
		suffix   string
		expected bool
	}{
		{"", "a", false},
		{"", "", true},
		{"a", "a", true},
		{"ab", "a", false},
		{"a", "b", false},
		{"a", "A", false},
		{"ab", "b", true},
	} {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, test.expected, pie.Suffix(test.suffix)(test.s))
		})
	}
}
