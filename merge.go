package main

import "log/slog"

type Interval struct {
	Start int
	End   int
}

// Merge takes a list of intervals, merges overlapping intervals, keeps non-overlapping intervals untouched and returns the merged list of all intervals.
// The input interval list must be ordered ascending by each Interval's Start and End values
func Merge(input []Interval) (output []Interval) {
	if len(input) <= 1 { // For 0 or 1 elements, the input intervals are already merged.
		return input
	}
	// Add the first element, as it is already the lowest by sorting
	output = append(output, input[0])
	input = input[1:]
	currentOutputIndex := 0

	for _, interval := range input {
		// Case: no overlap
		//  output:  x─x
		//   input:       x─x
		if interval.Start > output[currentOutputIndex].End {
			output = append(output, interval)
			currentOutputIndex++
		}
		// --> next start on or before current end

		// Case: next end after current end
		//  output:  x─────x
		//   input:    x─────x
		if interval.End > output[currentOutputIndex].End {
			output[currentOutputIndex].End = interval.End
			continue
		}

		// Case: next end before current end
		//  output:  x───────x
		//   input:     x──x
		// --> ignore, as no modification is required
	}
	slog.Debug("Output after merge", slog.Any("intervals", output))
	return
}
