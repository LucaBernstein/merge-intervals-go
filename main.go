package main

import (
	"fmt"
	"log/slog"

	"github.com/LucaBernstein/merge-intervals-go/interval"
)

func main() {
	slog.SetLogLoggerLevel(slog.LevelWarn) // For debugging purposes, set the level to LevelDebug instead of LevelWarn

	sanitizedInput := interval.LoadInputIntervals()
	intervals, err := interval.ParseInput(sanitizedInput)
	if err != nil {
		slog.Error("Error occurred while parsing input", slog.String("err", err.Error()))
	}
	output := interval.Merge(intervals)
	fmt.Println(interval.FormatOutput(output))
}
