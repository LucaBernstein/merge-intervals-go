package main

type Interval struct {
	Start int
	End   int
}

func Merge(input []Interval) (output []Interval) {
	if len(input) < 1 {
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
	return
}
