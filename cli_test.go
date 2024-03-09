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

	// Multiple intervals, at least one with zero length
	output, err = ParseInputArgs("[3,4][3,3]")
	assert.Nil(t, err)
	assert.Equal(t, []Interval{{Start: 3, End: 3}, {Start: 3, End: 4}}, output) // Sorted by Start, End
}

func BenchmarkParseInputArgs(b *testing.B) {
	var intervals []Interval
	var err error
	for n := 0; n < b.N; n++ {
		intervals, err = ParseInputArgs("[7,18][64000,128000][1,100000][2112,4224][0,1][25,117][1998,2401][2024,2025][15,20][117122,122117]")
	}
	// Assert recorded values from benchmarking test to avoid compiler optimizations (and possibly elimination of the function at test)
	if err != nil {
		b.Errorf("error in function while running benchmark: %s, intervals: %v", err.Error(), intervals)
	}
}

func TestFormatOutput(t *testing.T) {
	assert.Equal(t, "[7,15] [26,200]", FormatOutput([]Interval{{Start: 7, End: 15}, {Start: 26, End: 200}}))
}
