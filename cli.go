package main

import (
	"fmt"
	"log/slog"
	"os"
	"sort"
	"strconv"
	"strings"
)

const envIntervalsFile = "INTERVALS_FILE"

func GetSanitizedInputArgs() string {
	inputArgs := os.Args[1:] // skip first arg (program)

	// For larger inputs, the cli args don't work anymore. Hence, this option enables loading from file passed via env var
	if len(inputArgs) == 0 && os.Getenv(envIntervalsFile) != "" {
		data, err := os.ReadFile(os.Getenv(envIntervalsFile))
		if err != nil {
			slog.Error("Error while loading intervals from file", slog.String("err", err.Error()))
			os.Exit(1)
		}
		inputArgs = []string{string(data)}
	}

	// sanitize input: remove all blanks and join probably multiple args together
	sanitizedInput := strings.ReplaceAll(strings.Join(inputArgs, ""), " ", "")
	slog.Debug("Input", slog.Any("args", sanitizedInput))
	return sanitizedInput
}

func ParseInputArgs(input string) (output []Interval, err error) {
	if len(input) < 2 {
		return nil, fmt.Errorf("at least one complete number interval must be provided (found: '%s' instead)", input)
	}
	// Ensure outer brackets in input
	if input[0] != '[' || input[len(input)-1] != ']' {
		return nil, fmt.Errorf("unexpected interval format. Expected '[ <int> , <int> ]', found '%s' instead", input)
	}
	innerValue := input[1 : len(input)-1] // remove outer brackets
	intervals := strings.Split(innerValue, "][")
	for _, interval := range intervals {
		intervalRange := strings.SplitN(interval, ",", 2)
		if len(intervalRange) != 2 {
			return nil, fmt.Errorf("unexpected interval format. Expected '[ <int> , <int> ]', found '%s' instead", interval)
		}
		// assumption: interval definition itself is ordered
		start, errStart := strconv.Atoi(intervalRange[0])
		end, errEnd := strconv.Atoi(intervalRange[1])
		if errStart != nil || errEnd != nil {
			return nil, fmt.Errorf("unexpected interval start or end number format. Expected '[ <int> , <int> ]', found '%s' instead", interval)
		}
		output = append(output, Interval{Start: start, End: end})
	}
	// sort provided intervals to simplify overlap detection later on
	sort.Slice(output, func(i, j int) bool {
		// Sort by Start, End
		return output[i].Start < output[j].Start || (output[i].Start == output[j].Start && output[i].End < output[j].End)
	})
	slog.Debug("Parsed input", slog.Any("output intervals", output))
	return
}

func FormatOutput(intervals []Interval) (output string) {
	var stringifiedIntervals []string
	for _, interval := range intervals {
		stringifiedIntervals = append(stringifiedIntervals, fmt.Sprintf("[%d,%d]", interval.Start, interval.End))
	}
	return strings.Join(stringifiedIntervals, " ")
}
