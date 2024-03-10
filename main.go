package main

import (
	"fmt"
	"log/slog"
)

func main() {
	slog.SetLogLoggerLevel(slog.LevelWarn) // For debugging purposes, set the level to LevelDebug instead of LevelWarn

	sanitizedInput := LoadInputIntervals()
	intervals, err := ParseInput(sanitizedInput)
	if err != nil {
		slog.Error("Error occurred while parsing input", slog.String("err", err.Error()))
	}
	output := Merge(intervals)
	fmt.Println(FormatOutput(output))
}
