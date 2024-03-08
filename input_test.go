package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseInputArgs(t *testing.T) {
	var output []Interval
	var err error

	// No input
	output, err = ParseInputArgs("")
	assert.Nil(t, output)
	assert.NotNil(t, err)

	// Missing brackets
	output, _ = ParseInputArgs("3, 4")
	assert.NotNil(t, err)
	assert.Nil(t, output)

	// Single interval input
	output, _ = ParseInputArgs("[1,2]")
	assert.Equal(t, []Interval{{Start: 1, End: 2}}, output)

	// Multiple intervals, sorted
	output, _ = ParseInputArgs("[1,2][3,4]")
	assert.Equal(t, []Interval{{Start: 1, End: 2}, {Start: 3, End: 4}}, output)

	// Multiple intervals, unsorted
	output, err = ParseInputArgs("[3,4][1,2]")
	assert.Nil(t, err)
	assert.Equal(t, []Interval{{Start: 1, End: 2}, {Start: 3, End: 4}}, output)
}
