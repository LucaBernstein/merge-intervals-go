package interval_test

import (
	"log/slog"
	"os"
	"testing"

	"github.com/LucaBernstein/merge-intervals-go/interval"
	"github.com/stretchr/testify/assert"
)

func TestLoadFromFile(t *testing.T) {
	err := os.WriteFile("test-interval-file.txt", []byte("[1,7] [14,19] [6,12]"), os.ModePerm)
	assert.Nil(t, err)
	assert.Equal(t, "[1,7][14,19][6,12]", interval.LoadFromFile("test-interval-file.txt"))
}

func TestLoadFromFileNotExist(t *testing.T) {
	slog.SetLogLoggerLevel(slog.LevelError + 1) // Don't log expected error
	assert.Equal(t, "", interval.LoadFromFile("test-interval-file-does-not-exist.txt"))
}

func TestLoadFromArgs(t *testing.T) {
	// Note former args and restore after test execution
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()

	os.Args = []string{"merge-intervals-go", "[1, 2] [3, 4]", "[7,8]"} // Intervals can be passed in multiple arg splits

	assert.Equal(t, "[1,2][3,4][7,8]", interval.LoadInputIntervals())
}

func TestLoadIntervalsFallbackEnv(t *testing.T) {
	// Note former args and restore after test execution
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()

	os.Args = []string{"merge-intervals-go"}
	os.Setenv(interval.EnvIntervalsFile, "test-intervals-fallback-env.txt")

	err := os.WriteFile("test-intervals-fallback-env.txt", []byte("[ 100 , 700 ]"), os.ModePerm)
	assert.Nil(t, err)
	assert.Equal(t, "[100,700]", interval.LoadInputIntervals())
}

func TestParseInputArgs(t *testing.T) {
	var output []interval.Interval
	var err error

	// No input
	output, err = interval.ParseInput("")
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "complete number interval")
	assert.Nil(t, output)

	// Missing brackets
	output, err = interval.ParseInput("3, 4")
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "interval format")
	assert.Nil(t, output)

	// Non-conforming interval definition
	output, err = interval.ParseInput("[1,2,3]")
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "start or end number")
	assert.Nil(t, output)

	// Single interval input
	output, err = interval.ParseInput("[1,2]")
	assert.Nil(t, err)
	assert.Equal(t, []interval.Interval{{Start: 1, End: 2}}, output)

	// Multiple intervals, sorted
	output, err = interval.ParseInput("[1,2][3,4]")
	assert.Nil(t, err)
	assert.Equal(t, []interval.Interval{{Start: 1, End: 2}, {Start: 3, End: 4}}, output)

	// Multiple intervals, unsorted
	output, err = interval.ParseInput("[3,4][1,2]")
	assert.Nil(t, err)
	assert.Equal(t, []interval.Interval{{Start: 1, End: 2}, {Start: 3, End: 4}}, output)

	// Multiple intervals, at least one with zero length
	output, err = interval.ParseInput("[3,4][3,3]")
	assert.Nil(t, err)
	assert.Equal(t, []interval.Interval{{Start: 3, End: 3}, {Start: 3, End: 4}}, output) // Sorted by Start, End

	// Multiple intervals with negative numbers, unsorted
	output, err = interval.ParseInput("[-1,2][-11,-3]")
	assert.Nil(t, err)
	assert.Equal(t, []interval.Interval{{Start: -11, End: -3}, {Start: -1, End: 2}}, output)
}

func BenchmarkParseInputArgs(b *testing.B) {
	var intervals []interval.Interval
	var err error
	for n := 0; n < b.N; n++ {
		intervals, err = interval.ParseInput("[7,18][64000,128000][1,100000][2112,4224][0,1][25,117][1998,2401][2024,2025][15,20][117122,122117]")
	}
	// Assert recorded values from benchmarking test to avoid compiler optimizations (and possibly elimination of the function at test)
	if err != nil {
		b.Errorf("error in function while running benchmark: %s, intervals: %v", err.Error(), intervals)
	}
}

func TestFormatOutput(t *testing.T) {
	assert.Equal(t, "[7,15] [26,200]", interval.FormatOutput([]interval.Interval{{Start: 7, End: 15}, {Start: 26, End: 200}}))
}
