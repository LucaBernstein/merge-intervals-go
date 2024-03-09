package main

import (
	"fmt"
	"log/slog"
	"os"
	"strings"
)

func main() {
	slog.SetLogLoggerLevel(slog.LevelWarn) // For debugging purposes, set the level to LevelDebug instead of LevelWarn

	inputArgs := os.Args[1:] // skip first arg (program)
	// sanitize input: remove all blanks and join probably multiple args together
	sanitizedInput := strings.ReplaceAll(strings.Join(inputArgs, ""), " ", "")

	slog.Debug("Input", slog.Any("args", sanitizedInput))

	intervals, err := ParseInputArgs(sanitizedInput)
	if err != nil {
		slog.Error("Error occurred while parsing input", slog.String("err", err.Error()))
	}

	slog.Debug("Parsed input", slog.Any("intervals", intervals))

	output := Merge(intervals)

	slog.Debug("Output after merge", slog.Any("intervals", output))

	fmt.Println(output) // TODO: Format properly
}
