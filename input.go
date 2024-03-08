package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func ParseInputArgs(input string) (output []Interval, err error) {
	if len(input) < 2 {
		return nil, fmt.Errorf("at least one complete number interval must be provided (found: '%s' instead)", input)
	}
	innerValue := input[1 : len(input)-1] // remove outer brackets
	intervals := strings.Split(innerValue, "][")
	for _, interval := range intervals {
		intervalRange := strings.SplitN(interval, ",", 2)
		if len(intervalRange) != 2 {
			return nil, fmt.Errorf("unexpected interval format. Expected '[ <int> , <int> ]', found '%s' instead", interval)
		}
		start, errStart := strconv.Atoi(intervalRange[0])
		end, errEnd := strconv.Atoi(intervalRange[1])
		if errStart != nil || errEnd != nil {
			return nil, fmt.Errorf("unexpected interval start or end number format. Expected '[ <int> , <int> ]', found '%s' instead", interval)
		}
		output = append(output, Interval{Start: start, End: end})
	}
	// sort provided intervals to simplify overlap detection later on
	sort.Slice(output, func(i, j int) bool {
		return output[i].Start < output[j].Start
	})
	return
}
