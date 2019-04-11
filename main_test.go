package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

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
