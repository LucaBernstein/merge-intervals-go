package interval

import (
	"fmt"
	"log/slog"
	"os"
	"sort"
	"strconv"
	"strings"
)

const EnvIntervalsFile = "INTERVALS_FILE"

func LoadFromArgs() string {
	inputArgs := os.Args[1:] // skip first arg (program)
	return sanitizeInput(strings.Join(inputArgs, ""))
}

func LoadFromFile(filename string) string {
	data, err := os.ReadFile(filename)
	if err != nil {
		slog.Error("Error while loading intervals from file. Returning empty string.", slog.String("err", err.Error()))
		return ""
	}
	return sanitizeInput(string(data))
}

// sanitizeInput removes all blanks and joins probably multiple args together
func sanitizeInput(input string) string {
	sanitizedInput := strings.ReplaceAll(input, " ", "")
	slog.Debug("Input", slog.Any("args", sanitizedInput))
	return sanitizedInput
}

// LoadInputIntervals tries to load intervals to merge from cli args
//
// If no args are passed, it loads intervals from file
// Filename to load from is passed via env var
func LoadInputIntervals() string {
	// First option: Load intervals to merge from cli args
	input := LoadFromArgs()

	// Second option: For larger inputs, the cli args don't work anymore.
	// Hence, this option enables loading from file passed via env var
	if len(input) == 0 && os.Getenv(EnvIntervalsFile) != "" {
		return LoadFromFile(os.Getenv(EnvIntervalsFile))
	}

	return input
}

// ParseInput takes intervals as an argument, parses them into the internal Interval structure and sorts them by Start, End params
func ParseInput(input string) (output []Interval, err error) {
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
