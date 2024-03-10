package interval_test

import (
	"testing"

	"github.com/LucaBernstein/merge-intervals-go/interval"
	"github.com/stretchr/testify/assert"
)

func TestMerge(t *testing.T) {
	assert.Equal(t, []interval.Interval{}, interval.Merge([]interval.Interval{}))

	// Idempotency: An already merged interval should stay the same
	assert.Equal(t, []interval.Interval{{Start: 1, End: 2}}, interval.Merge([]interval.Interval{{Start: 1, End: 2}}))
	assert.Equal(t, []interval.Interval{{Start: 1, End: 2}, {Start: 7, End: 8}}, interval.Merge([]interval.Interval{{Start: 1, End: 2}, {Start: 7, End: 8}}))

	// Merge two identical intervals
	assert.Equal(t, []interval.Interval{{Start: 1, End: 2}}, interval.Merge([]interval.Interval{{Start: 1, End: 2}, {Start: 1, End: 2}}))

	// Merge two intervals where one interval has 0 length
	assert.Equal(t, []interval.Interval{{Start: 1, End: 2}}, interval.Merge([]interval.Interval{{Start: 1, End: 1}, {Start: 1, End: 2}}))

	// Merge completely overlapping intervals
	assert.Equal(t, []interval.Interval{{Start: 1, End: 5}}, interval.Merge([]interval.Interval{{Start: 1, End: 5}, {Start: 2, End: 4}}))

	// Merge partially overlapping intervals
	assert.Equal(t, []interval.Interval{{Start: 1, End: 6}}, interval.Merge([]interval.Interval{{Start: 1, End: 5}, {Start: 2, End: 6}}))

	// Merge touching intervals
	assert.Equal(t, []interval.Interval{{Start: 1, End: 6}}, interval.Merge([]interval.Interval{{Start: 1, End: 5}, {Start: 5, End: 6}}))

	// Don't merge neighboring intervals
	assert.Equal(t, []interval.Interval{{Start: 1, End: 6}, {Start: 7, End: 10}}, interval.Merge([]interval.Interval{{Start: 1, End: 6}, {Start: 7, End: 10}}))

	// Merge intervals with negative numbers
	assert.Equal(t, []interval.Interval{{Start: -11, End: -3}, {Start: -1, End: 1}}, interval.Merge([]interval.Interval{{Start: -11, End: -3}, {Start: -1, End: 0}, {Start: 0, End: 1}}))
}

func BenchmarkMerge(b *testing.B) {
	sampleInput, _ := interval.ParseInput("[0,1][1,2][2,3][3,100][4,101][5,102][600,700][800,999][1000,1000][1001,1002][1002,1004][1003,10000]")
	var intervals []interval.Interval
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		intervals = interval.Merge(sampleInput)
	}
	b.StopTimer()
	formattedOutput := interval.FormatOutput(intervals)
	if formattedOutput != "[0,102] [600,700] [800,999] [1000,1000] [1001,10000]" {
		b.Errorf("unexpected output: %s", formattedOutput)
	}
}

func fillIntersectingInterval(size int, overlap bool) (sample []interval.Interval) {
	for i := range size {
		if overlap {
			sample = append(sample, interval.Interval{Start: i, End: i + 2})
			continue
		}
		sample = append(sample, interval.Interval{Start: -i, End: -i})
	}
	return
}

func benchmarkMergeSkeleton(intervalCount int, overlap bool, b *testing.B) {
	sampleInput := fillIntersectingInterval(intervalCount, overlap)
	var intervals []interval.Interval
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		intervals = interval.Merge(sampleInput)
	}
	b.StopTimer()
	if overlap {
		expectedInterval := interval.Interval{Start: 0, End: intervalCount + 1}
		if len(intervals) != 1 || intervals[0] != expectedInterval {
			b.Errorf("unexpected output: %v", intervals)
		}
	} else {
		if len(intervals) != intervalCount {
			b.Errorf("unexpected output interval count: %d", len(intervals))
		}
	}
}

func BenchmarkMergeOverlap1000(b *testing.B)    { benchmarkMergeSkeleton(1000, true, b) }
func BenchmarkMergeOverlap10000(b *testing.B)   { benchmarkMergeSkeleton(10000, true, b) }
func BenchmarkMergeOverlap100000(b *testing.B)  { benchmarkMergeSkeleton(100000, true, b) }
func BenchmarkMergeOverlap1000000(b *testing.B) { benchmarkMergeSkeleton(1000000, true, b) }

func BenchmarkMergeDistinct1000(b *testing.B)     { benchmarkMergeSkeleton(1000, true, b) }
func BenchmarkMergeDistinct10000(b *testing.B)    { benchmarkMergeSkeleton(10000, true, b) }
func BenchmarkMergeDistinct100000(b *testing.B)   { benchmarkMergeSkeleton(100000, true, b) }
func BenchmarkMergeDistinct1000000(b *testing.B)  { benchmarkMergeSkeleton(1000000, true, b) }
func BenchmarkMergeDistinct10000000(b *testing.B) { benchmarkMergeSkeleton(10000000, true, b) }
