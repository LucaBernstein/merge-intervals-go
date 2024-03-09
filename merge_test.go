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

func BenchmarkMerge(b *testing.B) {
	sampleInput, _ := ParseInputArgs("[0,1][1,2][2,3][3,100][4,101][5,102][600,700][800,999][1000,1000][1001,1002][1002,1004][1003,10000]")
	var intervals []Interval
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		intervals = Merge(sampleInput)
	}
	b.StopTimer()
	formattedOutput := FormatOutput(intervals)
	if formattedOutput != "[0,102] [600,700] [800,999] [1000,1000] [1001,10000]" {
		b.Errorf("unexpected output: %s", formattedOutput)
	}
}

func fillIntersectingInterval(size int) (sample []Interval) {
	for i := range size {
		sample = append(sample, Interval{Start: i, End: i + 2})
	}
	return
}

func benchmarkMergeSkeleton(intervalCount int, b *testing.B) {
	sampleInput := fillIntersectingInterval(intervalCount)
	var intervals []Interval
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		intervals = Merge(sampleInput)
	}
	b.StopTimer()
	expectedInterval := Interval{Start: 0, End: intervalCount + 1}
	if len(intervals) != 1 || intervals[0] != expectedInterval {
		b.Errorf("unexpected output: %v", intervals)
	}
}

func BenchmarkMerge1000(b *testing.B)   { benchmarkMergeSkeleton(1000, b) }
func BenchmarkMerge10000(b *testing.B)  { benchmarkMergeSkeleton(10000, b) }
func BenchmarkMerge100000(b *testing.B) { benchmarkMergeSkeleton(100000, b) }
