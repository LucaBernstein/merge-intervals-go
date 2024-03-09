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

	// Merge touching intervals
	assert.Equal(t, []Interval{{Start: 1, End: 6}}, Merge([]Interval{{Start: 1, End: 5}, {Start: 5, End: 6}}))

	// Don't merge neighboring intervals
	assert.Equal(t, []Interval{{Start: 1, End: 6}, {Start: 7, End: 10}}, Merge([]Interval{{Start: 1, End: 6}, {Start: 7, End: 10}}))

	// Merge intervals with negative numbers
	assert.Equal(t, []Interval{{Start: -11, End: -3}, {Start: -1, End: 1}}, Merge([]Interval{{Start: -11, End: -3}, {Start: -1, End: 0}, {Start: 0, End: 1}}))
}
