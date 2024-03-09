package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMerge(t *testing.T) {
	assert.Equal(t, []Interval{}, Merge([]Interval{}))

	// Idempotency: An already merged interval should stay the same
	assert.Equal(t, []Interval{{Start: 1, End: 2}}, Merge([]Interval{{Start: 1, End: 2}}))
	assert.Equal(t, []Interval{{Start: 1, End: 2}, {Start: 7, End: 8}}, Merge([]Interval{{Start: 1, End: 2}, {Start: 7, End: 8}}))

	// Merge two identical intervals
	assert.Equal(t, []Interval{{Start: 1, End: 2}}, Merge([]Interval{{Start: 1, End: 2}, {Start: 1, End: 2}}))

	// Merge two intervals where one interval has 0 length
	assert.Equal(t, []Interval{{Start: 1, End: 2}}, Merge([]Interval{{Start: 1, End: 1}, {Start: 1, End: 2}}))

	// Merge completely overlapping intervals
	assert.Equal(t, []Interval{{Start: 1, End: 5}}, Merge([]Interval{{Start: 1, End: 5}, {Start: 2, End: 4}}))

	// Merge partially overlapping intervals
	assert.Equal(t, []Interval{{Start: 1, End: 6}}, Merge([]Interval{{Start: 1, End: 5}, {Start: 2, End: 6}}))
}
