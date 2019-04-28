package pie

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// There tests are just to make sure that the select functions for myInts are
// generated. The more extensive tests for these functions are in ints_test.go

func TestMyInts_Average(t *testing.T) {
	assert.Equal(t, 0.0, myInts(nil).Average())
	assert.Equal(t, 4.333333333333333, myInts{1, 5, 7}.Average())
}

func TestMyInts_Sum(t *testing.T) {
	assert.Equal(t, 13, myInts{1, 5, 7}.Sum())
}
